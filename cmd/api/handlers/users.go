package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"icl-broker/cmd/api/helpers"
	"net/http"

	"github.com/go-chi/jwtauth"
)

type GetCurrentUserPayload struct {
	Email string `json:"email"`
}

type GetCurrentUserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	// call the service
	getUserUrl := fmt.Sprintf("http://auth-service/api/user/%s", claims["email"])
	response, err := http.Get(getUserUrl)
	if err != nil {
		helpers.WriteJSONError(w, errors.New(fmt.Sprintf("error callling auth service, status = %d", response.StatusCode)))
		return
	}

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusOK {
		helpers.WriteJSONError(w, errors.New(fmt.Sprintf("error callling auth service, status = %d", response.StatusCode)))
		return
	}

	// create variable we will read responseBody into
	var jsonFromService GetCurrentUserResponse

	// decode the response from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, jsonFromService)
}
