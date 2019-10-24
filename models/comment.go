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

// Comment comment
// swagger:model Comment
type Comment struct {

	// article Id
	// Required: true
	ArticleID *uint64 `json:"articleId"`

	// author
	// Required: true
	Author *Author `json:"author"`

	// body
	// Required: true
	Body *string `json:"body"`

	// children
	Children Comments `json:"children,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// id
	// Required: true
	ID *uint64 `json:"id"`

	// parent Id
	// Required: true
	ParentID *uint64 `json:"parentId"`

	// voites
	Voites int64 `json:"voites,omitempty"`
}

// Validate validates this comment
func (m *Comment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArticleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Comment) validateArticleID(formats strfmt.Registry) error {

	if err := validate.Required("articleId", "body", m.ArticleID); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("author", "body", m.Author); err != nil {
		return err
	}

	if m.Author != nil {
		if err := m.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

func (m *Comment) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateChildren(formats strfmt.Registry) error {

	if swag.IsZero(m.Children) { // not required
		return nil
	}

	if err := m.Children.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("children")
		}
		return err
	}

	return nil
}

func (m *Comment) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parentId", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Comment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Comment) UnmarshalBinary(b []byte) error {
	var res Comment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
