// Code generated by go-swagger; DO NOT EDIT.

package articles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostArticlesArticleIDVoteHandlerFunc turns a function with the right signature into a post articles article ID vote handler
type PostArticlesArticleIDVoteHandlerFunc func(PostArticlesArticleIDVoteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostArticlesArticleIDVoteHandlerFunc) Handle(params PostArticlesArticleIDVoteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostArticlesArticleIDVoteHandler interface for that can handle valid post articles article ID vote params
type PostArticlesArticleIDVoteHandler interface {
	Handle(PostArticlesArticleIDVoteParams, interface{}) middleware.Responder
}

// NewPostArticlesArticleIDVote creates a new http.Handler for the post articles article ID vote operation
func NewPostArticlesArticleIDVote(ctx *middleware.Context, handler PostArticlesArticleIDVoteHandler) *PostArticlesArticleIDVote {
	return &PostArticlesArticleIDVote{Context: ctx, Handler: handler}
}

/*PostArticlesArticleIDVote swagger:route POST /articles/{articleId}/vote articles postArticlesArticleIdVote

PostArticlesArticleIDVote post articles article ID vote API

*/
type PostArticlesArticleIDVote struct {
	Context *middleware.Context
	Handler PostArticlesArticleIDVoteHandler
}

func (o *PostArticlesArticleIDVote) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostArticlesArticleIDVoteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
