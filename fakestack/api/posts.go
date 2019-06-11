package api

import (
	"github.com/jamesbibby/swag/swagger"
	"github.com/jamesbibby/swag/endpoint"
	"github.com/dathanb/migrations/fakestack/handler"
	"github.com/dathanb/migrations/fakestack/utils"
)

func GetPostsEndpoints() []*swagger.Endpoint {
	return []*swagger.Endpoint{
		endpoint.New("put", "/posts", "Create or update post",
			endpoint.Handler(utils.WrapHandler(handler.CreatePostEndpoint)),
		),
	}
}
