package datasetapi

import (
	"context"

	stypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/staging"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/staging"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) CompressToZip(ctx context.Context, params staging.CompressToZipParams) middleware.Responder {
	rparams := stypes.CompressToZipParams{
		Parameters: stypes.CompressToZipBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
			Size:         params.Parameters.Size,
		},
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CompressToZip(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CompressToZip endpoint\n")
		switch err.(type) {
		case *stypes.CompressToZipBadRequest: // 400
			v := err.(*stypes.CompressToZipBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCompressToZipBadRequest().WithPayload(&payload)
		case *stypes.CompressToZipUnauthorized: // 401
			v := err.(*stypes.CompressToZipUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCompressToZipUnauthorized().WithPayload(&payload)
		case *stypes.CompressToZipForbidden: // 403
			v := err.(*stypes.CompressToZipForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCompressToZipForbidden().WithPayload(&payload)
		case *stypes.CompressToZipBadGateway: // 502
			v := err.(*stypes.CompressToZipBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCompressToZipBadGateway().WithPayload(&payload)
		case *stypes.CompressToZipServiceUnavailable: // 503
			v := err.(*stypes.CompressToZipServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCompressToZipServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCompressToZipServiceUnavailable().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return staging.NewCompressToZipCreated().WithPayload(&payload)
}

func (p *DataSetAPI) CheckCompressToZipStatus(ctx context.Context, params staging.CheckCompressToZipStatusParams) middleware.Responder {
	rparams := stypes.CheckCompressToZipStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CheckCompressToZipStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckCompressToZipStatus endpoint\n")
		switch err.(type) {
		case *stypes.CheckCompressToZipStatusBadRequest: // 400
			v := err.(*stypes.CheckCompressToZipStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCompressToZipStatusBadRequest().WithPayload(&payload)
		case *stypes.CheckCompressToZipStatusUnauthorized: // 401
			v := err.(*stypes.CheckCompressToZipStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCompressToZipStatusUnauthorized().WithPayload(&payload)
		case *stypes.CheckCompressToZipStatusNotFound: // 404
			v := err.(*stypes.CheckCompressToZipStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCompressToZipStatusNotFound().WithPayload(&payload)
		case *stypes.CheckCompressToZipStatusRequestURITooLong: // 414
			v := err.(*stypes.CheckCompressToZipStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCompressToZipStatusRequestURITooLong().WithPayload(&payload)
		case *stypes.CheckCompressToZipStatusInternalServerError: // 500
			v := err.(*stypes.CheckCompressToZipStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCompressToZipStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCheckCompressToZipStatusInternalServerError().WithPayload(payload)
		}
	}	
	payload := staging.CheckCompressToZipStatusOKBody{
		Status:      res.Payload.Status,
		TargetPaths:  res.Payload.TargetPaths,
	}
	return staging.NewCheckCompressToZipStatusOK().WithPayload(&payload)
}

