// Code generated by go-swagger; DO NOT EDIT.

package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "morg/models"
)

// GetProfilesUsernameOKCode is the HTTP code returned for type GetProfilesUsernameOK
const GetProfilesUsernameOKCode int = 200

/*GetProfilesUsernameOK OK

swagger:response getProfilesUsernameOK
*/
type GetProfilesUsernameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewGetProfilesUsernameOK creates GetProfilesUsernameOK with default headers values
func NewGetProfilesUsernameOK() *GetProfilesUsernameOK {

	return &GetProfilesUsernameOK{}
}

// WithPayload adds the payload to the get profiles username o k response
func (o *GetProfilesUsernameOK) WithPayload(payload *models.Profile) *GetProfilesUsernameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profiles username o k response
func (o *GetProfilesUsernameOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfilesUsernameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
