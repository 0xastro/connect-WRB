// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PermissionType permission type
// swagger:model permission_type
type PermissionType string

const (

	// PermissionTypeNONE captures enum value "NONE"
	PermissionTypeNONE PermissionType = "NONE"

	// PermissionTypeREAD captures enum value "READ"
	PermissionTypeREAD PermissionType = "READ"

	// PermissionTypeWRITE captures enum value "WRITE"
	PermissionTypeWRITE PermissionType = "WRITE"
)

// for schema
var permissionTypeEnum []interface{}

func init() {
	var res []PermissionType
	if err := json.Unmarshal([]byte(`["NONE","READ","WRITE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		permissionTypeEnum = append(permissionTypeEnum, v)
	}
}

func (m PermissionType) validatePermissionTypeEnum(path, location string, value PermissionType) error {
	if err := validate.Enum(path, location, value, permissionTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this permission type
func (m PermissionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePermissionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
