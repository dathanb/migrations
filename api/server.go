package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/jamesbibby/swag/swagger"
	"github.com/jamesbibby/swag"
)


func Serve(port int) {
	router := mux.NewRouter()

	apiHandler := generateApiHandler(router)

	http.ListenAndServe(fmt.Sprintf(":%d", port), apiHandler)
}

func generateApiHandler(router *mux.Router) http.Handler {
	var endpoints []*swagger.Endpoint

	endpoints = append(endpoints, GetUsersEndpoints()...)

	api := swag.New(
		swag.Title("Q&A"),
		swag.Endpoints(endpoints...),
		swag.Description("A Q&A service"),
		swag.BasePath("/api"),
		swag.ContactEmail("dathan@uacity.com"),
		swag.Tag("Name", "QA Demo"),
	)

	api.Walk(func(path string, ep *swagger.Endpoint) {
		router.Path(path).Methods(ep.Method).Handler(ep.Handler.(http.HandlerFunc))
	})

	return api.Handler(true)
}
