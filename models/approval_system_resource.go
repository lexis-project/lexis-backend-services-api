// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApprovalSystemResource approval system resource
//
// swagger:model ApprovalSystemResource
type ApprovalSystemResource struct {

	// host name
	HostName string `json:"HostName,omitempty"`

	// ID
	ID *int64 `json:"ID,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// performance coefficient
	PerformanceCoefficient *float64 `json:"PerformanceCoefficient,omitempty"`

	// queue list
	QueueList []*ApprovalSystemQueue `json:"QueueList"`
}

// Validate validates this approval system resource
func (m *ApprovalSystemResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueueList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApprovalSystemResource) validateQueueList(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueList) { // not required
		return nil
	}

	for i := 0; i < len(m.QueueList); i++ {
		if swag.IsZero(m.QueueList[i]) { // not required
			continue
		}

		if m.QueueList[i] != nil {
			if err := m.QueueList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("QueueList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApprovalSystemResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApprovalSystemResource) UnmarshalBinary(b []byte) error {
	var res ApprovalSystemResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
