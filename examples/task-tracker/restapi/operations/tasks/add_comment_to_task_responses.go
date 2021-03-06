package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-swagger/go-swagger/examples/task-tracker/models"
)

/*AddCommentToTaskCreated Comment added

swagger:response addCommentToTaskCreated
*/
type AddCommentToTaskCreated struct {
}

// NewAddCommentToTaskCreated creates AddCommentToTaskCreated with default headers values
func NewAddCommentToTaskCreated() *AddCommentToTaskCreated {
	return &AddCommentToTaskCreated{}
}

// WriteResponse to the client
func (o *AddCommentToTaskCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
}

/*AddCommentToTaskDefault Error response

swagger:response addCommentToTaskDefault
*/
type AddCommentToTaskDefault struct {
	_statusCode int
	/*
	  Required: true
	*/
	XErrorCode string `json:"X-Error-Code"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddCommentToTaskDefault creates AddCommentToTaskDefault with default headers values
func NewAddCommentToTaskDefault(code int) *AddCommentToTaskDefault {
	if code <= 0 {
		code = 500
	}

	return &AddCommentToTaskDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add comment to task default response
func (o *AddCommentToTaskDefault) WithStatusCode(code int) *AddCommentToTaskDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add comment to task default response
func (o *AddCommentToTaskDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXErrorCode adds the xErrorCode to the add comment to task default response
func (o *AddCommentToTaskDefault) WithXErrorCode(xErrorCode string) *AddCommentToTaskDefault {
	o.XErrorCode = xErrorCode
	return o
}

// SetXErrorCode sets the xErrorCode to the add comment to task default response
func (o *AddCommentToTaskDefault) SetXErrorCode(xErrorCode string) {
	o.XErrorCode = xErrorCode
}

// WithPayload adds the payload to the add comment to task default response
func (o *AddCommentToTaskDefault) WithPayload(payload *models.Error) *AddCommentToTaskDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add comment to task default response
func (o *AddCommentToTaskDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddCommentToTaskDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Error-Code

	xErrorCode := o.XErrorCode
	if xErrorCode != "" {
		rw.Header().Set("X-Error-Code", xErrorCode)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
