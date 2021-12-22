package heappeapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	remoteheappeapi "github.com/lexis-project/lexis-backend-services-api.git/client"
	heappe "github.com/lexis-project/lexis-backend-services-api.git/client/cluster_information"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/cluster_information"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

const (
	SEC_API_KEY   = "KEY"
	SEC_ORG       = "ORG"
	SEC_PRJ       = "PRJ"
	SEC_NAME      = "SHORTNAME"
	SEC_ROLE      = "ROLE"
	SEC_USER      = "USER"
	SEC_USERNAME  = "USERNAME"
	SEC_TOKEN     = "TOKEN"
	CHECK_API_KEY = "API_KEY"
	CHECK_SELF    = "SELF_SERVICE"
	CHECK_ORG     = "OWN_ORG"
	CHECK_PRJ     = "OWN_PRJ"
	ROLE_END_USER = "end_usr"
	NIL           = "00000000-0000-0000-0000-000000000000"
	ORG_ATT       = "org_read"
	PRJ_ATT       = "prj_list"
)

type HEAppEAPI struct {
	authzURL string
}

//Credentials structure for autenticating user to get sessionCode
type Credentials struct {
	Username string `json:"username"`
	Token    string `json:"OpenIdAccessToken"`
}

//CredentialsBag for storing login credentials to compose json for POST request
type CredentialsBag struct {
	Values Credentials `json:"credentials"`
}

func New() *HEAppEAPI {

	return &HEAppEAPI{
		authzURL: "/heappe/UserAndLimitationManagement/AuthenticateUserOpenId",
	}

}

func (p *HEAppEAPI) getHPC(endpoint *string) (string, *remoteheappeapi.LEXISPortalAPI, error) {

	if endpoint != nil {

		url, e := url.Parse(*endpoint)

		if e != nil {

			return "", nil, e

		}

		config := remoteheappeapi.Config{
			URL: url,
		}

		r := remoteheappeapi.New(config)

		authzURL := *endpoint + p.authzURL

		return authzURL, r, nil

	}

	return "", nil, errors.New("The HEAppE endpoint needs to be provided")

}

func (p *HEAppEAPI) authenticateUser(username, token, url string) (string, error) {

	loginDetails := Credentials{
		Username: username,
		Token:    token,
	}

	credentials := CredentialsBag{
		Values: loginDetails,
	}

	jsonData, e := json.Marshal(credentials)

	if e != nil {

		return string(0), e

	}

	resp, e := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if e != nil {

		return string(0), e

	}

	defer resp.Body.Close()

	body, e := ioutil.ReadAll(resp.Body)

	if e != nil {

		return string(0), e

	}

	if resp.StatusCode != 200 {
		var bodyMessage string
		json.Unmarshal([]byte(body), &bodyMessage)
		l.Error.Printf("[HEAppE] Auth error: " + bodyMessage)
		return string(0), errors.New("Returned status " + resp.Status)

	}

	bodyStr := strings.Trim(string(body), "\"")

	if len(bodyStr) != 36 {

		return string(0), errors.New("REST API authentication failed - bad response")

	}

	return bodyStr, e

}

func (p *HEAppEAPI) GetCommandTemplate(ctx context.Context, params cluster_information.GetCommandTemplateParams) middleware.Responder {

	authzData := ctx.Value(restapi.AuthKey).(map[string]string)

	// This is alright... I talked to Jan K. And he tell, we can send anything,
	// he already getting the username from keycloak.
	// The reason why I don't want to send the username is because the heappe
	// api doesn't accept @ in body
	//
	//username := authzData[SEC_USERNAME]
	// if params.Username != nil {
	// 	username = *params.Username
	// 	fmt.Printf("Username passed\n")
	// }

	username := "test"

	url, client, e := p.getHPC(params.Endpoint)

	if e != nil {

		v := "Get HPC [HEAppE]: Something went wrong. Error: " + e.Error()

		eValue := models.ErrorResponse{
			ErrorString: &v,
		}

		return cluster_information.NewGetCommandTemplateInternalServerError().WithPayload(&eValue)

	}

	sessionCode, e := p.authenticateUser(username, authzData[SEC_TOKEN], url)

	if e != nil {

		v := "Auth [HEAppE]: Something went wrong. Error: " + e.Error()

		eValue := models.ErrorResponse{
			ErrorString: &v,
		}

		return cluster_information.NewGetCommandTemplateInternalServerError().WithPayload(&eValue)

	}

	callrequest := params.Request
	callrequest.SessionCode = &sessionCode

	request_params := heappe.NewGetCommandTemplateParams().WithRequest(callrequest)

	r, e := client.ClusterInformation.GetCommandTemplate(ctx, request_params)

	if e != nil {

		l.Warning.Printf("[HEAppE] There was an error while getting the command template. Error: %v.\n", e)

		// Errors masked as 500 because HEAppE doesn't fulfill its own swagger definition

		////	switch e.(type) {

		////	case *heappe.GetCommandTemplateInternalServerError:

		////		v := "Something went wrong. Error: " + e.Error()

		////		eValue := models.ErrorResponse{
		////			ErrorString: &v,
		////		}

		////		return cluster_information.NewGetCommandTemplateInternalServerError().WithPayload(&eValue)

		////	case *heappe.GetCommandTemplateBadRequest:

		////		v := e.(*heappe.GetCommandTemplateBadRequest).GetPayload()

		////		return cluster_information.NewGetCommandTemplateBadRequest().WithPayload(v)

		////	case *heappe.GetCommandTemplateRequestEntityTooLarge:

		////		v := e.(*heappe.GetCommandTemplateRequestEntityTooLarge).GetPayload()

		////		return cluster_information.NewGetCommandTemplateRequestEntityTooLarge().WithPayload(v)

		////	case *heappe.GetCommandTemplateTooManyRequests:

		////		v := e.(*heappe.GetCommandTemplateTooManyRequests).GetPayload()

		////		return cluster_information.NewGetCommandTemplateTooManyRequests().WithPayload(v)

		////	default:

		v := "Something went wrong. Error: " + e.Error()

		eValue := models.ErrorResponse{
			ErrorString: &v,
		}

		return cluster_information.NewGetCommandTemplateInternalServerError().WithPayload(&eValue)

		////	}

	}

	return cluster_information.NewGetCommandTemplateOK().WithPayload(r.Payload)

}

func (p *HEAppEAPI) ListAvailableClusters(ctx context.Context, params cluster_information.ListAvailableClustersParams) middleware.Responder {

	_, client, e := p.getHPC(params.Endpoint)

	if e != nil {

		v := "Something went wrong. Error: " + e.Error()

		eValue := models.ErrorResponse{
			ErrorString: &v,
		}

		return cluster_information.NewListAvailableClustersInternalServerError().WithPayload(&eValue)

	}

	request_params := heappe.NewListAvailableClustersParams()

	r, e := client.ClusterInformation.ListAvailableClusters(ctx, request_params)

	if e != nil {

		l.Warning.Printf("[HEAppE] There was an error while getting the command template. Error: %v.\n", e)

		// Errors masked as 500 because HEAppE doesn't fulfill its own swagger definition

		////	switch e.(type) {

		////	case *heappe.ListAvailableClustersInternalServerError:

		////		v := "Something went wrong. Error: " + e.Error()

		////		eValue := models.ErrorResponse{
		////			ErrorString: &v,
		////		}

		////		return cluster_information.NewListAvailableClustersInternalServerError().WithPayload(&eValue)

		////	case *heappe.ListAvailableClustersBadRequest:

		////		v := e.(*heappe.ListAvailableClustersBadRequest).GetPayload()

		////		return cluster_information.NewListAvailableClustersBadRequest().WithPayload(v)

		////	case *heappe.ListAvailableClustersRequestEntityTooLarge:

		////		v := e.(*heappe.ListAvailableClustersRequestEntityTooLarge).GetPayload()

		////		return cluster_information.NewListAvailableClustersRequestEntityTooLarge().WithPayload(v)

		////	case *heappe.ListAvailableClustersTooManyRequests:

		////		v := e.(*heappe.ListAvailableClustersTooManyRequests).GetPayload()

		////		return cluster_information.NewListAvailableClustersTooManyRequests().WithPayload(v)

		////	default:

		v := "Something went wrong. Error: " + e.Error()

		eValue := models.ErrorResponse{
			ErrorString: &v,
		}

		return cluster_information.NewListAvailableClustersInternalServerError().WithPayload(&eValue)

		////	}

	}

	return cluster_information.NewListAvailableClustersOK().WithPayload(r.Payload)

}
