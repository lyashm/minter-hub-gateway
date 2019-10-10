// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteFeedsFeedIDHandlerFunc turns a function with the right signature into a delete feeds feed ID handler
type DeleteFeedsFeedIDHandlerFunc func(DeleteFeedsFeedIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFeedsFeedIDHandlerFunc) Handle(params DeleteFeedsFeedIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteFeedsFeedIDHandler interface for that can handle valid delete feeds feed ID params
type DeleteFeedsFeedIDHandler interface {
	Handle(DeleteFeedsFeedIDParams, interface{}) middleware.Responder
}

// NewDeleteFeedsFeedID creates a new http.Handler for the delete feeds feed ID operation
func NewDeleteFeedsFeedID(ctx *middleware.Context, handler DeleteFeedsFeedIDHandler) *DeleteFeedsFeedID {
	return &DeleteFeedsFeedID{Context: ctx, Handler: handler}
}

/*DeleteFeedsFeedID swagger:route DELETE /feeds/{feedId} feeds deleteFeedsFeedId

DeleteFeedsFeedID delete feeds feed ID API

*/
type DeleteFeedsFeedID struct {
	Context *middleware.Context
	Handler DeleteFeedsFeedIDHandler
}

func (o *DeleteFeedsFeedID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFeedsFeedIDParams()

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
