package userorgapi

import (
	"context"
	"net/http"
	"strings"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/user_org_management"
	remoteuserorgapi "github.com/lexis-project/lexis-backend-services-userorg-service.git/client"
	access "github.com/lexis-project/lexis-backend-services-userorg-service.git/client/access_management"
	hpcres "github.com/lexis-project/lexis-backend-services-userorg-service.git/client/hpc_management"
	orgs "github.com/lexis-project/lexis-backend-services-userorg-service.git/client/organization_management"
	projects "github.com/lexis-project/lexis-backend-services-userorg-service.git/client/project_management"
	users "github.com/lexis-project/lexis-backend-services-userorg-service.git/client/user_management"
	uomodels "github.com/lexis-project/lexis-backend-services-userorg-service.git/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

// A simple in memory CRUD on data
// In real life, this should be using a persistent storage.
type UserOrgAPI struct {
	c remoteuserorgapi.Config
}

// New returns a new Pet manager
func New(c remoteuserorgapi.Config) *UserOrgAPI {

	return &UserOrgAPI{
		c: c,
	}

}

func (p *UserOrgAPI) getClient(param *http.Request) *remoteuserorgapi.UserOrganizationAPI {

	config := p.c

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := remoteuserorgapi.New(config)

	return r

}

func (p *UserOrgAPI) AddRole(ctx context.Context, params user_org_management.AddRoleParams) middleware.Responder {

	role_params := access.NewAddRoleParams()

	client := p.getClient(params.HTTPRequest)

	role_params.WithUserID(params.UserID).WithOrganizationID(params.OrganizationID).WithRole(params.Role)

	if params.ProjectID != nil {

		role_params = role_params.WithProjectID(params.ProjectID)

	}

	if params.ProjectShortName != nil {

		role_params = role_params.WithProjectShortName(params.ProjectShortName)

	}

	r, e := client.AccessManagement.AddRole(ctx, role_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while adding the role to the user. Error: %v.\n", e)

		switch e.(type) {

		case *access.AddRoleInternalServerError:

			v := e.(*access.AddRoleInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddRoleInternalServerError().WithPayload(&eValue)

		case *access.AddRoleUnauthorized:

			v := e.(*access.AddRoleUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddRoleUnauthorized().WithPayload(&eValue)

		case *access.AddRoleForbidden:

			v := e.(*access.AddRoleForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddRoleForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewAddRoleInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.AuthorizationResponse{
		Message: r.Payload.Message,
	}

	return user_org_management.NewAddRoleOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) ClearRole(ctx context.Context, params user_org_management.ClearRoleParams) middleware.Responder {

	role_params := access.NewClearRoleParams()

	client := p.getClient(params.HTTPRequest)

	role_params.WithUserID(params.UserID).WithOrganizationID(params.OrganizationID)

	if params.ProjectID != nil {

		role_params = role_params.WithProjectID(params.ProjectID)

	}

	if params.ProjectShortName != nil {

		role_params = role_params.WithProjectShortName(params.ProjectShortName)

	}

	r, e := client.AccessManagement.ClearRole(ctx, role_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while clearing the role for the user. Error: %v.\n", e)

		switch e.(type) {

		case *access.ClearRoleInternalServerError:

			v := e.(*access.ClearRoleInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewClearRoleInternalServerError().WithPayload(&eValue)

		case *access.ClearRoleUnauthorized:

			v := e.(*access.ClearRoleUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewClearRoleUnauthorized().WithPayload(&eValue)

		case *access.ClearRoleForbidden:

			v := e.(*access.ClearRoleForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewClearRoleForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewClearRoleInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.AuthorizationResponse{
		Message: r.Payload.Message,
	}

	return user_org_management.NewClearRoleOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) CreateHPCResource(ctx context.Context, params user_org_management.CreateHPCResourceParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	o := uomodels.HPCResource{
		AssociatedLEXISProject: params.HPCResource.AssociatedLEXISProject,
		AssociatedHPCProject:   params.HPCResource.AssociatedHPCProject,
		ApprovalStatus:         params.HPCResource.ApprovalStatus,
		CloudNetworkName:       params.HPCResource.CloudNetworkName,
		ProjectNetworkName:     params.HPCResource.ProjectNetworkName,
		HEAppEEndpoint:         params.HPCResource.HEAppEEndpoint,
		HPCProvider:            params.HPCResource.HPCProvider,
		HPCResourceID:          params.HPCResource.HPCResourceID,
		OpenStackEndpoint:      params.HPCResource.OpenStackEndpoint,
		OpenStackProjectID:     params.HPCResource.OpenStackProjectID,
		ResourceType:           params.HPCResource.ResourceType,
		TermsConsent:           params.HPCResource.TermsConsent,
	}

	hpc_params := hpcres.NewCreateHPCResourceParams().WithHPCResource(&o)

	r, e := client.HpcManagement.CreateHPCResource(ctx, hpc_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while creating the HPCResource. Error: %v.\n", e)

		switch e.(type) {

		case *hpcres.CreateHPCResourceConflict:

			v := e.(*hpcres.CreateHPCResourceConflict)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateHPCResourceConflict().WithPayload(&eValue)

		case *hpcres.CreateHPCResourceInternalServerError:

			v := e.(*hpcres.CreateHPCResourceInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateHPCResourceInternalServerError().WithPayload(&eValue)

		case *hpcres.CreateHPCResourceUnauthorized:

			v := e.(*hpcres.CreateHPCResourceUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateHPCResourceUnauthorized().WithPayload(&eValue)

		case *hpcres.CreateHPCResourceForbidden:

			v := e.(*hpcres.CreateHPCResourceForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateHPCResourceForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewCreateHPCResourceInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.ItemCreatedResponse{
		ID:   r.Payload.ID,
		Link: &r.Payload.Link,
	}

	return user_org_management.NewCreateHPCResourceCreated().WithPayload(&rValue)

}

func (p *UserOrgAPI) CreateOrganization(ctx context.Context, params user_org_management.CreateOrganizationParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	o := uomodels.Organization{
		CreatedBy:                params.Organization.CreatedBy,
		CreationDate:             params.Organization.CreationDate,
		FormalName:               params.Organization.FormalName,
		ID:                       params.Organization.ID,
		OrganizationEmailAddress: params.Organization.OrganizationEmailAddress,
		PrimaryTelephoneNumber:   params.Organization.PrimaryTelephoneNumber,
		RegisteredAddress1:       params.Organization.RegisteredAddress1,
		RegisteredAddress2:       params.Organization.RegisteredAddress2,
		RegisteredAddress3:       params.Organization.RegisteredAddress3,
		RegisteredCountry:        params.Organization.RegisteredCountry,
		Website:                  params.Organization.Website,
		OrganizationStatus:       params.Organization.OrganizationStatus,
		VATRegistrationNumber:    params.Organization.VATRegistrationNumber,
	}

	org_params := orgs.NewCreateOrganizationParams().WithOrganization(&o)

	r, e := client.OrganizationManagement.CreateOrganization(ctx, org_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while creating the Organization. Error: %v.\n", e)

		switch e.(type) {

		case *orgs.CreateOrganizationConflict:

			v := e.(*orgs.CreateOrganizationConflict)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateOrganizationConflict().WithPayload(&eValue)

		case *orgs.CreateOrganizationInternalServerError:

			v := e.(*orgs.CreateOrganizationInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateOrganizationInternalServerError().WithPayload(&eValue)

		case *orgs.CreateOrganizationUnauthorized:

			v := e.(*orgs.CreateOrganizationUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateOrganizationUnauthorized().WithPayload(&eValue)

		case *orgs.CreateOrganizationForbidden:

			v := e.(*orgs.CreateOrganizationForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateOrganizationForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewCreateOrganizationInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.ItemCreatedResponse{
		ID:   r.Payload.ID,
		Link: &r.Payload.Link,
	}

	return user_org_management.NewCreateOrganizationCreated().WithPayload(&rValue)

}

func (p *UserOrgAPI) CreateProject(ctx context.Context, params user_org_management.CreateProjectParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	o := uomodels.Project{
		AllowedOrganizations:   params.Project.AllowedOrganizations,
		LinkedOrganization:     params.Project.LinkedOrganization,
		ProjectContactPerson:   params.Project.ProjectContactPerson,
		ProjectCreatedBy:       params.Project.ProjectCreatedBy,
		ProjectCreationTime:    params.Project.ProjectCreationTime,
		ProjectDescription:     params.Project.ProjectDescription,
		ProjectID:              params.Project.ProjectID,
		ProjectName:            params.Project.ProjectName,
		ProjectShortName:       params.Project.ProjectShortName,
		ProjectStatus:          params.Project.ProjectStatus,
		ProjectStartDate:       params.Project.ProjectStartDate,
		ProjectTerminationDate: params.Project.ProjectTerminationDate,
		NormCoreHours:          params.Project.NormCoreHours,
		ProjectMaxPrice:        params.Project.ProjectMaxPrice,
		ProjectContactEmail:    params.Project.ProjectContactEmail,
		ProjectDomain:          params.Project.ProjectDomain,
	}

	project_params := projects.NewCreateProjectParams().WithProject(&o)

	r, e := client.ProjectManagement.CreateProject(ctx, project_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while creating the Project. Error: %v.\n", e)

		switch e.(type) {

		case *projects.CreateProjectConflict:

			v := e.(*projects.CreateProjectConflict)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateProjectConflict().WithPayload(&eValue)

		case *projects.CreateProjectUnprocessableEntity:

			v := e.(*projects.CreateProjectUnprocessableEntity)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateProjectUnprocessableEntity().WithPayload(&eValue)

		case *projects.CreateProjectInternalServerError:

			v := e.(*projects.CreateProjectInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateProjectInternalServerError().WithPayload(&eValue)

		case *projects.CreateProjectUnauthorized:

			v := e.(*projects.CreateProjectUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateProjectUnauthorized().WithPayload(&eValue)

		case *projects.CreateProjectForbidden:

			v := e.(*projects.CreateProjectForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateProjectForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewCreateProjectInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.ItemCreatedResponse{
		ID:   r.Payload.ID,
		Link: &r.Payload.Link,
	}

	return user_org_management.NewCreateProjectCreated().WithPayload(&rValue)

}

func (p *UserOrgAPI) CreateUser(ctx context.Context, params user_org_management.CreateUserParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	u := uomodels.User{
		AllowedOrganizations:          params.User.AllowedOrganizations,
		Permissions:                   params.User.Permissions,
		EmailAddress:                  params.User.EmailAddress,
		FirstName:                     params.User.FirstName,
		Username:                      params.User.Username,
		ID:                            params.User.ID,
		LastName:                      params.User.LastName,
		OrganizationID:                params.User.OrganizationID,
		PGPKeyID:                      params.User.PGPKeyID,
		Projects:                      params.User.Projects,
		RegistrationDateTime:          params.User.RegistrationDateTime,
		UserStatus:                    params.User.UserStatus,
		AgreedToTermsOfUse:            params.User.AgreedToTermsOfUse,
		TermsOfUseVersion:             params.User.TermsOfUseVersion,
		DateOfAgreementToTermsOfUse:   params.User.DateOfAgreementToTermsOfUse,
		AgreeToUseOfCookies:           params.User.AgreeToUseOfCookies,
		DateOfAgreementToUseOfCookies: params.User.DateOfAgreementToUseOfCookies,
	}

	user_params := users.NewCreateUserParams().WithUser(&u)

	r, e := client.UserManagement.CreateUser(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while creating the User. Error: %v.\n", e)

		switch e.(type) {

		case *users.CreateUserConflict:

			v := e.(*users.CreateUserConflict)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateUserConflict().WithPayload(&eValue)

		case *users.CreateUserInternalServerError:

			v := e.(*users.CreateUserInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateUserInternalServerError().WithPayload(&eValue)

		case *users.CreateUserUnauthorized:

			v := e.(*users.CreateUserUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateUserUnauthorized().WithPayload(&eValue)

		case *users.CreateUserForbidden:

			v := e.(*users.CreateUserForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewCreateUserForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewCreateUserInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.ItemCreatedResponse{
		ID:   r.Payload.ID,
		Link: &r.Payload.Link,
	}

	return user_org_management.NewCreateUserCreated().WithPayload(&rValue)

}

func (p *UserOrgAPI) DeleteHPCResource(ctx context.Context, params user_org_management.DeleteHPCResourceParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	hpc_params := hpcres.NewDeleteHPCResourceParams().WithID(params.ID)

	r, e := client.HpcManagement.DeleteHPCResource(ctx, hpc_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while deleting the HPCResource. Error: %v.\n", e)

		switch e.(type) {

		case *hpcres.DeleteHPCResourceNotFound:

			v := e.(*hpcres.DeleteHPCResourceNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteHPCResourceNotFound().WithPayload(&eValue)

		case *hpcres.DeleteHPCResourceInternalServerError:

			v := e.(*hpcres.DeleteHPCResourceInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteHPCResourceInternalServerError().WithPayload(&eValue)

		case *hpcres.DeleteHPCResourceUnauthorized:

			v := e.(*hpcres.DeleteHPCResourceUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteHPCResourceUnauthorized().WithPayload(&eValue)

		case *hpcres.DeleteHPCResourceForbidden:

			v := e.(*hpcres.DeleteHPCResourceForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteHPCResourceForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewDeleteHPCResourceInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.DeletedResponse{
		ID:      r.Payload.ID,
		Message: r.Payload.Message,
	}

	return user_org_management.NewDeleteHPCResourceOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) DeleteOrganization(ctx context.Context, params user_org_management.DeleteOrganizationParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	org_params := orgs.NewDeleteOrganizationParams().WithID(params.ID)

	r, e := client.OrganizationManagement.DeleteOrganization(ctx, org_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while deleting the Organization. Error: %v.\n", e)

		switch e.(type) {

		case *orgs.DeleteOrganizationNotFound:

			v := e.(*orgs.DeleteOrganizationNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteOrganizationNotFound().WithPayload(&eValue)

		case *orgs.DeleteOrganizationInternalServerError:

			v := e.(*orgs.DeleteOrganizationInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteOrganizationInternalServerError().WithPayload(&eValue)

		case *orgs.DeleteOrganizationUnauthorized:

			v := e.(*orgs.DeleteOrganizationUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteOrganizationUnauthorized().WithPayload(&eValue)

		case *orgs.DeleteOrganizationForbidden:

			v := e.(*orgs.DeleteOrganizationForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteOrganizationForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewDeleteOrganizationInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.DeletedResponse{
		ID:      r.Payload.ID,
		Message: r.Payload.Message,
	}

	return user_org_management.NewDeleteOrganizationOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) DeleteProject(ctx context.Context, params user_org_management.DeleteProjectParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	project_params := projects.NewDeleteProjectParams().WithID(params.ID)

	r, e := client.ProjectManagement.DeleteProject(ctx, project_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while deleting the Project. Error: %v.\n", e)

		switch e.(type) {

		case *projects.DeleteProjectNotFound:

			v := e.(*projects.DeleteProjectNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteProjectNotFound().WithPayload(&eValue)

		case *projects.DeleteProjectInternalServerError:

			v := e.(*projects.DeleteProjectInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteProjectInternalServerError().WithPayload(&eValue)

		case *projects.DeleteProjectUnauthorized:

			v := e.(*projects.DeleteProjectUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteProjectUnauthorized().WithPayload(&eValue)

		case *projects.DeleteProjectForbidden:

			v := e.(*projects.DeleteProjectForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteProjectForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewDeleteProjectInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.DeletedResponse{
		ID:      r.Payload.ID,
		Message: r.Payload.Message,
	}

	return user_org_management.NewDeleteProjectOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) DeleteUser(ctx context.Context, params user_org_management.DeleteUserParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	user_params := users.NewDeleteUserParams().WithID(params.ID)

	r, e := client.UserManagement.DeleteUser(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while deleting the User. Error: %v.\n", e)

		switch e.(type) {

		case *users.DeleteUserNotFound:

			v := e.(*users.DeleteUserNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserNotFound().WithPayload(&eValue)

		case *users.DeleteUserInternalServerError:

			v := e.(*users.DeleteUserInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserInternalServerError().WithPayload(&eValue)

		case *users.DeleteUserUnauthorized:

			v := e.(*users.DeleteUserUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserUnauthorized().WithPayload(&eValue)

		case *users.DeleteUserForbidden:

			v := e.(*users.DeleteUserForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewDeleteUserInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.DeletedResponse{
		ID:      r.Payload.ID,
		Message: r.Payload.Message,
	}

	return user_org_management.NewDeleteUserOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) GetHPCResource(ctx context.Context, params user_org_management.GetHPCResourceParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	hpc_params := hpcres.NewGetHPCResourceParams().WithID(params.ID)

	r, e := client.HpcManagement.GetHPCResource(ctx, hpc_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while retrieving the HPCResource. Error: %v.\n", e)

		switch e.(type) {

		case *hpcres.GetHPCResourceNotFound:

			v := e.(*hpcres.GetHPCResourceNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetHPCResourceNotFound().WithPayload(&eValue)

		case *hpcres.GetHPCResourceInternalServerError:

			v := e.(*hpcres.GetHPCResourceInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetHPCResourceInternalServerError().WithPayload(&eValue)

		case *hpcres.GetHPCResourceUnauthorized:

			v := e.(*hpcres.GetHPCResourceUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetHPCResourceUnauthorized().WithPayload(&eValue)

		case *hpcres.GetHPCResourceForbidden:

			v := e.(*hpcres.GetHPCResourceForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetHPCResourceForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewGetHPCResourceInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.HPCResource{
		AssociatedLEXISProject: r.Payload.AssociatedLEXISProject,
		AssociatedHPCProject:   r.Payload.AssociatedHPCProject,
		ApprovalStatus:         r.Payload.ApprovalStatus,
		CloudNetworkName:       r.Payload.CloudNetworkName,
		ProjectNetworkName:     r.Payload.ProjectNetworkName,
		HEAppEEndpoint:         r.Payload.HEAppEEndpoint,
		HPCProvider:            r.Payload.HPCProvider,
		HPCResourceID:          r.Payload.HPCResourceID,
		OpenStackEndpoint:      r.Payload.OpenStackEndpoint,
		OpenStackProjectID:     r.Payload.OpenStackProjectID,
		ResourceType:           r.Payload.ResourceType,
		TermsConsent:           r.Payload.TermsConsent,
	}

	return user_org_management.NewGetHPCResourceOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) GetOrganization(ctx context.Context, params user_org_management.GetOrganizationParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	org_params := orgs.NewGetOrganizationParams().WithID(params.ID)

	r, e := client.OrganizationManagement.GetOrganization(ctx, org_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while retrieving the Organization. Error: %v.\n", e)

		switch e.(type) {

		case *orgs.GetOrganizationNotFound:

			v := e.(*orgs.GetOrganizationNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetOrganizationNotFound().WithPayload(&eValue)

		case *orgs.GetOrganizationInternalServerError:

			v := e.(*orgs.GetOrganizationInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetOrganizationInternalServerError().WithPayload(&eValue)

		case *orgs.GetOrganizationUnauthorized:

			v := e.(*orgs.GetOrganizationUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetOrganizationUnauthorized().WithPayload(&eValue)

		case *orgs.GetOrganizationForbidden:

			v := e.(*orgs.GetOrganizationForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetOrganizationForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewGetOrganizationInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.Organization{
		CreatedBy:                r.Payload.CreatedBy,
		CreationDate:             r.Payload.CreationDate,
		FormalName:               r.Payload.FormalName,
		ID:                       r.Payload.ID,
		OrganizationEmailAddress: r.Payload.OrganizationEmailAddress,
		PrimaryTelephoneNumber:   r.Payload.PrimaryTelephoneNumber,
		RegisteredAddress1:       r.Payload.RegisteredAddress1,
		RegisteredAddress2:       r.Payload.RegisteredAddress2,
		RegisteredAddress3:       r.Payload.RegisteredAddress3,
		RegisteredCountry:        r.Payload.RegisteredCountry,
		Website:                  r.Payload.Website,
		OrganizationStatus:       r.Payload.OrganizationStatus,
		VATRegistrationNumber:    r.Payload.VATRegistrationNumber,
	}

	return user_org_management.NewGetOrganizationOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) GetProject(ctx context.Context, params user_org_management.GetProjectParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	org_params := projects.NewGetProjectParams().WithID(params.ID)

	r, e := client.ProjectManagement.GetProject(ctx, org_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while retrieving the Project. Error: %v.\n", e)

		switch e.(type) {

		case *projects.GetProjectNotFound:

			v := e.(*projects.GetProjectNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetProjectNotFound().WithPayload(&eValue)

		case *projects.GetProjectInternalServerError:

			v := e.(*projects.GetProjectInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetProjectInternalServerError().WithPayload(&eValue)

		case *projects.GetProjectUnauthorized:

			v := e.(*projects.GetProjectUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetProjectUnauthorized().WithPayload(&eValue)

		case *projects.GetProjectForbidden:

			v := e.(*projects.GetProjectForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetProjectForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewGetProjectInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.Project{
		AllowedOrganizations:   r.Payload.AllowedOrganizations,
		LinkedOrganization:     r.Payload.LinkedOrganization,
		ProjectContactPerson:   r.Payload.ProjectContactPerson,
		ProjectCreatedBy:       r.Payload.ProjectCreatedBy,
		ProjectCreationTime:    r.Payload.ProjectCreationTime,
		ProjectDescription:     r.Payload.ProjectDescription,
		ProjectID:              r.Payload.ProjectID,
		ProjectName:            r.Payload.ProjectName,
		ProjectShortName:       r.Payload.ProjectShortName,
		ProjectStatus:          r.Payload.ProjectStatus,
		ProjectStartDate:       r.Payload.ProjectStartDate,
		ProjectTerminationDate: r.Payload.ProjectTerminationDate,
		NormCoreHours:          r.Payload.NormCoreHours,
		ProjectMaxPrice:        r.Payload.ProjectMaxPrice,
		ProjectContactEmail:    r.Payload.ProjectContactEmail,
		ProjectDomain:          r.Payload.ProjectDomain,
	}

	return user_org_management.NewGetProjectOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) GetUser(ctx context.Context, params user_org_management.GetUserParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	user_params := users.NewGetUserParams().WithID(params.ID)

	r, e := client.UserManagement.GetUser(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while retrieving the User. Error: %v.\n", e)

		switch e.(type) {

		case *users.GetUserNotFound:

			v := e.(*users.GetUserNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetUserNotFound().WithPayload(&eValue)

		case *users.GetUserInternalServerError:

			v := e.(*users.GetUserInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetUserInternalServerError().WithPayload(&eValue)

		case *users.GetUserUnauthorized:

			v := e.(*users.GetUserUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetUserUnauthorized().WithPayload(&eValue)

		case *users.GetUserForbidden:

			v := e.(*users.GetUserForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewGetUserForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewGetUserInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.User{
		AllowedOrganizations:          r.Payload.AllowedOrganizations,
		Permissions:                   r.Payload.Permissions,
		EmailAddress:                  r.Payload.EmailAddress,
		FirstName:                     r.Payload.FirstName,
		Username:                      r.Payload.Username,
		ID:                            r.Payload.ID,
		LastName:                      r.Payload.LastName,
		OrganizationID:                r.Payload.OrganizationID,
		PGPKeyID:                      r.Payload.PGPKeyID,
		Projects:                      r.Payload.Projects,
		RegistrationDateTime:          r.Payload.RegistrationDateTime,
		UserStatus:                    r.Payload.UserStatus,
		AgreedToTermsOfUse:            r.Payload.AgreedToTermsOfUse,
		TermsOfUseVersion:             r.Payload.TermsOfUseVersion,
		DateOfAgreementToTermsOfUse:   r.Payload.DateOfAgreementToTermsOfUse,
		AgreeToUseOfCookies:           r.Payload.AgreeToUseOfCookies,
		DateOfAgreementToUseOfCookies: r.Payload.DateOfAgreementToUseOfCookies,
	}

	return user_org_management.NewGetUserOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) ListHPCResources(ctx context.Context, params user_org_management.ListHPCResourcesParams) middleware.Responder {

	var rValue []*models.HPCResource

	client := p.getClient(params.HTTPRequest)

	hpc_params := hpcres.NewListHPCResourcesParams()

	if params.Scope != nil {

		hpc_params = hpc_params.WithScope(params.Scope)

	}

	r, e := client.HpcManagement.ListHPCResources(ctx, hpc_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while retrieving the list of HPC Resources. Error: %v.\n", e)

		switch e.(type) {

		case *hpcres.ListHPCResourcesUnauthorized:

			v := e.(*hpcres.ListHPCResourcesUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewListHPCResourcesUnauthorized().WithPayload(&eValue)

		case *hpcres.ListHPCResourcesForbidden:

			v := e.(*hpcres.ListHPCResourcesForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewListHPCResourcesForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewListHPCResourcesInternalServerError().WithPayload(&eValue)

		}

	}

	for _, o := range r.Payload {

		newHpc := models.HPCResource{
			AssociatedLEXISProject: o.AssociatedLEXISProject,
			AssociatedHPCProject:   o.AssociatedHPCProject,
			ApprovalStatus:         o.ApprovalStatus,
			CloudNetworkName:       o.CloudNetworkName,
			ProjectNetworkName:     o.ProjectNetworkName,
			HEAppEEndpoint:         o.HEAppEEndpoint,
			HPCProvider:            o.HPCProvider,
			HPCResourceID:          o.HPCResourceID,
			OpenStackEndpoint:      o.OpenStackEndpoint,
			OpenStackProjectID:     o.OpenStackProjectID,
			ResourceType:           o.ResourceType,
			TermsConsent:           o.TermsConsent,
		}

		rValue = append(rValue, &newHpc)

	}

	return user_org_management.NewListHPCResourcesOK().WithPayload(rValue)

}

func (p *UserOrgAPI) ListOrganizations(ctx context.Context, params user_org_management.ListOrganizationsParams) middleware.Responder {

	var rValue []*models.Organization

	client := p.getClient(params.HTTPRequest)

	org_params := orgs.NewListOrganizationsParams()

	if params.Scope != nil {

		org_params = org_params.WithScope(params.Scope)

	}

	r, e := client.OrganizationManagement.ListOrganizations(ctx, org_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while retrieving the list of Organizations. Error: %v.\n", e)

		switch e.(type) {

		case *orgs.ListOrganizationsUnauthorized:

			v := e.(*orgs.ListOrganizationsUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewListOrganizationsUnauthorized().WithPayload(&eValue)

		case *orgs.ListOrganizationsForbidden:

			v := e.(*orgs.ListOrganizationsForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewListOrganizationsForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewListOrganizationsInternalServerError().WithPayload(&eValue)

		}

	}

	for _, o := range r.Payload {

		newOrg := models.Organization{
			CreatedBy:                o.CreatedBy,
			CreationDate:             o.CreationDate,
			FormalName:               o.FormalName,
			ID:                       o.ID,
			OrganizationEmailAddress: o.OrganizationEmailAddress,
			PrimaryTelephoneNumber:   o.PrimaryTelephoneNumber,
			RegisteredAddress1:       o.RegisteredAddress1,
			RegisteredAddress2:       o.RegisteredAddress2,
			RegisteredAddress3:       o.RegisteredAddress3,
			RegisteredCountry:        o.RegisteredCountry,
			Website:                  o.Website,
			OrganizationStatus:       o.OrganizationStatus,
			VATRegistrationNumber:    o.VATRegistrationNumber,
		}

		rValue = append(rValue, &newOrg)

	}

	return user_org_management.NewListOrganizationsOK().WithPayload(rValue)

}

func (p *UserOrgAPI) ListProjects(ctx context.Context, params user_org_management.ListProjectsParams) middleware.Responder {

	var rValue []*models.Project

	client := p.getClient(params.HTTPRequest)

	project_params := projects.NewListProjectsParams()

	if params.Scope != nil {

		project_params = project_params.WithScope(params.Scope)

	}

	r, e := client.ProjectManagement.ListProjects(ctx, project_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while retrieving the list of Projects. Error: %v.\n", e)

		switch e.(type) {

		case *projects.ListProjectsUnauthorized:

			v := e.(*projects.ListProjectsUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewListProjectsUnauthorized().WithPayload(&eValue)

		case *projects.ListProjectsForbidden:

			v := e.(*projects.ListProjectsForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewListProjectsForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewListProjectsInternalServerError().WithPayload(&eValue)

		}

	}

	for _, o := range r.Payload {

		newProject := models.Project{
			AllowedOrganizations:   o.AllowedOrganizations,
			LinkedOrganization:     o.LinkedOrganization,
			ProjectContactPerson:   o.ProjectContactPerson,
			ProjectCreatedBy:       o.ProjectCreatedBy,
			ProjectCreationTime:    o.ProjectCreationTime,
			ProjectDescription:     o.ProjectDescription,
			ProjectID:              o.ProjectID,
			ProjectName:            o.ProjectName,
			ProjectShortName:       o.ProjectShortName,
			ProjectStatus:          o.ProjectStatus,
			ProjectStartDate:       o.ProjectStartDate,
			ProjectTerminationDate: o.ProjectTerminationDate,
			NormCoreHours:          o.NormCoreHours,
			ProjectMaxPrice:        o.ProjectMaxPrice,
			ProjectContactEmail:    o.ProjectContactEmail,
			ProjectDomain:          o.ProjectDomain,
		}

		rValue = append(rValue, &newProject)

	}

	return user_org_management.NewListProjectsOK().WithPayload(rValue)

}

func (p *UserOrgAPI) ListUsers(ctx context.Context, params user_org_management.ListUsersParams) middleware.Responder {

	var rValue []*models.User

	user_params := users.NewListUsersParams()

	client := p.getClient(params.HTTPRequest)

	if params.Email != nil {

		user_params = user_params.WithEmail(params.Email)

	}

	if params.Project != nil {

		user_params = user_params.WithProject(params.Project)

	}

	if params.Permissions != nil {

		user_params = user_params.WithPermissions(params.Permissions)

	}

	if params.Scope != nil {

		user_params = user_params.WithScope(params.Scope)

	}

	r, e := client.UserManagement.ListUsers(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while retrieving the list of Users. Error: %v.\n", e)

		switch e.(type) {

		case *users.ListUsersUnauthorized:

			v := e.(*users.ListUsersUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewListUsersUnauthorized().WithPayload(&eValue)

		case *users.ListUsersForbidden:

			v := e.(*users.ListUsersForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewListUsersForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewListUsersInternalServerError().WithPayload(&eValue)

		}

	}

	for _, u := range r.Payload {

		newUser := models.User{
			AllowedOrganizations:          u.AllowedOrganizations,
			Permissions:                   u.Permissions,
			EmailAddress:                  u.EmailAddress,
			FirstName:                     u.FirstName,
			Username:                      u.Username,
			ID:                            u.ID,
			LastName:                      u.LastName,
			OrganizationID:                u.OrganizationID,
			PGPKeyID:                      u.PGPKeyID,
			RegistrationDateTime:          u.RegistrationDateTime,
			UserStatus:                    u.UserStatus,
			Projects:                      u.Projects,
			AgreedToTermsOfUse:            u.AgreedToTermsOfUse,
			TermsOfUseVersion:             u.TermsOfUseVersion,
			DateOfAgreementToTermsOfUse:   u.DateOfAgreementToTermsOfUse,
			AgreeToUseOfCookies:           u.AgreeToUseOfCookies,
			DateOfAgreementToUseOfCookies: u.DateOfAgreementToUseOfCookies,
		}

		rValue = append(rValue, &newUser)

	}

	return user_org_management.NewListUsersOK().WithPayload(rValue)

}

func (p *UserOrgAPI) UpdateHPCResource(ctx context.Context, params user_org_management.UpdateHPCResourceParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	o := uomodels.HPCResource{
		AssociatedLEXISProject: params.HPCResource.AssociatedLEXISProject,
		AssociatedHPCProject:   params.HPCResource.AssociatedHPCProject,
		ApprovalStatus:         params.HPCResource.ApprovalStatus,
		CloudNetworkName:       params.HPCResource.CloudNetworkName,
		ProjectNetworkName:     params.HPCResource.ProjectNetworkName,
		HEAppEEndpoint:         params.HPCResource.HEAppEEndpoint,
		HPCProvider:            params.HPCResource.HPCProvider,
		HPCResourceID:          params.HPCResource.HPCResourceID,
		OpenStackEndpoint:      params.HPCResource.OpenStackEndpoint,
		OpenStackProjectID:     params.HPCResource.OpenStackProjectID,
		ResourceType:           params.HPCResource.ResourceType,
		TermsConsent:           params.HPCResource.TermsConsent,
	}

	hpc_params := hpcres.NewUpdateHPCResourceParams().WithHPCResource(&o).WithID(params.ID)

	r, e := client.HpcManagement.UpdateHPCResource(ctx, hpc_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while updating the HPCResource. Error: %v.\n", e)

		switch e.(type) {

		case *hpcres.UpdateHPCResourceNotFound:

			v := e.(*hpcres.UpdateHPCResourceNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateHPCResourceNotFound().WithPayload(&eValue)

		case *hpcres.UpdateHPCResourceInternalServerError:

			v := e.(*hpcres.UpdateHPCResourceInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateHPCResourceInternalServerError().WithPayload(&eValue)

		case *hpcres.UpdateHPCResourceUnauthorized:

			v := e.(*hpcres.UpdateHPCResourceUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateHPCResourceUnauthorized().WithPayload(&eValue)

		case *hpcres.UpdateHPCResourceForbidden:

			v := e.(*hpcres.UpdateHPCResourceForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateHPCResourceForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewUpdateHPCResourceInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.HPCResource{
		AssociatedLEXISProject: r.Payload.AssociatedLEXISProject,
		AssociatedHPCProject:   r.Payload.AssociatedHPCProject,
		ApprovalStatus:         r.Payload.ApprovalStatus,
		CloudNetworkName:       r.Payload.CloudNetworkName,
		ProjectNetworkName:     r.Payload.ProjectNetworkName,
		HEAppEEndpoint:         r.Payload.HEAppEEndpoint,
		HPCProvider:            r.Payload.HPCProvider,
		HPCResourceID:          r.Payload.HPCResourceID,
		OpenStackEndpoint:      r.Payload.OpenStackEndpoint,
		OpenStackProjectID:     r.Payload.OpenStackProjectID,
		ResourceType:           r.Payload.ResourceType,
		TermsConsent:           r.Payload.TermsConsent,
	}

	return user_org_management.NewUpdateHPCResourceOK().WithPayload(&rValue)
}
func (p *UserOrgAPI) UpdateOrganization(ctx context.Context, params user_org_management.UpdateOrganizationParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	o := uomodels.Organization{
		CreatedBy:                params.Organization.CreatedBy,
		CreationDate:             params.Organization.CreationDate,
		FormalName:               params.Organization.FormalName,
		ID:                       params.Organization.ID,
		OrganizationEmailAddress: params.Organization.OrganizationEmailAddress,
		PrimaryTelephoneNumber:   params.Organization.PrimaryTelephoneNumber,
		RegisteredAddress1:       params.Organization.RegisteredAddress1,
		RegisteredAddress2:       params.Organization.RegisteredAddress2,
		RegisteredAddress3:       params.Organization.RegisteredAddress3,
		RegisteredCountry:        params.Organization.RegisteredCountry,
		Website:                  params.Organization.Website,
		OrganizationStatus:       params.Organization.OrganizationStatus,
		VATRegistrationNumber:    params.Organization.VATRegistrationNumber,
	}

	org_params := orgs.NewUpdateOrganizationParams().WithOrganization(&o).WithID(params.ID)

	r, e := client.OrganizationManagement.UpdateOrganization(ctx, org_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while updating the Organization. Error: %v.\n", e)

		switch e.(type) {

		case *orgs.UpdateOrganizationNotFound:

			v := e.(*orgs.UpdateOrganizationNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateOrganizationNotFound().WithPayload(&eValue)

		case *orgs.UpdateOrganizationInternalServerError:

			v := e.(*orgs.UpdateOrganizationInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateOrganizationInternalServerError().WithPayload(&eValue)

		case *orgs.UpdateOrganizationUnauthorized:

			v := e.(*orgs.UpdateOrganizationUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateOrganizationUnauthorized().WithPayload(&eValue)

		case *orgs.UpdateOrganizationForbidden:

			v := e.(*orgs.UpdateOrganizationForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateOrganizationForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewUpdateOrganizationInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.Organization{
		CreatedBy:                r.Payload.CreatedBy,
		CreationDate:             r.Payload.CreationDate,
		FormalName:               r.Payload.FormalName,
		ID:                       r.Payload.ID,
		OrganizationEmailAddress: r.Payload.OrganizationEmailAddress,
		PrimaryTelephoneNumber:   r.Payload.PrimaryTelephoneNumber,
		RegisteredAddress1:       r.Payload.RegisteredAddress1,
		RegisteredAddress2:       r.Payload.RegisteredAddress2,
		RegisteredAddress3:       r.Payload.RegisteredAddress3,
		RegisteredCountry:        r.Payload.RegisteredCountry,
		Website:                  r.Payload.Website,
		OrganizationStatus:       r.Payload.OrganizationStatus,
		VATRegistrationNumber:    r.Payload.VATRegistrationNumber,
	}

	return user_org_management.NewUpdateOrganizationOK().WithPayload(&rValue)
}

func (p *UserOrgAPI) UpdateProject(ctx context.Context, params user_org_management.UpdateProjectParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	o := uomodels.Project{
		AllowedOrganizations:   params.Project.AllowedOrganizations,
		LinkedOrganization:     params.Project.LinkedOrganization,
		ProjectContactPerson:   params.Project.ProjectContactPerson,
		ProjectCreatedBy:       params.Project.ProjectCreatedBy,
		ProjectCreationTime:    params.Project.ProjectCreationTime,
		ProjectDescription:     params.Project.ProjectDescription,
		ProjectID:              params.Project.ProjectID,
		ProjectName:            params.Project.ProjectName,
		ProjectShortName:       params.Project.ProjectShortName,
		ProjectStatus:          params.Project.ProjectStatus,
		ProjectStartDate:       params.Project.ProjectStartDate,
		ProjectTerminationDate: params.Project.ProjectTerminationDate,
		NormCoreHours:          params.Project.NormCoreHours,
		ProjectMaxPrice:        params.Project.ProjectMaxPrice,
		ProjectContactEmail:    params.Project.ProjectContactEmail,
		ProjectDomain:          params.Project.ProjectDomain,
	}

	project_params := projects.NewUpdateProjectParams().WithProject(&o).WithID(params.ID)

	r, e := client.ProjectManagement.UpdateProject(ctx, project_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while updating the Project. Error: %v.\n", e)

		switch e.(type) {

		case *projects.UpdateProjectNotFound:

			v := e.(*projects.UpdateProjectNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateProjectNotFound().WithPayload(&eValue)

		case *projects.UpdateProjectUnprocessableEntity:

			v := e.(*projects.UpdateProjectUnprocessableEntity)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateProjectUnprocessableEntity().WithPayload(&eValue)

		case *projects.UpdateProjectInternalServerError:

			v := e.(*projects.UpdateProjectInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateProjectInternalServerError().WithPayload(&eValue)

		case *projects.UpdateProjectUnauthorized:

			v := e.(*projects.UpdateProjectUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateProjectUnauthorized().WithPayload(&eValue)

		case *projects.UpdateProjectForbidden:

			v := e.(*projects.UpdateProjectForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateProjectForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewUpdateProjectInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.Project{
		AllowedOrganizations:   r.Payload.AllowedOrganizations,
		LinkedOrganization:     r.Payload.LinkedOrganization,
		ProjectContactPerson:   r.Payload.ProjectContactPerson,
		ProjectCreatedBy:       r.Payload.ProjectCreatedBy,
		ProjectCreationTime:    r.Payload.ProjectCreationTime,
		ProjectDescription:     r.Payload.ProjectDescription,
		ProjectID:              r.Payload.ProjectID,
		ProjectName:            r.Payload.ProjectName,
		ProjectShortName:       r.Payload.ProjectShortName,
		ProjectStatus:          r.Payload.ProjectStatus,
		ProjectStartDate:       r.Payload.ProjectStartDate,
		ProjectTerminationDate: r.Payload.ProjectTerminationDate,
		NormCoreHours:          r.Payload.NormCoreHours,
		ProjectMaxPrice:        r.Payload.ProjectMaxPrice,
		ProjectContactEmail:    r.Payload.ProjectContactEmail,
		ProjectDomain:          r.Payload.ProjectDomain,
	}

	return user_org_management.NewUpdateProjectOK().WithPayload(&rValue)
}

func (p *UserOrgAPI) UpdateUser(ctx context.Context, params user_org_management.UpdateUserParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	u := uomodels.User{
		AllowedOrganizations:          params.User.AllowedOrganizations,
		Permissions:                   params.User.Permissions,
		EmailAddress:                  params.User.EmailAddress,
		FirstName:                     params.User.FirstName,
		Username:                      params.User.Username,
		ID:                            params.User.ID,
		LastName:                      params.User.LastName,
		OrganizationID:                params.User.OrganizationID,
		PGPKeyID:                      params.User.PGPKeyID,
		Projects:                      params.User.Projects,
		RegistrationDateTime:          params.User.RegistrationDateTime,
		UserStatus:                    params.User.UserStatus,
		AgreedToTermsOfUse:            params.User.AgreedToTermsOfUse,
		TermsOfUseVersion:             params.User.TermsOfUseVersion,
		DateOfAgreementToTermsOfUse:   params.User.DateOfAgreementToTermsOfUse,
		AgreeToUseOfCookies:           params.User.AgreeToUseOfCookies,
		DateOfAgreementToUseOfCookies: params.User.DateOfAgreementToUseOfCookies,
	}

	user_params := users.NewUpdateUserParams().WithUser(&u).WithID(params.ID)

	r, e := client.UserManagement.UpdateUser(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while updating the User. Error: %v.\n", e)

		switch e.(type) {

		case *users.UpdateUserNotFound:

			v := e.(*users.UpdateUserNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateUserNotFound().WithPayload(&eValue)

		case *users.UpdateUserInternalServerError:

			v := e.(*users.UpdateUserInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateUserInternalServerError().WithPayload(&eValue)

		case *users.UpdateUserUnauthorized:

			v := e.(*users.UpdateUserUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateUserUnauthorized().WithPayload(&eValue)

		case *users.UpdateUserForbidden:

			v := e.(*users.UpdateUserForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewUpdateUserForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewUpdateUserInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.User{
		AllowedOrganizations:          r.Payload.AllowedOrganizations,
		Permissions:                   r.Payload.Permissions,
		EmailAddress:                  r.Payload.EmailAddress,
		FirstName:                     r.Payload.FirstName,
		Username:                      r.Payload.Username,
		ID:                            r.Payload.ID,
		LastName:                      r.Payload.LastName,
		OrganizationID:                r.Payload.OrganizationID,
		PGPKeyID:                      r.Payload.PGPKeyID,
		RegistrationDateTime:          r.Payload.RegistrationDateTime,
		UserStatus:                    r.Payload.UserStatus,
		AgreedToTermsOfUse:            r.Payload.AgreedToTermsOfUse,
		TermsOfUseVersion:             r.Payload.TermsOfUseVersion,
		DateOfAgreementToTermsOfUse:   r.Payload.DateOfAgreementToTermsOfUse,
		AgreeToUseOfCookies:           r.Payload.AgreeToUseOfCookies,
		DateOfAgreementToUseOfCookies: r.Payload.DateOfAgreementToUseOfCookies,
	}

	return user_org_management.NewUpdateUserOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) AddUserToOrganization(ctx context.Context, params user_org_management.AddUserToOrganizationParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	user_params := orgs.NewAddUserToOrganizationParams().WithID(params.ID).WithUserID(params.UserID)

	r, e := client.OrganizationManagement.AddUserToOrganization(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while updating the User. Error: %v.\n", e)

		switch e.(type) {

		case *orgs.AddUserToOrganizationNotFound:

			v := e.(*orgs.AddUserToOrganizationNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddUserToOrganizationNotFound().WithPayload(&eValue)

		case *orgs.AddUserToOrganizationInternalServerError:

			v := e.(*orgs.AddUserToOrganizationInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddUserToOrganizationInternalServerError().WithPayload(&eValue)

		case *orgs.AddUserToOrganizationUnauthorized:

			v := e.(*orgs.AddUserToOrganizationUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddUserToOrganizationUnauthorized().WithPayload(&eValue)

		case *orgs.AddUserToOrganizationForbidden:

			v := e.(*orgs.AddUserToOrganizationForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddUserToOrganizationForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewAddUserToOrganizationInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.OKResponse{
		Message: r.Payload.Message,
	}

	return user_org_management.NewAddUserToOrganizationOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) AddUserToProject(ctx context.Context, params user_org_management.AddUserToProjectParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	user_params := projects.NewAddUserToProjectParams().WithID(params.ID).WithUserID(params.UserID)

	r, e := client.ProjectManagement.AddUserToProject(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while updating the User. Error: %v.\n", e)

		switch e.(type) {

		case *projects.AddUserToProjectNotFound:

			v := e.(*projects.AddUserToProjectNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddUserToProjectNotFound().WithPayload(&eValue)

		case *projects.AddUserToProjectInternalServerError:

			v := e.(*projects.AddUserToProjectInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddUserToProjectInternalServerError().WithPayload(&eValue)

		case *projects.AddUserToProjectUnauthorized:

			v := e.(*projects.AddUserToProjectUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddUserToProjectUnauthorized().WithPayload(&eValue)

		case *projects.AddUserToProjectForbidden:

			v := e.(*projects.AddUserToProjectForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewAddUserToProjectForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewAddUserToProjectInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.OKResponse{
		Message: r.Payload.Message,
	}

	return user_org_management.NewAddUserToProjectOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) DeleteUserFromOrganization(ctx context.Context, params user_org_management.DeleteUserFromOrganizationParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	user_params := orgs.NewDeleteUserFromOrganizationParams().WithID(params.ID).WithUserID(params.UserID)

	r, e := client.OrganizationManagement.DeleteUserFromOrganization(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while updating the User. Error: %v.\n", e)

		switch e.(type) {

		case *orgs.DeleteUserFromOrganizationNotFound:

			v := e.(*orgs.DeleteUserFromOrganizationNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserFromOrganizationNotFound().WithPayload(&eValue)

		case *orgs.DeleteUserFromOrganizationInternalServerError:

			v := e.(*orgs.DeleteUserFromOrganizationInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserFromOrganizationInternalServerError().WithPayload(&eValue)

		case *orgs.DeleteUserFromOrganizationUnauthorized:

			v := e.(*orgs.DeleteUserFromOrganizationUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserFromOrganizationUnauthorized().WithPayload(&eValue)

		case *orgs.DeleteUserFromOrganizationForbidden:

			v := e.(*orgs.DeleteUserFromOrganizationForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserFromOrganizationForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewDeleteUserFromOrganizationInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.DeletedResponse{
		Message: r.Payload.Message,
	}

	return user_org_management.NewDeleteUserFromOrganizationOK().WithPayload(&rValue)

}

func (p *UserOrgAPI) DeleteUserFromProject(ctx context.Context, params user_org_management.DeleteUserFromProjectParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	user_params := projects.NewDeleteUserFromProjectParams().WithID(params.ID).WithUserID(params.UserID)

	r, e := client.ProjectManagement.DeleteUserFromProject(ctx, user_params)

	if e != nil {

		l.Warning.Printf("[UO] There was an error while updating the User. Error: %v.\n", e)

		switch e.(type) {

		case *projects.DeleteUserFromProjectNotFound:

			v := e.(*projects.DeleteUserFromProjectNotFound)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserFromProjectNotFound().WithPayload(&eValue)

		case *projects.DeleteUserFromProjectInternalServerError:

			v := e.(*projects.DeleteUserFromProjectInternalServerError)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserFromProjectInternalServerError().WithPayload(&eValue)

		case *projects.DeleteUserFromProjectUnauthorized:

			v := e.(*projects.DeleteUserFromProjectUnauthorized)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserFromProjectUnauthorized().WithPayload(&eValue)

		case *projects.DeleteUserFromProjectForbidden:

			v := e.(*projects.DeleteUserFromProjectForbidden)

			eValue := models.ErrorResponse{
				ErrorString: &v.Payload.Message,
			}

			return user_org_management.NewDeleteUserFromProjectForbidden().WithPayload(&eValue)

		default:

			v := e.Error()

			eValue := models.ErrorResponse{
				ErrorString: &v,
			}

			return user_org_management.NewDeleteUserFromProjectInternalServerError().WithPayload(&eValue)

		}

	}

	rValue := models.DeletedResponse{
		Message: r.Payload.Message,
	}

	return user_org_management.NewDeleteUserFromProjectOK().WithPayload(&rValue)

}
