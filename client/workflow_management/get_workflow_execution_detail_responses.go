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

// GetWorkflowExecutionDetailReader is a Reader for the GetWorkflowExecutionDetail structure.
type GetWorkflowExecutionDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowExecutionDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowExecutionDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkflowExecutionDetailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowExecutionDetailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowExecutionDetailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowExecutionDetailOK creates a GetWorkflowExecutionDetailOK with default headers values
func NewGetWorkflowExecutionDetailOK() *GetWorkflowExecutionDetailOK {
	return &GetWorkflowExecutionDetailOK{}
}

/*GetWorkflowExecutionDetailOK handles this case with default header values.

Workflow Execution Detail
*/
type GetWorkflowExecutionDetailOK struct {
	Payload *models.WorkflowExecutionDetail
}

func (o *GetWorkflowExecutionDetailOK) Error() string {
	return fmt.Sprintf("[GET /workflow/{workflowId}/execution/{workflowExecutionId}][%d] getWorkflowExecutionDetailOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowExecutionDetailOK) GetPayload() *models.WorkflowExecutionDetail {
	return o.Payload
}

func (o *GetWorkflowExecutionDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkflowExecutionDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowExecutionDetailUnauthorized creates a GetWorkflowExecutionDetailUnauthorized with default headers values
func NewGetWorkflowExecutionDetailUnauthorized() *GetWorkflowExecutionDetailUnauthorized {
	return &GetWorkflowExecutionDetailUnauthorized{}
}

/*GetWorkflowExecutionDetailUnauthorized handles this case with default header values.

Unauthorized
*/
type GetWorkflowExecutionDetailUnauthorized struct {
	Payload *models.AuthorizationResponse
}

func (o *GetWorkflowExecutionDetailUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflow/{workflowId}/execution/{workflowExecutionId}][%d] getWorkflowExecutionDetailUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkflowExecutionDetailUnauthorized) GetPayload() *models.AuthorizationResponse {
	return o.Payload
}

func (o *GetWorkflowExecutionDetailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowExecutionDetailNotFound creates a GetWorkflowExecutionDetailNotFound with default headers values
func NewGetWorkflowExecutionDetailNotFound() *GetWorkflowExecutionDetailNotFound {
	return &GetWorkflowExecutionDetailNotFound{}
}

/*GetWorkflowExecutionDetailNotFound handles this case with default header values.

The application provided does not exist.
*/
type GetWorkflowExecutionDetailNotFound struct {
	Payload *models.MissingResponse
}

func (o *GetWorkflowExecutionDetailNotFound) Error() string {
	return fmt.Sprintf("[GET /workflow/{workflowId}/execution/{workflowExecutionId}][%d] getWorkflowExecutionDetailNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkflowExecutionDetailNotFound) GetPayload() *models.MissingResponse {
	return o.Payload
}

func (o *GetWorkflowExecutionDetailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MissingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowExecutionDetailInternalServerError creates a GetWorkflowExecutionDetailInternalServerError with default headers values
func NewGetWorkflowExecutionDetailInternalServerError() *GetWorkflowExecutionDetailInternalServerError {
	return &GetWorkflowExecutionDetailInternalServerError{}
}

/*GetWorkflowExecutionDetailInternalServerError handles this case with default header values.

unexpected error
*/
type GetWorkflowExecutionDetailInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetWorkflowExecutionDetailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workflow/{workflowId}/execution/{workflowExecutionId}][%d] getWorkflowExecutionDetailInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkflowExecutionDetailInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWorkflowExecutionDetailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}