package handler

import (
	"encoding/json"
	"github.com/ansel1/merry"
	"github.com/dathanb/migrations/fakestack/db"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type CreateUserRequest struct {
	Id          int    `json:"id"`
	DisplayName string `json:"display_name"`
}

var (
	putUserRequestCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fakestack_putUser_request_total",
		Help: "The total number of Users sent to the service",
	})
	putUserRequestTiming = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "fakestack_putUser_request_timing",
		Help: "the amount of time a request to create a new User takes",
	})
)

func RegisterUserEndpoint(request *http.Request, vars map[string]string) ([]byte, int, error) {
	var reqObject CreateUserRequest
	logrus.Debug("Handling create user request to %v", request.URL)

	putUserRequestCount.Inc()
	timer := prometheus.NewTimer(putUserRequestTiming)
	defer timer.ObserveDuration()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, 0, merry.WithHTTPCode(err, 500)
	}

	err = json.Unmarshal(body, &reqObject)
	if err != nil {
		return nil, 0, merry.WithHTTPCode(err, 400)
	}

	user, err := db.ApplicationDAL().Users().UpsertUser(request.Context(), reqObject.Id, reqObject.DisplayName)
	if err != nil {
		return nil, 0, merry.WithHTTPCode(err, 500)
	}

	data, err := json.Marshal(user)
	if err != nil {
		return nil, 0, merry.WithHTTPCode(err, 500)
	}

	return data, 200, nil
}
