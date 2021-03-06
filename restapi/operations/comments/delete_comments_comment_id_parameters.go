// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteCommentsCommentIDParams creates a new DeleteCommentsCommentIDParams object
// no default values defined in spec.
func NewDeleteCommentsCommentIDParams() DeleteCommentsCommentIDParams {

	return DeleteCommentsCommentIDParams{}
}

// DeleteCommentsCommentIDParams contains all the bound params for the delete comments comment ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteCommentsCommentID
type DeleteCommentsCommentIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	CommentID uint64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteCommentsCommentIDParams() beforehand.
func (o *DeleteCommentsCommentIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rCommentID, rhkCommentID, _ := route.Params.GetOK("commentId")
	if err := o.bindCommentID(rCommentID, rhkCommentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCommentID binds and validates parameter CommentID from path.
func (o *DeleteCommentsCommentIDParams) bindCommentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("commentId", "path", "uint64", raw)
	}
	o.CommentID = value

	return nil
}
