package datasetapi

import (
	"context"

	stypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/staging"
//	dtypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
//	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/staging"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)


func (p *DataSetAPI) CheckDuplicationStatus(ctx context.Context, params staging.CheckDuplicationStatusParams) middleware.Responder {

	rparams := stypes.CheckDuplicationStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CheckDuplicationStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckDuplicationStatus endpoint\n")
		switch err.(type) {
		case *stypes.CheckDuplicationStatusBadRequest: // 400
			v := err.(*stypes.CheckDuplicationStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDuplicationStatusBadRequest().WithPayload(&payload)
		case *stypes.CheckDuplicationStatusUnauthorized: // 401
			v := err.(*stypes.CheckDuplicationStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDuplicationStatusUnauthorized().WithPayload(&payload)
		case *stypes.CheckDuplicationStatusNotFound: // 404
			v := err.(*stypes.CheckDuplicationStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDuplicationStatusNotFound().WithPayload(&payload)
		case *stypes.CheckDuplicationStatusRequestURITooLong: // 414
			v := err.(*stypes.CheckDuplicationStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDuplicationStatusRequestURITooLong().WithPayload(&payload)
		case *stypes.CheckDuplicationStatusInternalServerError: // 500
			v := err.(*stypes.CheckDuplicationStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDuplicationStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCheckDuplicationStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := staging.CheckDuplicationStatusOKBody{
		Status: res.Payload.Status,
		TargetPath: res.Payload.TargetPath,
	}
	return staging.NewCheckDuplicationStatusOK().WithPayload(&payload)

}

func (p *DataSetAPI) Duplicate(ctx context.Context, params staging.DuplicateParams) middleware.Responder {
	rparams := stypes.DuplicateParams{
		Parameters: stypes.DuplicateBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
			TargetSystem: params.Parameters.TargetSystem,
			TargetPath:   params.Parameters.TargetPath,
			Title:        params.Parameters.Title,
		},
	}

	res, err := p.getClient(params.HTTPRequest).Staging.Duplicate(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Duplicate endpoint\n")
		switch err.(type) {
		case *stypes.DuplicateBadRequest: // 400
			v := err.(*stypes.DuplicateBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDuplicateBadRequest().WithPayload(&payload)
		case *stypes.DuplicateUnauthorized: // 401
			v := err.(*stypes.DuplicateUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDuplicateUnauthorized().WithPayload(&payload)
		case *stypes.DuplicateNotFound: // 404
			v := err.(*stypes.DuplicateNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDuplicateNotFound().WithPayload(&payload)
		case *stypes.DuplicateRequestURITooLong: // 414
			v := err.(*stypes.DuplicateRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDuplicateRequestURITooLong().WithPayload(&payload)
		case *stypes.DuplicateInternalServerError: // 500
			v := err.(*stypes.DuplicateInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDuplicateInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewDuplicateInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return staging.NewDuplicateCreated().WithPayload(&payload)
}
