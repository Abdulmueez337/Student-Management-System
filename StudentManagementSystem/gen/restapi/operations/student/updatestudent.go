// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdatestudentHandlerFunc turns a function with the right signature into a updatestudent handler
type UpdatestudentHandlerFunc func(UpdatestudentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatestudentHandlerFunc) Handle(params UpdatestudentParams) middleware.Responder {
	return fn(params)
}

// UpdatestudentHandler interface for that can handle valid updatestudent params
type UpdatestudentHandler interface {
	Handle(UpdatestudentParams) middleware.Responder
}

// NewUpdatestudent creates a new http.Handler for the updatestudent operation
func NewUpdatestudent(ctx *middleware.Context, handler UpdatestudentHandler) *Updatestudent {
	return &Updatestudent{Context: ctx, Handler: handler}
}

/* Updatestudent swagger:route PUT /Student/updaterecoard/{email} Student updatestudent

Updated student recoard

Enter the valid email to update the student recoard

*/
type Updatestudent struct {
	Context *middleware.Context
	Handler UpdatestudentHandler
}

func (o *Updatestudent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdatestudentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
