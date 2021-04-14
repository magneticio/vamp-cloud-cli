// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplicationInput application input
//
// swagger:model ApplicationInput
type ApplicationInput struct {

	// cluster ID
	// Required: true
	ClusterID *int64 `json:"clusterID"`

	// description
	Description string `json:"description,omitempty"`

	// ingress type
	// Required: true
	// Enum: [NGINX CONTOUR]
	IngressType *string `json:"ingressType"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`
}

// Validate validates this application input
func (m *ApplicationInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationInput) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterID", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

var applicationInputTypeIngressTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NGINX","CONTOUR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationInputTypeIngressTypePropEnum = append(applicationInputTypeIngressTypePropEnum, v)
	}
}

const (

	// ApplicationInputIngressTypeNGINX captures enum value "NGINX"
	ApplicationInputIngressTypeNGINX string = "NGINX"

	// ApplicationInputIngressTypeCONTOUR captures enum value "CONTOUR"
	ApplicationInputIngressTypeCONTOUR string = "CONTOUR"
)

// prop value enum
func (m *ApplicationInput) validateIngressTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationInputTypeIngressTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationInput) validateIngressType(formats strfmt.Registry) error {

	if err := validate.Required("ingressType", "body", m.IngressType); err != nil {
		return err
	}

	// value enum
	if err := m.validateIngressTypeEnum("ingressType", "body", *m.IngressType); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationInput) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationInput) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationInput) UnmarshalBinary(b []byte) error {
	var res ApplicationInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
