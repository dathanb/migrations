package api

import (
	"github.com/jamesbibby/swag/swagger"
	"github.com/jamesbibby/swag/endpoint"
	"github.com/udacity/migration-demo/utils"
	"github.com/udacity/migration-demo/handler"
)

func GetProductsEndpoint() []*swagger.Endpoint {
	return []*swagger.Endpoint{
		endpoint.New("get", "/products", "Search products",
			endpoint.Handler(utils.WrapHandler(handler.SearchProductsEndpoint)),
		),
	}
}
