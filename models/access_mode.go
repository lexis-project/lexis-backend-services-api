// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AccessMode access mode
//
// swagger:model AccessMode
type AccessMode string

const (

	// AccessModePublic captures enum value "public"
	AccessModePublic AccessMode = "public"

	// AccessModeProject captures enum value "project"
	AccessModeProject AccessMode = "project"

	// AccessModeUser captures enum value "user"
	AccessModeUser AccessMode = "user"
)

// for schema
var accessModeEnum []interface{}

func init() {
	var res []AccessMode
	if err := json.Unmarshal([]byte(`["public","project","user"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessModeEnum = append(accessModeEnum, v)
	}
}

func (m AccessMode) validateAccessModeEnum(path, location string, value AccessMode) error {
	if err := validate.EnumCase(path, location, value, accessModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this access mode
func (m AccessMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccessModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}