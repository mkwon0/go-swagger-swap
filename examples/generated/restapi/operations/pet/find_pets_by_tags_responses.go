package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-swagger/go-swagger/examples/generated/models"
)

/*FindPetsByTagsOK successful operation

swagger:response findPetsByTagsOK
*/
type FindPetsByTagsOK struct {

	// In: body
	Payload []*models.Pet `json:"body,omitempty"`
}

// NewFindPetsByTagsOK creates FindPetsByTagsOK with default headers values
func NewFindPetsByTagsOK() *FindPetsByTagsOK {
	return &FindPetsByTagsOK{}
}

// WithPayload adds the payload to the find pets by tags o k response
func (o *FindPetsByTagsOK) WithPayload(payload []*models.Pet) *FindPetsByTagsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pets by tags o k response
func (o *FindPetsByTagsOK) SetPayload(payload []*models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPetsByTagsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindPetsByTagsBadRequest Invalid tag value

swagger:response findPetsByTagsBadRequest
*/
type FindPetsByTagsBadRequest struct {
}

// NewFindPetsByTagsBadRequest creates FindPetsByTagsBadRequest with default headers values
func NewFindPetsByTagsBadRequest() *FindPetsByTagsBadRequest {
	return &FindPetsByTagsBadRequest{}
}

// WriteResponse to the client
func (o *FindPetsByTagsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}