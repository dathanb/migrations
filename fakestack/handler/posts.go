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

type CreatePostRequest struct {
	Id       int    `json:"id"`
	PostType int    `json:"post_type_id"`
	UserId   int    `json:"user_id"`
	Body     string `json:"body"`
}

var (
	putPostRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fakestack_putPosts_requests_total",
		Help: "The total number of Posts sent to the service",
	})
)

func CreatePostEndpoint(request *http.Request, vars map[string]string) ([]byte, int, error) {
	var reqObject CreatePostRequest
	logrus.Debug("Handling create post request to %v", request.URL)

	putPostRequests.Inc()

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, 0, merry.Wrap(err)
	}

	err = json.Unmarshal(body, &reqObject)
	if err != nil {
		return nil, 0, merry.Wrap(err)
	}

	post, err := db.ApplicationDAL().Posts().UpsertPost(request.Context(), reqObject.Id, reqObject.PostType,
		reqObject.UserId, reqObject.Body)
	if err != nil {
		return nil, 0, merry.Wrap(err)
	}

	data, err := json.Marshal(post)
	if err != nil {
		return nil, 0, merry.Wrap(err)
	}

	return data, 200, nil
}
