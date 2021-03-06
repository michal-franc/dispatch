///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subscriber subscriber
// swagger:model Subscriber

type Subscriber struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// type
	// Required: true
	Type *string `json:"type"`
}

/* polymorph Subscriber name false */

/* polymorph Subscriber type false */

// Validate validates this subscriber
func (m *Subscriber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscriber) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var subscriberTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["function","event"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriberTypeTypePropEnum = append(subscriberTypeTypePropEnum, v)
	}
}

const (
	// SubscriberTypeFunction captures enum value "function"
	SubscriberTypeFunction string = "function"
	// SubscriberTypeEvent captures enum value "event"
	SubscriberTypeEvent string = "event"
)

// prop value enum
func (m *Subscriber) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriberTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscriber) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscriber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscriber) UnmarshalBinary(b []byte) error {
	var res Subscriber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
