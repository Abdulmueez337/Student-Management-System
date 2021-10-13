// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetstudentByNameHandlerFunc turns a function with the right signature into a getstudent by name handler
type GetstudentByNameHandlerFunc func(GetstudentByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetstudentByNameHandlerFunc) Handle(params GetstudentByNameParams) middleware.Responder {
	return fn(params)
}

// GetstudentByNameHandler interface for that can handle valid getstudent by name params
type GetstudentByNameHandler interface {
	Handle(GetstudentByNameParams) middleware.Responder
}

// NewGetstudentByName creates a new http.Handler for the getstudent by name operation
func NewGetstudentByName(ctx *middleware.Context, handler GetstudentByNameHandler) *GetstudentByName {
	return &GetstudentByName{Context: ctx, Handler: handler}
}

/* GetstudentByName swagger:route GET /Student/getrecoard/{email} Student getstudentByName

Get student recoard

*/
type GetstudentByName struct {
	Context *middleware.Context
	Handler GetstudentByNameHandler
}

func (o *GetstudentByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetstudentByNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
