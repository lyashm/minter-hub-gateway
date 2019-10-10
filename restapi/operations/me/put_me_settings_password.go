// Code generated by go-swagger; DO NOT EDIT.

package me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutMeSettingsPasswordHandlerFunc turns a function with the right signature into a put me settings password handler
type PutMeSettingsPasswordHandlerFunc func(PutMeSettingsPasswordParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutMeSettingsPasswordHandlerFunc) Handle(params PutMeSettingsPasswordParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutMeSettingsPasswordHandler interface for that can handle valid put me settings password params
type PutMeSettingsPasswordHandler interface {
	Handle(PutMeSettingsPasswordParams, interface{}) middleware.Responder
}

// NewPutMeSettingsPassword creates a new http.Handler for the put me settings password operation
func NewPutMeSettingsPassword(ctx *middleware.Context, handler PutMeSettingsPasswordHandler) *PutMeSettingsPassword {
	return &PutMeSettingsPassword{Context: ctx, Handler: handler}
}

/*PutMeSettingsPassword swagger:route PUT /me/settings/password me putMeSettingsPassword

PutMeSettingsPassword put me settings password API

*/
type PutMeSettingsPassword struct {
	Context *middleware.Context
	Handler PutMeSettingsPasswordHandler
}

func (o *PutMeSettingsPassword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutMeSettingsPasswordParams()

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
