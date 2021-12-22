package approvalapi

import (
	"context"
	"net/http"
	"strings"

	remoteapprovalapi "github.com/lexis-project/lexis-backend-services-interface-approval-system.git/client"
	remoteapprovalresources "github.com/lexis-project/lexis-backend-services-interface-approval-system.git/client/hpc_resource_management"
	remoteapprovalmodels "github.com/lexis-project/lexis-backend-services-interface-approval-system.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/approval_system_management"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

// A simple in memory CRUD on data
// In real life, this should be using a persistent storage.
type ApprovalSystemAPI struct {
	c remoteapprovalapi.Config
}

// New returns a new Pet manager
func New(c remoteapprovalapi.Config) *ApprovalSystemAPI {

	return &ApprovalSystemAPI{
		c: c,
	}

}

func (p *ApprovalSystemAPI) getClient(param *http.Request) *remoteapprovalapi.ApprovalSystemInterfaceAPI {

	config := p.c

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := remoteapprovalapi.New(config)

	return r

}

func (p *ApprovalSystemAPI) ListResources(ctx context.Context, params approval_system_management.ListResourcesParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	resource_params := remoteapprovalresources.NewListHPCResourcesParams()

	r, e := client.HpcResourceManagement.ListHPCResources(ctx, resource_params)

	if e != nil {

		l.Warning.Printf("[AS] There was an error while retrieving the list of available HPC Resources. Error: %v.\n", e)

		switch e.(type) {
		case *remoteapprovalresources.ListHPCResourcesNotFound:
			v := e.(*remoteapprovalresources.ListHPCResourcesNotFound)
			eValue := models.ApprovalSystemMissingResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewListResourcesNotFound().WithPayload(&eValue)
		default:
			v := e.(*remoteapprovalresources.ListHPCResourcesInternalServerError)
			eValue := models.ApprovalSystemErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewListResourcesInternalServerError().WithPayload(&eValue)
		}
	}

	var rValue []*models.ApprovalSystemResource
	for _, o := range r.Payload {

		var qValue []*models.ApprovalSystemQueue
		for _, q := range o.QueueList {

			newQueue := models.ApprovalSystemQueue{
				ID:            q.ID,
				Name:          q.Name,
				Description:   q.Description,
				Type:          q.Type,
				MaxWallTime:   q.MaxWallTime,
				NumberOfNodes: q.NumberOfNodes,
				CoresPerNode:  q.CoresPerNode,
			}

			qValue = append(qValue, &newQueue)
		}

		newResource := models.ApprovalSystemResource{
			ID:                     o.ID,
			Name:                   o.Name,
			HostName:               o.HostName,
			PerformanceCoefficient: o.PerformanceCoefficient,
			QueueList:              qValue,
		}

		rValue = append(rValue, &newResource)

	}

	return approval_system_management.NewListResourcesOK().WithPayload(rValue)

}

func (p *ApprovalSystemAPI) ListProjectHPCResourceRequest(ctx context.Context, params approval_system_management.ListProjectHPCResourceRequestParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	resource_params := remoteapprovalresources.NewListProjectHPCResourceRequestParams()
	resource_params.AssociatedLEXISProject = params.AssociatedLEXISProject

	r, e := client.HpcResourceManagement.ListProjectHPCResourceRequest(ctx, resource_params)

	if e != nil {

		l.Warning.Printf("[AS] There was an error while retrieving the list of Resource Requests by ID of associated LEXIS project. Error: %v.\n", e)

		switch e.(type) {
		case *remoteapprovalresources.ListProjectHPCResourceRequestNotFound:
			v := e.(*remoteapprovalresources.ListProjectHPCResourceRequestNotFound)
			eValue := models.ApprovalSystemMissingResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewListProjectHPCResourceRequestNotFound().WithPayload(&eValue)
		default:
			v := e.(*remoteapprovalresources.ListProjectHPCResourceRequestInternalServerError)
			eValue := models.ApprovalSystemErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewListProjectHPCResourceRequestInternalServerError().WithPayload(&eValue)
		}
	}

	var rValue []*models.ApprovalSystemResourceRequest
	for _, o := range r.Payload {

		var iValue []*models.ApprovalSystemResourceRequestItem
		for _, q := range o.Resources {

			resourceRequestItem := models.ApprovalSystemResourceRequestItem{
				ClusterID:   q.ClusterID,
				ClusterName: q.ClusterName,
				QueueID:     q.QueueID,
				QueueName:   q.QueueName,
			}

			iValue = append(iValue, &resourceRequestItem)
		}

		newResourceRequest := models.ApprovalSystemResourceRequest{
			HPCResourceID:              o.HPCResourceID,
			AssociatedLEXISProject:     o.AssociatedLEXISProject,
			AssociatedLEXISProjectName: o.AssociatedLEXISProjectName,
			ProjectContactEmail:        o.ProjectContactEmail,
			PrimaryInvestigator:        o.PrimaryInvestigator,
			CoreHoursExpected:          o.CoreHoursExpected,
			Budget:                     o.Budget,
			DateStart:                  o.DateStart,
			DateEnd:                    o.DateEnd,
			ApprovalStatus:             o.ApprovalStatus,
			TermsConsent:               o.TermsConsent,
			ApprovalObjections:         o.ApprovalObjections,
			AssociatedHPCProject:       o.AssociatedHPCProject,
			CloudNetworkName:           o.CloudNetworkName,
			HEAppEEndpoint:             o.HEAppEEndpoint,
			HPCProvider:                o.HPCProvider,
			OpenStackEndpoint:          o.OpenStackEndpoint,
			ResourceType:               o.ResourceType,
			Resources:                  iValue,
		}

		rValue = append(rValue, &newResourceRequest)

	}

	return approval_system_management.NewListProjectHPCResourceRequestOK().WithPayload(rValue)
}

func (p *ApprovalSystemAPI) ListProjectHPCApprovedResourceRequest(ctx context.Context, params approval_system_management.ListProjectHPCApprovedResourceRequestParams) middleware.Responder {
	client := p.getClient(params.HTTPRequest)

	resource_params := remoteapprovalresources.NewListProjectHPCApprovedResourceRequestParams()
	resource_params.AssociatedLEXISProject = params.AssociatedLEXISProject

	r, e := client.HpcResourceManagement.ListProjectHPCApprovedResourceRequest(ctx, resource_params)

	if e != nil {

		l.Warning.Printf("[AS] There was an error while retrieving the list of Approved Resource Requests by ID of associated LEXIS project. Error: %v.\n", e)

		switch e.(type) {
		case *remoteapprovalresources.ListProjectHPCApprovedResourceRequestNotFound:
			v := e.(*remoteapprovalresources.ListProjectHPCApprovedResourceRequestNotFound)
			eValue := models.ApprovalSystemMissingResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewListProjectHPCApprovedResourceRequestNotFound().WithPayload(&eValue)
		default:
			v := e.(*remoteapprovalresources.ListProjectHPCApprovedResourceRequestInternalServerError)
			eValue := models.ApprovalSystemErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewListProjectHPCApprovedResourceRequestInternalServerError().WithPayload(&eValue)
		}
	}

	var rValue []*models.ApprovalSystemApprovedResourceRequest
	for _, o := range r.Payload {

		newResourceRequest := models.ApprovalSystemApprovedResourceRequest{
			HPCResourceID:              o.HPCResourceID,
			ApprovalObjections:         o.ApprovalObjections,
			ApprovalStatus:             o.ApprovalStatus,
			AssociatedHPCProject:       o.AssociatedHPCProject,
			AssociatedLEXISProject:     o.AssociatedLEXISProject,
			AssociatedLEXISProjectName: o.AssociatedLEXISProjectName,
			CloudNetworkName:           o.CloudNetworkName,
			HEAppEEndpoint:             o.HEAppEEndpoint,
			HPCProvider:                o.HPCProvider,
			OpenStackEndpoint:          o.OpenStackEndpoint,
			PrimaryInvestigator:        o.PrimaryInvestigator,
			ProjectContactEmail:        o.ProjectContactEmail,
			ResourceType:               o.ResourceType,
			TermsConsent:               o.TermsConsent,
		}

		rValue = append(rValue, &newResourceRequest)
	}

	return approval_system_management.NewListProjectHPCApprovedResourceRequestOK().WithPayload(rValue)
}

func (h *ApprovalSystemAPI) CreateResourceRequest(ctx context.Context, params approval_system_management.CreateResourceRequestParams) middleware.Responder {

	var rValue []*remoteapprovalmodels.HPCResourceRequestItem
	for _, r := range params.ResourceRequest.Resources {

		resourceRequestItem := remoteapprovalmodels.HPCResourceRequestItem{
			ClusterID:   r.ClusterID,
			ClusterName: r.ClusterName,
			QueueID:     r.QueueID,
			QueueName:   r.QueueName,
		}

		rValue = append(rValue, &resourceRequestItem)
	}

	newResourceRequest := remoteapprovalmodels.HPCResourceRequest{
		AssociatedLEXISProject:     params.ResourceRequest.AssociatedLEXISProject,
		AssociatedLEXISProjectName: params.ResourceRequest.AssociatedLEXISProjectName,
		ProjectContactEmail:        params.ResourceRequest.ProjectContactEmail,
		PrimaryInvestigator:        params.ResourceRequest.PrimaryInvestigator,
		CoreHoursExpected:          params.ResourceRequest.CoreHoursExpected,
		Budget:                     params.ResourceRequest.Budget,
		DateStart:                  params.ResourceRequest.DateStart,
		DateEnd:                    params.ResourceRequest.DateEnd,
		ApprovalStatus:             params.ResourceRequest.ApprovalStatus,
		TermsConsent:               params.ResourceRequest.TermsConsent,
		ApprovalObjections:         params.ResourceRequest.ApprovalObjections,
		AssociatedHPCProject:       params.ResourceRequest.AssociatedHPCProject,
		CloudNetworkName:           params.ResourceRequest.CloudNetworkName,
		HEAppEEndpoint:             params.ResourceRequest.HEAppEEndpoint,
		HPCProvider:                params.ResourceRequest.HPCProvider,
		OpenStackEndpoint:          params.ResourceRequest.OpenStackEndpoint,
		ResourceType:               params.ResourceRequest.ResourceType,
		Resources:                  rValue,
	}

	client := h.getClient(params.HTTPRequest)

	as_params := remoteapprovalresources.NewCreateResourceRequestParams().WithResourceRequest(&newResourceRequest)

	_, e := client.HpcResourceManagement.CreateResourceRequest(ctx, as_params)

	if e != nil {

		l.Warning.Printf("[AS] There was an error while creating the Resource Request. Error: %v.\n", e)

		switch e.(type) {
		case *remoteapprovalresources.CreateResourceRequestBadRequest:
			v := e.(*remoteapprovalresources.CreateResourceRequestBadRequest)
			eValue := models.ApprovalSystemInvalidResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewCreateResourceRequestBadRequest().WithPayload(&eValue)
		case *remoteapprovalresources.CreateResourceRequestNotFound:
			v := e.(*remoteapprovalresources.CreateResourceRequestNotFound)
			eValue := models.ApprovalSystemMissingResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewCreateResourceRequestNotFound().WithPayload(&eValue)
		default:
			v := e.(*remoteapprovalresources.CreateResourceRequestInternalServerError)
			eValue := models.ApprovalSystemErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewCreateResourceRequestInternalServerError().WithPayload(&eValue)
		}
	}

	return approval_system_management.NewCreateResourceRequestCreated()
}

func (h *ApprovalSystemAPI) CreateApprovedResourceRequest(ctx context.Context, params approval_system_management.CreateApprovedResourceRequestParams) middleware.Responder {

	newResourceRequest := remoteapprovalmodels.HPCApprovedResourcesRequest{
		ApprovalObjections:         params.ApprovedResourceRequest.ApprovalObjections,
		ApprovalStatus:             params.ApprovedResourceRequest.ApprovalStatus,
		AssociatedHPCProject:       params.ApprovedResourceRequest.AssociatedHPCProject,
		AssociatedLEXISProject:     params.ApprovedResourceRequest.AssociatedLEXISProject,
		AssociatedLEXISProjectName: params.ApprovedResourceRequest.AssociatedLEXISProjectName,
		CloudNetworkName:           params.ApprovedResourceRequest.CloudNetworkName,
		HEAppEEndpoint:             params.ApprovedResourceRequest.HEAppEEndpoint,
		HPCProvider:                params.ApprovedResourceRequest.HPCProvider,
		OpenStackEndpoint:          params.ApprovedResourceRequest.OpenStackEndpoint,
		PrimaryInvestigator:        params.ApprovedResourceRequest.PrimaryInvestigator,
		ProjectContactEmail:        params.ApprovedResourceRequest.ProjectContactEmail,
		ResourceType:               params.ApprovedResourceRequest.ResourceType,
		TermsConsent:               params.ApprovedResourceRequest.TermsConsent,
	}

	client := h.getClient(params.HTTPRequest)

	as_params := remoteapprovalresources.NewCreateApprovedResourceRequestParams().WithApprovedResourceRequest(&newResourceRequest)

	_, e := client.HpcResourceManagement.CreateApprovedResourceRequest(ctx, as_params)

	if e != nil {

		l.Warning.Printf("[AS] There was an error while creating the Approved Resource Request. Error: %v.\n", e)

		switch e.(type) {
		case *remoteapprovalresources.CreateApprovedResourceRequestBadRequest:
			v := e.(*remoteapprovalresources.CreateApprovedResourceRequestBadRequest)
			eValue := models.ApprovalSystemInvalidResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewCreateApprovedResourceRequestBadRequest().WithPayload(&eValue)
		case *remoteapprovalresources.CreateApprovedResourceRequestNotFound:
			v := e.(*remoteapprovalresources.CreateApprovedResourceRequestNotFound)
			eValue := models.ApprovalSystemMissingResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewCreateApprovedResourceRequestNotFound().WithPayload(&eValue)
		default:
			v := e.(*remoteapprovalresources.CreateApprovedResourceRequestInternalServerError)
			eValue := models.ApprovalSystemErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewCreateApprovedResourceRequestInternalServerError().WithPayload(&eValue)
		}
	}

	return approval_system_management.NewCreateApprovedResourceRequestCreated()
}

func (p *ApprovalSystemAPI) HPCResourceRequest(ctx context.Context, params approval_system_management.HPCResourceRequestParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	resource_params := remoteapprovalresources.NewHPCResourceRequestParams()
	resource_params.HPCResourceID = params.HPCResourceID

	r, e := client.HpcResourceManagement.HPCResourceRequest(ctx, resource_params)

	if e != nil {

		l.Warning.Printf("[AS] There was an error while retrieving the Resource Request (by its HPCResourceID). Error: %v.\n", e)

		switch e.(type) {
		case *remoteapprovalresources.HPCResourceRequestNotFound:
			v := e.(*remoteapprovalresources.HPCResourceRequestNotFound)
			eValue := models.ApprovalSystemMissingResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewHPCResourceRequestNotFound().WithPayload(&eValue)
		default:
			v := e.(*remoteapprovalresources.HPCResourceRequestInternalServerError)
			eValue := models.ApprovalSystemErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewHPCResourceRequestInternalServerError().WithPayload(&eValue)
		}
	}

	var iValue []*models.ApprovalSystemResourceRequestItem
	for _, q := range r.Payload.Resources {

		resourceRequestItem := models.ApprovalSystemResourceRequestItem{
			ClusterID:   q.ClusterID,
			ClusterName: q.ClusterName,
			QueueID:     q.QueueID,
			QueueName:   q.QueueName,
		}

		iValue = append(iValue, &resourceRequestItem)
	}

	newResourceRequest := models.ApprovalSystemResourceRequest{
		HPCResourceID:              r.Payload.HPCResourceID,
		AssociatedLEXISProject:     r.Payload.AssociatedLEXISProject,
		AssociatedLEXISProjectName: r.Payload.AssociatedLEXISProjectName,
		ProjectContactEmail:        r.Payload.ProjectContactEmail,
		PrimaryInvestigator:        r.Payload.PrimaryInvestigator,
		CoreHoursExpected:          r.Payload.CoreHoursExpected,
		Budget:                     r.Payload.Budget,
		DateStart:                  r.Payload.DateStart,
		DateEnd:                    r.Payload.DateEnd,
		ApprovalStatus:             r.Payload.ApprovalStatus,
		TermsConsent:               r.Payload.TermsConsent,
		ApprovalObjections:         r.Payload.ApprovalObjections,
		AssociatedHPCProject:       r.Payload.AssociatedHPCProject,
		CloudNetworkName:           r.Payload.CloudNetworkName,
		HEAppEEndpoint:             r.Payload.HEAppEEndpoint,
		HPCProvider:                r.Payload.HPCProvider,
		OpenStackEndpoint:          r.Payload.OpenStackEndpoint,
		ResourceType:               r.Payload.ResourceType,
		Resources:                  iValue,
	}

	return approval_system_management.NewHPCResourceRequestOK().WithPayload(&newResourceRequest)
}

func (p *ApprovalSystemAPI) HPCApprovedResourceRequest(ctx context.Context, params approval_system_management.HPCApprovedResourceRequestParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	resource_params := remoteapprovalresources.NewHPCApprovedResourceRequestParams()
	resource_params.HPCResourceID = params.HPCResourceID

	r, e := client.HpcResourceManagement.HPCApprovedResourceRequest(ctx, resource_params)

	if e != nil {

		l.Warning.Printf("[AS] There was an error while retrieving the Approved Resource Request (by its HPCResourceID). Error: %v.\n", e)

		switch e.(type) {
		case *remoteapprovalresources.HPCApprovedResourceRequestNotFound:
			v := e.(*remoteapprovalresources.HPCApprovedResourceRequestNotFound)
			eValue := models.ApprovalSystemMissingResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewHPCApprovedResourceRequestNotFound().WithPayload(&eValue)
		default:
			v := e.(*remoteapprovalresources.HPCApprovedResourceRequestInternalServerError)
			eValue := models.ApprovalSystemErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return approval_system_management.NewHPCApprovedResourceRequestInternalServerError().WithPayload(&eValue)
		}

	}

	newResourceRequest := models.ApprovalSystemApprovedResourceRequest{
		HPCResourceID:              r.Payload.HPCResourceID,
		ApprovalObjections:         r.Payload.ApprovalObjections,
		ApprovalStatus:             r.Payload.ApprovalStatus,
		AssociatedHPCProject:       r.Payload.AssociatedHPCProject,
		AssociatedLEXISProject:     r.Payload.AssociatedLEXISProject,
		AssociatedLEXISProjectName: r.Payload.AssociatedLEXISProjectName,
		CloudNetworkName:           r.Payload.CloudNetworkName,
		HEAppEEndpoint:             r.Payload.HEAppEEndpoint,
		HPCProvider:                r.Payload.HPCProvider,
		OpenStackEndpoint:          r.Payload.OpenStackEndpoint,
		PrimaryInvestigator:        r.Payload.PrimaryInvestigator,
		ProjectContactEmail:        r.Payload.ProjectContactEmail,
		ResourceType:               r.Payload.ResourceType,
		TermsConsent:               r.Payload.TermsConsent,
	}

	return approval_system_management.NewHPCApprovedResourceRequestOK().WithPayload(&newResourceRequest)
}
