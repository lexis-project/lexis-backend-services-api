// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// CheckDuplicationStatusReader is a Reader for the CheckDuplicationStatus structure.
type CheckDuplicationStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckDuplicationStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckDuplicationStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckDuplicationStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckDuplicationStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckDuplicationStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewCheckDuplicationStatusRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckDuplicationStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckDuplicationStatusOK creates a CheckDuplicationStatusOK with default headers values
func NewCheckDuplicationStatusOK() *CheckDuplicationStatusOK {
	return &CheckDuplicationStatusOK{}
}

/*CheckDuplicationStatusOK handles this case with default header values.

This means that the status has been returned to the user in the response body.
*/
type CheckDuplicationStatusOK struct {
	Payload *CheckDuplicationStatusOKBody
}

func (o *CheckDuplicationStatusOK) Error() string {
	return fmt.Sprintf("[GET /dataset/duplicate/{request_id}][%d] checkDuplicationStatusOK  %+v", 200, o.Payload)
}

func (o *CheckDuplicationStatusOK) GetPayload() *CheckDuplicationStatusOKBody {
	return o.Payload
}

func (o *CheckDuplicationStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckDuplicationStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDuplicationStatusBadRequest creates a CheckDuplicationStatusBadRequest with default headers values
func NewCheckDuplicationStatusBadRequest() *CheckDuplicationStatusBadRequest {
	return &CheckDuplicationStatusBadRequest{}
}

/*CheckDuplicationStatusBadRequest handles this case with default header values.

This means that the request ID given by the user is incorrect.
*/
type CheckDuplicationStatusBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CheckDuplicationStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /dataset/duplicate/{request_id}][%d] checkDuplicationStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckDuplicationStatusBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDuplicationStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDuplicationStatusUnauthorized creates a CheckDuplicationStatusUnauthorized with default headers values
func NewCheckDuplicationStatusUnauthorized() *CheckDuplicationStatusUnauthorized {
	return &CheckDuplicationStatusUnauthorized{}
}

/*CheckDuplicationStatusUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and duplication can't be triggered unless the user first log in with a valid user
*/
type CheckDuplicationStatusUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckDuplicationStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dataset/duplicate/{request_id}][%d] checkDuplicationStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckDuplicationStatusUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDuplicationStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDuplicationStatusNotFound creates a CheckDuplicationStatusNotFound with default headers values
func NewCheckDuplicationStatusNotFound() *CheckDuplicationStatusNotFound {
	return &CheckDuplicationStatusNotFound{}
}

/*CheckDuplicationStatusNotFound handles this case with default header values.

This means that the ID doesn't exist and thus a status can't be returned.
*/
type CheckDuplicationStatusNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CheckDuplicationStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /dataset/duplicate/{request_id}][%d] checkDuplicationStatusNotFound  %+v", 404, o.Payload)
}

func (o *CheckDuplicationStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDuplicationStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDuplicationStatusRequestURITooLong creates a CheckDuplicationStatusRequestURITooLong with default headers values
func NewCheckDuplicationStatusRequestURITooLong() *CheckDuplicationStatusRequestURITooLong {
	return &CheckDuplicationStatusRequestURITooLong{}
}

/*CheckDuplicationStatusRequestURITooLong handles this case with default header values.

This means that the the request ID is longer than the server is willing to interpret.
*/
type CheckDuplicationStatusRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *CheckDuplicationStatusRequestURITooLong) Error() string {
	return fmt.Sprintf("[GET /dataset/duplicate/{request_id}][%d] checkDuplicationStatusRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *CheckDuplicationStatusRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDuplicationStatusRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDuplicationStatusInternalServerError creates a CheckDuplicationStatusInternalServerError with default headers values
func NewCheckDuplicationStatusInternalServerError() *CheckDuplicationStatusInternalServerError {
	return &CheckDuplicationStatusInternalServerError{}
}

/*CheckDuplicationStatusInternalServerError handles this case with default header values.

This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.
*/
type CheckDuplicationStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CheckDuplicationStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dataset/duplicate/{request_id}][%d] checkDuplicationStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckDuplicationStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckDuplicationStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckDuplicationStatusOKBody check duplication status o k body
swagger:model CheckDuplicationStatusOKBody
*/
type CheckDuplicationStatusOKBody struct {

	// status
	// Required: true
	Status *string `json:"status"`

	// Single path for staging endpoints
	TargetPath string `json:"target_path,omitempty"`
}

// Validate validates this check duplication status o k body
func (o *CheckDuplicationStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckDuplicationStatusOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("checkDuplicationStatusOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckDuplicationStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckDuplicationStatusOKBody) UnmarshalBinary(b []byte) error {
	var res CheckDuplicationStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
