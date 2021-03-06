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

// ListProjectHPCApprovedResourceRequestReader is a Reader for the ListProjectHPCApprovedResourceRequest structure.
type ListProjectHPCApprovedResourceRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectHPCApprovedResourceRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectHPCApprovedResourceRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListProjectHPCApprovedResourceRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProjectHPCApprovedResourceRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProjectHPCApprovedResourceRequestOK creates a ListProjectHPCApprovedResourceRequestOK with default headers values
func NewListProjectHPCApprovedResourceRequestOK() *ListProjectHPCApprovedResourceRequestOK {
	return &ListProjectHPCApprovedResourceRequestOK{}
}

/*ListProjectHPCApprovedResourceRequestOK handles this case with default header values.

list of HPC approved resource requests returned for particular LEXIS project
*/
type ListProjectHPCApprovedResourceRequestOK struct {
	Payload []*models.ApprovalSystemApprovedResourceRequest
}

func (o *ListProjectHPCApprovedResourceRequestOK) Error() string {
	return fmt.Sprintf("[GET /approval_system/projectApprovedResourceRequest/{AssociatedLEXISProject}][%d] listProjectHPCApprovedResourceRequestOK  %+v", 200, o.Payload)
}

func (o *ListProjectHPCApprovedResourceRequestOK) GetPayload() []*models.ApprovalSystemApprovedResourceRequest {
	return o.Payload
}

func (o *ListProjectHPCApprovedResourceRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectHPCApprovedResourceRequestNotFound creates a ListProjectHPCApprovedResourceRequestNotFound with default headers values
func NewListProjectHPCApprovedResourceRequestNotFound() *ListProjectHPCApprovedResourceRequestNotFound {
	return &ListProjectHPCApprovedResourceRequestNotFound{}
}

/*ListProjectHPCApprovedResourceRequestNotFound handles this case with default header values.

The AssociatedLEXISProject ID provided does not exist.
*/
type ListProjectHPCApprovedResourceRequestNotFound struct {
	Payload *models.ApprovalSystemMissingResponse
}

func (o *ListProjectHPCApprovedResourceRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /approval_system/projectApprovedResourceRequest/{AssociatedLEXISProject}][%d] listProjectHPCApprovedResourceRequestNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectHPCApprovedResourceRequestNotFound) GetPayload() *models.ApprovalSystemMissingResponse {
	return o.Payload
}

func (o *ListProjectHPCApprovedResourceRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemMissingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectHPCApprovedResourceRequestInternalServerError creates a ListProjectHPCApprovedResourceRequestInternalServerError with default headers values
func NewListProjectHPCApprovedResourceRequestInternalServerError() *ListProjectHPCApprovedResourceRequestInternalServerError {
	return &ListProjectHPCApprovedResourceRequestInternalServerError{}
}

/*ListProjectHPCApprovedResourceRequestInternalServerError handles this case with default header values.

unexpected error
*/
type ListProjectHPCApprovedResourceRequestInternalServerError struct {
	Payload *models.ApprovalSystemErrorResponse
}

func (o *ListProjectHPCApprovedResourceRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /approval_system/projectApprovedResourceRequest/{AssociatedLEXISProject}][%d] listProjectHPCApprovedResourceRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectHPCApprovedResourceRequestInternalServerError) GetPayload() *models.ApprovalSystemErrorResponse {
	return o.Payload
}

func (o *ListProjectHPCApprovedResourceRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
