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

// Article article
// swagger:model Article
type Article struct {

	// author
	// Required: true
	Author *Author `json:"author"`

	// comments
	// Required: true
	Comments *uint64 `json:"comments"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// description
	// Required: true
	Description *string `json:"description"`

	// id
	// Required: true
	ID *uint64 `json:"id"`

	// image
	// Required: true
	Image *string `json:"image"`

	// tags
	// Required: true
	// Unique: true
	Tags []string `json:"tags"`

	// title
	// Required: true
	Title *string `json:"title"`

	// voites
	// Required: true
	Voites *int64 `json:"voites"`

	// voted
	// Required: true
	// Enum: [1 0 -1]
	Voted *int32 `json:"voted"`
}

// Validate validates this article
func (m *Article) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Article) validateAuthor(formats strfmt.Registry) error {

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

func (m *Article) validateComments(formats strfmt.Registry) error {

	if err := validate.Required("comments", "body", m.Comments); err != nil {
		return err
	}

	return nil
}

func (m *Article) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Article) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Article) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Article) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *Article) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

func (m *Article) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Article) validateVoites(formats strfmt.Registry) error {

	if err := validate.Required("voites", "body", m.Voites); err != nil {
		return err
	}

	return nil
}

var articleTypeVotedPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[1,0,-1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		articleTypeVotedPropEnum = append(articleTypeVotedPropEnum, v)
	}
}

// prop value enum
func (m *Article) validateVotedEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, articleTypeVotedPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Article) validateVoted(formats strfmt.Registry) error {

	if err := validate.Required("voted", "body", m.Voted); err != nil {
		return err
	}

	// value enum
	if err := m.validateVotedEnum("voted", "body", *m.Voted); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Article) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Article) UnmarshalBinary(b []byte) error {
	var res Article
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
