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

// Release release
//
// swagger:model Release
type Release struct {

	// ID
	// Min Length: 1
	ID string `json:"ID,omitempty"`

	// policy ID
	PolicyID int64 `json:"PolicyID,omitempty"`

	// source version ID
	SourceVersionID int64 `json:"SourceVersionID,omitempty"`

	// target version
	TargetVersion int64 `json:"TargetVersion,omitempty"`
}

// Validate validates this release
func (m *Release) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Release) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("ID", "body", string(m.ID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Release) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Release) UnmarshalBinary(b []byte) error {
	var res Release
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
