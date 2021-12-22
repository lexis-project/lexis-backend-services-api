package datasetapi

import (
	"bytes"
	dstypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	models "github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"context"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
	"io/ioutil"
)

func (u *DataSetAPI) PostDatasetStagingDownload(ctx context.Context, params data_set_management.PostDatasetStagingDownloadParams) middleware.Responder {

	rparams := dstypes.PostStagingDownloadParams{
		Parameters: dstypes.PostStagingDownloadBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
		},
	}
	var buffermemory bytes.Buffer
	buffer := ioutil.NopCloser(&buffermemory)
	_, err := u.getClient(params.HTTPRequest).DataSetManagement.PostStagingDownload(ctx, &rparams, &buffermemory)

	if err != nil {
		l.Info.Printf("Error calling PostStagingDownload endpoint\n")
		switch err.(type) {
		case *dstypes.PostStagingDownloadBadRequest: // 400
			v := err.(*dstypes.PostStagingDownloadBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPostDatasetStagingDownloadBadRequest().WithPayload(&payload)
		case *dstypes.PostStagingDownloadUnauthorized: // 401
			v := err.(*dstypes.PostStagingDownloadUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPostDatasetStagingDownloadUnauthorized().WithPayload(&payload)
		case *dstypes.PostStagingDownloadForbidden: // 403
			v := err.(*dstypes.PostStagingDownloadForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPostDatasetStagingDownloadForbidden().WithPayload(&payload)
		case *dstypes.PostStagingDownloadBadGateway: // 502
			v := err.(*dstypes.PostStagingDownloadBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPostDatasetStagingDownloadBadGateway().WithPayload(&payload)
		case *dstypes.PostStagingDownloadServiceUnavailable: // 503
			v := err.(*dstypes.PostStagingDownloadServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPostDatasetStagingDownloadServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewPostDatasetStagingDownloadServiceUnavailable().WithPayload(payload)
		}
	}

	// need to perform a type transformation here...
	//return data_set_management.NewPostStagingDownloadCreated().WithPayload(resp.Payload)
	return data_set_management.NewPostDatasetStagingDownloadOK().WithPayload(buffer)
}
