package api

import (
	"github.com/dathanb/migrations/fakestack/handler"
	"github.com/dathanb/migrations/fakestack/utils"
	"github.com/jamesbibby/swag/endpoint"
	"github.com/jamesbibby/swag/swagger"
)

func GetPostsEndpoints() []*swagger.Endpoint {
	return []*swagger.Endpoint{
		endpoint.New("put", "/posts", "Create or update post",
			endpoint.Handler(utils.WrapHandler(handler.CreatePostEndpoint)),
		),
	}
}
