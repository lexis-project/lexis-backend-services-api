package datasetapi

import (
	"context"

	stypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/staging"
	remotedatamodel "github.com/lexis-project/lexis-backend-services-interface-datasets.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/staging"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) CheckDeletionStatus(ctx context.Context, params staging.CheckDeletionStatusParams) middleware.Responder {
	rparams := stypes.CheckDeletionStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CheckDeletionStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckDeletionStatus endpoint\n")
		switch err.(type) {
		case *stypes.CheckDeletionStatusBadRequest: // 400
			v := err.(*stypes.CheckDeletionStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDeletionStatusBadRequest().WithPayload(&payload)
		case *stypes.CheckDeletionStatusUnauthorized: // 401
			v := err.(*stypes.CheckDeletionStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDeletionStatusUnauthorized().WithPayload(&payload)
		case *stypes.CheckDeletionStatusNotFound: // 404
			v := err.(*stypes.CheckDeletionStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDeletionStatusNotFound().WithPayload(&payload)
		case *stypes.CheckDeletionStatusRequestURITooLong: // 414
			v := err.(*stypes.CheckDeletionStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDeletionStatusRequestURITooLong().WithPayload(&payload)
		case *stypes.CheckDeletionStatusInternalServerError: // 500
			v := err.(*stypes.CheckDeletionStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckDeletionStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCheckDeletionStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := staging.CheckDeletionStatusOKBody{
		Status: res.Payload.Status,
	}
	return staging.NewCheckDeletionStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CheckStageStatus(ctx context.Context, params staging.CheckStageStatusParams) middleware.Responder {
	rparams := stypes.CheckStageStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).Staging.CheckStageStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckStageStatus endpoint\n")
		switch err.(type) {
		case *stypes.CheckStageStatusBadRequest: // 400
			v := err.(*stypes.CheckStageStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckStageStatusBadRequest().WithPayload(&payload)
		case *stypes.CheckStageStatusUnauthorized: // 401
			v := err.(*stypes.CheckStageStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckStageStatusUnauthorized().WithPayload(&payload)
		case *stypes.CheckStageStatusNotFound: // 404
			v := err.(*stypes.CheckStageStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckStageStatusNotFound().WithPayload(&payload)
		case *stypes.CheckStageStatusRequestURITooLong: // 414
			v := err.(*stypes.CheckStageStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckStageStatusRequestURITooLong().WithPayload(&payload)
		case *stypes.CheckStageStatusInternalServerError: // 500
			v := err.(*stypes.CheckStageStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewCheckStageStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewCheckStageStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := staging.CheckStageStatusOKBody{
		Status:      res.Payload.Status,
		TargetPath:  res.Payload.TargetPath,
	}
	return staging.NewCheckStageStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) Delete(ctx context.Context, params staging.DeleteParams) middleware.Responder {
	rparams := stypes.DeleteParams{
		Parameters: stypes.DeleteBody{
			TargetSystem: params.Parameters.TargetSystem,
			TargetPath:   params.Parameters.TargetPath,
		},
	}

	res, err := p.getClient(params.HTTPRequest).Staging.Delete(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Delete endpoint\n")
		switch err.(type) {
		case *stypes.DeleteBadRequest: // 400
			v := err.(*stypes.DeleteBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDeleteBadRequest().WithPayload(&payload)
		case *stypes.DeleteUnauthorized: // 401
			v := err.(*stypes.DeleteUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDeleteUnauthorized().WithPayload(&payload)
		case *stypes.DeleteForbidden: // 403
			v := err.(*stypes.DeleteForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDeleteForbidden().WithPayload(&payload)
		case *stypes.DeleteNotFound: // 404
			v := err.(*stypes.DeleteNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDeleteNotFound().WithPayload(&payload)
		case *stypes.DeleteRequestURITooLong: // 414
			v := err.(*stypes.DeleteRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDeleteRequestURITooLong().WithPayload(&payload)
		case *stypes.DeleteTooManyRequests: // 429
			v := err.(*stypes.DeleteTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDeleteTooManyRequests().WithPayload(&payload)
		case *stypes.DeleteInternalServerError: // 500
			v := err.(*stypes.DeleteInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewDeleteInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewDeleteInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return staging.NewDeleteCreated().WithPayload(&payload)
}

func (p *DataSetAPI) Stage(ctx context.Context, params staging.StageParams) middleware.Responder {
	rparams := stypes.StageParams{
		Parameters: stypes.StageBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
			TargetSystem: params.Parameters.TargetSystem,
			TargetPath:   params.Parameters.TargetPath,
			JobID:        params.Parameters.JobID,
			TaskID:       params.Parameters.TaskID,
                        HeappeURL:    params.Parameters.HeappeURL,
                        Encryption:   params.Parameters.Encryption,
                        Compression:  params.Parameters.Compression,
		},
	}
	if params.Parameters.Metadata != nil {
		rparams.Parameters.Metadata = &remotedatamodel.Metadata{
			CustomMetadata:       params.Parameters.Metadata.CustomMetadata,
			CustomMetadataSchema: params.Parameters.Metadata.CustomMetadataSchema,
			Creator:              params.Parameters.Metadata.Creator,
			Contributor:          params.Parameters.Metadata.Contributor,
			Publisher:            params.Parameters.Metadata.Publisher,
			Owner:                params.Parameters.Metadata.Owner,
			Identifier:           params.Parameters.Metadata.Identifier,
			PublicationYear:      params.Parameters.Metadata.PublicationYear,
			ResourceType:         params.Parameters.Metadata.ResourceType,
			ResourceTypeGeneral:  params.Parameters.Metadata.ResourceTypeGeneral,
			Title:                params.Parameters.Metadata.Title,
			RelatedIdentifier:    params.Parameters.Metadata.RelatedIdentifier,
			Rights:               params.Parameters.Metadata.Rights,
			RightsIdentifier:     params.Parameters.Metadata.RightsIdentifier,
			RightsURI:            params.Parameters.Metadata.RightsURI,
		}
	}

	res, err := p.getClient(params.HTTPRequest).Staging.Stage(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Stage endpoint\n")
		switch err.(type) {
		case *stypes.StageBadRequest: // 400
			v := err.(*stypes.StageBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewStageBadRequest().WithPayload(&payload)
		case *stypes.StageUnauthorized: // 401
			v := err.(*stypes.StageUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewStageUnauthorized().WithPayload(&payload)
		case *stypes.StageForbidden: // 403
			v := err.(*stypes.StageForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewStageForbidden().WithPayload(&payload)
		case *stypes.StageNotFound: // 404
			v := err.(*stypes.StageNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewStageNotFound().WithPayload(&payload)
		case *stypes.StageRequestURITooLong: // 414
			v := err.(*stypes.StageRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewStageRequestURITooLong().WithPayload(&payload)
		case *stypes.StageTooManyRequests: // 429
			v := err.(*stypes.StageTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewStageTooManyRequests().WithPayload(&payload)
		case *stypes.StageInternalServerError: // 500
			v := err.(*stypes.StageInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return staging.NewStageInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return staging.NewStageInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return staging.NewStageCreated().WithPayload(&payload)
}

func (p *DataSetAPI) StagingInfo(ctx context.Context, params staging.StagingInfoParams) middleware.Responder {
	rparams := stypes.StagingInfoParams{}

	res, err := p.getClient(params.HTTPRequest).Staging.StagingInfo(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling StagingInfo endpoint\n")
		//can't happen
		switch err.(type) {
		default:
			//payload := fillErrorResponse(err)
			return nil
		}
	}
	var payload []string
	payload = res.Payload

	return staging.NewStagingInfoOK().WithPayload(payload)
}
