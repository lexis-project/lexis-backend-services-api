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

// ListProjectHPCResourceRequestReader is a Reader for the ListProjectHPCResourceRequest structure.
type ListProjectHPCResourceRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectHPCResourceRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectHPCResourceRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListProjectHPCResourceRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProjectHPCResourceRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProjectHPCResourceRequestOK creates a ListProjectHPCResourceRequestOK with default headers values
func NewListProjectHPCResourceRequestOK() *ListProjectHPCResourceRequestOK {
	return &ListProjectHPCResourceRequestOK{}
}

/*ListProjectHPCResourceRequestOK handles this case with default header values.

list of HPC resource requests returned for particular LEXIS project
*/
type ListProjectHPCResourceRequestOK struct {
	Payload []*models.ApprovalSystemResourceRequest
}

func (o *ListProjectHPCResourceRequestOK) Error() string {
	return fmt.Sprintf("[GET /approval_system/projectResourceRequest/{AssociatedLEXISProject}][%d] listProjectHPCResourceRequestOK  %+v", 200, o.Payload)
}

func (o *ListProjectHPCResourceRequestOK) GetPayload() []*models.ApprovalSystemResourceRequest {
	return o.Payload
}

func (o *ListProjectHPCResourceRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectHPCResourceRequestNotFound creates a ListProjectHPCResourceRequestNotFound with default headers values
func NewListProjectHPCResourceRequestNotFound() *ListProjectHPCResourceRequestNotFound {
	return &ListProjectHPCResourceRequestNotFound{}
}

/*ListProjectHPCResourceRequestNotFound handles this case with default header values.

The AssociatedLEXISProject ID provided does not exist.
*/
type ListProjectHPCResourceRequestNotFound struct {
	Payload *models.ApprovalSystemMissingResponse
}

func (o *ListProjectHPCResourceRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /approval_system/projectResourceRequest/{AssociatedLEXISProject}][%d] listProjectHPCResourceRequestNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectHPCResourceRequestNotFound) GetPayload() *models.ApprovalSystemMissingResponse {
	return o.Payload
}

func (o *ListProjectHPCResourceRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemMissingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectHPCResourceRequestInternalServerError creates a ListProjectHPCResourceRequestInternalServerError with default headers values
func NewListProjectHPCResourceRequestInternalServerError() *ListProjectHPCResourceRequestInternalServerError {
	return &ListProjectHPCResourceRequestInternalServerError{}
}

/*ListProjectHPCResourceRequestInternalServerError handles this case with default header values.

unexpected error
*/
type ListProjectHPCResourceRequestInternalServerError struct {
	Payload *models.ApprovalSystemErrorResponse
}

func (o *ListProjectHPCResourceRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /approval_system/projectResourceRequest/{AssociatedLEXISProject}][%d] listProjectHPCResourceRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectHPCResourceRequestInternalServerError) GetPayload() *models.ApprovalSystemErrorResponse {
	return o.Payload
}

func (o *ListProjectHPCResourceRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}