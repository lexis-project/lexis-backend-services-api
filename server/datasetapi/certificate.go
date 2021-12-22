package datasetapi

import (
	"bytes"
	"context"
	"io/ioutil"

	dstypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	models "github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

func (p *DataSetAPI) Certificate(ctx context.Context,
	params data_set_management.CertificateParams) middleware.Responder {

	// no params
	rparams := dstypes.NewCertificateParams()
	//https://stackoverflow.com/questions/40316052/in-memory-file-for-testing
	//https://stackoverflow.com/questions/23454940/getting-bytes-buffer-does-not-implement-io-writer-error-message/23454941
	var buffermemory bytes.Buffer
	buffer := ioutil.NopCloser(&buffermemory)
	_, err := p.getClient(params.HTTPRequest).DataSetManagement.Certificate(ctx, rparams, &buffermemory)

	if err != nil {
		l.Info.Printf("Error calling Certificate endpoint\n")
		switch err.(type) {
		case *dstypes.CertificateUnauthorized: // 401
			v := err.(*dstypes.CertificateUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCertificateUnauthorized().WithPayload(&payload)
		case *dstypes.CertificateInternalServerError: // 500
			v := err.(*dstypes.CertificateInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCertificateInternalServerError().WithPayload(&payload)
		case *dstypes.CertificateBadGateway: // 502
			v := err.(*dstypes.CertificateBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCertificateBadGateway().WithPayload(&payload)
		case *dstypes.CertificateServiceUnavailable: // 503
			v := err.(*dstypes.CertificateServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCertificateServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCertificateServiceUnavailable().WithPayload(payload)
		}
	}

	// need to perform a type transformation here...
	//return data_set_management.NewDownloadDatasetCreated().WithPayload(resp.Payload)
	return data_set_management.NewCertificateOK().WithPayload(buffer)
}
