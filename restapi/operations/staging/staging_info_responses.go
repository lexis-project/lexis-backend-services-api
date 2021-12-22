// Code generated by go-swagger; DO NOT EDIT.

package staging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// StagingInfoOKCode is the HTTP code returned for type StagingInfoOK
const StagingInfoOKCode int = 200

/*StagingInfoOK List of possible target / source

swagger:response stagingInfoOK
*/
type StagingInfoOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewStagingInfoOK creates StagingInfoOK with default headers values
func NewStagingInfoOK() *StagingInfoOK {

	return &StagingInfoOK{}
}

// WithPayload adds the payload to the staging info o k response
func (o *StagingInfoOK) WithPayload(payload []string) *StagingInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the staging info o k response
func (o *StagingInfoOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StagingInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}