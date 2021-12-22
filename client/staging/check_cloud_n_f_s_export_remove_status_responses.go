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

// CheckCloudNFSExportRemoveStatusReader is a Reader for the CheckCloudNFSExportRemoveStatus structure.
type CheckCloudNFSExportRemoveStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckCloudNFSExportRemoveStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckCloudNFSExportRemoveStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckCloudNFSExportRemoveStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckCloudNFSExportRemoveStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckCloudNFSExportRemoveStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 414:
		result := NewCheckCloudNFSExportRemoveStatusRequestURITooLong()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckCloudNFSExportRemoveStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckCloudNFSExportRemoveStatusOK creates a CheckCloudNFSExportRemoveStatusOK with default headers values
func NewCheckCloudNFSExportRemoveStatusOK() *CheckCloudNFSExportRemoveStatusOK {
	return &CheckCloudNFSExportRemoveStatusOK{}
}

/*CheckCloudNFSExportRemoveStatusOK handles this case with default header values.

This means that the status has been returned to the user in the response body.
*/
type CheckCloudNFSExportRemoveStatusOK struct {
	Payload *CheckCloudNFSExportRemoveStatusOKBody
}

func (o *CheckCloudNFSExportRemoveStatusOK) Error() string {
	return fmt.Sprintf("[GET /dataset/cloud/remove/{param}][%d] checkCloudNFSExportRemoveStatusOK  %+v", 200, o.Payload)
}

func (o *CheckCloudNFSExportRemoveStatusOK) GetPayload() *CheckCloudNFSExportRemoveStatusOKBody {
	return o.Payload
}

func (o *CheckCloudNFSExportRemoveStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckCloudNFSExportRemoveStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCloudNFSExportRemoveStatusBadRequest creates a CheckCloudNFSExportRemoveStatusBadRequest with default headers values
func NewCheckCloudNFSExportRemoveStatusBadRequest() *CheckCloudNFSExportRemoveStatusBadRequest {
	return &CheckCloudNFSExportRemoveStatusBadRequest{}
}

/*CheckCloudNFSExportRemoveStatusBadRequest handles this case with default header values.

This means that the request ID given by the user is incorrect.
*/
type CheckCloudNFSExportRemoveStatusBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CheckCloudNFSExportRemoveStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /dataset/cloud/remove/{param}][%d] checkCloudNFSExportRemoveStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckCloudNFSExportRemoveStatusBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCloudNFSExportRemoveStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCloudNFSExportRemoveStatusUnauthorized creates a CheckCloudNFSExportRemoveStatusUnauthorized with default headers values
func NewCheckCloudNFSExportRemoveStatusUnauthorized() *CheckCloudNFSExportRemoveStatusUnauthorized {
	return &CheckCloudNFSExportRemoveStatusUnauthorized{}
}

/*CheckCloudNFSExportRemoveStatusUnauthorized handles this case with default header values.

This means that the user is not authenticated with keycloak and data transfer can't be triggered unless the user first log in with a valid user
*/
type CheckCloudNFSExportRemoveStatusUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CheckCloudNFSExportRemoveStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dataset/cloud/remove/{param}][%d] checkCloudNFSExportRemoveStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckCloudNFSExportRemoveStatusUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCloudNFSExportRemoveStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCloudNFSExportRemoveStatusNotFound creates a CheckCloudNFSExportRemoveStatusNotFound with default headers values
func NewCheckCloudNFSExportRemoveStatusNotFound() *CheckCloudNFSExportRemoveStatusNotFound {
	return &CheckCloudNFSExportRemoveStatusNotFound{}
}

/*CheckCloudNFSExportRemoveStatusNotFound handles this case with default header values.

This means that the ID doesn't exist and thus a status can't be returned.
*/
type CheckCloudNFSExportRemoveStatusNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CheckCloudNFSExportRemoveStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /dataset/cloud/remove/{param}][%d] checkCloudNFSExportRemoveStatusNotFound  %+v", 404, o.Payload)
}

func (o *CheckCloudNFSExportRemoveStatusNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCloudNFSExportRemoveStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCloudNFSExportRemoveStatusRequestURITooLong creates a CheckCloudNFSExportRemoveStatusRequestURITooLong with default headers values
func NewCheckCloudNFSExportRemoveStatusRequestURITooLong() *CheckCloudNFSExportRemoveStatusRequestURITooLong {
	return &CheckCloudNFSExportRemoveStatusRequestURITooLong{}
}

/*CheckCloudNFSExportRemoveStatusRequestURITooLong handles this case with default header values.

This means that the the request ID is longer than the server is willing to interpret.
*/
type CheckCloudNFSExportRemoveStatusRequestURITooLong struct {
	Payload *models.ErrorResponse
}

func (o *CheckCloudNFSExportRemoveStatusRequestURITooLong) Error() string {
	return fmt.Sprintf("[GET /dataset/cloud/remove/{param}][%d] checkCloudNFSExportRemoveStatusRequestUriTooLong  %+v", 414, o.Payload)
}

func (o *CheckCloudNFSExportRemoveStatusRequestURITooLong) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCloudNFSExportRemoveStatusRequestURITooLong) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckCloudNFSExportRemoveStatusInternalServerError creates a CheckCloudNFSExportRemoveStatusInternalServerError with default headers values
func NewCheckCloudNFSExportRemoveStatusInternalServerError() *CheckCloudNFSExportRemoveStatusInternalServerError {
	return &CheckCloudNFSExportRemoveStatusInternalServerError{}
}

/*CheckCloudNFSExportRemoveStatusInternalServerError handles this case with default header values.

This means that something has gone wrong on the data transfer steering server. The user is advised to wait and send the request again.
*/
type CheckCloudNFSExportRemoveStatusInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CheckCloudNFSExportRemoveStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dataset/cloud/remove/{param}][%d] checkCloudNFSExportRemoveStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckCloudNFSExportRemoveStatusInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CheckCloudNFSExportRemoveStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckCloudNFSExportRemoveStatusOKBody check cloud n f s export remove status o k body
swagger:model CheckCloudNFSExportRemoveStatusOKBody
*/
type CheckCloudNFSExportRemoveStatusOKBody struct {

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this check cloud n f s export remove status o k body
func (o *CheckCloudNFSExportRemoveStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckCloudNFSExportRemoveStatusOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("checkCloudNFSExportRemoveStatusOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckCloudNFSExportRemoveStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckCloudNFSExportRemoveStatusOKBody) UnmarshalBinary(b []byte) error {
	var res CheckCloudNFSExportRemoveStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
