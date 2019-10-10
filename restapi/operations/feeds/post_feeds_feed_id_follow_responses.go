// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostFeedsFeedIDFollowOKCode is the HTTP code returned for type PostFeedsFeedIDFollowOK
const PostFeedsFeedIDFollowOKCode int = 200

/*PostFeedsFeedIDFollowOK OK

swagger:response postFeedsFeedIdFollowOK
*/
type PostFeedsFeedIDFollowOK struct {
}

// NewPostFeedsFeedIDFollowOK creates PostFeedsFeedIDFollowOK with default headers values
func NewPostFeedsFeedIDFollowOK() *PostFeedsFeedIDFollowOK {

	return &PostFeedsFeedIDFollowOK{}
}

// WriteResponse to the client
func (o *PostFeedsFeedIDFollowOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}