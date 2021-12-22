package datasetapi

import (
	"context"

	stypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) CheckPIDStatus(ctx context.Context, params data_set_management.CheckPIDStatusParams) middleware.Responder {
	rparams := stypes.CheckPIDStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckPIDStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckPIDStatus endpoint\n")
		switch err.(type) {
		case *stypes.CheckPIDStatusBadRequest: // 400
			v := err.(*stypes.CheckPIDStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPIDStatusBadRequest().WithPayload(&payload)
		case *stypes.CheckPIDStatusUnauthorized: // 401
			v := err.(*stypes.CheckPIDStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPIDStatusUnauthorized().WithPayload(&payload)
		case *stypes.CheckPIDStatusNotFound: // 404
			v := err.(*stypes.CheckPIDStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPIDStatusNotFound().WithPayload(&payload)
		case *stypes.CheckPIDStatusRequestURITooLong: // 414
			v := err.(*stypes.CheckPIDStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPIDStatusRequestURITooLong().WithPayload(&payload)
		case *stypes.CheckPIDStatusInternalServerError: // 500
			v := err.(*stypes.CheckPIDStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPIDStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckPIDStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := models.DataReplication{
		Status: res.Payload.Status,
		PID:    res.Payload.PID,
		TargetPath: res.Payload.TargetPath,
	}
	return data_set_management.NewCheckPIDStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CheckReplicateStatus(ctx context.Context, params data_set_management.CheckReplicateStatusParams) middleware.Responder {
	rparams := stypes.CheckReplicateStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckReplicateStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckReplicateStatus endpoint\n")
		switch err.(type) {
		case *stypes.CheckReplicateStatusBadRequest: // 400
			v := err.(*stypes.CheckReplicateStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckReplicateStatusBadRequest().WithPayload(&payload)
		case *stypes.CheckReplicateStatusUnauthorized: // 401
			v := err.(*stypes.CheckReplicateStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckReplicateStatusUnauthorized().WithPayload(&payload)
		case *stypes.CheckReplicateStatusNotFound: // 404
			v := err.(*stypes.CheckReplicateStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckReplicateStatusNotFound().WithPayload(&payload)
		case *stypes.CheckReplicateStatusRequestURITooLong: // 414
			v := err.(*stypes.CheckReplicateStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckReplicateStatusRequestURITooLong().WithPayload(&payload)
		case *stypes.CheckReplicateStatusInternalServerError: // 500
			v := err.(*stypes.CheckReplicateStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckReplicateStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckReplicateStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := models.DataReplication{
		Status: res.Payload.Status,
		PID:    res.Payload.PID,
	}
	return data_set_management.NewCheckReplicateStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) PID(ctx context.Context, params data_set_management.PIDParams) middleware.Responder {
	rparams := stypes.PIDParams{
		Parameters: stypes.PIDBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
			ParentPid:    params.Parameters.ParentPid,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.PID(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling PID endpoint\n")
		switch err.(type) {
		case *stypes.PIDBadRequest: // 400
			v := err.(*stypes.PIDBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPIDBadRequest().WithPayload(&payload)
		case *stypes.PIDUnauthorized: // 401
			v := err.(*stypes.PIDUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPIDUnauthorized().WithPayload(&payload)
		case *stypes.PIDNotFound: // 404
			v := err.(*stypes.PIDNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPIDNotFound().WithPayload(&payload)
		case *stypes.PIDRequestURITooLong: // 414
			v := err.(*stypes.PIDRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPIDRequestURITooLong().WithPayload(&payload)
		case *stypes.PIDInternalServerError: // 500
			v := err.(*stypes.PIDInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPIDInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewPIDInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewPIDCreated().WithPayload(&payload)
}

func (p *DataSetAPI) Replicate(ctx context.Context, params data_set_management.ReplicateParams) middleware.Responder {
	rparams := stypes.ReplicateParams{
		Parameters: stypes.ReplicateBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
			TargetSystem: params.Parameters.TargetSystem,
			TargetPath:   params.Parameters.TargetPath,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.Replicate(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Replicate endpoint\n")
		switch err.(type) {
		case *stypes.ReplicateBadRequest: // 400
			v := err.(*stypes.ReplicateBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewReplicateBadRequest().WithPayload(&payload)
		case *stypes.ReplicateUnauthorized: // 401
			v := err.(*stypes.ReplicateUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewReplicateUnauthorized().WithPayload(&payload)
		case *stypes.ReplicateNotFound: // 404
			v := err.(*stypes.ReplicateNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewReplicateNotFound().WithPayload(&payload)
		case *stypes.ReplicateRequestURITooLong: // 414
			v := err.(*stypes.ReplicateRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewReplicateRequestURITooLong().WithPayload(&payload)
		case *stypes.ReplicateInternalServerError: // 500
			v := err.(*stypes.ReplicateInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewReplicateInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewReplicateInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewReplicateCreated().WithPayload(&payload)
}
