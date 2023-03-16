package handlers

import (
	"errors"
	"fmt"
	"icl-broker/cmd/api/helpers"
	"log"
	"net/http"
	"net/rpc"
)

type RPCPayload struct {
	PostId    uint
	ImageBody string
}

type RPCResponse struct {
	Id    string
	Error bool
}

type uploadImagePayload struct {
	PostId  uint   `json:"postId"`
	Content string `json:"content"`
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	var requestPayload uploadImagePayload
	err := helpers.ReadJSON(w, r, &requestPayload)
	if err != nil {
		helpers.WriteJSONError(w, err)
		return
	}

	loggerRpcClient, err := rpc.Dial("tcp", "images-service:5001")
	if err != nil {
		helpers.WriteJSONError(w, errors.New("Error when dialing rpc service"))
		return
	}

	rpcPayload := RPCPayload{
		PostId:    requestPayload.PostId,
		ImageBody: requestPayload.Content,
	}
	log.Println("Log via RPC", rpcPayload)

	var result RPCResponse
	// RPCServer - is a struct created in the logger service
	// LogInfo MUST start with a capital letter (i.e it must me exported)
	err = loggerRpcClient.Call("RPCServer.SaveImage", rpcPayload, &result)
	if err != nil {
		fmt.Println("error calling RPCServer.SaveImage", err)
		helpers.WriteJSONError(w, errors.New("error calling RPCServer.SaveImage"), http.StatusInternalServerError)
	}

	payload := helpers.JsonResponse{
		Error:   result.Error,
		Message: result.Id,
	}

	helpers.WriteJSON(w, http.StatusAccepted, payload)
}
