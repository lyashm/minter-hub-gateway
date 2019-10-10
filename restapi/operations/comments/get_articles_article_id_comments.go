// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetArticlesArticleIDCommentsHandlerFunc turns a function with the right signature into a get articles article ID comments handler
type GetArticlesArticleIDCommentsHandlerFunc func(GetArticlesArticleIDCommentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetArticlesArticleIDCommentsHandlerFunc) Handle(params GetArticlesArticleIDCommentsParams) middleware.Responder {
	return fn(params)
}

// GetArticlesArticleIDCommentsHandler interface for that can handle valid get articles article ID comments params
type GetArticlesArticleIDCommentsHandler interface {
	Handle(GetArticlesArticleIDCommentsParams) middleware.Responder
}

// NewGetArticlesArticleIDComments creates a new http.Handler for the get articles article ID comments operation
func NewGetArticlesArticleIDComments(ctx *middleware.Context, handler GetArticlesArticleIDCommentsHandler) *GetArticlesArticleIDComments {
	return &GetArticlesArticleIDComments{Context: ctx, Handler: handler}
}

/*GetArticlesArticleIDComments swagger:route GET /articles/{articleId}/comments comments getArticlesArticleIdComments

GetArticlesArticleIDComments get articles article ID comments API

*/
type GetArticlesArticleIDComments struct {
	Context *middleware.Context
	Handler GetArticlesArticleIDCommentsHandler
}

func (o *GetArticlesArticleIDComments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetArticlesArticleIDCommentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}