// Code generated by go-swagger; DO NOT EDIT.

package articles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "morg/models"
)

// GetFeedsFeedIDArticlesOKCode is the HTTP code returned for type GetFeedsFeedIDArticlesOK
const GetFeedsFeedIDArticlesOKCode int = 200

/*GetFeedsFeedIDArticlesOK OK

swagger:response getFeedsFeedIdArticlesOK
*/
type GetFeedsFeedIDArticlesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Article `json:"body,omitempty"`
}

// NewGetFeedsFeedIDArticlesOK creates GetFeedsFeedIDArticlesOK with default headers values
func NewGetFeedsFeedIDArticlesOK() *GetFeedsFeedIDArticlesOK {

	return &GetFeedsFeedIDArticlesOK{}
}

// WithPayload adds the payload to the get feeds feed Id articles o k response
func (o *GetFeedsFeedIDArticlesOK) WithPayload(payload []*models.Article) *GetFeedsFeedIDArticlesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get feeds feed Id articles o k response
func (o *GetFeedsFeedIDArticlesOK) SetPayload(payload []*models.Article) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFeedsFeedIDArticlesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Article, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
