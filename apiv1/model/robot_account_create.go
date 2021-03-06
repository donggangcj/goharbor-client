// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RobotAccountCreate robot account create
//
// swagger:model RobotAccountCreate
type RobotAccountCreate struct {

	// The permission of robot account
	Access []*RobotAccountAccess `json:"access"`

	// The description of robot account
	Description string `json:"description,omitempty"`

	// The name of robot account
	Name string `json:"name,omitempty"`
}

// Validate validates this robot account create
func (m *RobotAccountCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RobotAccountCreate) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	for i := 0; i < len(m.Access); i++ {
		if swag.IsZero(m.Access[i]) { // not required
			continue
		}

		if m.Access[i] != nil {
			if err := m.Access[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this robot account create based on the context it is used
func (m *RobotAccountCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RobotAccountCreate) contextValidateAccess(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Access); i++ {

		if m.Access[i] != nil {
			if err := m.Access[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RobotAccountCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RobotAccountCreate) UnmarshalBinary(b []byte) error {
	var res RobotAccountCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
