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

// CreateResourceRequestReader is a Reader for the CreateResourceRequest structure.
type CreateResourceRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResourceRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateResourceRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateResourceRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateResourceRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateResourceRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateResourceRequestCreated creates a CreateResourceRequestCreated with default headers values
func NewCreateResourceRequestCreated() *CreateResourceRequestCreated {
	return &CreateResourceRequestCreated{}
}

/*CreateResourceRequestCreated handles this case with default header values.

resource request created
*/
type CreateResourceRequestCreated struct {
	Payload *models.ApprovalSystemCreatedResponse
}

func (o *CreateResourceRequestCreated) Error() string {
	return fmt.Sprintf("[POST /approval_system/resourceRequest][%d] createResourceRequestCreated  %+v", 201, o.Payload)
}

func (o *CreateResourceRequestCreated) GetPayload() *models.ApprovalSystemCreatedResponse {
	return o.Payload
}

func (o *CreateResourceRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemCreatedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceRequestBadRequest creates a CreateResourceRequestBadRequest with default headers values
func NewCreateResourceRequestBadRequest() *CreateResourceRequestBadRequest {
	return &CreateResourceRequestBadRequest{}
}

/*CreateResourceRequestBadRequest handles this case with default header values.

invalid input, object invalid
*/
type CreateResourceRequestBadRequest struct {
	Payload *models.ApprovalSystemInvalidResponse
}

func (o *CreateResourceRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /approval_system/resourceRequest][%d] createResourceRequestBadRequest  %+v", 400, o.Payload)
}

func (o *CreateResourceRequestBadRequest) GetPayload() *models.ApprovalSystemInvalidResponse {
	return o.Payload
}

func (o *CreateResourceRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemInvalidResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceRequestNotFound creates a CreateResourceRequestNotFound with default headers values
func NewCreateResourceRequestNotFound() *CreateResourceRequestNotFound {
	return &CreateResourceRequestNotFound{}
}

/*CreateResourceRequestNotFound handles this case with default header values.

Required stuff for resource request creation hasn't been met.
*/
type CreateResourceRequestNotFound struct {
	Payload *models.ApprovalSystemMissingResponse
}

func (o *CreateResourceRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /approval_system/resourceRequest][%d] createResourceRequestNotFound  %+v", 404, o.Payload)
}

func (o *CreateResourceRequestNotFound) GetPayload() *models.ApprovalSystemMissingResponse {
	return o.Payload
}

func (o *CreateResourceRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemMissingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceRequestInternalServerError creates a CreateResourceRequestInternalServerError with default headers values
func NewCreateResourceRequestInternalServerError() *CreateResourceRequestInternalServerError {
	return &CreateResourceRequestInternalServerError{}
}

/*CreateResourceRequestInternalServerError handles this case with default header values.

unexpected error
*/
type CreateResourceRequestInternalServerError struct {
	Payload *models.ApprovalSystemErrorResponse
}

func (o *CreateResourceRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /approval_system/resourceRequest][%d] createResourceRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateResourceRequestInternalServerError) GetPayload() *models.ApprovalSystemErrorResponse {
	return o.Payload
}

func (o *CreateResourceRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApprovalSystemErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}