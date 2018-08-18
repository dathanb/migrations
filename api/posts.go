package api

import (
	"github.com/jamesbibby/swag/swagger"
	"github.com/jamesbibby/swag/endpoint"
	"github.com/udacity/migration-demo/handler"
	"github.com/udacity/migration-demo/utils"
)

func GetPostsEndpoints() []*swagger.Endpoint {
	return []*swagger.Endpoint{
		endpoint.New("put", "/posts", "Create or update post",
			endpoint.Handler(utils.WrapHandler(handler.CreatePostEndpoint)),
		),
	}
}
