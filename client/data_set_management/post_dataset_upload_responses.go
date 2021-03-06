// Code generated by go-swagger; DO NOT EDIT.

package data_set_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/lexis-project/lexis-backend-services-api.git/models"
)

// PostDatasetUploadReader is a Reader for the PostDatasetUpload structure.
type PostDatasetUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDatasetUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostDatasetUploadCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDatasetUploadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDatasetUploadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPostDatasetUploadPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostDatasetUploadRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostDatasetUploadUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 460:
		result := NewPostDatasetUploadStatus460()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostDatasetUploadServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDatasetUploadCreated creates a PostDatasetUploadCreated with default headers values
func NewPostDatasetUploadCreated() *PostDatasetUploadCreated {
	return &PostDatasetUploadCreated{}
}

/*PostDatasetUploadCreated handles this case with default header values.

Created
*/
type PostDatasetUploadCreated struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.

	*/
	AccessControlExposeHeaders string
	/*Url of the created resource.
	 */
	Location string

	TusResumable string
	/*Added by the Creation With Upload Extension in combination with the expiration extension. The Upload-Expires response header indicates the time after which the unfinished upload expires. A Server MAY wish to remove incomplete uploads after a given period of time to prevent abandoned uploads from taking up extra storage. The Client SHOULD use this header to determine if an upload is still valid before attempting to resume the upload. This header MUST be included in every PATCH response if the upload is going to expire. If the expiration is known at the creation, the Upload-Expires header MUST be included in the response to the initial POST request. Its value MAY change over time. If a Client does attempt to resume an upload which has since been removed by the Server, the Server SHOULD respond with the 404 Not Found or 410 Gone status. The latter one SHOULD be used if the Server is keeping track of expired uploads. In both cases the Client SHOULD start a new upload. The value of the Upload-Expires header MUST be in RFC 7231 datetime format.
	 */
	UploadExpires string
	/*Added by the Creation With Upload Extension. The Upload-Offset request and response header indicates a byte offset within a resource. The value MUST be a non-negative integer.
	 */
	UploadOffset int64
}

func (o *PostDatasetUploadCreated) Error() string {
	return fmt.Sprintf("[POST /dataset/upload/][%d] postDatasetUploadCreated ", 201)
}

func (o *PostDatasetUploadCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Expose-Headers
	o.AccessControlExposeHeaders = response.GetHeader("Access-Control-Expose-Headers")

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header Tus-Resumable
	o.TusResumable = response.GetHeader("Tus-Resumable")

	// response header Upload-Expires
	o.UploadExpires = response.GetHeader("Upload-Expires")

	// response header Upload-Offset
	uploadOffset, err := swag.ConvertInt64(response.GetHeader("Upload-Offset"))
	if err != nil {
		return errors.InvalidType("Upload-Offset", "header", "int64", response.GetHeader("Upload-Offset"))
	}
	o.UploadOffset = uploadOffset

	return nil
}

// NewPostDatasetUploadBadRequest creates a PostDatasetUploadBadRequest with default headers values
func NewPostDatasetUploadBadRequest() *PostDatasetUploadBadRequest {
	return &PostDatasetUploadBadRequest{}
}

/*PostDatasetUploadBadRequest handles this case with default header values.

Added by the Creation With Upload Extension in combination with the checksum extension. The checksum algorithm is not supported by the server
*/
type PostDatasetUploadBadRequest struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.

	*/
	AccessControlExposeHeaders string
	/*Protocol version
	 */
	TusResumable string
}

func (o *PostDatasetUploadBadRequest) Error() string {
	return fmt.Sprintf("[POST /dataset/upload/][%d] postDatasetUploadBadRequest ", 400)
}

func (o *PostDatasetUploadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Expose-Headers
	o.AccessControlExposeHeaders = response.GetHeader("Access-Control-Expose-Headers")

	// response header Tus-Resumable
	o.TusResumable = response.GetHeader("Tus-Resumable")

	return nil
}

// NewPostDatasetUploadUnauthorized creates a PostDatasetUploadUnauthorized with default headers values
func NewPostDatasetUploadUnauthorized() *PostDatasetUploadUnauthorized {
	return &PostDatasetUploadUnauthorized{}
}

/*PostDatasetUploadUnauthorized handles this case with default header values.

Authorization failed
*/
type PostDatasetUploadUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *PostDatasetUploadUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dataset/upload/][%d] postDatasetUploadUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDatasetUploadUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostDatasetUploadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDatasetUploadPreconditionFailed creates a PostDatasetUploadPreconditionFailed with default headers values
func NewPostDatasetUploadPreconditionFailed() *PostDatasetUploadPreconditionFailed {
	return &PostDatasetUploadPreconditionFailed{}
}

/*PostDatasetUploadPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type PostDatasetUploadPreconditionFailed struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.

	*/
	AccessControlExposeHeaders string

	TusResumable string
	/*The Tus-Version response header MUST be a comma-separated list of protocol versions supported by the Server. The list MUST be sorted by Server's preference where the first one is the most preferred one.
	 */
	TusVersion string
}

func (o *PostDatasetUploadPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /dataset/upload/][%d] postDatasetUploadPreconditionFailed ", 412)
}

func (o *PostDatasetUploadPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Expose-Headers
	o.AccessControlExposeHeaders = response.GetHeader("Access-Control-Expose-Headers")

	// response header Tus-Resumable
	o.TusResumable = response.GetHeader("Tus-Resumable")

	// response header Tus-Version
	o.TusVersion = response.GetHeader("Tus-Version")

	return nil
}

// NewPostDatasetUploadRequestEntityTooLarge creates a PostDatasetUploadRequestEntityTooLarge with default headers values
func NewPostDatasetUploadRequestEntityTooLarge() *PostDatasetUploadRequestEntityTooLarge {
	return &PostDatasetUploadRequestEntityTooLarge{}
}

/*PostDatasetUploadRequestEntityTooLarge handles this case with default header values.

If the length of the upload exceeds the maximum, which MAY be specified using the Tus-Max-Size header, the Server MUST respond with the 413 Request Entity Too Large status.
*/
type PostDatasetUploadRequestEntityTooLarge struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.

	*/
	AccessControlExposeHeaders string

	TusResumable string
}

func (o *PostDatasetUploadRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /dataset/upload/][%d] postDatasetUploadRequestEntityTooLarge ", 413)
}

func (o *PostDatasetUploadRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Expose-Headers
	o.AccessControlExposeHeaders = response.GetHeader("Access-Control-Expose-Headers")

	// response header Tus-Resumable
	o.TusResumable = response.GetHeader("Tus-Resumable")

	return nil
}

// NewPostDatasetUploadUnsupportedMediaType creates a PostDatasetUploadUnsupportedMediaType with default headers values
func NewPostDatasetUploadUnsupportedMediaType() *PostDatasetUploadUnsupportedMediaType {
	return &PostDatasetUploadUnsupportedMediaType{}
}

/*PostDatasetUploadUnsupportedMediaType handles this case with default header values.

Added by the Creation With Upload Extension. Content-Type was not application/offset+octet-stream
*/
type PostDatasetUploadUnsupportedMediaType struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.

	*/
	AccessControlExposeHeaders string
	/*Protocol version
	 */
	TusResumable string
}

func (o *PostDatasetUploadUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /dataset/upload/][%d] postDatasetUploadUnsupportedMediaType ", 415)
}

func (o *PostDatasetUploadUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Expose-Headers
	o.AccessControlExposeHeaders = response.GetHeader("Access-Control-Expose-Headers")

	// response header Tus-Resumable
	o.TusResumable = response.GetHeader("Tus-Resumable")

	return nil
}

// NewPostDatasetUploadStatus460 creates a PostDatasetUploadStatus460 with default headers values
func NewPostDatasetUploadStatus460() *PostDatasetUploadStatus460 {
	return &PostDatasetUploadStatus460{}
}

/*PostDatasetUploadStatus460 handles this case with default header values.

Added by the Creation With Upload Extension in combination with the checksum extension. Checksums mismatch
*/
type PostDatasetUploadStatus460 struct {
	/*Needed to make browsers accept the additional headers used by
	the tus protocol.

	*/
	AccessControlExposeHeaders string
	/*Protocol version
	 */
	TusResumable string
}

func (o *PostDatasetUploadStatus460) Error() string {
	return fmt.Sprintf("[POST /dataset/upload/][%d] postDatasetUploadStatus460 ", 460)
}

func (o *PostDatasetUploadStatus460) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Expose-Headers
	o.AccessControlExposeHeaders = response.GetHeader("Access-Control-Expose-Headers")

	// response header Tus-Resumable
	o.TusResumable = response.GetHeader("Tus-Resumable")

	return nil
}

// NewPostDatasetUploadServiceUnavailable creates a PostDatasetUploadServiceUnavailable with default headers values
func NewPostDatasetUploadServiceUnavailable() *PostDatasetUploadServiceUnavailable {
	return &PostDatasetUploadServiceUnavailable{}
}

/*PostDatasetUploadServiceUnavailable handles this case with default header values.

Service Unavailable
*/
type PostDatasetUploadServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *PostDatasetUploadServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /dataset/upload/][%d] postDatasetUploadServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostDatasetUploadServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PostDatasetUploadServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
