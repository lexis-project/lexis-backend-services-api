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

// GetCommandTemplateReader is a Reader for the GetCommandTemplate structure.
type GetCommandTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCommandTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCommandTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCommandTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetCommandTemplateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCommandTemplateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCommandTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCommandTemplateOK creates a GetCommandTemplateOK with default headers values
func NewGetCommandTemplateOK() *GetCommandTemplateOK {
	return &GetCommandTemplateOK{}
}

/*GetCommandTemplateOK handles this case with default header values.

Success
*/
type GetCommandTemplateOK struct {
	Payload []string
}

func (o *GetCommandTemplateOK) Error() string {
	return fmt.Sprintf("[POST /heappe/ClusterInformation/GetCommandTemplateParametersName][%d] getCommandTemplateOK  %+v", 200, o.Payload)
}

func (o *GetCommandTemplateOK) GetPayload() []string {
	return o.Payload
}

func (o *GetCommandTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommandTemplateBadRequest creates a GetCommandTemplateBadRequest with default headers values
func NewGetCommandTemplateBadRequest() *GetCommandTemplateBadRequest {
	return &GetCommandTemplateBadRequest{}
}

/*GetCommandTemplateBadRequest handles this case with default header values.

Bad Request
*/
type GetCommandTemplateBadRequest struct {
	Payload *models.HeappeBadRequest
}

func (o *GetCommandTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /heappe/ClusterInformation/GetCommandTemplateParametersName][%d] getCommandTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *GetCommandTemplateBadRequest) GetPayload() *models.HeappeBadRequest {
	return o.Payload
}

func (o *GetCommandTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HeappeBadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommandTemplateRequestEntityTooLarge creates a GetCommandTemplateRequestEntityTooLarge with default headers values
func NewGetCommandTemplateRequestEntityTooLarge() *GetCommandTemplateRequestEntityTooLarge {
	return &GetCommandTemplateRequestEntityTooLarge{}
}

/*GetCommandTemplateRequestEntityTooLarge handles this case with default header values.

Client Error
*/
type GetCommandTemplateRequestEntityTooLarge struct {
	Payload *models.HeappeError
}

func (o *GetCommandTemplateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /heappe/ClusterInformation/GetCommandTemplateParametersName][%d] getCommandTemplateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCommandTemplateRequestEntityTooLarge) GetPayload() *models.HeappeError {
	return o.Payload
}

func (o *GetCommandTemplateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HeappeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommandTemplateTooManyRequests creates a GetCommandTemplateTooManyRequests with default headers values
func NewGetCommandTemplateTooManyRequests() *GetCommandTemplateTooManyRequests {
	return &GetCommandTemplateTooManyRequests{}
}

/*GetCommandTemplateTooManyRequests handles this case with default header values.

Client Error
*/
type GetCommandTemplateTooManyRequests struct {
	Payload *models.HeappeError
}

func (o *GetCommandTemplateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /heappe/ClusterInformation/GetCommandTemplateParametersName][%d] getCommandTemplateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCommandTemplateTooManyRequests) GetPayload() *models.HeappeError {
	return o.Payload
}

func (o *GetCommandTemplateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HeappeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommandTemplateInternalServerError creates a GetCommandTemplateInternalServerError with default headers values
func NewGetCommandTemplateInternalServerError() *GetCommandTemplateInternalServerError {
	return &GetCommandTemplateInternalServerError{}
}

/*GetCommandTemplateInternalServerError handles this case with default header values.

Server Error
*/
type GetCommandTemplateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetCommandTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /heappe/ClusterInformation/GetCommandTemplateParametersName][%d] getCommandTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCommandTemplateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCommandTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}