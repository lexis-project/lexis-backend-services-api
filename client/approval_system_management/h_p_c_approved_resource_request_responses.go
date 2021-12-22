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

// HPCApprovedResourceRequestReader is a Reader for the HPCApprovedResourceRequest structure.
type HPCApprovedResourceRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HPCApprovedResourceRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHPCApprovedResourceRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewHPCApprovedResourceRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHPCApprovedResourceRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHPCApprovedResourceRequestOK creates a HPCApprovedResourceRequestOK with default headers values
func NewHPCApprovedResourceRequestOK() *HPCApprovedResourceRequestOK {
	return &HPCApprovedResourceRequestOK{}
}

/*HPCApprovedResourceRequestOK handles this case with default header values.

HPC approved resource request
*/
type HPCApprovedResourceRequestOK struct {
	Payload *models.ApprovalSystemApprovedResourceRequest
}

func (o *HPCApprovedResourceRequestOK) Error() string {
	return fmt.Sprintf("[GET /approval_system/approvedResourceRequest/{HPCResourceID}][%d] hPCApprovedResourceRequestOK  %+v", 200, o.Payload)
}

func (o *HPCApprovedResourceRequestOK) GetPayload() *models.ApprovalSystemApprovedResourceRequest {
	return o.Payload
}

func (o *HPCApprovedResourceRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemApprovedResourceRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHPCApprovedResourceRequestNotFound creates a HPCApprovedResourceRequestNotFound with default headers values
func NewHPCApprovedResourceRequestNotFound() *HPCApprovedResourceRequestNotFound {
	return &HPCApprovedResourceRequestNotFound{}
}

/*HPCApprovedResourceRequestNotFound handles this case with default header values.

Resource request not found.
*/
type HPCApprovedResourceRequestNotFound struct {
	Payload *models.ApprovalSystemMissingResponse
}

func (o *HPCApprovedResourceRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /approval_system/approvedResourceRequest/{HPCResourceID}][%d] hPCApprovedResourceRequestNotFound  %+v", 404, o.Payload)
}

func (o *HPCApprovedResourceRequestNotFound) GetPayload() *models.ApprovalSystemMissingResponse {
	return o.Payload
}

func (o *HPCApprovedResourceRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemMissingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHPCApprovedResourceRequestInternalServerError creates a HPCApprovedResourceRequestInternalServerError with default headers values
func NewHPCApprovedResourceRequestInternalServerError() *HPCApprovedResourceRequestInternalServerError {
	return &HPCApprovedResourceRequestInternalServerError{}
}

/*HPCApprovedResourceRequestInternalServerError handles this case with default header values.

unexpected error
*/
type HPCApprovedResourceRequestInternalServerError struct {
	Payload *models.ApprovalSystemErrorResponse
}

func (o *HPCApprovedResourceRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /approval_system/approvedResourceRequest/{HPCResourceID}][%d] hPCApprovedResourceRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *HPCApprovedResourceRequestInternalServerError) GetPayload() *models.ApprovalSystemErrorResponse {
	return o.Payload
}

func (o *HPCApprovedResourceRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
