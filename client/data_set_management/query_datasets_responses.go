// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// QueryDatasetsReader is a Reader for the QueryDatasets structure.
type QueryDatasetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDatasetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDatasetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryDatasetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryDatasetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryDatasetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewQueryDatasetsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewQueryDatasetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryDatasetsOK creates a QueryDatasetsOK with default headers values
func NewQueryDatasetsOK() *QueryDatasetsOK {
	return &QueryDatasetsOK{}
}

/*QueryDatasetsOK handles this case with default header values.

Metadata for matching datasets returned
*/
type QueryDatasetsOK struct {
	Payload []*models.DatasetMetadataQueryResponse
}

func (o *QueryDatasetsOK) Error() string {
	return fmt.Sprintf("[POST /dataset/search/metadata][%d] queryDatasetsOK  %+v", 200, o.Payload)
}

func (o *QueryDatasetsOK) GetPayload() []*models.DatasetMetadataQueryResponse {
	return o.Payload
}

func (o *QueryDatasetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatasetsBadRequest creates a QueryDatasetsBadRequest with default headers values
func NewQueryDatasetsBadRequest() *QueryDatasetsBadRequest {
	return &QueryDatasetsBadRequest{}
}

/*QueryDatasetsBadRequest handles this case with default header values.

Malformed Request
*/
type QueryDatasetsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *QueryDatasetsBadRequest) Error() string {
	return fmt.Sprintf("[POST /dataset/search/metadata][%d] queryDatasetsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryDatasetsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QueryDatasetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatasetsUnauthorized creates a QueryDatasetsUnauthorized with default headers values
func NewQueryDatasetsUnauthorized() *QueryDatasetsUnauthorized {
	return &QueryDatasetsUnauthorized{}
}

/*QueryDatasetsUnauthorized handles this case with default header values.

Authorization failed
*/
type QueryDatasetsUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *QueryDatasetsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dataset/search/metadata][%d] queryDatasetsUnauthorized  %+v", 401, o.Payload)
}

func (o *QueryDatasetsUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QueryDatasetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatasetsInternalServerError creates a QueryDatasetsInternalServerError with default headers values
func NewQueryDatasetsInternalServerError() *QueryDatasetsInternalServerError {
	return &QueryDatasetsInternalServerError{}
}

/*QueryDatasetsInternalServerError handles this case with default header values.

Internal error processing request
*/
type QueryDatasetsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *QueryDatasetsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dataset/search/metadata][%d] queryDatasetsInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryDatasetsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QueryDatasetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatasetsBadGateway creates a QueryDatasetsBadGateway with default headers values
func NewQueryDatasetsBadGateway() *QueryDatasetsBadGateway {
	return &QueryDatasetsBadGateway{}
}

/*QueryDatasetsBadGateway handles this case with default header values.

Bad Gateway
*/
type QueryDatasetsBadGateway struct {
	Payload *models.ErrorResponse
}

func (o *QueryDatasetsBadGateway) Error() string {
	return fmt.Sprintf("[POST /dataset/search/metadata][%d] queryDatasetsBadGateway  %+v", 502, o.Payload)
}

func (o *QueryDatasetsBadGateway) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QueryDatasetsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDatasetsServiceUnavailable creates a QueryDatasetsServiceUnavailable with default headers values
func NewQueryDatasetsServiceUnavailable() *QueryDatasetsServiceUnavailable {
	return &QueryDatasetsServiceUnavailable{}
}

/*QueryDatasetsServiceUnavailable handles this case with default header values.

unexpected error connecting to further backends
*/
type QueryDatasetsServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *QueryDatasetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dataset/search/metadata][%d] queryDatasetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *QueryDatasetsServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QueryDatasetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
