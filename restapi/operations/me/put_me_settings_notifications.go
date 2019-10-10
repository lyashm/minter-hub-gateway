// Code generated by go-swagger; DO NOT EDIT.

package me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutMeSettingsNotificationsHandlerFunc turns a function with the right signature into a put me settings notifications handler
type PutMeSettingsNotificationsHandlerFunc func(PutMeSettingsNotificationsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutMeSettingsNotificationsHandlerFunc) Handle(params PutMeSettingsNotificationsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutMeSettingsNotificationsHandler interface for that can handle valid put me settings notifications params
type PutMeSettingsNotificationsHandler interface {
	Handle(PutMeSettingsNotificationsParams, interface{}) middleware.Responder
}

// NewPutMeSettingsNotifications creates a new http.Handler for the put me settings notifications operation
func NewPutMeSettingsNotifications(ctx *middleware.Context, handler PutMeSettingsNotificationsHandler) *PutMeSettingsNotifications {
	return &PutMeSettingsNotifications{Context: ctx, Handler: handler}
}

/*PutMeSettingsNotifications swagger:route PUT /me/settings/notifications me putMeSettingsNotifications

PutMeSettingsNotifications put me settings notifications API

*/
type PutMeSettingsNotifications struct {
	Context *middleware.Context
	Handler PutMeSettingsNotificationsHandler
}

func (o *PutMeSettingsNotifications) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutMeSettingsNotificationsParams()

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