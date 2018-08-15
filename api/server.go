package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/jamesbibby/swag/swagger"
	"github.com/jamesbibby/swag"
	log "github.com/sirupsen/logrus"
)


func Serve(port int) {
	router := mux.NewRouter()

	apiHandler := generateApiHandler(router)
	router.Path("/api/v1/swagger").Methods("GET").Handler(apiHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func generateApiHandler(router *mux.Router) http.Handler {
	var endpoints []*swagger.Endpoint

	endpoints = append(endpoints, GetUsersEndpoints()...)

	api := swag.New(
		swag.Title("Q&A"),
		swag.Endpoints(endpoints...),
		swag.Description("A Q&A service"),
		swag.BasePath("/api/v1"),
		swag.ContactEmail("dathan@uacity.com"),
		swag.Tag("Name", "QA Demo"),
	)

	api.Walk(func(path string, ep *swagger.Endpoint) {
		h := ep.Handler.(http.HandlerFunc)
		if log.GetLevel() >= log.DebugLevel {
			log.WithField("OperationID", ep.OperationID).Debug("Adding endpoint to router")
		}
		router.Path(path).Methods(ep.Method).Handler(h)
	})

	return api.Handler(true)
}
