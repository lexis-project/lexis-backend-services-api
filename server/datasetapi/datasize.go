package datasetapi

import (
	"context"

	dtypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
//	remotedatamodel "github.com/lexis-project/lexis-backend-services-interface-datasets.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) CheckSizeStatus(ctx context.Context, params data_set_management.CheckSizeStatusParams) middleware.Responder {
	rparams := dtypes.CheckSizeStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckSizeStatus(ctx, &rparams)
	if err != nil {
		l.Info.Printf("Error calling CheckSizeStatus endpoint\n")
		switch err.(type) {
		case *dtypes.CheckSizeStatusBadRequest: // 400
			v := err.(*dtypes.CheckSizeStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckSizeStatusBadRequest().WithPayload(&payload)
		case *dtypes.CheckSizeStatusUnauthorized: // 401
			v := err.(*dtypes.CheckSizeStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckSizeStatusUnauthorized().WithPayload(&payload)
		case *dtypes.CheckSizeStatusForbidden: // 403
			v := err.(*dtypes.CheckSizeStatusForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckSizeStatusForbidden().WithPayload(&payload)
		case *dtypes.CheckSizeStatusNotFound: // 404
			v := err.(*dtypes.CheckSizeStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckSizeStatusNotFound().WithPayload(&payload)
		case *dtypes.CheckSizeStatusRequestURITooLong: // 414
			v := err.(*dtypes.CheckSizeStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckSizeStatusRequestURITooLong().WithPayload(&payload)
		case *dtypes.CheckSizeStatusTooManyRequests: // 429
			v := err.(*dtypes.CheckSizeStatusTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckSizeStatusTooManyRequests().WithPayload(&payload)
		case *dtypes.CheckSizeStatusInternalServerError: // 500
			v := err.(*dtypes.CheckSizeStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckSizeStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckSizeStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := models.DataSize{
		Result: res.Payload.Result,
		Size: res.Payload.Size,
		Totalfiles: res.Payload.Totalfiles,
		Smallfiles: res.Payload.Smallfiles,
	}
	return data_set_management.NewCheckSizeStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) Size(ctx context.Context, params data_set_management.SizeParams) middleware.Responder {
	rparams := dtypes.SizeParams{
		Parameters: dtypes.SizeBody{
			TargetSystem: params.Parameters.TargetSystem,
			TargetPath:   params.Parameters.TargetPath,
		},
	}
	res, err := p.getClient(params.HTTPRequest).DataSetManagement.Size(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Size endpoint\n")
		switch err.(type) {
		case *dtypes.SizeBadRequest: // 400
			v := err.(*dtypes.SizeBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewSizeBadRequest().WithPayload(&payload)
		case *dtypes.SizeUnauthorized: // 401
			v := err.(*dtypes.SizeUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewSizeUnauthorized().WithPayload(&payload)
		case *dtypes.SizeForbidden: // 403
			v := err.(*dtypes.SizeForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewSizeForbidden().WithPayload(&payload)
		case *dtypes.SizeNotFound: // 404
			v := err.(*dtypes.SizeNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewSizeNotFound().WithPayload(&payload)
		case *dtypes.SizeRequestURITooLong: // 414
			v := err.(*dtypes.SizeRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewSizeRequestURITooLong().WithPayload(&payload)
		case *dtypes.SizeTooManyRequests: // 429
			v := err.(*dtypes.SizeTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewSizeTooManyRequests().WithPayload(&payload)
		case *dtypes.SizeInternalServerError: // 500
			v := err.(*dtypes.SizeInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewSizeInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewSizeInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewSizeCreated().WithPayload(&payload)
}
