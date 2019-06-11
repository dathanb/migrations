package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/jamesbibby/swag/swagger"
	"github.com/jamesbibby/swag"
	log "github.com/sirupsen/logrus"
	"net/http/httptest"
	"net/url"
)


func Serve(port int) {
	router := mux.NewRouter()

	apiHandler := generateApiHandler(router)
	router.Path("/api/v1/swagger").Methods("GET").Handler(apiHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func SwaggerApi() *swagger.API {
	var endpoints []*swagger.Endpoint

	endpoints = append(endpoints, GetUsersEndpoints()...)
	endpoints = append(endpoints, GetPostsEndpoints()...)

	return swag.New(
		swag.Title("Q&A"),
		swag.Endpoints(endpoints...),
		swag.Description("A Q&A service"),
		swag.BasePath("/api/v1"),
		swag.ContactEmail("dathan@gmail.com"),
		swag.Tag("Name", "QA Demo"),
	)
}

func SwaggerDefinition(api *swagger.API) (string, error) {
	handler := api.Handler(true)
	swaggerUrl := url.URL{
		Host:"localhost",
		Path:"/",
		Scheme:"http",
	}

	request := http.Request{
		Method:"GET",
		URL: &swaggerUrl,
		Host:"localhost",
	}
	responseWriter := httptest.NewRecorder()
	handler(responseWriter, &request)

	return responseWriter.Body.String(), nil
}

func generateApiHandler(router *mux.Router) http.Handler {
	api := SwaggerApi()

	api.Walk(func(path string, ep *swagger.Endpoint) {
		h := ep.Handler.(http.HandlerFunc)
		if log.GetLevel() >= log.DebugLevel {
			log.WithField("OperationID", ep.OperationID).Debug("Adding endpoint to router")
		}
		router.Path(path).Methods(ep.Method).Handler(h)
	})

	return api.Handler(true)
}
