package datasetapi

import (
	"context"

	stypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/staging"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/staging"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) CheckCloudNFSExportAddStatus(ctx context.Context, params staging.CheckCloudNFSExportAddStatusParams) middleware.Responder {
	rparams := stypes.CheckCloudNFSExportAddStatusParams{
		Param: params.Param,
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CheckCloudNFSExportAddStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckCloudNFSExportAddStatus endpoint\n")
		switch err.(type) {
		case *stypes.CheckCloudNFSExportAddStatusBadRequest: // 400
			v := err.(*stypes.CheckCloudNFSExportAddStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportAddStatusBadRequest().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportAddStatusUnauthorized: // 401
			v := err.(*stypes.CheckCloudNFSExportAddStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportAddStatusUnauthorized().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportAddStatusForbidden: // 403
			v := err.(*stypes.CheckCloudNFSExportAddStatusForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportAddStatusForbidden().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportAddStatusNotFound: // 404
			v := err.(*stypes.CheckCloudNFSExportAddStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportAddStatusNotFound().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportAddStatusRequestURITooLong: // 414
			v := err.(*stypes.CheckCloudNFSExportAddStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportAddStatusRequestURITooLong().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportAddStatusInternalServerError: // 500
			v := err.(*stypes.CheckCloudNFSExportAddStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportAddStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCheckCloudNFSExportAddStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := staging.CheckCloudNFSExportAddStatusOKBody{
		Status: res.Payload.Status,
	}
	return staging.NewCheckCloudNFSExportAddStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CheckCloudNFSExportRemoveStatus(ctx context.Context, params staging.CheckCloudNFSExportRemoveStatusParams) middleware.Responder {
	rparams := stypes.CheckCloudNFSExportRemoveStatusParams{
		Param: params.Param,
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CheckCloudNFSExportRemoveStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckCloudNFSExportRemoveStatus endpoint\n")
		switch err.(type) {
		case *stypes.CheckCloudNFSExportRemoveStatusBadRequest: // 400
			v := err.(*stypes.CheckCloudNFSExportRemoveStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportRemoveStatusBadRequest().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportRemoveStatusUnauthorized: // 401
			v := err.(*stypes.CheckCloudNFSExportRemoveStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportRemoveStatusUnauthorized().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportRemoveStatusNotFound: // 404
			v := err.(*stypes.CheckCloudNFSExportRemoveStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportRemoveStatusNotFound().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportRemoveStatusRequestURITooLong: // 414
			v := err.(*stypes.CheckCloudNFSExportRemoveStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportRemoveStatusRequestURITooLong().WithPayload(&payload)
		case *stypes.CheckCloudNFSExportRemoveStatusInternalServerError: // 500
			v := err.(*stypes.CheckCloudNFSExportRemoveStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckCloudNFSExportRemoveStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCheckCloudNFSExportRemoveStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := staging.CheckCloudNFSExportRemoveStatusOKBody{
		Status: res.Payload.Status,
	}
	return staging.NewCheckCloudNFSExportRemoveStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CloudNFSExportAdd(ctx context.Context, params staging.CloudNFSExportAddParams) middleware.Responder {
	rparams := stypes.CloudNFSExportAddParams{
		Param: params.Param,
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CloudNFSExportAdd(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CloudNFSExportAdd endpoint\n")
		switch err.(type) {
		case *stypes.CloudNFSExportAddBadRequest: // 400
			v := err.(*stypes.CloudNFSExportAddBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportAddBadRequest().WithPayload(&payload)
		case *stypes.CloudNFSExportAddUnauthorized: // 401
			v := err.(*stypes.CloudNFSExportAddUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportAddUnauthorized().WithPayload(&payload)
		case *stypes.CloudNFSExportAddForbidden: // 403
			v := err.(*stypes.CloudNFSExportAddForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportAddForbidden().WithPayload(&payload)
		case *stypes.CloudNFSExportAddRequestURITooLong: // 414
			v := err.(*stypes.CloudNFSExportAddRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportAddRequestURITooLong().WithPayload(&payload)
		case *stypes.CloudNFSExportAddInternalServerError: // 500
			v := err.(*stypes.CloudNFSExportAddInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportAddInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCloudNFSExportAddInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return staging.NewCloudNFSExportAddCreated().WithPayload(&payload)
}

func (p *DataSetAPI) CloudNFSExportRemove(ctx context.Context, params staging.CloudNFSExportRemoveParams) middleware.Responder {
	rparams := stypes.CloudNFSExportRemoveParams{
		Param: params.Param,
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CloudNFSExportRemove(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CloudNFSExportRemove endpoint\n")
		switch err.(type) {
		case *stypes.CloudNFSExportRemoveBadRequest: // 400
			v := err.(*stypes.CloudNFSExportRemoveBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportRemoveBadRequest().WithPayload(&payload)
		case *stypes.CloudNFSExportRemoveUnauthorized: // 401
			v := err.(*stypes.CloudNFSExportRemoveUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportRemoveUnauthorized().WithPayload(&payload)
		case *stypes.CloudNFSExportRemoveRequestURITooLong: // 414
			v := err.(*stypes.CloudNFSExportRemoveRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportRemoveRequestURITooLong().WithPayload(&payload)
		case *stypes.CloudNFSExportRemoveInternalServerError: // 500
			v := err.(*stypes.CloudNFSExportRemoveInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCloudNFSExportRemoveInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCloudNFSExportRemoveInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return staging.NewCloudNFSExportRemoveCreated().WithPayload(&payload)
}
