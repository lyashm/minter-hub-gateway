// Code generated by go-swagger; DO NOT EDIT.

package me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutMeSettingsHandlerFunc turns a function with the right signature into a put me settings handler
type PutMeSettingsHandlerFunc func(PutMeSettingsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutMeSettingsHandlerFunc) Handle(params PutMeSettingsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutMeSettingsHandler interface for that can handle valid put me settings params
type PutMeSettingsHandler interface {
	Handle(PutMeSettingsParams, interface{}) middleware.Responder
}

// NewPutMeSettings creates a new http.Handler for the put me settings operation
func NewPutMeSettings(ctx *middleware.Context, handler PutMeSettingsHandler) *PutMeSettings {
	return &PutMeSettings{Context: ctx, Handler: handler}
}

/*PutMeSettings swagger:route PUT /me/settings me putMeSettings

PutMeSettings put me settings API

*/
type PutMeSettings struct {
	Context *middleware.Context
	Handler PutMeSettingsHandler
}

func (o *PutMeSettings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutMeSettingsParams()

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
