// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ItemCreatedResponse item created response
//
// swagger:model ItemCreatedResponse
type ItemCreatedResponse struct {

	// ID
	ID string `json:"ID,omitempty"`

	// link
	// Required: true
	Link *string `json:"Link"`
}

// Validate validates this item created response
func (m *ItemCreatedResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemCreatedResponse) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("Link", "body", m.Link); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemCreatedResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemCreatedResponse) UnmarshalBinary(b []byte) error {
	var res ItemCreatedResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}