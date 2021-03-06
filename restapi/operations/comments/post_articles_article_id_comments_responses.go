// Code generated by go-swagger; DO NOT EDIT.

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "morg/models"
)

// PostArticlesArticleIDCommentsOKCode is the HTTP code returned for type PostArticlesArticleIDCommentsOK
const PostArticlesArticleIDCommentsOKCode int = 200

/*PostArticlesArticleIDCommentsOK OK

swagger:response postArticlesArticleIdCommentsOK
*/
type PostArticlesArticleIDCommentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Comment `json:"body,omitempty"`
}

// NewPostArticlesArticleIDCommentsOK creates PostArticlesArticleIDCommentsOK with default headers values
func NewPostArticlesArticleIDCommentsOK() *PostArticlesArticleIDCommentsOK {

	return &PostArticlesArticleIDCommentsOK{}
}

// WithPayload adds the payload to the post articles article Id comments o k response
func (o *PostArticlesArticleIDCommentsOK) WithPayload(payload *models.Comment) *PostArticlesArticleIDCommentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post articles article Id comments o k response
func (o *PostArticlesArticleIDCommentsOK) SetPayload(payload *models.Comment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostArticlesArticleIDCommentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
