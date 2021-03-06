// Code generated by go-swagger; DO NOT EDIT.

package articles

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

// NewPostArticlesArticleIDVoteParams creates a new PostArticlesArticleIDVoteParams object
// no default values defined in spec.
func NewPostArticlesArticleIDVoteParams() PostArticlesArticleIDVoteParams {

	return PostArticlesArticleIDVoteParams{}
}

// PostArticlesArticleIDVoteParams contains all the bound params for the post articles article ID vote operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostArticlesArticleIDVote
type PostArticlesArticleIDVoteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ArticleID uint64
	/*
	  Required: true
	  In: query
	*/
	Vote int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostArticlesArticleIDVoteParams() beforehand.
func (o *PostArticlesArticleIDVoteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rArticleID, rhkArticleID, _ := route.Params.GetOK("articleId")
	if err := o.bindArticleID(rArticleID, rhkArticleID, route.Formats); err != nil {
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

// bindArticleID binds and validates parameter ArticleID from path.
func (o *PostArticlesArticleIDVoteParams) bindArticleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("articleId", "path", "uint64", raw)
	}
	o.ArticleID = value

	return nil
}

// bindVote binds and validates parameter Vote from query.
func (o *PostArticlesArticleIDVoteParams) bindVote(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PostArticlesArticleIDVoteParams) validateVote(formats strfmt.Registry) error {

	if err := validate.Enum("vote", "query", o.Vote, []interface{}{-1, 0, 1}); err != nil {
		return err
	}

	return nil
}
