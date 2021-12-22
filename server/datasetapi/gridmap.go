package datasetapi

import (
	"context"

	dstypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) AddGridmapEntry(ctx context.Context, params data_set_management.AddGridmapEntryParams) middleware.Responder { // need to include params here..
	rparams := dstypes.AddGridmapEntryParams{
		Parameters: dstypes.AddGridmapEntryBody{
			Dn:   params.Parameters.Dn,
			User: params.Parameters.User,
		},
	}

	_, err := p.getClient(params.HTTPRequest).DataSetManagement.AddGridmapEntry(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling AddGridmapEntry endpoint\n")
		switch err.(type) {
		case *dstypes.AddGridmapEntryUnauthorized: // 401
			v := err.(*dstypes.AddGridmapEntryUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewAddGridmapEntryUnauthorized().WithPayload(&payload)
		case *dstypes.AddGridmapEntryForbidden: // 403
			v := err.(*dstypes.AddGridmapEntryForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewAddGridmapEntryForbidden().WithPayload(&payload)
		case *dstypes.AddGridmapEntryServiceUnavailable: // 503
			v := err.(*dstypes.AddGridmapEntryServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewAddGridmapEntryServiceUnavailable().WithPayload(&payload)
		case *dstypes.AddGridmapEntryBadGateway: // 502
			v := err.(*dstypes.AddGridmapEntryBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewAddGridmapEntryBadGateway().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewAddGridmapEntryServiceUnavailable().WithPayload(payload)
		}
	}
	return data_set_management.NewAddGridmapEntryCreated()
}

func (p *DataSetAPI) RemoveGridmapEntry(ctx context.Context, params data_set_management.RemoveGridmapEntryParams) middleware.Responder { // need to include params here..
	rparams := dstypes.RemoveGridmapEntryParams{
		Parameters: dstypes.RemoveGridmapEntryBody{
			User: params.Parameters.User,
		},
	}

	_, err := p.getClient(params.HTTPRequest).DataSetManagement.RemoveGridmapEntry(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling RemoveGridmapEntry endpoint\n")
		switch err.(type) {
		case *dstypes.RemoveGridmapEntryUnauthorized: // 401
			v := err.(*dstypes.RemoveGridmapEntryUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewRemoveGridmapEntryUnauthorized().WithPayload(&payload)
		case *dstypes.RemoveGridmapEntryForbidden: // 403
			v := err.(*dstypes.RemoveGridmapEntryForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewRemoveGridmapEntryForbidden().WithPayload(&payload)
		case *dstypes.RemoveGridmapEntryServiceUnavailable: // 503
			v := err.(*dstypes.RemoveGridmapEntryServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewRemoveGridmapEntryServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewRemoveGridmapEntryServiceUnavailable().WithPayload(payload)
		}
	}
	return data_set_management.NewRemoveGridmapEntryNoContent()
}
