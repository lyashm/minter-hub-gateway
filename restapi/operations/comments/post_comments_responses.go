// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "morg/models"
)

// PostCommentsOKCode is the HTTP code returned for type PostCommentsOK
const PostCommentsOKCode int = 200

/*PostCommentsOK OK

swagger:response postCommentsOK
*/
type PostCommentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Comment `json:"body,omitempty"`
}

// NewPostCommentsOK creates PostCommentsOK with default headers values
func NewPostCommentsOK() *PostCommentsOK {

	return &PostCommentsOK{}
}

// WithPayload adds the payload to the post comments o k response
func (o *PostCommentsOK) WithPayload(payload *models.Comment) *PostCommentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments o k response
func (o *PostCommentsOK) SetPayload(payload *models.Comment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
