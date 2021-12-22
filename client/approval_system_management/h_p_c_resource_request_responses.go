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

// HPCResourceRequestReader is a Reader for the HPCResourceRequest structure.
type HPCResourceRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HPCResourceRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHPCResourceRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewHPCResourceRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHPCResourceRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHPCResourceRequestOK creates a HPCResourceRequestOK with default headers values
func NewHPCResourceRequestOK() *HPCResourceRequestOK {
	return &HPCResourceRequestOK{}
}

/*HPCResourceRequestOK handles this case with default header values.

HPC resource request
*/
type HPCResourceRequestOK struct {
	Payload *models.ApprovalSystemResourceRequest
}

func (o *HPCResourceRequestOK) Error() string {
	return fmt.Sprintf("[GET /approval_system/resourceRequest/{HPCResourceID}][%d] hPCResourceRequestOK  %+v", 200, o.Payload)
}

func (o *HPCResourceRequestOK) GetPayload() *models.ApprovalSystemResourceRequest {
	return o.Payload
}

func (o *HPCResourceRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemResourceRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHPCResourceRequestNotFound creates a HPCResourceRequestNotFound with default headers values
func NewHPCResourceRequestNotFound() *HPCResourceRequestNotFound {
	return &HPCResourceRequestNotFound{}
}

/*HPCResourceRequestNotFound handles this case with default header values.

Resource request not found.
*/
type HPCResourceRequestNotFound struct {
	Payload *models.ApprovalSystemMissingResponse
}

func (o *HPCResourceRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /approval_system/resourceRequest/{HPCResourceID}][%d] hPCResourceRequestNotFound  %+v", 404, o.Payload)
}

func (o *HPCResourceRequestNotFound) GetPayload() *models.ApprovalSystemMissingResponse {
	return o.Payload
}

func (o *HPCResourceRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemMissingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHPCResourceRequestInternalServerError creates a HPCResourceRequestInternalServerError with default headers values
func NewHPCResourceRequestInternalServerError() *HPCResourceRequestInternalServerError {
	return &HPCResourceRequestInternalServerError{}
}

/*HPCResourceRequestInternalServerError handles this case with default header values.

unexpected error
*/
type HPCResourceRequestInternalServerError struct {
	Payload *models.ApprovalSystemErrorResponse
}

func (o *HPCResourceRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /approval_system/resourceRequest/{HPCResourceID}][%d] hPCResourceRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *HPCResourceRequestInternalServerError) GetPayload() *models.ApprovalSystemErrorResponse {
	return o.Payload
}

func (o *HPCResourceRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}