package datasetapi

import (
	"context"
	dstypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) CheckCompressionEncryptionStatus(ctx context.Context, params data_set_management.CheckCompressionEncryptionStatusParams) middleware.Responder {
	rparams := dstypes.CheckCompressionEncryptionStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckCompressionEncryptionStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckCompressionEncryptionStatus endpoint\n")
		switch err.(type) {
		case *dstypes.CheckCompressionEncryptionStatusBadRequest: // 400
			v := err.(*dstypes.CheckCompressionEncryptionStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionEncryptionStatusBadRequest().WithPayload(&payload)
		case *dstypes.CheckCompressionEncryptionStatusUnauthorized: // 401
			v := err.(*dstypes.CheckCompressionEncryptionStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionEncryptionStatusUnauthorized().WithPayload(&payload)
		case *dstypes.CheckCompressionEncryptionStatusNotFound: // 404
			v := err.(*dstypes.CheckCompressionEncryptionStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionEncryptionStatusNotFound().WithPayload(&payload)
		case *dstypes.CheckCompressionEncryptionStatusRequestURITooLong: // 414
			v := err.(*dstypes.CheckCompressionEncryptionStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionEncryptionStatusRequestURITooLong().WithPayload(&payload)
		case *dstypes.CheckCompressionEncryptionStatusInternalServerError: // 500
			v := err.(*dstypes.CheckCompressionEncryptionStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionEncryptionStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckCompressionEncryptionStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := data_set_management.CheckCompressionEncryptionStatusOKBody{
		Status: res.Payload.Status,
		TargetPath: res.Payload.TargetPath,
	}
	return data_set_management.NewCheckCompressionEncryptionStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CheckDecryptionStatus(ctx context.Context, params data_set_management.CheckDecryptionStatusParams) middleware.Responder {
	rparams := dstypes.CheckDecryptionStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckDecryptionStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckDecryptionStatus endpoint\n")
		switch err.(type) {
		case *dstypes.CheckDecryptionStatusBadRequest: // 400
			v := err.(*dstypes.CheckDecryptionStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionStatusBadRequest().WithPayload(&payload)
		case *dstypes.CheckDecryptionStatusUnauthorized: // 401
			v := err.(*dstypes.CheckDecryptionStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionStatusUnauthorized().WithPayload(&payload)
		case *dstypes.CheckDecryptionStatusNotFound: // 404
			v := err.(*dstypes.CheckDecryptionStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionStatusNotFound().WithPayload(&payload)
		case *dstypes.CheckDecryptionStatusRequestURITooLong: // 414
			v := err.(*dstypes.CheckDecryptionStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionStatusRequestURITooLong().WithPayload(&payload)
		case *dstypes.CheckDecryptionStatusInternalServerError: // 500
			v := err.(*dstypes.CheckDecryptionStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckDecryptionStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := data_set_management.CheckDecryptionStatusOKBody{
		Status: res.Payload.Status,
		TargetPath: res.Payload.TargetPath,
	}
	return data_set_management.NewCheckDecryptionStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CheckDecompressionStatus(ctx context.Context, params data_set_management.CheckDecompressionStatusParams) middleware.Responder {
	rparams := dstypes.CheckDecompressionStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckDecompressionStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckDecompressionStatus endpoint\n")
		switch err.(type) {
		case *dstypes.CheckDecompressionStatusBadRequest: // 400
			v := err.(*dstypes.CheckDecompressionStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecompressionStatusBadRequest().WithPayload(&payload)
		case *dstypes.CheckDecompressionStatusUnauthorized: // 401
			v := err.(*dstypes.CheckDecompressionStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecompressionStatusUnauthorized().WithPayload(&payload)
		case *dstypes.CheckDecompressionStatusNotFound: // 404
			v := err.(*dstypes.CheckDecompressionStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecompressionStatusNotFound().WithPayload(&payload)
		case *dstypes.CheckDecompressionStatusRequestURITooLong: // 414
			v := err.(*dstypes.CheckDecompressionStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecompressionStatusRequestURITooLong().WithPayload(&payload)
		case *dstypes.CheckDecompressionStatusInternalServerError: // 500
			v := err.(*dstypes.CheckDecompressionStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecompressionStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckDecompressionStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := data_set_management.CheckDecompressionStatusOKBody{
		Status: res.Payload.Status,
		TargetPath: res.Payload.TargetPath,
	}
	return data_set_management.NewCheckDecompressionStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CheckDecryptionDecompressionStatus(ctx context.Context, params data_set_management.CheckDecryptionDecompressionStatusParams) middleware.Responder {
	rparams := dstypes.CheckDecryptionDecompressionStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckDecryptionDecompressionStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckDecryptionDecompressionStatus endpoint\n")
		switch err.(type) {
		case *dstypes.CheckDecryptionDecompressionStatusBadRequest: // 400
			v := err.(*dstypes.CheckDecryptionDecompressionStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionDecompressionStatusBadRequest().WithPayload(&payload)
		case *dstypes.CheckDecryptionDecompressionStatusUnauthorized: // 401
			v := err.(*dstypes.CheckDecryptionDecompressionStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionDecompressionStatusUnauthorized().WithPayload(&payload)
		case *dstypes.CheckDecryptionDecompressionStatusNotFound: // 404
			v := err.(*dstypes.CheckDecryptionDecompressionStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionDecompressionStatusNotFound().WithPayload(&payload)
		case *dstypes.CheckDecryptionDecompressionStatusRequestURITooLong: // 414
			v := err.(*dstypes.CheckDecryptionDecompressionStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionDecompressionStatusRequestURITooLong().WithPayload(&payload)
		case *dstypes.CheckDecryptionDecompressionStatusInternalServerError: // 500
			v := err.(*dstypes.CheckDecryptionDecompressionStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckDecryptionDecompressionStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckDecryptionDecompressionStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := data_set_management.CheckDecryptionDecompressionStatusOKBody{
		Status: res.Payload.Status,
		TargetPath: res.Payload.TargetPath,
	}
	return data_set_management.NewCheckDecryptionDecompressionStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CheckCompressionStatus(ctx context.Context, params data_set_management.CheckCompressionStatusParams) middleware.Responder {
	rparams := dstypes.CheckCompressionStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckCompressionStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckCompressionStatus endpoint\n")
		switch err.(type) {
		case *dstypes.CheckCompressionStatusBadRequest: // 400
			v := err.(*dstypes.CheckCompressionStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionStatusBadRequest().WithPayload(&payload)
		case *dstypes.CheckCompressionStatusUnauthorized: // 401
			v := err.(*dstypes.CheckCompressionStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionStatusUnauthorized().WithPayload(&payload)
		case *dstypes.CheckCompressionStatusNotFound: // 404
			v := err.(*dstypes.CheckCompressionStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionStatusNotFound().WithPayload(&payload)
		case *dstypes.CheckCompressionStatusRequestURITooLong: // 414
			v := err.(*dstypes.CheckCompressionStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionStatusRequestURITooLong().WithPayload(&payload)
		case *dstypes.CheckCompressionStatusInternalServerError: // 500
			v := err.(*dstypes.CheckCompressionStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckCompressionStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckCompressionStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := data_set_management.CheckCompressionStatusOKBody{
		Status: res.Payload.Status,
		TargetPath: res.Payload.TargetPath,
	}
	return data_set_management.NewCheckCompressionStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) CheckEncryptionStatus(ctx context.Context, params data_set_management.CheckEncryptionStatusParams) middleware.Responder {
	rparams := dstypes.CheckEncryptionStatusParams{
		RequestID: params.RequestID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckEncryptionStatus(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckEncryptionStatus endpoint\n")
		switch err.(type) {
		case *dstypes.CheckEncryptionStatusBadRequest: // 400
			v := err.(*dstypes.CheckEncryptionStatusBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckEncryptionStatusBadRequest().WithPayload(&payload)
		case *dstypes.CheckEncryptionStatusUnauthorized: // 401
			v := err.(*dstypes.CheckEncryptionStatusUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckEncryptionStatusUnauthorized().WithPayload(&payload)
		case *dstypes.CheckEncryptionStatusNotFound: // 404
			v := err.(*dstypes.CheckEncryptionStatusNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckEncryptionStatusNotFound().WithPayload(&payload)
		case *dstypes.CheckEncryptionStatusRequestURITooLong: // 414
			v := err.(*dstypes.CheckEncryptionStatusRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckEncryptionStatusRequestURITooLong().WithPayload(&payload)
		case *dstypes.CheckEncryptionStatusInternalServerError: // 500
			v := err.(*dstypes.CheckEncryptionStatusInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckEncryptionStatusInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckEncryptionStatusInternalServerError().WithPayload(payload)
		}
	}
	payload := data_set_management.CheckEncryptionStatusOKBody{
		Status: res.Payload.Status,
		TargetPath: res.Payload.TargetPath,
	}
	return data_set_management.NewCheckEncryptionStatusOK().WithPayload(&payload)
}

func (p *DataSetAPI) Compress(ctx context.Context, params data_set_management.CompressParams) middleware.Responder {
	rparams := dstypes.CompressParams{
		Parameters: dstypes.CompressBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.Compress(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Compress endpoint\n")
		switch err.(type) {
		case *dstypes.CompressBadRequest: // 400
			v := err.(*dstypes.CompressBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressBadRequest().WithPayload(&payload)
		case *dstypes.CompressUnauthorized: // 401
			v := err.(*dstypes.CompressUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressUnauthorized().WithPayload(&payload)
		case *dstypes.CompressForbidden: // 403
			v := err.(*dstypes.CompressForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressForbidden().WithPayload(&payload)
		case *dstypes.CompressNotFound: // 404
			v := err.(*dstypes.CompressNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressNotFound().WithPayload(&payload)
		case *dstypes.CompressRequestURITooLong: // 414
			v := err.(*dstypes.CompressRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressRequestURITooLong().WithPayload(&payload)
		case *dstypes.CompressTooManyRequests: // 429
			v := err.(*dstypes.CompressTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressTooManyRequests().WithPayload(&payload)
		case *dstypes.CompressInternalServerError: // 500
			v := err.(*dstypes.CompressInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCompressInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewCompressCreated().WithPayload(&payload)
}

func (p *DataSetAPI) CompressEncrypt(ctx context.Context, params data_set_management.CompressEncryptParams) middleware.Responder {
	rparams := dstypes.CompressEncryptParams{
		Parameters: dstypes.CompressEncryptBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
			Project:      params.Parameters.Project,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CompressEncrypt(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CompressEncrypt endpoint\n")
		switch err.(type) {
		case *dstypes.CompressEncryptBadRequest: // 400
			v := err.(*dstypes.CompressEncryptBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressEncryptBadRequest().WithPayload(&payload)
		case *dstypes.CompressEncryptUnauthorized: // 401
			v := err.(*dstypes.CompressEncryptUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressEncryptUnauthorized().WithPayload(&payload)
		case *dstypes.CompressEncryptForbidden: // 403
			v := err.(*dstypes.CompressEncryptForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressEncryptForbidden().WithPayload(&payload)
		case *dstypes.CompressEncryptNotFound: // 404
			v := err.(*dstypes.CompressEncryptNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressEncryptNotFound().WithPayload(&payload)
		case *dstypes.CompressEncryptRequestURITooLong: // 414
			v := err.(*dstypes.CompressEncryptRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressEncryptRequestURITooLong().WithPayload(&payload)
		case *dstypes.CompressEncryptTooManyRequests: // 429
			v := err.(*dstypes.CompressEncryptTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressEncryptTooManyRequests().WithPayload(&payload)
		case *dstypes.CompressEncryptInternalServerError: // 500
			v := err.(*dstypes.CompressEncryptInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCompressEncryptInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCompressEncryptInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewCompressEncryptCreated().WithPayload(&payload)
}

func (p *DataSetAPI) Decompress(ctx context.Context, params data_set_management.DecompressParams) middleware.Responder {
	rparams := dstypes.DecompressParams{
		Parameters: dstypes.DecompressBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.Decompress(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Decompress endpoint\n")
		switch err.(type) {
		case *dstypes.DecompressBadRequest: // 400
			v := err.(*dstypes.DecompressBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecompressBadRequest().WithPayload(&payload)
		case *dstypes.DecompressUnauthorized: // 401
			v := err.(*dstypes.DecompressUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecompressUnauthorized().WithPayload(&payload)
		case *dstypes.DecompressForbidden: // 403
			v := err.(*dstypes.DecompressForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecompressForbidden().WithPayload(&payload)
		case *dstypes.DecompressNotFound: // 404
			v := err.(*dstypes.DecompressNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecompressNotFound().WithPayload(&payload)
		case *dstypes.DecompressRequestURITooLong: // 414
			v := err.(*dstypes.DecompressRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecompressRequestURITooLong().WithPayload(&payload)
		case *dstypes.DecompressTooManyRequests: // 429
			v := err.(*dstypes.DecompressTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecompressTooManyRequests().WithPayload(&payload)
		case *dstypes.DecompressInternalServerError: // 500
			v := err.(*dstypes.DecompressInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecompressInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewDecompressInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewDecompressCreated().WithPayload(&payload)
}

func (p *DataSetAPI) Decrypt(ctx context.Context, params data_set_management.DecryptParams) middleware.Responder {
	rparams := dstypes.DecryptParams{
		Parameters: dstypes.DecryptBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
			Project:      params.Parameters.Project,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.Decrypt(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Decrypt endpoint\n")
		switch err.(type) {
		case *dstypes.DecryptBadRequest: // 400
			v := err.(*dstypes.DecryptBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptBadRequest().WithPayload(&payload)
		case *dstypes.DecryptUnauthorized: // 401
			v := err.(*dstypes.DecryptUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptUnauthorized().WithPayload(&payload)
		case *dstypes.DecryptForbidden: // 403
			v := err.(*dstypes.DecryptForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptForbidden().WithPayload(&payload)
		case *dstypes.DecryptNotFound: // 404
			v := err.(*dstypes.DecryptNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptNotFound().WithPayload(&payload)
		case *dstypes.DecryptRequestURITooLong: // 414
			v := err.(*dstypes.DecryptRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptRequestURITooLong().WithPayload(&payload)
		case *dstypes.DecryptTooManyRequests: // 429
			v := err.(*dstypes.DecryptTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptTooManyRequests().WithPayload(&payload)
		case *dstypes.DecryptInternalServerError: // 500
			v := err.(*dstypes.DecryptInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewDecryptInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewDecryptCreated().WithPayload(&payload)
}

func (p *DataSetAPI) DecryptDecompress(ctx context.Context, params data_set_management.DecryptDecompressParams) middleware.Responder {
	rparams := dstypes.DecryptDecompressParams{
		Parameters: dstypes.DecryptDecompressBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
                        Project:      params.Parameters.Project,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.DecryptDecompress(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling DecryptDecompress endpoint\n")
		switch err.(type) {
		case *dstypes.DecryptDecompressBadRequest: // 400
			v := err.(*dstypes.DecryptDecompressBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptDecompressBadRequest().WithPayload(&payload)
		case *dstypes.DecryptDecompressUnauthorized: // 401
			v := err.(*dstypes.DecryptDecompressUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptDecompressUnauthorized().WithPayload(&payload)
		case *dstypes.DecryptDecompressForbidden: // 403
			v := err.(*dstypes.DecryptDecompressForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptDecompressForbidden().WithPayload(&payload)
		case *dstypes.DecryptDecompressNotFound: // 404
			v := err.(*dstypes.DecryptDecompressNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptDecompressNotFound().WithPayload(&payload)
		case *dstypes.DecryptDecompressRequestURITooLong: // 414
			v := err.(*dstypes.DecryptDecompressRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptDecompressRequestURITooLong().WithPayload(&payload)
		case *dstypes.DecryptDecompressTooManyRequests: // 429
			v := err.(*dstypes.DecryptDecompressTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptDecompressTooManyRequests().WithPayload(&payload)
		case *dstypes.DecryptDecompressInternalServerError: // 500
			v := err.(*dstypes.DecryptDecompressInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDecryptDecompressInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewDecryptDecompressInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewDecryptDecompressCreated().WithPayload(&payload)
}

func (p *DataSetAPI) Encrypt(ctx context.Context, params data_set_management.EncryptParams) middleware.Responder {
	rparams := dstypes.EncryptParams{
		Parameters: dstypes.EncryptBody{
			SourceSystem: params.Parameters.SourceSystem,
			SourcePath:   params.Parameters.SourcePath,
                        Project:      params.Parameters.Project,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.Encrypt(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling Encrypt endpoint\n")
		switch err.(type) {
		case *dstypes.EncryptBadRequest: // 400
			v := err.(*dstypes.EncryptBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewEncryptBadRequest().WithPayload(&payload)
		case *dstypes.EncryptUnauthorized: // 401
			v := err.(*dstypes.EncryptUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewEncryptUnauthorized().WithPayload(&payload)
		case *dstypes.EncryptForbidden: // 403
			v := err.(*dstypes.EncryptForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewEncryptForbidden().WithPayload(&payload)
		case *dstypes.EncryptNotFound: // 404
			v := err.(*dstypes.EncryptNotFound)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewEncryptNotFound().WithPayload(&payload)
		case *dstypes.EncryptRequestURITooLong: // 414
			v := err.(*dstypes.EncryptRequestURITooLong)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewEncryptRequestURITooLong().WithPayload(&payload)
		case *dstypes.EncryptTooManyRequests: // 429
			v := err.(*dstypes.EncryptTooManyRequests)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewEncryptTooManyRequests().WithPayload(&payload)
		case *dstypes.EncryptInternalServerError: // 500
			v := err.(*dstypes.EncryptInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewEncryptInternalServerError().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewEncryptInternalServerError().WithPayload(payload)
		}
	}
	payload := models.SteeringRequestID{
		RequestID: res.Payload.RequestID,
	}
	return data_set_management.NewEncryptCreated().WithPayload(&payload)
}

