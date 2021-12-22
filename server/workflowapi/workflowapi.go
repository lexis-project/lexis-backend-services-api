package workflowapi

import (
	"context"
	"net/http"
	"strings"

	remoteworkflowapi "code.it4i.cz/lexis/wp4/alien4cloud-interface/client"
	workflowExecutionManagement "code.it4i.cz/lexis/wp4/alien4cloud-interface/client/workflow_execution_management"
	workflowManagement "code.it4i.cz/lexis/wp4/alien4cloud-interface/client/workflow_management"
	workflowTemplateManagement "code.it4i.cz/lexis/wp4/alien4cloud-interface/client/workflow_template_management"
	wfmodels "code.it4i.cz/lexis/wp4/alien4cloud-interface/models"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/workflow_management"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
)

// A simple in memory CRUD on data
// In real life, this should be using a persistent storage.
type WorkflowAPI struct {
	//db dbManager.DbParameter
	c remoteworkflowapi.Config
}

// New returns a new Pet manager
func New(c remoteworkflowapi.Config) *WorkflowAPI {

	return &WorkflowAPI{
		c: c,
	}

}

func (p *WorkflowAPI) getClient(param *http.Request) *remoteworkflowapi.Alien4CloudInterface {

	config := p.c

	if len(param.Header.Get("Authorization")) > 0 {

		token := strings.Fields(param.Header.Get("Authorization"))[1]

		config.AuthInfo = httptransport.BearerToken(token)

	}

	r := remoteworkflowapi.New(config)

	return r

}

// Workflow Template Management

func (p *WorkflowAPI) GetWorkflowTemplates(ctx context.Context, params workflow_management.GetWorkflowTemplatesParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wt_params := workflowTemplateManagement.NewGetWorkflowTemplatesParams()
	r, e := client.WorkflowTemplateManagement.GetWorkflowTemplates(ctx, wt_params)

	if e != nil {
		switch e.(type) {
		case *workflowTemplateManagement.GetWorkflowTemplatesUnauthorized:
			v := e.(*workflowTemplateManagement.GetWorkflowTemplatesUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowTemplatesUnauthorized().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewGetWorkflowTemplatesInternalServerError().WithPayload(&eValue)
		}
	}

	workflowTemplates := []*models.WorkflowTemplate{}
	for i := range r.Payload {
		workflowTemplates = append(workflowTemplates, &models.WorkflowTemplate{
			WorkflowTemplateName: r.Payload[i].WorkflowTemplateName,
			WorkflowTemplateID:   r.Payload[i].WorkflowTemplateID,
			Description:          r.Payload[i].Description,
		})
	}

	return workflow_management.NewGetWorkflowTemplatesOK().WithPayload(workflowTemplates)
}

func (p *WorkflowAPI) GetWorkflowTemplate(ctx context.Context, params workflow_management.GetWorkflowTemplateParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wt_params := workflowTemplateManagement.NewGetWorkflowTemplateParams().WithWorkflowTemplateID(params.WorkflowTemplateID)
	r, e := client.WorkflowTemplateManagement.GetWorkflowTemplate(ctx, wt_params)

	if e != nil {

		switch e.(type) {
		case *workflowTemplateManagement.GetWorkflowTemplateUnauthorized:
			v := e.(*workflowTemplateManagement.GetWorkflowTemplateUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowTemplateUnauthorized().WithPayload(&eValue)
		case *workflowTemplateManagement.GetWorkflowTemplateNotFound:
			v := e.(*workflowTemplateManagement.GetWorkflowTemplateNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowTemplateNotFound().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewGetWorkflowTemplateInternalServerError().WithPayload(&eValue)
		}
	}

	tempFiles := []*models.InputFile{}
	for i := range r.Payload.InputFiles {
		tempFiles = append(tempFiles, &models.InputFile{
			InputFileName: r.Payload.InputFiles[i].InputFileName,
			InputFileType: r.Payload.InputFiles[i].InputFileType,
			Description:   r.Payload.InputFiles[i].Description,
		})
	}

	tempParams := []*models.InputParameter{}
	for i := range r.Payload.InputParameters {
		tempParams = append(tempParams, &models.InputParameter{
			InputParamName:         r.Payload.InputParameters[i].InputParamName,
			InputParamType:         r.Payload.InputParameters[i].InputParamType,
			InputParamRequired:     r.Payload.InputParameters[i].InputParamRequired,
			InputParamDefaultValue: r.Payload.InputParameters[i].InputParamDefaultValue,
			Description:            r.Payload.InputParameters[i].Description,
			Task:                   r.Payload.InputParameters[i].Task,
			IsDataset:              r.Payload.InputParameters[i].IsDataset,
			IsDatasetID:            r.Payload.InputParameters[i].IsDatasetID,
			IsDatasetPath:          r.Payload.InputParameters[i].IsDatasetPath,
			DisplayName:            r.Payload.InputParameters[i].DisplayName,
		})
	}

	tempNodes := []*models.NodeTemplate{}
	for i := range r.Payload.NodeTemplates {
		tempTags := []*models.Tag{}
		for j := range r.Payload.NodeTemplates[i].Tags {
			tempTags = append(tempTags, &models.Tag{
				Key:   r.Payload.NodeTemplates[i].Tags[j].Key,
				Value: r.Payload.NodeTemplates[i].Tags[j].Value,
			})
		}
		tempNodes = append(tempNodes, &models.NodeTemplate{
			NodeName: r.Payload.NodeTemplates[i].NodeName,
			NodeType: r.Payload.NodeTemplates[i].NodeType,
			Tags:     tempTags,
		})
	}

	workflowTemplate := models.WorkflowTemplate{
		WorkflowTemplateName: r.Payload.WorkflowTemplateName,
		Description:          r.Payload.Description,
		WorkflowTemplateID:   r.Payload.WorkflowTemplateID,
		InputFiles:           tempFiles,
		InputParameters:      tempParams,
		NodeTemplates:        tempNodes,
	}

	return workflow_management.NewGetWorkflowTemplateOK().WithPayload(&workflowTemplate)
}

func (p *WorkflowAPI) UploadWorkflowTemplate(ctx context.Context, params workflow_management.UploadWorkflowTemplateParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	//Create a runtime.NamedReadCloser
	namedFile := runtime.NamedReader("WorkflowTemplateFile", params.WorkflowTemplateFile)

	//Set Workflow Template Upload Params
	wtu_params := workflowTemplateManagement.NewUploadWorkflowTemplateParams().WithWorkflowTemplateFile(namedFile)

	r, e := client.WorkflowTemplateManagement.UploadWorkflowTemplate(ctx, wtu_params)

	if e != nil {
		switch e.(type) {
		case *workflowTemplateManagement.UploadWorkflowTemplateUnauthorized:
			v := e.(*workflowTemplateManagement.UploadWorkflowTemplateUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewUploadWorkflowTemplateUnauthorized().WithPayload(&eValue)
		case *workflowTemplateManagement.UploadWorkflowTemplateBadRequest:
			v := e.(*workflowTemplateManagement.UploadWorkflowTemplateBadRequest)
			eValue := models.InvalidResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewUploadWorkflowTemplateBadRequest().WithPayload(&eValue)
		case *workflowTemplateManagement.UploadWorkflowTemplateConflict:
			v := e.(*workflowTemplateManagement.UploadWorkflowTemplateConflict)
			eValue := models.ConflictResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewUploadWorkflowTemplateConflict().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewUploadWorkflowTemplateInternalServerError().WithPayload(&eValue)
		}
	}

	rValue := models.ItemCreatedResponse{
		ID:   r.Payload.ID,
		Link: &r.Payload.Link,
	}

	return workflow_management.NewUploadWorkflowTemplateCreated().WithPayload(&rValue)
}

// Workflow Management

func (p *WorkflowAPI) GetWorkflows(ctx context.Context, params workflow_management.GetWorkflowsParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wf_params := workflowManagement.NewGetWorkflowsParams()
	r, e := client.WorkflowManagement.GetWorkflows(ctx, wf_params)

	if e != nil {
		switch e.(type) {
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewGetWorkflowsInternalServerError().WithPayload(&eValue)
		}
	}

	workflows := []*models.Workflow{}
	for i := range r.Payload {
		workflows = append(workflows, &models.Workflow{
			WorkflowID:         r.Payload[i].WorkflowID,
			WorkflowName:       r.Payload[i].WorkflowName,
			ProjectID:          r.Payload[i].ProjectID,
			ProjectName:        r.Payload[i].ProjectName,
			WorkflowTemplateID: r.Payload[i].WorkflowTemplateID,
			Description:        r.Payload[i].Description,
			CreationTime:       r.Payload[i].CreationTime,
			CreatedBy:          r.Payload[i].CreatedBy,
		})
	}

	return workflow_management.NewGetWorkflowsOK().WithPayload(workflows)

}

func (p *WorkflowAPI) CreateWorkflow(ctx context.Context, params workflow_management.CreateWorkflowParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	// TODO: The below functionality needs tested on full system, to check if it works as expected by uploading the token if required
	//tempParams := make(map[string]interface{})
	//for k, v := range params.WorkflowRequest.InputParameters {
	//	tempParams[k] = v
	//}

	//// Need to check if the token is required in input parameters and add if so
	//wt_params := workflowTemplateManagement.NewGetWorkflowTemplateParams().WithWorkflowTemplateID(*params.WorkflowRequest.WorkflowTemplateID)
	//workflowTemplate, e := client.WorkflowTemplateManagement.GetWorkflowTemplate(ctx, wt_params)

	//if e != nil {

	//	v := e.Error()
	//	eValue := models.ErrorResponse{
	//		ErrorString: &v,
	//	}

	//	return workflow_management.NewGetWorkflowTemplateInternalServerError().WithPayload(&eValue)

	//}

	//// Need to check if this is the correct token
	//for i := range workflowTemplate.Payload.InputParameters {
	//	if workflowTemplate.Payload.InputParameters[i].InputParamName == "token" {
	//		//tempParams["token"] = p.c.AuthInfo
	//		tempParams["token"] = "123456"
	//	}
	//}

	//tempFiles := []*wfmodels.InputFile{}
	//for i := range params.WorkflowRequest.InputFiles {
	//	tempFiles = append(tempFiles, &wfmodels.InputFile{
	//		InputFileName: params.WorkflowRequest.InputFiles[i].InputFileName,
	//		Path:          params.WorkflowRequest.InputFiles[i].Path,
	//	})
	//}

	//TODO: Ensure params are returned correctly, why are they not htrowing same error?

	w := wfmodels.WorkflowRequest{
		WorkflowName:       params.WorkflowRequest.WorkflowName,
		WorkflowTemplateID: params.WorkflowRequest.WorkflowTemplateID,
		ProjectID:          params.WorkflowRequest.ProjectID,
		Description:        params.WorkflowRequest.Description,
		//InputFiles:           tempFiles,
		//InputParameters:      tempParams,
	}

	wf_params := workflowManagement.NewCreateWorkflowParams().WithWorkflowRequest(&w)

	r, e := client.WorkflowManagement.CreateWorkflow(ctx, wf_params)

	if e != nil {
		switch e.(type) {
		case *workflowManagement.CreateWorkflowConflict:
			v := e.(*workflowManagement.CreateWorkflowConflict)
			eValue := models.ConflictResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowConflict().WithPayload(&eValue)

		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewCreateWorkflowInternalServerError().WithPayload(&eValue)
		}
	}
	rValue := models.ItemCreatedResponse{
		ID:   r.Payload.ID,
		Link: &r.Payload.Link,
	}

	return workflow_management.NewCreateWorkflowCreated().WithPayload(&rValue)
}

func (p *WorkflowAPI) GetWorkflow(ctx context.Context, params workflow_management.GetWorkflowParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wf_params := workflowManagement.NewGetWorkflowParams().WithWorkflowID(params.WorkflowID)
	r, e := client.WorkflowManagement.GetWorkflow(ctx, wf_params)

	if e != nil {
		switch e.(type) {
		case *workflowManagement.GetWorkflowUnauthorized:
			v := e.(*workflowManagement.GetWorkflowUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowUnauthorized().WithPayload(&eValue)
		case *workflowManagement.GetWorkflowNotFound:
			v := e.(*workflowManagement.GetWorkflowNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowNotFound().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewGetWorkflowInternalServerError().WithPayload(&eValue)
		}
	}

	tempFiles := []*models.InputFile{}
	for i := range r.Payload.InputFiles {
		tempFiles = append(tempFiles, &models.InputFile{
			InputFileName: r.Payload.InputFiles[i].InputFileName,
			InputFileType: r.Payload.InputFiles[i].InputFileType,
			Description:   r.Payload.InputFiles[i].Description,
		})
	}

	//Remove the token from the returned inputs as this should be hidden from the user
	tempInputs := []*models.InputParameter{}
	for i := range r.Payload.InputParameters {
		if r.Payload.InputParameters[i].InputParamName != "token" && r.Payload.InputParameters[i].InputParamName != "project_id" && r.Payload.InputParameters[i].InputParamName != "project_short_name" {
			tempInputs = append(tempInputs, &models.InputParameter{
				InputParamName:         r.Payload.InputParameters[i].InputParamName,
				InputParamType:         r.Payload.InputParameters[i].InputParamType,
				InputParamRequired:     r.Payload.InputParameters[i].InputParamRequired,
				InputParamDefaultValue: r.Payload.InputParameters[i].InputParamDefaultValue,
				Description:            r.Payload.InputParameters[i].Description,
				Task:                   r.Payload.InputParameters[i].Task,
				IsDataset:              r.Payload.InputParameters[i].IsDataset,
				IsDatasetID:            r.Payload.InputParameters[i].IsDatasetID,
				IsDatasetPath:          r.Payload.InputParameters[i].IsDatasetPath,
				DisplayName:            r.Payload.InputParameters[i].DisplayName,
			})
		}
	}

	tempNodes := []*models.NodeTemplate{}
	for i := range r.Payload.NodeTemplates {
		tempTags := []*models.Tag{}
		for j := range r.Payload.NodeTemplates[i].Tags {
			tempTags = append(tempTags, &models.Tag{
				Key:   r.Payload.NodeTemplates[i].Tags[j].Key,
				Value: r.Payload.NodeTemplates[i].Tags[j].Value,
			})
		}
		tempNodes = append(tempNodes, &models.NodeTemplate{
			NodeName: r.Payload.NodeTemplates[i].NodeName,
			NodeType: r.Payload.NodeTemplates[i].NodeType,
			Tags:     tempTags,
		})
	}

	workflow := models.WorkflowDetail{
		WorkflowID:         r.Payload.WorkflowID,
		WorkflowName:       r.Payload.WorkflowName,
		ProjectID:          r.Payload.ProjectID,
		ProjectName:        r.Payload.ProjectName,
		ProjectShortName:   r.Payload.ProjectShortName,
		WorkflowTemplateID: r.Payload.WorkflowTemplateID,
		Description:        r.Payload.Description,
		CreatedBy:          r.Payload.CreatedBy,
		CreationTime:       r.Payload.CreationTime,
		InputFiles:         tempFiles,
		NodeTemplates:      tempNodes,
		InputParameters:    tempInputs,
	}

	return workflow_management.NewGetWorkflowOK().WithPayload(&workflow)
}

func (p *WorkflowAPI) DeleteWorkflow(ctx context.Context, params workflow_management.DeleteWorkflowParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wf_params := workflowManagement.NewDeleteWorkflowParams().WithWorkflowID(params.WorkflowID)
	_, e := client.WorkflowManagement.DeleteWorkflow(ctx, wf_params)

	if e != nil {
		switch e.(type) {
		case *workflowManagement.DeleteWorkflowUnauthorized:
			v := e.(*workflowManagement.DeleteWorkflowUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewDeleteWorkflowUnauthorized().WithPayload(&eValue)
		case *workflowManagement.DeleteWorkflowNotFound:
			v := e.(*workflowManagement.DeleteWorkflowNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewDeleteWorkflowNotFound().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewDeleteWorkflowInternalServerError().WithPayload(&eValue)
		}
	}

	return workflow_management.NewDeleteWorkflowOK()
}

// Workflow Execution Management

func (p *WorkflowAPI) ListWorkflowExecutions(ctx context.Context, params workflow_management.ListWorkflowExecutionsParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wfe_params := workflowExecutionManagement.NewListWorkflowExecutionsParams().WithWorkflowID(params.WorkflowID)
	r, e := client.WorkflowExecutionManagement.ListWorkflowExecutions(ctx, wfe_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.ListWorkflowExecutionsUnauthorized:
			v := e.(*workflowExecutionManagement.ListWorkflowExecutionsUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewListWorkflowExecutionsUnauthorized().WithPayload(&eValue)
		case *workflowExecutionManagement.ListWorkflowExecutionsNotFound:
			v := e.(*workflowExecutionManagement.ListWorkflowExecutionsNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewListWorkflowExecutionsNotFound().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewListWorkflowExecutionsInternalServerError().WithPayload(&eValue)
		}
	}

	workflowExecutions := []*models.WorkflowExecution{}
	for i := range r.Payload {
		workflowExecutions = append(workflowExecutions, &models.WorkflowExecution{
			WorkflowExecutionID:     r.Payload[i].WorkflowExecutionID,
			WorkflowExecutionName:   r.Payload[i].WorkflowExecutionName,
			WorkflowExecutionStatus: r.Payload[i].WorkflowExecutionStatus,
			WorkflowID:              r.Payload[i].WorkflowID,
			WorkflowName:            r.Payload[i].WorkflowName,
			CreationTime:            r.Payload[i].CreationTime,
		})
	}

	return workflow_management.NewListWorkflowExecutionsOK().WithPayload(workflowExecutions)
}

func (p *WorkflowAPI) CreateWorkflowExecution(ctx context.Context, params workflow_management.CreateWorkflowExecutionParams) middleware.Responder {
	client := p.getClient(params.HTTPRequest)

	wf_params := workflowManagement.NewGetWorkflowParams().WithWorkflowID(params.WorkflowID)
	workflow, e := client.WorkflowManagement.GetWorkflow(ctx, wf_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.CreateWorkflowExecutionUnauthorized:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionUnauthorized().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionBadRequest:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionBadRequest)
			eValue := models.InvalidResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionBadRequest().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionNotFound:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionNotFound().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewCreateWorkflowExecutionInternalServerError().WithPayload(&eValue)
		}
	}

	wfe_params := workflowExecutionManagement.NewCreateWorkflowExecutionParams().WithWorkflowID(params.WorkflowID)
	if params.WorkflowExecutionRequest != nil {
		tempParams := make(map[string]interface{})
		for k, v := range params.WorkflowExecutionRequest.InputParameters {
			tempParams[k] = v
		}
		for i := range workflow.Payload.InputParameters {
			if workflow.Payload.InputParameters[i].InputParamName == "token" {
				/* Need to check if this is the correct token */
				tempParams["token"] = strings.Fields(params.HTTPRequest.Header.Get("Authorization"))[1]
			} else if workflow.Payload.InputParameters[i].InputParamName == "project_id" {
				/* Add the project ID as input parameter if required */
				tempParams["project_id"] = workflow.Payload.ProjectID
			} else if workflow.Payload.InputParameters[i].InputParamName == "project_short_name" {
				/* Add the project short name as input parameter if required */
				tempParams["project_short_name"] = workflow.Payload.ProjectShortName
			}
		}

		tempFiles := []*wfmodels.InputFile{}
		for i := range params.WorkflowExecutionRequest.InputFiles {
			tempFiles = append(tempFiles, &wfmodels.InputFile{
				InputFileName: params.WorkflowExecutionRequest.InputFiles[i].InputFileName,
				Path:          params.WorkflowExecutionRequest.InputFiles[i].Path,
			})
		}

		w := wfmodels.CreateWorkflowExecutionRequest{
			WorkflowTemplateID: params.WorkflowExecutionRequest.WorkflowTemplateID,
			IsCronJob:          params.WorkflowExecutionRequest.IsCronJob,
			IsScheduledJob:     params.WorkflowExecutionRequest.IsScheduledJob,
			IsBatchJob:         params.WorkflowExecutionRequest.IsBatchJob,
			InputFiles:         tempFiles,
			InputParameters:    tempParams,
		}
		wfe_params = wfe_params.WithWorkflowExecutionRequest(&w)
	}

	r, e := client.WorkflowExecutionManagement.CreateWorkflowExecution(ctx, wfe_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.CreateWorkflowExecutionUnauthorized:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionUnauthorized().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionNotFound:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionNotFound().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionBadRequest:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionBadRequest)
			eValue := models.InvalidResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionBadRequest().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewCreateWorkflowExecutionInternalServerError().WithPayload(&eValue)
		}
	}

	rValue := models.ItemCreatedResponse{
		ID:   r.Payload.ID,
		Link: &r.Payload.Link,
	}

	return workflow_management.NewCreateWorkflowExecutionCreated().WithPayload(&rValue)
}

func (p *WorkflowAPI) CreateWorkflowExecutions(ctx context.Context, params workflow_management.CreateWorkflowExecutionsParams) middleware.Responder {
	client := p.getClient(params.HTTPRequest)

	wf_params := workflowManagement.NewGetWorkflowParams().WithWorkflowID(params.WorkflowID)
	workflow, e := client.WorkflowManagement.GetWorkflow(ctx, wf_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.CreateWorkflowExecutionsUnauthorized:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionsUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionsUnauthorized().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionsBadRequest:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionsBadRequest)
			eValue := models.InvalidResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionsBadRequest().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionsNotFound:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionsNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionsNotFound().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewCreateWorkflowExecutionsInternalServerError().WithPayload(&eValue)
		}
	}

	wfes_params := workflowExecutionManagement.NewCreateWorkflowExecutionsParams().WithWorkflowID(params.WorkflowID)
	if params.WorkflowExecutionRequests != nil {
		additionalParams := make(map[string]interface{})
		for i := range workflow.Payload.InputParameters {
			if workflow.Payload.InputParameters[i].InputParamName == "token" {
				/* Need to check if this is the correct token */
				additionalParams["token"] = strings.Fields(params.HTTPRequest.Header.Get("Authorization"))[1]
			} else if workflow.Payload.InputParameters[i].InputParamName == "project_id" {
				/* Add the project ID as input parameter if required */
				additionalParams["project_id"] = workflow.Payload.ProjectID
			} else if workflow.Payload.InputParameters[i].InputParamName == "project_short_name" {
				/* Add the project short name as input parameter if required */
				additionalParams["project_short_name"] = workflow.Payload.ProjectShortName
			}
		}

		for _, wfereq := range params.WorkflowExecutionRequests {
			tempParams := make(map[string]interface{})
			for k, v := range wfereq.InputParameters {
				tempParams[k] = v
			}

			for k, v := range additionalParams {
				tempParams[k] = v
			}

			tempFiles := []*wfmodels.InputFile{}
			for i := range wfereq.InputFiles {
				tempFiles = append(tempFiles, &wfmodels.InputFile{
					InputFileName: wfereq.InputFiles[i].InputFileName,
					Path:          wfereq.InputFiles[i].Path,
				})
			}

			w := wfmodels.CreateWorkflowExecutionRequest{
				WorkflowTemplateID: wfereq.WorkflowTemplateID,
				IsCronJob:          wfereq.IsCronJob,
				IsScheduledJob:     wfereq.IsScheduledJob,
				IsBatchJob:         wfereq.IsBatchJob,
				InputFiles:         tempFiles,
				InputParameters:    tempParams,
			}
			wfes_params.WorkflowExecutionRequests = append(wfes_params.WorkflowExecutionRequests, &w)
		}
	}

	r, e := client.WorkflowExecutionManagement.CreateWorkflowExecutions(ctx, wfes_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.CreateWorkflowExecutionsUnauthorized:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionsUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionsUnauthorized().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionsNotFound:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionsNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionsNotFound().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionsBadRequest:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionsBadRequest)
			eValue := models.InvalidResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionsBadRequest().WithPayload(&eValue)
		case *workflowExecutionManagement.CreateWorkflowExecutionsConflict:
			v := e.(*workflowExecutionManagement.CreateWorkflowExecutionsConflict)
			eValue := models.ConflictResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCreateWorkflowExecutionsConflict().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewCreateWorkflowExecutionsInternalServerError().WithPayload(&eValue)
		}
	}

	rValue := []*models.ItemCreatedResponse{}
	for i := range r.Payload {
		rValue = append(rValue, &models.ItemCreatedResponse{
			ID:   r.Payload[i].ID,
			Link: &r.Payload[i].Link,
		})
	}

	return workflow_management.NewCreateWorkflowExecutionsCreated().WithPayload(rValue)
}

func (p *WorkflowAPI) CancelWorkflowExecution(ctx context.Context, params workflow_management.CancelWorkflowExecutionParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wfe_params := workflowExecutionManagement.NewCancelWorkflowExecutionParams().WithWorkflowID(params.WorkflowID)
	wfe_params = wfe_params.WithWorkflowExecutionID(params.WorkflowExecutionID)
	_, e := client.WorkflowExecutionManagement.CancelWorkflowExecution(ctx, wfe_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.CancelWorkflowExecutionNotFound:
			v := e.(*workflowExecutionManagement.CancelWorkflowExecutionNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewCancelWorkflowExecutionNotFound().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewCancelWorkflowExecutionInternalServerError().WithPayload(&eValue)
		}
	}

	return workflow_management.NewCancelWorkflowExecutionOK()
}

func (p *WorkflowAPI) DeleteWorkflowExecution(ctx context.Context, params workflow_management.DeleteWorkflowExecutionParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wfe_params := workflowExecutionManagement.NewDeleteWorkflowExecutionParams().WithWorkflowID(params.WorkflowID)
	wfe_params = wfe_params.WithWorkflowExecutionID(params.WorkflowExecutionID)
	_, e := client.WorkflowExecutionManagement.DeleteWorkflowExecution(ctx, wfe_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.DeleteWorkflowExecutionNotFound:
			v := e.(*workflowExecutionManagement.DeleteWorkflowExecutionNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewDeleteWorkflowExecutionNotFound().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewDeleteWorkflowExecutionInternalServerError().WithPayload(&eValue)
		}
	}

	return workflow_management.NewDeleteWorkflowExecutionOK()
}

func (p *WorkflowAPI) GetWorkflowExecutionLogs(ctx context.Context, params workflow_management.GetWorkflowExecutionLogsParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wfe_params := workflowExecutionManagement.NewGetWorkflowExecutionLogsParams().WithWorkflowID(params.WorkflowID)
	wfe_params = wfe_params.WithWorkflowExecutionID(params.WorkflowExecutionID)
	r, e := client.WorkflowExecutionManagement.GetWorkflowExecutionLogs(ctx, wfe_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.GetWorkflowExecutionLogsNotFound:
			v := e.(*workflowExecutionManagement.GetWorkflowExecutionLogsNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowExecutionLogsNotFound().WithPayload(&eValue)
		case *workflowExecutionManagement.GetWorkflowExecutionLogsUnauthorized:
			v := e.(*workflowExecutionManagement.GetWorkflowExecutionLogsUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowExecutionLogsUnauthorized().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewGetWorkflowExecutionLogsInternalServerError().WithPayload(&eValue)
		}
	}

	logs := []*models.Logs{}
	for i := range r.Payload {
		logs = append(logs, &models.Logs{
			DeploymentID:           r.Payload[i].DeploymentID,
			DeploymentPaaSID:       r.Payload[i].DeploymentPaaSID,
			Level:                  r.Payload[i].Level,
			Timestamp:              r.Payload[i].Timestamp,
			WorkflowExecutionStage: r.Payload[i].WorkflowExecutionStage,
			ExecutionID:            r.Payload[i].ExecutionID,
			NodeID:                 r.Payload[i].NodeID,
			InstanceID:             r.Payload[i].InstanceID,
			InterfaceName:          r.Payload[i].InterfaceName,
			OperationName:          r.Payload[i].OperationName,
			Content:                r.Payload[i].Content,
		})
	}

	return workflow_management.NewGetWorkflowExecutionLogsOK().WithPayload(logs)
}

func (p *WorkflowAPI) GetWorkflowExecutionStepStatus(ctx context.Context, params workflow_management.GetWorkflowExecutionStepStatusParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wfe_params := workflowExecutionManagement.NewGetWorkflowExecutionStepStatusParams().WithWorkflowID(params.WorkflowID)
	wfe_params = wfe_params.WithWorkflowExecutionID(params.WorkflowExecutionID)
	r, e := client.WorkflowExecutionManagement.GetWorkflowExecutionStepStatus(ctx, wfe_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.GetWorkflowExecutionStepStatusNotFound:
			v := e.(*workflowExecutionManagement.GetWorkflowExecutionStepStatusNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowExecutionStepStatusNotFound().WithPayload(&eValue)
		case *workflowExecutionManagement.GetWorkflowExecutionStepStatusUnauthorized:
			v := e.(*workflowExecutionManagement.GetWorkflowExecutionStepStatusUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowExecutionStepStatusUnauthorized().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewGetWorkflowExecutionStepStatusInternalServerError().WithPayload(&eValue)
		}
	}

	workflowExecutionStepStatus := []*models.WorkflowExecutionStepStatus{}
	for i := range r.Payload {
		workflowExecutionStepStatus = append(workflowExecutionStepStatus, &models.WorkflowExecutionStepStatus{
			Step:            r.Payload[i].Step,
			PrecedingSteps:  r.Payload[i].PrecedingSteps,
			SucceedingSteps: r.Payload[i].SucceedingSteps,
			Status:          r.Payload[i].Status,
			NodeName:        r.Payload[i].NodeName,
			NodeType:        r.Payload[i].NodeType,
			Task:            r.Payload[i].Task,
			Location:        r.Payload[i].Location,
			ActivityType:    r.Payload[i].ActivityType,
		})
	}

	return workflow_management.NewGetWorkflowExecutionStepStatusOK().WithPayload(workflowExecutionStepStatus)
}

func (p *WorkflowAPI) GetWorkflowExecutionDetail(ctx context.Context, params workflow_management.GetWorkflowExecutionDetailParams) middleware.Responder {

	client := p.getClient(params.HTTPRequest)

	wfe_params := workflowExecutionManagement.NewGetWorkflowExecutionDetailParams().WithWorkflowID(params.WorkflowID)
	wfe_params = wfe_params.WithWorkflowExecutionID(params.WorkflowExecutionID)
	r, e := client.WorkflowExecutionManagement.GetWorkflowExecutionDetail(ctx, wfe_params)

	if e != nil {
		switch e.(type) {
		case *workflowExecutionManagement.GetWorkflowExecutionDetailNotFound:
			v := e.(*workflowExecutionManagement.GetWorkflowExecutionDetailNotFound)
			eValue := models.MissingResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowExecutionDetailNotFound().WithPayload(&eValue)
		case *workflowExecutionManagement.GetWorkflowExecutionDetailUnauthorized:
			v := e.(*workflowExecutionManagement.GetWorkflowExecutionDetailUnauthorized)
			eValue := models.AuthorizationResponse{
				Message: v.Payload.Message,
			}
			return workflow_management.NewGetWorkflowExecutionDetailUnauthorized().WithPayload(&eValue)
		default:
			v := e.Error()
			eValue := models.ErrorResponse{
				ErrorString: &v,
			}
			return workflow_management.NewGetWorkflowExecutionDetailInternalServerError().WithPayload(&eValue)
		}
	}

	tempFiles := []*models.InputFile{}
	for i := range r.Payload.UploadedInputFiles {
		tempFiles = append(tempFiles, &models.InputFile{
			InputFileName: r.Payload.UploadedInputFiles[i].InputFileName,
			InputFileType: r.Payload.UploadedInputFiles[i].InputFileType,
			Description:   r.Payload.UploadedInputFiles[i].Description,
		})
	}

	//Remove the token from the returned inputs as this should be hidden from the user
	tempInputs := []*models.InputParameter{}
	for i := range r.Payload.InputParameters {
		if r.Payload.InputParameters[i].InputParamName != "token" {
			tempInputs = append(tempInputs, &models.InputParameter{
				InputParamName:  r.Payload.InputParameters[i].InputParamName,
				InputParamValue: r.Payload.InputParameters[i].InputParamValue,
			})
		}
	}

	tempNodes := []*models.NodeTemplate{}
	for i := range r.Payload.NodeTemplates {
		tempTags := []*models.Tag{}
		for j := range r.Payload.NodeTemplates[i].Tags {
			tempTags = append(tempTags, &models.Tag{
				Key:   r.Payload.NodeTemplates[i].Tags[j].Key,
				Value: r.Payload.NodeTemplates[i].Tags[j].Value,
			})
		}
		tempNodes = append(tempNodes, &models.NodeTemplate{
			NodeName: r.Payload.NodeTemplates[i].NodeName,
			NodeType: r.Payload.NodeTemplates[i].NodeType,
			Tags:     tempTags,
		})
	}

	outputProps := []*models.OutputProperty{}
	for i := range r.Payload.OutputProperties {
		outputProps = append(outputProps, &models.OutputProperty{
			NodeName:       r.Payload.OutputProperties[i].NodeName,
			AttributeName:  r.Payload.OutputProperties[i].AttributeName,
			AttributeValue: r.Payload.OutputProperties[i].AttributeValue,
		})
	}
	workflowExecution := models.WorkflowExecutionDetail{
		WorkflowID:                   r.Payload.WorkflowID,
		WorkflowExecutionStatus:      r.Payload.WorkflowExecutionStatus,
		WorkflowExecutionID:          r.Payload.WorkflowExecutionID,
		WorkflowExecutionName:        r.Payload.WorkflowExecutionName,
		WorkflowExecutionStage:       r.Payload.WorkflowExecutionStage,
		WorkflowExecutionStageStatus: r.Payload.WorkflowExecutionStageStatus,
		CreationTime:                 r.Payload.CreationTime,
		CreatedBy:                    r.Payload.CreatedBy,
		OutputProperties:             outputProps,
		InputParameters:              tempInputs,
		UploadedInputFiles:           tempFiles,
		NodeTemplates:                tempNodes,
	}

	return workflow_management.NewGetWorkflowExecutionDetailOK().WithPayload(&workflowExecution)
}
