package handler

import (
	"net/http"
	"github.com/udacity/go-errors"
	"encoding/json"
	"io/ioutil"
	"github.com/dathanb/fakestack/db"
	"github.com/sirupsen/logrus"
)

type CreatePostRequest struct {
	Id       int    `json:"id"`
	PostType int    `json:"post_type_id"`
	UserId   int    `json:"user_id"`
	Body     string `json:"body"`
}

func CreatePostEndpoint(request *http.Request, vars map[string]string) ([]byte, int, error) {
	var reqObject CreatePostRequest
	logrus.Debug("Handling create post request to %v", request.URL)

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.HTTPError, err)
	}

	err = json.Unmarshal(body, &reqObject)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.HTTPBadRequestError, err)
	}

	post, err := db.ApplicationDAL().Posts().UpsertPost(request.Context(), reqObject.Id, reqObject.PostType,
		reqObject.UserId, reqObject.Body)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.SQLCommitError, err)
	}

	data, err := json.Marshal(post)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.JSONMarshalingError, err)
	}

	return data, 200, nil
}
