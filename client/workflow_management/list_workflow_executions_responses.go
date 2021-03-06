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

// ListWorkflowExecutionsReader is a Reader for the ListWorkflowExecutions structure.
type ListWorkflowExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListWorkflowExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListWorkflowExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListWorkflowExecutionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListWorkflowExecutionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListWorkflowExecutionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListWorkflowExecutionsOK creates a ListWorkflowExecutionsOK with default headers values
func NewListWorkflowExecutionsOK() *ListWorkflowExecutionsOK {
	return &ListWorkflowExecutionsOK{}
}

/*ListWorkflowExecutionsOK handles this case with default header values.

List of available LEXIS Workflow Executions
*/
type ListWorkflowExecutionsOK struct {
	Payload []*models.WorkflowExecution
}

func (o *ListWorkflowExecutionsOK) Error() string {
	return fmt.Sprintf("[GET /workflow/{workflowId}/execution][%d] listWorkflowExecutionsOK  %+v", 200, o.Payload)
}

func (o *ListWorkflowExecutionsOK) GetPayload() []*models.WorkflowExecution {
	return o.Payload
}

func (o *ListWorkflowExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListWorkflowExecutionsUnauthorized creates a ListWorkflowExecutionsUnauthorized with default headers values
func NewListWorkflowExecutionsUnauthorized() *ListWorkflowExecutionsUnauthorized {
	return &ListWorkflowExecutionsUnauthorized{}
}

/*ListWorkflowExecutionsUnauthorized handles this case with default header values.

Authorization information is missing or invalid.
*/
type ListWorkflowExecutionsUnauthorized struct {
	Payload *models.AuthorizationResponse
}

func (o *ListWorkflowExecutionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workflow/{workflowId}/execution][%d] listWorkflowExecutionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListWorkflowExecutionsUnauthorized) GetPayload() *models.AuthorizationResponse {
	return o.Payload
}

func (o *ListWorkflowExecutionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListWorkflowExecutionsNotFound creates a ListWorkflowExecutionsNotFound with default headers values
func NewListWorkflowExecutionsNotFound() *ListWorkflowExecutionsNotFound {
	return &ListWorkflowExecutionsNotFound{}
}

/*ListWorkflowExecutionsNotFound handles this case with default header values.

List Worklow Executions reuired dependencies not found
*/
type ListWorkflowExecutionsNotFound struct {
	Payload *models.MissingResponse
}

func (o *ListWorkflowExecutionsNotFound) Error() string {
	return fmt.Sprintf("[GET /workflow/{workflowId}/execution][%d] listWorkflowExecutionsNotFound  %+v", 404, o.Payload)
}

func (o *ListWorkflowExecutionsNotFound) GetPayload() *models.MissingResponse {
	return o.Payload
}

func (o *ListWorkflowExecutionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MissingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListWorkflowExecutionsInternalServerError creates a ListWorkflowExecutionsInternalServerError with default headers values
func NewListWorkflowExecutionsInternalServerError() *ListWorkflowExecutionsInternalServerError {
	return &ListWorkflowExecutionsInternalServerError{}
}

/*ListWorkflowExecutionsInternalServerError handles this case with default header values.

Unexpected error.
*/
type ListWorkflowExecutionsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ListWorkflowExecutionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workflow/{workflowId}/execution][%d] listWorkflowExecutionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListWorkflowExecutionsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListWorkflowExecutionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
