package datasetapi

import (
	"context"

	dstypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) CreateSSHFSExport(ctx context.Context, params data_set_management.CreateSSHFSExportParams) middleware.Responder { // need to include params here..
	rparams := dstypes.CreateSSHFSExportParams{
		Parameters: dstypes.CreateSSHFSExportBody{
			Host:   params.Parameters.Host,
			Pubkey: params.Parameters.Pubkey,
			Path:   params.Parameters.Path,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CreateSSHFSExport(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CreateSSHFSExport endpoint\n")
		switch err.(type) {
		case *dstypes.CreateSSHFSExportBadRequest: // 400
			v := err.(*dstypes.CreateSSHFSExportBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateSSHFSExportBadRequest().WithPayload(&payload)
		case *dstypes.CreateSSHFSExportUnauthorized: // 401
			v := err.(*dstypes.CreateSSHFSExportUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateSSHFSExportUnauthorized().WithPayload(&payload)
		case *dstypes.CreateSSHFSExportForbidden: // 403
			v := err.(*dstypes.CreateSSHFSExportForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateSSHFSExportForbidden().WithPayload(&payload)
		case *dstypes.CreateSSHFSExportServiceUnavailable: // 503
			v := err.(*dstypes.CreateSSHFSExportServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateSSHFSExportServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCreateSSHFSExportServiceUnavailable().WithPayload(payload)
		}
	}
	ret := data_set_management.CreateSSHFSExportCreatedBody{
		User:  res.Payload.User,
		Sshfs: res.Payload.Sshfs,
	}
	return data_set_management.NewCreateSSHFSExportCreated().WithPayload(&ret)
}

func (p *DataSetAPI) DeleteSSHFSExport(ctx context.Context, params data_set_management.DeleteSSHFSExportParams) middleware.Responder { // need to include params here..
	rparams := dstypes.DeleteSSHFSExportParams{
		Parameters: dstypes.DeleteSSHFSExportBody{
			User: params.Parameters.User,
			Path: params.Parameters.Path,
		},
	}

	_, err := p.getClient(params.HTTPRequest).DataSetManagement.DeleteSSHFSExport(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling DeleteSSHFSExport endpoint\n")
		switch err.(type) {
		case *dstypes.DeleteSSHFSExportBadRequest: // 400
			v := err.(*dstypes.DeleteSSHFSExportBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteSSHFSExportBadRequest().WithPayload(&payload)
		case *dstypes.DeleteSSHFSExportUnauthorized: // 401
			v := err.(*dstypes.DeleteSSHFSExportUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteSSHFSExportUnauthorized().WithPayload(&payload)
		case *dstypes.DeleteSSHFSExportForbidden: // 403
			v := err.(*dstypes.DeleteSSHFSExportForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteSSHFSExportForbidden().WithPayload(&payload)
		case *dstypes.DeleteSSHFSExportNotFound: // 404
			v := err.(*dstypes.DeleteSSHFSExportNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteSSHFSExportNotFound().WithPayload(&payload)
		case *dstypes.DeleteSSHFSExportBadGateway: // 502
			v := err.(*dstypes.DeleteSSHFSExportBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteSSHFSExportBadGateway().WithPayload(&payload)
		case *dstypes.DeleteSSHFSExportServiceUnavailable: // 503
			v := err.(*dstypes.DeleteSSHFSExportServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteSSHFSExportServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewDeleteSSHFSExportServiceUnavailable().WithPayload(payload)
		}
	}

	return data_set_management.NewDeleteSSHFSExportNoContent()
}
