// Code generated by go-swagger; DO NOT EDIT.

package workflow_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// GetWorkflowsReader is a Reader for the GetWorkflows structure.
type GetWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkflowsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowsOK creates a GetWorkflowsOK with default headers values
func NewGetWorkflowsOK() *GetWorkflowsOK {
	return &GetWorkflowsOK{}
}

/*GetWorkflowsOK handles this case with default header values.

List of available LEXIS Workflows
*/
type GetWorkflowsOK struct {
	Payload []*models.Workflow
}

func (o *GetWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /workflow][%d] getWorkflowsOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowsOK) GetPayload() []*models.Workflow {
	return o.Payload
}

func (o *GetWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowsUnauthorized creates a GetWorkflowsUnauthorized with default headers values
func NewGetWorkflowsUnauthorized() *GetWorkflowsUnauthorized {
	return &GetWorkflowsUnauthorized{}
}

/*GetWorkflowsUnauthorized handles this case with default header values.

Authorization information is missing or invalid.
*/
type GetWorkflowsUnauthorized struct {
	Payload *models.AuthorizationResponse
}

func (o *GetWorkflowsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflow][%d] getWorkflowsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkflowsUnauthorized) GetPayload() *models.AuthorizationResponse {
	return o.Payload
}

func (o *GetWorkflowsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowsInternalServerError creates a GetWorkflowsInternalServerError with default headers values
func NewGetWorkflowsInternalServerError() *GetWorkflowsInternalServerError {
	return &GetWorkflowsInternalServerError{}
}

/*GetWorkflowsInternalServerError handles this case with default header values.

Unexpected error.
*/
type GetWorkflowsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetWorkflowsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workflow][%d] getWorkflowsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkflowsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWorkflowsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
