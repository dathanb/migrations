package handler

import (
	"net/http"
	"github.com/udacity/go-errors"
	"encoding/json"
	"io/ioutil"
	"github.com/udacity/migration-demo/db/dal"
)

type CreateUserRequest struct {
	Id int `json:"id"`
	DisplayName string `json:"display_name"`
}

func RegisterUserEndpoint(request *http.Request, vars map[string]string) ([]byte, int, error) {
	var reqObject CreateUserRequest

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.HTTPError, err)
	}
	err = json.Unmarshal(body, &reqObject)
	if err != nil {
		return nil, 0, errors.WithRootCause(errors.HTTPBadRequestError, err)
	}

	// TODO: insert into DB
	dal.NewUsersDAL()

	return nil, 0, errors.NotImplementedError
}
