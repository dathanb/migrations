package handler

import (
	"net/http"
	"github.com/udacity/go-errors"
	"encoding/json"
	"io/ioutil"
	"github.com/udacity/migration-demo/db"
	"github.com/sirupsen/logrus"
)

type CreateUserRequest struct {
	Id int `json:"id"`
	DisplayName string `json:"display_name"`
}

func RegisterUserEndpoint(request *http.Request, vars map[string]string) ([]byte, int, error) {
	var reqObject CreateUserRequest
	logrus.Debug("Handling create user request to %v", request.URL)

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.HTTPError, err)
	}
	err = json.Unmarshal(body, &reqObject)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.HTTPBadRequestError, err)
	}

	user, err := db.ApplicationDAL().Users().UpsertUser(request.Context(), reqObject.Id, reqObject.DisplayName)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.SQLCommitError, err)
	}

	data, err := json.Marshal(user)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.JSONMarshalingError, err)
	}

	return data, 200, nil
}
