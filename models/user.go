// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
// swagger:model User
type User struct {

	// bio
	Bio string `json:"bio,omitempty"`

	// email
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// id
	ID uint64 `json:"id,omitempty"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// status
	Status uint32 `json:"status,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// username
	// Minimum: 6
	Username string `json:"username,omitempty"`

	// wallet
	Wallet string `json:"wallet,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateUsername(formats strfmt.Registry) error {

	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.Minimum("username", "body", float64(m.Username), 6, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
