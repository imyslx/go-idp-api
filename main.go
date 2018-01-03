//go:generate goagen bootstrap -d github.com/imyslx/go-idp-api/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/imyslx/go-idp-api/app"
	. "github.com/imyslx/go-idp-api/controller"
)

func main() {
	// Create service
	service := goa.New("imyslx/go-idp-api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "hosts" controller
	c := NewHostsController(service)
	app.MountHostsController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)
	// Mount "swaggerui" controller
	c3 := NewSwaggeruiController(service)
	app.MountSwaggeruiController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
