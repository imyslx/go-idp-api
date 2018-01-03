package controller

import (
	"github.com/goadesign/goa"
	"github.com/imyslx/go-idp-api/app"
)

// HostsController implements the hosts resource.
type HostsController struct {
	*goa.Controller
}

// NewHostsController creates a hosts controller.
func NewHostsController(service *goa.Service) *HostsController {
	return &HostsController{Controller: service.NewController("HostsController")}
}

// List runs the list action.
func (c *HostsController) List(ctx *app.ListHostsContext) error {
	// HostsController_List: start_implement

	// Put your logic here
	mt := &app.Result{}
	mta := []*app.Result{}
	mta = append(mta, mt)
	return ctx.OK(mta)

	// HostsController_List: end_implement
	return nil
}
