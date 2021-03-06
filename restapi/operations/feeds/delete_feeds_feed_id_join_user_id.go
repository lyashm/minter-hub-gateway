// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteFeedsFeedIDJoinUserIDHandlerFunc turns a function with the right signature into a delete feeds feed ID join user ID handler
type DeleteFeedsFeedIDJoinUserIDHandlerFunc func(DeleteFeedsFeedIDJoinUserIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFeedsFeedIDJoinUserIDHandlerFunc) Handle(params DeleteFeedsFeedIDJoinUserIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteFeedsFeedIDJoinUserIDHandler interface for that can handle valid delete feeds feed ID join user ID params
type DeleteFeedsFeedIDJoinUserIDHandler interface {
	Handle(DeleteFeedsFeedIDJoinUserIDParams, interface{}) middleware.Responder
}

// NewDeleteFeedsFeedIDJoinUserID creates a new http.Handler for the delete feeds feed ID join user ID operation
func NewDeleteFeedsFeedIDJoinUserID(ctx *middleware.Context, handler DeleteFeedsFeedIDJoinUserIDHandler) *DeleteFeedsFeedIDJoinUserID {
	return &DeleteFeedsFeedIDJoinUserID{Context: ctx, Handler: handler}
}

/*DeleteFeedsFeedIDJoinUserID swagger:route DELETE /feeds/{feedId}/join/{userId} feeds deleteFeedsFeedIdJoinUserId

DeleteFeedsFeedIDJoinUserID delete feeds feed ID join user ID API

*/
type DeleteFeedsFeedIDJoinUserID struct {
	Context *middleware.Context
	Handler DeleteFeedsFeedIDJoinUserIDHandler
}

func (o *DeleteFeedsFeedIDJoinUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFeedsFeedIDJoinUserIDParams()

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
