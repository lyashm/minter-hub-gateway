// Code generated by go-swagger; DO NOT EDIT.

package me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostMeNotificationsOKCode is the HTTP code returned for type PostMeNotificationsOK
const PostMeNotificationsOKCode int = 200

/*PostMeNotificationsOK OK

swagger:response postMeNotificationsOK
*/
type PostMeNotificationsOK struct {
}

// NewPostMeNotificationsOK creates PostMeNotificationsOK with default headers values
func NewPostMeNotificationsOK() *PostMeNotificationsOK {

	return &PostMeNotificationsOK{}
}

// WriteResponse to the client
func (o *PostMeNotificationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
