// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostCommentsCommentIDVoteParams creates a new PostCommentsCommentIDVoteParams object
// with the default values initialized.
func NewPostCommentsCommentIDVoteParams() PostCommentsCommentIDVoteParams {

	var (
		// initialize parameters with default values

		voteDefault = int64(0)
	)

	return PostCommentsCommentIDVoteParams{
		Vote: voteDefault,
	}
}

// PostCommentsCommentIDVoteParams contains all the bound params for the post comments comment ID vote operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostCommentsCommentIDVote
type PostCommentsCommentIDVoteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	CommentID uint64
	/*
	  Required: true
	  In: query
	  Default: 0
	*/
	Vote int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostCommentsCommentIDVoteParams() beforehand.
func (o *PostCommentsCommentIDVoteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rCommentID, rhkCommentID, _ := route.Params.GetOK("commentId")
	if err := o.bindCommentID(rCommentID, rhkCommentID, route.Formats); err != nil {
		res = append(res, err)
	}

	qVote, qhkVote, _ := qs.GetOK("vote")
	if err := o.bindVote(qVote, qhkVote, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCommentID binds and validates parameter CommentID from path.
func (o *PostCommentsCommentIDVoteParams) bindCommentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindVote binds and validates parameter Vote from query.
func (o *PostCommentsCommentIDVoteParams) bindVote(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("vote", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("vote", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("vote", "query", "int64", raw)
	}
	o.Vote = value

	if err := o.validateVote(formats); err != nil {
		return err
	}

	return nil
}

// validateVote carries on validations for parameter Vote
func (o *PostCommentsCommentIDVoteParams) validateVote(formats strfmt.Registry) error {

	if err := validate.Enum("vote", "query", o.Vote, []interface{}{-1, 0, 1}); err != nil {
		return err
	}

	return nil
}