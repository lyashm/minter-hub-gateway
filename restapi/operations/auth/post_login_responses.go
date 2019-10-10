// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "morg/models"
)

// PostLoginOKCode is the HTTP code returned for type PostLoginOK
const PostLoginOKCode int = 200

/*PostLoginOK OK

swagger:response postLoginOK
*/
type PostLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewPostLoginOK creates PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {

	return &PostLoginOK{}
}

// WithPayload adds the payload to the post login o k response
func (o *PostLoginOK) WithPayload(payload *models.Profile) *PostLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login o k response
func (o *PostLoginOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}