package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-swagger/go-swagger/examples/task-tracker/models"
)

/*GetTaskDetailsOK Task details

swagger:response getTaskDetailsOK
*/
type GetTaskDetailsOK struct {

	// In: body
	Payload *models.Task `json:"body,omitempty"`
}

// NewGetTaskDetailsOK creates GetTaskDetailsOK with default headers values
func NewGetTaskDetailsOK() *GetTaskDetailsOK {
	return &GetTaskDetailsOK{}
}

// WithPayload adds the payload to the get task details o k response
func (o *GetTaskDetailsOK) WithPayload(payload *models.Task) *GetTaskDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task details o k response
func (o *GetTaskDetailsOK) SetPayload(payload *models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTaskDetailsUnprocessableEntity Validation error

swagger:response getTaskDetailsUnprocessableEntity
*/
type GetTaskDetailsUnprocessableEntity struct {

	// In: body
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewGetTaskDetailsUnprocessableEntity creates GetTaskDetailsUnprocessableEntity with default headers values
func NewGetTaskDetailsUnprocessableEntity() *GetTaskDetailsUnprocessableEntity {
	return &GetTaskDetailsUnprocessableEntity{}
}

// WithPayload adds the payload to the get task details unprocessable entity response
func (o *GetTaskDetailsUnprocessableEntity) WithPayload(payload *models.ValidationError) *GetTaskDetailsUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task details unprocessable entity response
func (o *GetTaskDetailsUnprocessableEntity) SetPayload(payload *models.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskDetailsUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTaskDetailsDefault Error response

swagger:response getTaskDetailsDefault
*/
type GetTaskDetailsDefault struct {
	_statusCode int
	/*
	  Required: true
	*/
	XErrorCode string `json:"X-Error-Code"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTaskDetailsDefault creates GetTaskDetailsDefault with default headers values
func NewGetTaskDetailsDefault(code int) *GetTaskDetailsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTaskDetailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get task details default response
func (o *GetTaskDetailsDefault) WithStatusCode(code int) *GetTaskDetailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get task details default response
func (o *GetTaskDetailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXErrorCode adds the xErrorCode to the get task details default response
func (o *GetTaskDetailsDefault) WithXErrorCode(xErrorCode string) *GetTaskDetailsDefault {
	o.XErrorCode = xErrorCode
	return o
}

// SetXErrorCode sets the xErrorCode to the get task details default response
func (o *GetTaskDetailsDefault) SetXErrorCode(xErrorCode string) {
	o.XErrorCode = xErrorCode
}

// WithPayload adds the payload to the get task details default response
func (o *GetTaskDetailsDefault) WithPayload(payload *models.Error) *GetTaskDetailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task details default response
func (o *GetTaskDetailsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskDetailsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Error-Code
	rw.Header().Add("X-Error-Code", fmt.Sprintf("%v", o.XErrorCode))

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
