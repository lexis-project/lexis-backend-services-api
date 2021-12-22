package datasetapi

import (
	"context"

	dstypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	"github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
	"strings"
)

func (p *DataSetAPI) FilePatch (ctx context.Context, params data_set_management.FilePatchParams) middleware.Responder { 

	rparams := dstypes.FilePatchParams{
		Body: params.Body,
		ContentLength: params.ContentLength,
		TusResumable: params.TusResumable,
		UploadChecksum: params.UploadChecksum,
		UploadOffset: params.UploadOffset,
		ID: params.ID,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.
		FilePatch(ctx, &rparams)
	if err != nil {
		l.Info.Printf("Error calling FilePatch endpoint\n")
		//non 2XX status codes are returned in err here!
		switch err.(type) {
		case *dstypes.FilePatchUnauthorized: // 401
			v := err.(*dstypes.FilePatchUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewFilePatchUnauthorized().WithPayload(&payload)
		case *dstypes.FilePatchUnsupportedMediaType: // 415
			r := data_set_management.NewFilePatchUnsupportedMediaType()
			r.TusResumable = err.(*dstypes.FilePatchForbidden).TusResumable
			r.AccessControlExposeHeaders = "Tus-Resumable"
			return r
		case *dstypes.FilePatchNotFound: // 404
			r := data_set_management.NewFilePatchNotFound()
			r.TusResumable = err.(*dstypes.FilePatchNotFound).TusResumable
			r.AccessControlExposeHeaders = "Tus-Resumable"
			return r
		case *dstypes.FilePatchForbidden: // 403
			r := data_set_management.NewFilePatchForbidden()
			r.TusResumable = err.(*dstypes.FilePatchForbidden).TusResumable
			r.AccessControlExposeHeaders = "Tus-Resumable"
			return r
		case *dstypes.FilePatchGone: // 410
			r := data_set_management.NewFilePatchGone()
			r.TusResumable = err.(*dstypes.FilePatchGone).TusResumable
			r.AccessControlExposeHeaders = "Tus-Resumable"
			return r
		case *dstypes.FilePatchBadRequest: // 400
			r := data_set_management.NewFilePatchBadRequest()
			r.TusResumable = err.(*dstypes.FilePatchBadRequest).TusResumable
			r.AccessControlExposeHeaders = "Tus-Resumable"
			return r
		case *dstypes.FilePatchStatus460: // 460
			r := data_set_management.NewFilePatchStatus460()
			r.TusResumable = err.(*dstypes.FilePatchStatus460).TusResumable
			r.AccessControlExposeHeaders = "Tus-Resumable"
			return r
		case *dstypes.FilePatchPreconditionFailed: // 412
			r := data_set_management.NewFilePatchPreconditionFailed()
			r.TusResumable = err.(*dstypes.FilePatchPreconditionFailed).TusResumable
			r.AccessControlExposeHeaders = "Tus-Resumable"
			return r
		case *dstypes.FilePatchServiceUnavailable: // 503
			v := err.(*dstypes.FilePatchServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewFilePatchServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewFilePatchServiceUnavailable().WithPayload(payload)
		}
	}
	r := data_set_management.FilePatchNoContent {
		TusResumable : res.TusResumable,
		UploadExpires : res.UploadExpires,
		UploadOffset : res.UploadOffset,
		AccessControlExposeHeaders : "Tus-Resumable, Upload-Offset, Upload-Expires",
	}
	return &r
}

func (p *DataSetAPI) FilesDelete (ctx context.Context, params data_set_management.FilesDeleteParams) middleware.Responder { 
	rparams := dstypes.FilesDeleteParams{
		TusResumable: params.TusResumable,
		ID: params.ID,
	}
	res, err := p.getClient(params.HTTPRequest).DataSetManagement.
		FilesDelete(ctx, &rparams)
	if err != nil {
		l.Info.Printf("Error calling FilesDelete endpoint\n")
		//non 2XX status codes are returned in err here!
		switch err.(type) {
		case *dstypes.FilesDeleteUnauthorized: // 401
			v := err.(*dstypes.FilesDeleteUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewFilesDeleteUnauthorized().WithPayload(&payload)
		case *dstypes.FilesDeleteServiceUnavailable: // 503
			v := err.(*dstypes.FilesDeleteServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewFilesDeleteServiceUnavailable().WithPayload(&payload)
		case *dstypes.FilesDeletePreconditionFailed: // 412
			r := data_set_management.NewFilesDeletePreconditionFailed()
			r.TusResumable = err.(*dstypes.FilesDeletePreconditionFailed).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewFilesDeleteServiceUnavailable().WithPayload(payload)
		}
	}
	r := data_set_management.FilesDeleteNoContent {
		TusResumable : res.TusResumable,
		AccessControlExposeHeaders: "TusResumable",
	}
	return &r
}

func (p *DataSetAPI) FilesHead (ctx context.Context, params data_set_management.FilesHeadParams) middleware.Responder { 
	rparams := dstypes.FilesHeadParams{
		TusResumable: params.TusResumable,
		ID: params.ID,
	}
	res, err := p.getClient(params.HTTPRequest).DataSetManagement.
		FilesHead(ctx, &rparams)
	if err != nil {
		l.Info.Printf("Error calling FilesHead endpoint\n")
		//non 2XX status codes are returned in err here!
		switch err.(type) {
		case *dstypes.FilesHeadUnauthorized: // 401
			v := err.(*dstypes.FilesHeadUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewFilesHeadUnauthorized().WithPayload(&payload)
		case *dstypes.FilesHeadNotFound: // 404
			r := data_set_management.NewFilesHeadNotFound()
			r.TusResumable = err.(*dstypes.FilesHeadNotFound).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		case *dstypes.FilesHeadForbidden: // 403
			r := data_set_management.NewFilesHeadForbidden()
			r.TusResumable = err.(*dstypes.FilesHeadForbidden).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		case *dstypes.FilesHeadGone: // 410
			r := data_set_management.NewFilesHeadGone()
			r.TusResumable = err.(*dstypes.FilesHeadGone).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		case *dstypes.FilesHeadPreconditionFailed: // 412
			r := data_set_management.NewFilesHeadPreconditionFailed()
			r.TusResumable = err.(*dstypes.FilesHeadPreconditionFailed).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		case *dstypes.FilesHeadServiceUnavailable: // 503
			v := err.(*dstypes.FilesHeadServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewFilesHeadServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewFilesHeadServiceUnavailable().WithPayload(payload)
		}
	}
	r := data_set_management.FilesHeadOK {
		TusResumable : res.TusResumable,
		UploadOffset : res.UploadOffset,
		UploadLength : res.UploadLength,
		CacheControl: res.CacheControl,
		AccessControlExposeHeaders : "Upload-Offset, Upload-Length, Tus-Resumable",
	}
	return &r
}

func (p *DataSetAPI) OptionsDatasetUpload (ctx context.Context, params data_set_management.OptionsDatasetUploadParams) middleware.Responder { 

	rparams := dstypes.OptionsDatasetUploadParams{
		TusResumable: params.TusResumable,
	}

	resOK, resNoContent, err := p.getClient(params.HTTPRequest).DataSetManagement.
		OptionsDatasetUpload(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling OptionsDatasetUpload endpoint\n")
		//non 2XX status codes are returned in err here!
		switch err.(type) {
		case *dstypes.OptionsDatasetUploadUnauthorized: // 401
			v := err.(*dstypes.OptionsDatasetUploadUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewOptionsDatasetUploadUnauthorized().WithPayload(&payload)
		case *dstypes.OptionsDatasetUploadServiceUnavailable: // 503
			v := err.(*dstypes.OptionsDatasetUploadServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewOptionsDatasetUploadServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewOptionsDatasetUploadServiceUnavailable().WithPayload(payload)
		}
	}
	if resOK != nil {
		r := data_set_management.OptionsDatasetUploadOK {
			TusResumable : resOK.TusResumable,
			TusVersion : resOK.TusVersion,
			TusMaxSize: resOK.TusMaxSize,
			TusExtension: resOK.TusExtension,
			TusChecksumAlgorithm: resOK.TusChecksumAlgorithm,
			AccessControlExposeHeaders: "Tus-Version, Tus-Resumable, Tus-Max-Size, Tus-Extension, Tus-Checksum-Algorithm",
		}
		return &r
	}
	if resNoContent != nil {
		r := data_set_management.OptionsDatasetUploadNoContent {
			TusResumable : resNoContent.TusResumable,
			TusVersion : resNoContent.TusVersion,
			TusMaxSize: resNoContent.TusMaxSize,
			TusExtension: resNoContent.TusExtension,
			TusChecksumAlgorithm: resNoContent.TusChecksumAlgorithm,
			AccessControlExposeHeaders: "Tus-Version, Tus-Resumable, Tus-Max-Size, Tus-Extension, Tus-Checksum-Algorithm",
		}
		return &r
	}
	return nil // can't happen
}

func (p *DataSetAPI) PostDatasetUpload (ctx context.Context, params data_set_management.PostDatasetUploadParams) middleware.Responder { 

	rparams := dstypes.PostDatasetUploadParams{
		ContentLength: params.ContentLength,
		TusResumable: params.TusResumable,
		UploadChecksum: params.UploadChecksum,
		UploadConcat: params.UploadConcat,
		UploadDeferLength: params.UploadDeferLength,
		UploadMetadata: params.UploadMetadata,
		UploadOffset: params.UploadOffset,
		UploadLength: params.UploadLength,
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.
		PostDatasetUpload(ctx, &rparams)
	if err != nil {
		l.Info.Printf("Error calling PostDatasetUpload endpoint\n")
		//non 2XX status codes are returned in err here!
		switch err.(type) {
		case *dstypes.PostDatasetUploadUnauthorized: // 401
			v := err.(*dstypes.PostDatasetUploadUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPostDatasetUploadUnauthorized().WithPayload(&payload)
		case *dstypes.PostDatasetUploadRequestEntityTooLarge: // 413
			r := data_set_management.NewPostDatasetUploadRequestEntityTooLarge()
			r.TusResumable = err.(*dstypes.PostDatasetUploadRequestEntityTooLarge).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		case *dstypes.PostDatasetUploadUnsupportedMediaType: // 415
			r := data_set_management.NewPostDatasetUploadUnsupportedMediaType()
			r.TusResumable = err.(*dstypes.PostDatasetUploadUnsupportedMediaType).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		case *dstypes.PostDatasetUploadStatus460: // 460
			r := data_set_management.NewPostDatasetUploadStatus460()
			r.TusResumable = err.(*dstypes.PostDatasetUploadStatus460).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		case *dstypes.PostDatasetUploadPreconditionFailed: // 412
			r := data_set_management.NewPostDatasetUploadPreconditionFailed()
			r.TusResumable = err.(*dstypes.PostDatasetUploadPreconditionFailed).TusResumable
			r.AccessControlExposeHeaders = "TusResumable"
			return r
		case *dstypes.PostDatasetUploadServiceUnavailable: // 503
			v := err.(*dstypes.PostDatasetUploadServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewPostDatasetUploadServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewPostDatasetUploadServiceUnavailable().WithPayload(payload)
		}
	}
	l := strings.Split(res.Location, "/upload/")[1]
	if p.baseurl[len(p.baseurl)-1] =='/' { //avoid double /
		l = p.baseurl + params.HTTPRequest.URL.Path[1:] + l
	} else {
		l = p.baseurl + params.HTTPRequest.URL.Path + l
	}


	r := data_set_management.PostDatasetUploadCreated {
		Location: l,
		TusResumable : res.TusResumable,
		UploadExpires : res.UploadExpires,
		UploadOffset : res.UploadOffset,
		AccessControlExposeHeaders : "Location, Tus-Resumable, Upload-Offset, Upload-Expires",
	}
	return &r
}
