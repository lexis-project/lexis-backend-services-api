package datasetapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	remotedataapi "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client"
	dstypes "github.com/lexis-project/lexis-backend-services-interface-datasets.git/client/data_set_management"
	remotedatamodel "github.com/lexis-project/lexis-backend-services-interface-datasets.git/models"
	models "github.com/lexis-project/lexis-backend-services-api.git/models"
	"github.com/lexis-project/lexis-backend-services-api.git/restapi/operations/data_set_management"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
	l "gitlab.com/cyclops-utilities/logging"
)

// A simple in memory CRUD on data
// In real life, this should be using a persistent storage.
type DataSetAPI struct {
	c remotedataapi.Config
	baseurl string
}

func New(c remotedataapi.Config, baseurl string) *DataSetAPI {
	return &DataSetAPI{
		c: c,
		baseurl: baseurl,
	}
}

func (p *DataSetAPI) getClient(param *http.Request) *remotedataapi.WP3BackedDataServiceAPI {
	config := p.c
	if len(param.Header.Get("Authorization")) > 0 {
		token := strings.Fields(param.Header.Get("Authorization"))[1]
		config.AuthInfo = httptransport.BearerToken(token)
	}
	r := remotedataapi.New(config)
	return r
}

func fillErrorResponse(e error) *models.ErrorResponse {
	s := e.Error()
	returnValError := models.ErrorResponse{
		ErrorString: &s,
	}
	return &returnValError
}

func copyListingResult(a *models.DatasetContent, b *remotedatamodel.DatasetContent) {
	a.Checksum = b.Checksum
	a.CreateTime = b.CreateTime
	a.Name = b.Name
	a.Size = b.Size
	a.Type = b.Type
	for i := 0; i < len(b.Contents); i++ {
		tmp := models.DatasetContent{}
		copyListingResult(&tmp, b.Contents[i])
		(*a).Contents = append((*a).Contents, &tmp)
	}
}

func (p *DataSetAPI) Listing(ctx context.Context, params data_set_management.ListingParams) middleware.Responder {

	// need to include params here..
	rparams := dstypes.ListingParams{
		JSON: dstypes.ListingBody{
			Access:     remotedatamodel.AccessMode(params.JSON.Access),
			InternalID: params.JSON.InternalID,
			Path:       params.JSON.Path,
			Project:    params.JSON.Project,
			Recursive:  params.JSON.Recursive,
			Zone:	    params.JSON.Zone,
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.Listing(ctx, &rparams)
	if err != nil {
		//non 2XX status codes are returned in err here!
		l.Info.Printf("Error calling Listing endpoint\n")
		switch err.(type) {
		case *dstypes.ListingBadRequest: // 400
			v := err.(*dstypes.ListingBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewListingBadRequest().WithPayload(&payload)
		case *dstypes.ListingUnauthorized: // 401
			v := err.(*dstypes.ListingUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewListingUnauthorized().WithPayload(&payload)
		case *dstypes.ListingForbidden: // 403
			v := err.(*dstypes.ListingForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewListingForbidden().WithPayload(&payload)
		case *dstypes.ListingBadGateway: // 502
			v := err.(*dstypes.ListingBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewListingBadGateway().WithPayload(&payload)
		case *dstypes.ListingServiceUnavailable: // 503
			v := err.(*dstypes.ListingServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			//Irods permissions too restrictive for this user is also 503, return empty list of datasets
			if strings.HasPrefix(*payload.ErrorString, "Irods permissions too restrictive for this user") {
				emptyPayload := models.DatasetContent{}
				return data_set_management.NewListingOK().WithPayload(&emptyPayload)
			}
			return data_set_management.NewListingServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewListingServiceUnavailable().WithPayload(payload)
		}
	}

	payload := models.DatasetContent{}
	copyListingResult(&payload, res.Payload)
	return data_set_management.NewListingOK().WithPayload(&payload)
}

func (p *DataSetAPI) CreateDataset(ctx context.Context, params data_set_management.CreateDatasetParams) middleware.Responder {

	// need to include params here..
	rparams := dstypes.CreateDatasetParams{
		Dataset: dstypes.CreateDatasetBody{
			Path:           params.Dataset.Path,
			File:           params.Dataset.File,
			Name:           params.Dataset.Name,
			InternalID:     params.Dataset.InternalID, //string
			Access:         remotedatamodel.AccessMode(params.Dataset.Access),
			Project:        params.Dataset.Project, //string*
			PushMethod:     remotedatamodel.PushMethod(params.Dataset.PushMethod),
			CompressMethod: remotedatamodel.CompressMethod(params.Dataset.CompressMethod),
			Enc:            params.Dataset.Enc,
			Comp:           params.Dataset.Comp,
			Zone:		params.Dataset.Zone,
		},
	}

	if params.Dataset.Metadata != nil {
		rparams.Dataset.Metadata =
			&remotedatamodel.Metadata{
				CustomMetadata:       params.Dataset.Metadata.CustomMetadata,
				CustomMetadataSchema: params.Dataset.Metadata.CustomMetadataSchema,
				Creator:              params.Dataset.Metadata.Creator,
				Contributor:          params.Dataset.Metadata.Contributor,
				Publisher:            params.Dataset.Metadata.Publisher,
				Owner:                params.Dataset.Metadata.Owner,
				Identifier:           params.Dataset.Metadata.Identifier,
				AlternateIdentifier:  params.Dataset.Metadata.AlternateIdentifier,
				PublicationYear:      params.Dataset.Metadata.PublicationYear,
				ResourceType:         params.Dataset.Metadata.ResourceType,
				ResourceTypeGeneral:  params.Dataset.Metadata.ResourceTypeGeneral,
				Title:                params.Dataset.Metadata.Title,
				RelatedIdentifier:    params.Dataset.Metadata.RelatedIdentifier,
				Rights:               params.Dataset.Metadata.Rights,
				RightsIdentifier:     params.Dataset.Metadata.RightsIdentifier,
				RightsURI:            params.Dataset.Metadata.RightsURI,
				Scope:		      params.Dataset.Metadata.Scope,
				Format:               params.Dataset.Metadata.Format,
				RelatedSoftware:      params.Dataset.Metadata.RelatedSoftware,
				Description:          params.Dataset.Metadata.Description,
				CreationDate:         params.Dataset.Metadata.CreationDate,
			}
	}

	resok, rescr, err := p.getClient(params.HTTPRequest).DataSetManagement.CreateDataset(ctx, &rparams)

	if err != nil {
		//non 2XX status codes are returned in err here!
		l.Info.Printf("Error calling CreateDataset endpoint\n")
		switch err.(type) {
		case *dstypes.CreateDatasetBadRequest: // 400
			v := err.(*dstypes.CreateDatasetBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateDatasetBadRequest().WithPayload(&payload)
		case *dstypes.CreateDatasetUnauthorized: // 401
			v := err.(*dstypes.CreateDatasetUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateDatasetUnauthorized().WithPayload(&payload)
		case *dstypes.CreateDatasetForbidden: // 401
			v := err.(*dstypes.CreateDatasetForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateDatasetForbidden().WithPayload(&payload)
		case *dstypes.CreateDatasetInternalServerError: // 500
			v := err.(*dstypes.CreateDatasetInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateDatasetInternalServerError().WithPayload(&payload)
		case *dstypes.CreateDatasetBadGateway: // 502
			v := err.(*dstypes.CreateDatasetBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateDatasetBadGateway().WithPayload(&payload)
		case *dstypes.CreateDatasetServiceUnavailable: // 503
			v := err.(*dstypes.CreateDatasetServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCreateDatasetServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCreateDatasetServiceUnavailable().WithPayload(payload)
		}
	}


	if resok !=nil {
		payload := models.DatasetItemCreatedResponse{
			InternalID: resok.Payload.InternalID,
			Status:     resok.Payload.Status,
		}
		return data_set_management.NewCreateDatasetCreated().WithPayload(&payload)
	} else {
		payload := models.DatasetItemCreatedResponse{
			InternalID: rescr.Payload.InternalID,
			Status:     rescr.Payload.Status,
		}
		return data_set_management.NewCreateDatasetOK().WithPayload(&payload)
	}
}
func (p *DataSetAPI) DownloadDataset(ctx context.Context, params data_set_management.DownloadDatasetParams) middleware.Responder {

	// need to include params here..
	rparams := dstypes.DownloadDatasetParams{
		JSON: dstypes.DownloadDatasetBody{
			InternalID:  params.JSON.InternalID, //string
			Access:      remotedatamodel.AccessMode(params.JSON.Access),
			Project:     params.JSON.Project, //string*
			PushMethod:  params.JSON.PushMethod,
			Archivetype: remotedatamodel.ArchiveType(params.JSON.Archivetype),
			Path:        params.JSON.Path,
                        Zone:        params.JSON.Zone,
		},
	}
	//https://stackoverflow.com/questions/40316052/in-memory-file-for-testing
	//https://stackoverflow.com/questions/23454940/getting-bytes-buffer-does-not-implement-io-writer-error-message/23454941
	var buffermemory bytes.Buffer
	buffer := ioutil.NopCloser(&buffermemory)

	_, err := p.getClient(params.HTTPRequest).DataSetManagement.DownloadDataset(ctx, &rparams, &buffermemory)

	if err != nil {
		l.Info.Printf("Error calling DownloadDataset endpoint\n")
		switch err.(type) {
		case *dstypes.DownloadDatasetBadRequest: // 400
			v := err.(*dstypes.DownloadDatasetBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDownloadDatasetBadRequest().WithPayload(&payload)
		case *dstypes.DownloadDatasetUnauthorized: // 401
			v := err.(*dstypes.DownloadDatasetUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDownloadDatasetUnauthorized().WithPayload(&payload)
		case *dstypes.DownloadDatasetForbidden: // 403
			v := err.(*dstypes.DownloadDatasetForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDownloadDatasetForbidden().WithPayload(&payload)
		case *dstypes.DownloadDatasetInternalServerError: // 500
			v := err.(*dstypes.DownloadDatasetInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDownloadDatasetInternalServerError().WithPayload(&payload)
		case *dstypes.DownloadDatasetBadGateway: // 502
			v := err.(*dstypes.DownloadDatasetBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDownloadDatasetBadGateway().WithPayload(&payload)
		case *dstypes.DownloadDatasetServiceUnavailable: // 503
			v := err.(*dstypes.DownloadDatasetServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDownloadDatasetServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewDownloadDatasetServiceUnavailable().WithPayload(payload)
		}
	}

	// need to perform a type transformation here...
	//return data_set_management.NewDownloadDatasetCreated().WithPayload(resp.Payload)
	return data_set_management.NewDownloadDatasetOK().WithPayload(buffer)
}

func (p *DataSetAPI) DeleteDataset(ctx context.Context, params data_set_management.DeleteDatasetParams) middleware.Responder {

	// need to include params here..
	rparams := dstypes.DeleteDatasetParams{
		JSON: dstypes.DeleteDatasetBody{
			InternalID: params.JSON.InternalID, //string
			Access:     remotedatamodel.AccessMode(params.JSON.Access),
			Project:    params.JSON.Project,
			Path:       params.JSON.Path,
                        Zone:       params.JSON.Zone,
		},
	}

	res, _, err := p.getClient(params.HTTPRequest).DataSetManagement.
		DeleteDataset(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling DeleteDataset endpoint\n")
		//non 2XX status codes are returned in err here!
		switch err.(type) {
		case *dstypes.DeleteDatasetBadRequest: // 400
			v := err.(*dstypes.DeleteDatasetBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetBadRequest().WithPayload(&payload)
		case *dstypes.DeleteDatasetUnauthorized: // 401
			v := err.(*dstypes.DeleteDatasetUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetUnauthorized().WithPayload(&payload)
		case *dstypes.DeleteDatasetForbidden: // 403
			v := err.(*dstypes.DeleteDatasetForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetForbidden().WithPayload(&payload)
		case *dstypes.DeleteDatasetInternalServerError: // 500
			v := err.(*dstypes.DeleteDatasetInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetInternalServerError().WithPayload(&payload)
		case *dstypes.DeleteDatasetBadGateway: // 502
			v := err.(*dstypes.DeleteDatasetBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetBadGateway().WithPayload(&payload)
		case *dstypes.DeleteDatasetServiceUnavailable: // 503
			v := err.(*dstypes.DeleteDatasetServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewDeleteDatasetServiceUnavailable().WithPayload(payload)
		}
	}
	payload := data_set_management.DeleteDatasetCreatedBody{
		InternalID: res.Payload.InternalID,
		RequestID:  res.Payload.RequestID,
		StagingAPI: res.Payload.StagingAPI,
	}
	// need to perform a type transformation here...
	//return data_set_management.NewDeleteDatasetDeleted().WithPayload(resp.Payload)
	return data_set_management.NewDeleteDatasetCreated().WithPayload(&payload)
}

func (p *DataSetAPI) DeleteDatasetByMetadata(ctx context.Context,
	params data_set_management.DeleteDatasetByMetadataParams) middleware.Responder {

	// need to include params here..
	rparams := dstypes.DeleteDatasetByMetadataParams{
		MetadataQuery: &remotedatamodel.MetadataQuery{
			CustomMetadata:          params.MetadataQuery.CustomMetadata,
			Access:                  remotedatamodel.AccessMode(params.MetadataQuery.Access),
			Project:                 params.MetadataQuery.Project,
			Creator:                 params.MetadataQuery.Creator, //string
			Contributor:             params.MetadataQuery.Contributor,
			Publisher:               params.MetadataQuery.Publisher,
			Owner:                   params.MetadataQuery.Owner,
			Identifier:              params.MetadataQuery.Identifier,
			InternalID:              params.MetadataQuery.InternalID,
			AlternateIdentifier:     params.MetadataQuery.AlternateIdentifier,
			AlternateIdentifierType: params.MetadataQuery.AlternateIdentifierType,
			PublicationYear:         params.MetadataQuery.PublicationYear,
			ResourceType:            params.MetadataQuery.ResourceType,
			ResourceTypeGeneral:     params.MetadataQuery.ResourceTypeGeneral,
			Title:                   params.MetadataQuery.Title,
			RelatedIdentifier:       params.MetadataQuery.RelatedIdentifier,
			Rights:                  params.MetadataQuery.Rights,
			RightsIdentifier:        params.MetadataQuery.RightsIdentifier,
			RightsURI:               params.MetadataQuery.RightsURI,
			EUDATFIO:                params.MetadataQuery.EUDATFIO,
			EUDATFIXEDCONTENT:       params.MetadataQuery.EUDATFIXEDCONTENT,
			EUDATPARENT:             params.MetadataQuery.EUDATPARENT,
			EUDATREPLICA:            params.MetadataQuery.EUDATREPLICA,
			EUDATROR:                params.MetadataQuery.EUDATROR,
			PID:                     params.MetadataQuery.PID,
			Scope:			 params.MetadataQuery.Scope,
			Format:                  params.MetadataQuery.Format,
			RelatedSoftware:         params.MetadataQuery.RelatedSoftware,
			Description:             params.MetadataQuery.Description,
			CreationDate:            params.MetadataQuery.CreationDate,
		},
	}

	_, err := p.getClient(params.HTTPRequest).DataSetManagement.DeleteDatasetByMetadata(ctx, &rparams)

	if err != nil {
		//non 2XX status codes are returned in err here!
		l.Info.Printf("Error calling DeleteDatasetByMetadata endpoint\n")
		switch err.(type) {
		case *dstypes.DeleteDatasetByMetadataBadRequest: // 400
			v := err.(*dstypes.DeleteDatasetByMetadataBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetByMetadataBadRequest().WithPayload(&payload)
		case *dstypes.DeleteDatasetByMetadataUnauthorized: // 401
			v := err.(*dstypes.DeleteDatasetByMetadataUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetByMetadataUnauthorized().WithPayload(&payload)
		case *dstypes.DeleteDatasetByMetadataForbidden: //403
			v := err.(*dstypes.DeleteDatasetByMetadataForbidden)
			payload := data_set_management.DeleteDatasetByMetadataForbiddenBody{
				ErrorString: v.Payload.ErrorString,
			}
			//rgh FIXME add this later.
			////		var items []*data_set_management.DeleteDatasetByMetadataForbiddenBodyPermissionErrorItems0
			////		for i := 0; i < len(v.Payload.PermissionError); i++ {
			////			items = append(items, &data_set_management.DeleteDatasetByMetadataForbiddenBodyPermissionErrorItems0{
			////				InternalID: v.Payload.PermissionError[i].InternalID,
			////			})
			////		}
			////		payload.PermissionError = items
			return data_set_management.NewDeleteDatasetByMetadataForbidden().WithPayload(&payload)
		case *dstypes.DeleteDatasetByMetadataInternalServerError: // 500
			v := err.(*dstypes.DeleteDatasetByMetadataInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetByMetadataInternalServerError().WithPayload(&payload)
		case *dstypes.DeleteDatasetByMetadataBadGateway: // 502
			v := err.(*dstypes.DeleteDatasetByMetadataBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetByMetadataBadGateway().WithPayload(&payload)
		case *dstypes.DeleteDatasetByMetadataServiceUnavailable: // 503
			v := err.(*dstypes.DeleteDatasetByMetadataServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewDeleteDatasetByMetadataServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewDeleteDatasetByMetadataServiceUnavailable().WithPayload(payload)
		}
	}
	return data_set_management.NewDeleteDatasetByMetadataNoContent()
}

func (p *DataSetAPI) QueryDatasets(ctx context.Context, params data_set_management.QueryDatasetsParams) middleware.Responder {
	// need to include params here..
	rparams := dstypes.QueryDatasetsParams{
		MetadataQuery: &remotedatamodel.MetadataQuery{
			CustomMetadata:          params.MetadataQuery.CustomMetadata,
			Access:                  remotedatamodel.AccessMode(params.MetadataQuery.Access),
			Project:                 params.MetadataQuery.Project,
			Creator:                 params.MetadataQuery.Creator, //string
			Contributor:             params.MetadataQuery.Contributor,
			Publisher:               params.MetadataQuery.Publisher,
			Owner:                   params.MetadataQuery.Owner,
			Identifier:              params.MetadataQuery.Identifier,
			InternalID:              params.MetadataQuery.InternalID,
			AlternateIdentifier:     params.MetadataQuery.AlternateIdentifier,
			AlternateIdentifierType: params.MetadataQuery.AlternateIdentifierType,
			PublicationYear:         params.MetadataQuery.PublicationYear,
			ResourceType:            params.MetadataQuery.ResourceType,
			ResourceTypeGeneral:     params.MetadataQuery.ResourceTypeGeneral,
			Title:                   params.MetadataQuery.Title,
			RelatedIdentifier:       params.MetadataQuery.RelatedIdentifier,
			Rights:                  params.MetadataQuery.Rights,
			RightsIdentifier:        params.MetadataQuery.RightsIdentifier,
			RightsURI:               params.MetadataQuery.RightsURI,
			EUDATFIO:                params.MetadataQuery.EUDATFIO,
			EUDATFIXEDCONTENT:       params.MetadataQuery.EUDATFIXEDCONTENT,
			EUDATPARENT:             params.MetadataQuery.EUDATPARENT,
			EUDATREPLICA:            params.MetadataQuery.EUDATREPLICA,
			EUDATROR:                params.MetadataQuery.EUDATROR,
			PID:                     params.MetadataQuery.PID,
                        Scope:                   params.MetadataQuery.Scope,
                        Format:                  params.MetadataQuery.Format,
                        RelatedSoftware:         params.MetadataQuery.RelatedSoftware,
                        Description:             params.MetadataQuery.Description,
                        CreationDate:            params.MetadataQuery.CreationDate,
		},
	}
	res, err := p.getClient(params.HTTPRequest).DataSetManagement.QueryDatasets(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling QueryDatasets endpoint\n")
		switch err.(type) {
		case *dstypes.QueryDatasetsBadRequest: // 400
			v := err.(*dstypes.QueryDatasetsBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewQueryDatasetsBadRequest().WithPayload(&payload)
		case *dstypes.QueryDatasetsUnauthorized: // 401
			v := err.(*dstypes.QueryDatasetsUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewQueryDatasetsUnauthorized().WithPayload(&payload)
		case *dstypes.QueryDatasetsInternalServerError: // 500
			v := err.(*dstypes.QueryDatasetsInternalServerError)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewQueryDatasetsInternalServerError().WithPayload(&payload)
		case *dstypes.QueryDatasetsBadGateway: // 502
			v := err.(*dstypes.QueryDatasetsBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewQueryDatasetsBadGateway().WithPayload(&payload)
		case *dstypes.QueryDatasetsServiceUnavailable: // 503
			v := err.(*dstypes.QueryDatasetsServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewQueryDatasetsServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewQueryDatasetsServiceUnavailable().WithPayload(payload)
		}
	}
	var items []*models.DatasetMetadataQueryResponse
	for i := range res.Payload {
		items = append(items, &models.DatasetMetadataQueryResponse{
			Location: &models.Location{
				InternalID: res.Payload[i].Location.InternalID, //string
				Access:     models.AccessMode(res.Payload[i].Location.Access),
				Project:    res.Payload[i].Location.Project,
				Zone: 	    res.Payload[i].Location.Zone,
			},
			Metadata: &models.DatasetMetadata{
				CustomMetadata:       res.Payload[i].Metadata.CustomMetadata,
				CustomMetadataSchema: res.Payload[i].Metadata.CustomMetadataSchema,
				Creator:              res.Payload[i].Metadata.Creator,
				Contributor:          res.Payload[i].Metadata.Contributor,
				Publisher:            res.Payload[i].Metadata.Publisher,
				Owner:                res.Payload[i].Metadata.Owner,
				Identifier:           res.Payload[i].Metadata.Identifier,
				AlternateIdentifier:  res.Payload[i].Metadata.AlternateIdentifier,
				PublicationYear:      res.Payload[i].Metadata.PublicationYear,
				ResourceType:         res.Payload[i].Metadata.ResourceType,
				ResourceTypeGeneral:  res.Payload[i].Metadata.ResourceTypeGeneral,
				Title:                res.Payload[i].Metadata.Title,
				RelatedIdentifier:    res.Payload[i].Metadata.RelatedIdentifier,
				Rights:               res.Payload[i].Metadata.Rights,
				RightsIdentifier:     res.Payload[i].Metadata.RightsIdentifier,
				RightsURI:            res.Payload[i].Metadata.RightsURI,
                                Scope:                res.Payload[i].Metadata.Scope,
                                Format:               res.Payload[i].Metadata.Format,
                                RelatedSoftware:      res.Payload[i].Metadata.RelatedSoftware,
                                Description:          res.Payload[i].Metadata.Description,
                                CreationDate:         res.Payload[i].Metadata.CreationDate,
			},
			Eudat: &models.Eudat{
				EUDATFIO:          res.Payload[i].Eudat.EUDATFIO,
				EUDATFIXEDCONTENT: res.Payload[i].Eudat.EUDATFIXEDCONTENT,
				EUDATPARENT:       res.Payload[i].Eudat.EUDATPARENT,
				EUDATREPLICA:      res.Payload[i].Eudat.EUDATREPLICA,
				EUDATROR:          res.Payload[i].Eudat.EUDATROR,
				PID:               res.Payload[i].Eudat.PID,
			},
                        Flags: &models.DatasetFlags{
				Encryption:	res.Payload[i].Flags.Encryption,
				Compression:	res.Payload[i].Flags.Compression,
			},
		})
	}
	return data_set_management.NewQueryDatasetsOK().WithPayload(items)
}

func (p *DataSetAPI) CheckPermission(ctx context.Context, params data_set_management.CheckPermissionParams) middleware.Responder {
	// need to include params here..
	rparams := dstypes.CheckPermissionParams{
		Access: dstypes.CheckPermissionBody{
			Access:  remotedatamodel.AccessMode(params.Access.Access),
			Project: params.Access.Project, //string*
		},
	}

	res, err := p.getClient(params.HTTPRequest).DataSetManagement.CheckPermission(ctx, &rparams)

	if err != nil {
		l.Info.Printf("Error calling CheckPermission endpoint\n")
		switch err.(type) {
		case *dstypes.CheckPermissionBadRequest: // 400
			v := err.(*dstypes.CheckPermissionBadRequest)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPermissionBadRequest().WithPayload(&payload)
		case *dstypes.CheckPermissionUnauthorized: // 401
			v := err.(*dstypes.CheckPermissionUnauthorized)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPermissionUnauthorized().WithPayload(&payload)
		case *dstypes.CheckPermissionForbidden: // 403
			v := err.(*dstypes.CheckPermissionForbidden)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPermissionForbidden().WithPayload(&payload)
		case *dstypes.CheckPermissionBadGateway: // 502
			v := err.(*dstypes.CheckPermissionBadGateway)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPermissionBadGateway().WithPayload(&payload)
		case *dstypes.CheckPermissionServiceUnavailable: // 503
			v := err.(*dstypes.CheckPermissionServiceUnavailable)
			payload := models.ErrorResponse{
				ErrorString: v.Payload.ErrorString,
			}
			return data_set_management.NewCheckPermissionServiceUnavailable().WithPayload(&payload)
		default:
			payload := fillErrorResponse(err)
			return data_set_management.NewCheckPermissionServiceUnavailable().WithPayload(payload)
		}
	}

	//return data_set_management.NewCheckPermissionOK().WithPayload(resp.Payload)
	payload := data_set_management.CheckPermissionOKBody{
		Status: res.Payload.Status,
	}
	return data_set_management.NewCheckPermissionOK().WithPayload(&payload)
}
