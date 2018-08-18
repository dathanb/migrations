package api

import (
	"github.com/jamesbibby/swag/swagger"
	"github.com/jamesbibby/swag/endpoint"
	"github.com/udacity/migration-demo/handler"
	"github.com/udacity/migration-demo/utils"
)

func GetUsersEndpoints() []*swagger.Endpoint {
	return []*swagger.Endpoint{
		endpoint.New("put", "/users", "Register user",
			endpoint.Handler(utils.WrapHandler(handler.RegisterUserEndpoint)),
		),
	}
}
