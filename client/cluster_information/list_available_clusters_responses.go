// Code generated by go-swagger; DO NOT EDIT.

package cluster_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// ListAvailableClustersReader is a Reader for the ListAvailableClusters structure.
type ListAvailableClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAvailableClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAvailableClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAvailableClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewListAvailableClustersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListAvailableClustersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAvailableClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAvailableClustersOK creates a ListAvailableClustersOK with default headers values
func NewListAvailableClustersOK() *ListAvailableClustersOK {
	return &ListAvailableClustersOK{}
}

/*ListAvailableClustersOK handles this case with default header values.

list of clusters retruend
*/
type ListAvailableClustersOK struct {
	Payload []*models.HeappeCluster
}

func (o *ListAvailableClustersOK) Error() string {
	return fmt.Sprintf("[GET /heappe/ClusterInformation/ListAvailableClusters][%d] listAvailableClustersOK  %+v", 200, o.Payload)
}

func (o *ListAvailableClustersOK) GetPayload() []*models.HeappeCluster {
	return o.Payload
}

func (o *ListAvailableClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAvailableClustersBadRequest creates a ListAvailableClustersBadRequest with default headers values
func NewListAvailableClustersBadRequest() *ListAvailableClustersBadRequest {
	return &ListAvailableClustersBadRequest{}
}

/*ListAvailableClustersBadRequest handles this case with default header values.

Bad Request
*/
type ListAvailableClustersBadRequest struct {
	Payload *models.HeappeBadRequest
}

func (o *ListAvailableClustersBadRequest) Error() string {
	return fmt.Sprintf("[GET /heappe/ClusterInformation/ListAvailableClusters][%d] listAvailableClustersBadRequest  %+v", 400, o.Payload)
}

func (o *ListAvailableClustersBadRequest) GetPayload() *models.HeappeBadRequest {
	return o.Payload
}

func (o *ListAvailableClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HeappeBadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAvailableClustersRequestEntityTooLarge creates a ListAvailableClustersRequestEntityTooLarge with default headers values
func NewListAvailableClustersRequestEntityTooLarge() *ListAvailableClustersRequestEntityTooLarge {
	return &ListAvailableClustersRequestEntityTooLarge{}
}

/*ListAvailableClustersRequestEntityTooLarge handles this case with default header values.

Client Error
*/
type ListAvailableClustersRequestEntityTooLarge struct {
	Payload *models.HeappeError
}

func (o *ListAvailableClustersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /heappe/ClusterInformation/ListAvailableClusters][%d] listAvailableClustersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListAvailableClustersRequestEntityTooLarge) GetPayload() *models.HeappeError {
	return o.Payload
}

func (o *ListAvailableClustersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HeappeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAvailableClustersTooManyRequests creates a ListAvailableClustersTooManyRequests with default headers values
func NewListAvailableClustersTooManyRequests() *ListAvailableClustersTooManyRequests {
	return &ListAvailableClustersTooManyRequests{}
}

/*ListAvailableClustersTooManyRequests handles this case with default header values.

Client Error
*/
type ListAvailableClustersTooManyRequests struct {
	Payload *models.HeappeError
}

func (o *ListAvailableClustersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /heappe/ClusterInformation/ListAvailableClusters][%d] listAvailableClustersTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListAvailableClustersTooManyRequests) GetPayload() *models.HeappeError {
	return o.Payload
}

func (o *ListAvailableClustersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HeappeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAvailableClustersInternalServerError creates a ListAvailableClustersInternalServerError with default headers values
func NewListAvailableClustersInternalServerError() *ListAvailableClustersInternalServerError {
	return &ListAvailableClustersInternalServerError{}
}

/*ListAvailableClustersInternalServerError handles this case with default header values.

Server Error
*/
type ListAvailableClustersInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ListAvailableClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /heappe/ClusterInformation/ListAvailableClusters][%d] listAvailableClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAvailableClustersInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAvailableClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
