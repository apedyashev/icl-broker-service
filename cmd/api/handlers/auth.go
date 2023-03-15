package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"icl-broker/cmd/api/helpers"
	"net/http"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var requestPayload LoginPayload
	err := helpers.ReadJSON(w, r, &requestPayload)
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}
	jsonData, _ := json.Marshal(requestPayload)

	// call the service
	request, err := http.NewRequest("POST", "http://auth-service/api/auth/login", bytes.NewBuffer(jsonData))
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		helpers.WriteJSONError(w, errors.New("Invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		helpers.WriteJSONError(w, errors.New(fmt.Sprintf("error callling auth service, status = %d", response.StatusCode)))
		return
	}

	// // create variable we will read responseBody into
	// var jsonFromService AuthLoginResponse

	// // decode the response from the auth service
	// err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	// if err != nil {
	// 	helpers.WriteJSONError(w, err)
	// 	return
	// }

	// create a token
	token, err := helpers.GenerateToken(requestPayload.Email)
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}
	jsonResponse := AuthLoginResponse{
		Token: token,
	}

	helpers.WriteJSON(w, http.StatusOK, jsonResponse)
}

type RegisterPayload struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRegisterResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var requestPayload RegisterPayload
	err := helpers.ReadJSON(w, r, &requestPayload)
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}
	jsonData, _ := json.Marshal(requestPayload)

	// call the service
	request, err := http.NewRequest("POST", "http://auth-service/api/auth/register", bytes.NewBuffer(jsonData))
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusCreated {
		helpers.WriteJSONError(w, errors.New(fmt.Sprintf("error callling auth service, status = %d", response.StatusCode)))
		return
	}

	// create variable we will read responseBody into
	var jsonFromService AuthRegisterResponse

	// decode the response from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, jsonFromService)
}
