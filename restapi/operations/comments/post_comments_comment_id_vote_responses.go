// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostCommentsCommentIDVoteOKCode is the HTTP code returned for type PostCommentsCommentIDVoteOK
const PostCommentsCommentIDVoteOKCode int = 200

/*PostCommentsCommentIDVoteOK OK

swagger:response postCommentsCommentIdVoteOK
*/
type PostCommentsCommentIDVoteOK struct {

	/*
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewPostCommentsCommentIDVoteOK creates PostCommentsCommentIDVoteOK with default headers values
func NewPostCommentsCommentIDVoteOK() *PostCommentsCommentIDVoteOK {

	return &PostCommentsCommentIDVoteOK{}
}

// WithPayload adds the payload to the post comments comment Id vote o k response
func (o *PostCommentsCommentIDVoteOK) WithPayload(payload int64) *PostCommentsCommentIDVoteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments comment Id vote o k response
func (o *PostCommentsCommentIDVoteOK) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsCommentIDVoteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}