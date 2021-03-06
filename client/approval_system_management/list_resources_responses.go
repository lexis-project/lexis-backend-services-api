// Code generated by go-swagger; DO NOT EDIT.

package approval_system_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// ListResourcesReader is a Reader for the ListResources structure.
type ListResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListResourcesOK creates a ListResourcesOK with default headers values
func NewListResourcesOK() *ListResourcesOK {
	return &ListResourcesOK{}
}

/*ListResourcesOK handles this case with default header values.

list of HPC resources returned from from approval system
*/
type ListResourcesOK struct {
	Payload []*models.ApprovalSystemResource
}

func (o *ListResourcesOK) Error() string {
	return fmt.Sprintf("[GET /approval_system/resource][%d] listResourcesOK  %+v", 200, o.Payload)
}

func (o *ListResourcesOK) GetPayload() []*models.ApprovalSystemResource {
	return o.Payload
}

func (o *ListResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourcesNotFound creates a ListResourcesNotFound with default headers values
func NewListResourcesNotFound() *ListResourcesNotFound {
	return &ListResourcesNotFound{}
}

/*ListResourcesNotFound handles this case with default header values.

List of available HPC resources not found
*/
type ListResourcesNotFound struct {
	Payload *models.ApprovalSystemMissingResponse
}

func (o *ListResourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /approval_system/resource][%d] listResourcesNotFound  %+v", 404, o.Payload)
}

func (o *ListResourcesNotFound) GetPayload() *models.ApprovalSystemMissingResponse {
	return o.Payload
}

func (o *ListResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemMissingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourcesInternalServerError creates a ListResourcesInternalServerError with default headers values
func NewListResourcesInternalServerError() *ListResourcesInternalServerError {
	return &ListResourcesInternalServerError{}
}

/*ListResourcesInternalServerError handles this case with default header values.

unexpected error
*/
type ListResourcesInternalServerError struct {
	Payload *models.ApprovalSystemErrorResponse
}

func (o *ListResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /approval_system/resource][%d] listResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListResourcesInternalServerError) GetPayload() *models.ApprovalSystemErrorResponse {
	return o.Payload
}

func (o *ListResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
