// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostFeedsFeedIDJoinOKCode is the HTTP code returned for type PostFeedsFeedIDJoinOK
const PostFeedsFeedIDJoinOKCode int = 200

/*PostFeedsFeedIDJoinOK OK

swagger:response postFeedsFeedIdJoinOK
*/
type PostFeedsFeedIDJoinOK struct {
}

// NewPostFeedsFeedIDJoinOK creates PostFeedsFeedIDJoinOK with default headers values
func NewPostFeedsFeedIDJoinOK() *PostFeedsFeedIDJoinOK {

	return &PostFeedsFeedIDJoinOK{}
}

// WriteResponse to the client
func (o *PostFeedsFeedIDJoinOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
