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

// Location location
//
// swagger:model Location
type Location struct {

	// access
	// Required: true
	Access AccessMode `json:"access"`

	// internal ID
	// Required: true
	InternalID *string `json:"internalID"`

	// project
	// Required: true
	Project *string `json:"project"`

	// zone
	// Required: true
	Zone *string `json:"zone"`
}

// Validate validates this location
func (m *Location) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Location) validateAccess(formats strfmt.Registry) error {

	if err := m.Access.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		}
		return err
	}

	return nil
}

func (m *Location) validateInternalID(formats strfmt.Registry) error {

	if err := validate.Required("internalID", "body", m.InternalID); err != nil {
		return err
	}

	return nil
}

func (m *Location) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

func (m *Location) validateZone(formats strfmt.Registry) error {

	if err := validate.Required("zone", "body", m.Zone); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Location) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Location) UnmarshalBinary(b []byte) error {
	var res Location
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}