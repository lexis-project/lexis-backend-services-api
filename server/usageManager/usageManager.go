package usageManager

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/encoding/json"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/usage_management"
	userorgapi "github.com/lexis-project/lexis-backend-services-userorg-service.git/client"
	hpcresources "github.com/lexis-project/lexis-backend-services-userorg-service.git/client/hpc_management"
	projects "github.com/lexis-project/lexis-backend-services-userorg-service.git/client/project_management"
	userorgmodels "github.com/lexis-project/lexis-backend-services-userorg-service.git/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	cdrapi "gitlab.com/cyclops-community/cdr-client-interface/client"
	cdrusage "gitlab.com/cyclops-community/cdr-client-interface/client/usage_management"
	cdrmodels "gitlab.com/cyclops-community/cdr-client-interface/models"
	csapi "gitlab.com/cyclops-community/cs-client-interface/client"
	cscredit "gitlab.com/cyclops-community/cs-client-interface/client/credit_management"
	udrapi "gitlab.com/cyclops-community/udr-client-interface/client"
	datamodels "gitlab.com/cyclops-utilities/datamodels"
	l "gitlab.com/cyclops-utilities/logging"
)

// A simple in memory CRUD on data
// In real life, this should be using a persistent storage.
type UsageManager struct {
	userorgConfig userorgapi.Config
	udrConfig     udrapi.Config
	cdrConfig     cdrapi.Config
	csConfig      csapi.Config
}

// New returns a new Pet manager
func New(userorgConfig userorgapi.Config, udrConfig udrapi.Config, cdrConfig cdrapi.Config, csConfig csapi.Config) *UsageManager {

	return &UsageManager{
		userorgConfig: userorgConfig,
		udrConfig:     udrConfig,
		cdrConfig:     cdrConfig,
		csConfig:      csConfig,
	}

}

func (m *UsageManager) getClientUO(param *http.Request) *userorgapi.UserOrganizationAPI {

	config := m.userorgConfig

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := userorgapi.New(config)

	return r

}

func (m *UsageManager) getClientUDR(param *http.Request) *udrapi.UDRManagementAPI {

	config := m.udrConfig

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := udrapi.New(config)

	return r

}

func (m *UsageManager) getClientCS(param *http.Request) *csapi.CreditSystemManagementAPI {

	config := m.csConfig

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := csapi.New(config)

	return r

}

func (m *UsageManager) getClientCDR(param *http.Request) *cdrapi.CDRManagementAPI {

	config := m.cdrConfig

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := cdrapi.New(config)

	return r

}

func (m *UsageManager) GetUsage(ctx context.Context, params usage_management.GetUsageParams) middleware.Responder {

	var cdr []*cdrmodels.CReport
	var hpcs []*userorgmodels.HPCResource

	to := strfmt.DateTime(time.Now())

	if params.To != nil {

		to = *params.To

	}

	// Get hpcresources of project
	UO := m.getClientUO(params.HTTPRequest)

	hpc_params := hpcresources.NewListHPCResourcesParams()

	h, e := UO.HpcManagement.ListHPCResources(ctx, hpc_params)

	if e != nil {

		v := e.Error()
		eValue := models.ErrorResponse{
			ErrorString: &v,
		}

		return usage_management.NewGetUsageInternalServerError().WithPayload(&eValue)

	}

	for i := range h.Payload {

		if h.Payload[i].AssociatedLEXISProject == params.ID {

			hpcs = append(hpcs, h.Payload[i])

		}

	}

	projectParams := projects.NewGetProjectParams().WithID(params.ID)

	p, e := UO.ProjectManagement.GetProject(ctx, projectParams)

	if e != nil {

		v := e.Error()
		eValue := models.ErrorResponse{
			ErrorString: &v,
		}

		return usage_management.NewGetUsageInternalServerError().WithPayload(&eValue)

	}

	project := *p.Payload

	from := project.ProjectStartDate

	if params.From != nil {

		from = *params.From

	}

	var pList []string

	for _, p := range hpcs {

		pList = append(pList, p.OpenStackProjectID, p.HPCResourceID)

	}

	idlist := strings.Join(pList, ",")

	// Get the CDR usage for projects
	C := m.getClientCDR(params.HTTPRequest)

	cdr_params := cdrusage.NewGetSystemUsageParams().WithIdlist(&idlist).WithTo(&to).WithFrom(&from)

	c, e := C.UsageManagement.GetSystemUsage(ctx, cdr_params)

	cdr = c.Payload

	// Processing starts

	rValue := models.Usage{
		ProjectID:   params.ID,
		ProjectName: project.ProjectName,
		TimeTo:      to,
		TimeFrom:    from,
	}

	var hpcuses []*models.HPCUsage

	for i := range hpcs {

		newUsage := models.HPCUsage{
			HPCProjectID: hpcs[i].HPCResourceID,
		}

		var newAccs []datamodels.JSONdb

		for id := range cdr {

			if cdr[id].AccountID != hpcs[i].HPCResourceID && cdr[id].AccountID != hpcs[i].OpenStackProjectID {

				continue

			}

			for ids := range cdr[id].Usage {

				newAc := make(datamodels.JSONdb)

				newAc["resource_id"] = cdr[id].Usage[ids].ResourceID
				newAc["resource_name"] = cdr[id].Usage[ids].ResourceName
				newAc["resource_type"] = cdr[id].Usage[ids].ResourceType
				newAc["description"] = fmt.Sprintf("%v: [%v] (%v)", cdr[id].Usage[ids].ResourceType, cdr[id].Usage[ids].ResourceName, cdr[id].Usage[ids].Unit)
				newAc["unit"] = cdr[id].Usage[ids].Unit
				newAc["usage"] = cdr[id].Usage[ids].UsageBreakup

				if value, exists := cdr[id].Usage[ids].Metadata["flavorname"]; exists {

					newAc["flavor"] = value

				}

				if value, exists := cdr[id].Usage[ids].Metadata["size"]; exists {

					newAc["size"] = value

				}

				if value, exists := cdr[id].Usage[ids].Cost["netTotal"]; exists {

					newAc["cost"] = value

				}

				newAccs = append(newAccs, newAc)

			}

		}

		var cleanAccs []datamodels.JSONdb

		for _, ac := range newAccs {

			exists := false

			for i, ca := range cleanAccs {

				if ac["resource_id"].(string) == ca["resource_id"].(string) &&
					ac["description"].(string) == ca["description"].(string) {

					exists = true

					uses := make(datamodels.JSONdb)

					for _, use := range []string{"active", "inactive", "error", "used", "terminated"} {

						var newUse, oldUse float64

						if v, exists := ca["usage"].(datamodels.JSONdb)[use]; exists {

							newUse = m.getFloat(v)

						}

						if v, exists := ac["usage"].(datamodels.JSONdb)[use]; exists {

							oldUse = m.getFloat(v)

						}

						if total := newUse + oldUse; total > float64(0) {

							uses[use] = total

						}

					}

					newCost := m.getFloat(ca["cost"])
					oldCost := m.getFloat(ac["cost"])

					cleanAccs[i]["cost"] = newCost + oldCost
					cleanAccs[i]["usage"] = uses

				}

			}

			if !exists {

				cleanAccs = append(cleanAccs, ac)

			}

		}

		if len(cleanAccs) < 1 {

			continue

		}

		var summaryAccs []datamodels.JSONdb

		for _, ac := range cleanAccs {

			exists := false

			for i, ca := range summaryAccs {

				if ac["resource_type"].(string) == ca["resource_type"].(string) &&
					(ac["flavor"] == nil || ac["flavor"].(string) == ca["flavor"].(string)) {

					exists = true

					uses := make(datamodels.JSONdb)

					for _, use := range []string{"active", "inactive", "error", "used", "terminated"} {

						var newUse, oldUse float64

						if v, exists := ca["usage"].(datamodels.JSONdb)[use]; exists {

							newUse = m.getFloat(v)

						}

						if v, exists := ac["usage"].(datamodels.JSONdb)[use]; exists {

							oldUse = m.getFloat(v)

						}

						if total := newUse + oldUse; total > float64(0) {

							uses[use] = total

						}

					}

					newCost := m.getFloat(ca["cost"])
					oldCost := m.getFloat(ac["cost"])

					summaryAccs[i]["cost"] = math.Round((newCost+oldCost)*100) / 100
					summaryAccs[i]["usage"] = uses
					summaryAccs[i]["resources_amount"] = summaryAccs[i]["resources_amount"].(int64) + 1

					if value, exists := ac["size"]; exists {

						size, _ := strconv.ParseInt(value.(string), 10, 0)
						summaryAccs[i]["size"] = summaryAccs[i]["size"].(int64) + size

					}

				}

			}

			if !exists {

				newAC := make(datamodels.JSONdb)

				newAC["usage"] = ac["usage"]
				newAC["resource_type"] = ac["resource_type"]

				if value, exists := ac["flavor"]; exists {

					newAC["flavor"] = value

				}

				if value, exists := ac["size"]; exists {

					size, _ := strconv.ParseInt(value.(string), 10, 0)
					newAC["size"] = size

				}

				if value, exists := ac["cost"]; exists {

					newAC["cost"] = math.Round(m.getFloat(value)*100) / 100

				}

				newAC["resources_amount"] = int64(1)

				summaryAccs = append(summaryAccs, newAC)

			}

		}

		if len(summaryAccs) < 1 {

			continue

		}

		newUsage.AccountingData = summaryAccs

		hpcuses = append(hpcuses, &newUsage)

	}

	rValue.HPCProjects = hpcuses

	return usage_management.NewGetUsageOK().WithPayload(&rValue)

}

// getFloat
func (m *UsageManager) getFloat(i interface{}) (f float64) {

	var e error

	if i == nil {

		f = float64(0)

		return

	}

	if v := reflect.ValueOf(i); v.Kind() == reflect.Float64 {

		f = i.(float64)

	} else {

		f, e = i.(json.Number).Float64()

		if e != nil {

			l.Trace.Printf("[UsageManager] GetFloat failed to convert [ %v ]. Error: %v\n", i, e)

		}

	}

	return

}

func (m *UsageManager) GetCredit(ctx context.Context, params usage_management.GetCreditParams) middleware.Responder {

	client := m.getClientCS(params.HTTPRequest)

	cs_params := cscredit.NewGetCreditParams().WithID(params.ID.String())

	r, e := client.CreditManagement.GetCredit(ctx, cs_params)

	if e != nil {

		l.Warning.Printf("[CYC] There was an error while retrieving the Credit Status. Error: %v.\n", e)

		switch e.(type) {

		case *cscredit.GetCreditNotFound:

			v := e.(*cscredit.GetCreditNotFound)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewGetCreditNotFound().WithPayload(&eValue)

		case *cscredit.GetCreditInternalServerError:

			v := e.(*cscredit.GetCreditInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewGetCreditInternalServerError().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return usage_management.NewGetCreditInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.CreditStatus{
		AccountID:       r.Payload.AccountID,
		AvailableCredit: r.Payload.AvailableCredit,
		AvailableCash:   r.Payload.AvailableCash,
		LastUpdate:      r.Payload.LastUpdate,
	}

	return usage_management.NewGetCreditOK().WithPayload(&rValue)

}

func (m *UsageManager) GetHistory(ctx context.Context, params usage_management.GetHistoryParams) middleware.Responder {

	client := m.getClientCS(params.HTTPRequest)

	cs_params := cscredit.NewGetHistoryParams().WithID(params.ID.String())

	if params.FilterSystem != nil {

		cs_params.WithFilterSystem(params.FilterSystem)
	}

	if params.Medium != nil {

		cs_params.WithMedium(params.Medium)

	}

	r, e := client.CreditManagement.GetHistory(ctx, cs_params)

	if e != nil {

		l.Warning.Printf("[CYC] There was an error while retrieving the Credit Status. Error: %v.\n", e)

		switch e.(type) {

		case *cscredit.GetHistoryNotFound:

			v := e.(*cscredit.GetHistoryNotFound)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewGetHistoryNotFound().WithPayload(&eValue)

		case *cscredit.GetHistoryInternalServerError:

			v := e.(*cscredit.GetHistoryInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewGetHistoryInternalServerError().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return usage_management.NewGetHistoryInternalServerError().WithPayload(&eValue)

		}

	}

	var events []*models.Event

	for i := range r.Payload.Events {

		event := models.Event{
			AuthorizedBy: r.Payload.Events[i].AuthorizedBy,
			Delta:        r.Payload.Events[i].Delta,
			EventType:    r.Payload.Events[i].EventType,
			Timestamp:    r.Payload.Events[i].Timestamp,
			Medium:       r.Payload.Events[i].Medium,
		}

		events = append(events, &event)

	}

	rValue := models.CreditHistory{
		AccountID: r.Payload.AccountID,
		Events:    events,
	}

	return usage_management.NewGetHistoryOK().WithPayload(&rValue)

}

func (m *UsageManager) DecreaseCredit(ctx context.Context, params usage_management.DecreaseCreditParams) middleware.Responder {

	client := m.getClientCS(params.HTTPRequest)

	cs_params := cscredit.NewDecreaseCreditParams().WithID(params.ID.String()).WithMedium(params.Medium).WithAmount(params.Amount)

	r, e := client.CreditManagement.DecreaseCredit(ctx, cs_params)

	if e != nil {

		l.Warning.Printf("[CYC] There was an error while decreasing the Credit. Error: %v.\n", e)

		switch e.(type) {

		case *cscredit.DecreaseCreditNotFound:

			v := e.(*cscredit.DecreaseCreditNotFound)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewDecreaseCreditNotFound().WithPayload(&eValue)

		case *cscredit.DecreaseCreditInternalServerError:

			v := e.(*cscredit.DecreaseCreditInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewDecreaseCreditInternalServerError().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return usage_management.NewDecreaseCreditInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.CreditStatus{
		AccountID:       r.Payload.AccountID,
		AvailableCredit: r.Payload.AvailableCredit,
		AvailableCash:   r.Payload.AvailableCash,
		LastUpdate:      r.Payload.LastUpdate,
	}

	return usage_management.NewDecreaseCreditOK().WithPayload(&rValue)

}

func (m *UsageManager) IncreaseCredit(ctx context.Context, params usage_management.IncreaseCreditParams) middleware.Responder {

	client := m.getClientCS(params.HTTPRequest)

	cs_params := cscredit.NewIncreaseCreditParams().WithID(params.ID.String()).WithMedium(params.Medium).WithAmount(params.Amount)

	r, e := client.CreditManagement.IncreaseCredit(ctx, cs_params)

	if e != nil {

		l.Warning.Printf("[CYC] There was an error while increasing the Credit. Error: %v.\n", e)

		switch e.(type) {

		case *cscredit.IncreaseCreditNotFound:

			v := e.(*cscredit.IncreaseCreditNotFound)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewIncreaseCreditNotFound().WithPayload(&eValue)

		case *cscredit.IncreaseCreditInternalServerError:

			v := e.(*cscredit.IncreaseCreditInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewIncreaseCreditInternalServerError().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return usage_management.NewIncreaseCreditInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.CreditStatus{
		AccountID:       r.Payload.AccountID,
		AvailableCredit: r.Payload.AvailableCredit,
		AvailableCash:   r.Payload.AvailableCash,
		LastUpdate:      r.Payload.LastUpdate,
	}

	return usage_management.NewIncreaseCreditOK().WithPayload(&rValue)

}

func (m *UsageManager) AddConsumption(ctx context.Context, params usage_management.AddConsumptionParams) middleware.Responder {

	client := m.getClientCS(params.HTTPRequest)

	cs_params := cscredit.NewAddConsumptionParams().WithID(params.ID.String()).WithMedium(params.Medium).WithAmount(params.Amount)

	r, e := client.CreditManagement.AddConsumption(ctx, cs_params)

	if e != nil {

		l.Warning.Printf("[CYC] There was an error while generating a Credit consumption. Error: %v.\n", e)

		switch e.(type) {

		case *cscredit.AddConsumptionNotFound:

			v := e.(*cscredit.AddConsumptionNotFound)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewAddConsumptionNotFound().WithPayload(&eValue)

		case *cscredit.AddConsumptionInternalServerError:

			v := e.(*cscredit.AddConsumptionInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}

			return usage_management.NewAddConsumptionInternalServerError().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return usage_management.NewAddConsumptionInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.CreditStatus{
		AccountID:       r.Payload.AccountID,
		AvailableCredit: r.Payload.AvailableCredit,
		AvailableCash:   r.Payload.AvailableCash,
		LastUpdate:      r.Payload.LastUpdate,
	}

	return usage_management.NewAddConsumptionOK().WithPayload(&rValue)

}
