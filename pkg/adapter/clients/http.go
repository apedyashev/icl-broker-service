package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BeforeDecode func(response *http.Response) error

type GetConfig struct {
	QueryParams  *url.Values
	BeforeDecode BeforeDecode
}

type PostConfig struct {
	Body       []byte
	OnResponse BeforeDecode
}

type PutConfig PostConfig

type HttpClient interface {
	Get(url string, config *GetConfig, responseData any) error
	Put(url string, config *PutConfig, responseData any) error
}

type httpClient struct {
}

func NewHttpClient() HttpClient {
	return &httpClient{}
}

func (c *httpClient) Get(url string, config *GetConfig, responseData any) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	fmt.Println("config.QueryParams", config.QueryParams)
	request.Header.Set("Content-Type", "application/json")
	if config.QueryParams != nil {
		request.URL.RawQuery = config.QueryParams.Encode()
	}
	fmt.Println("Creating client")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("do request error", err)
		return err
	}
	defer response.Body.Close()

	fmt.Println("Before decode")
	err = config.BeforeDecode(response)
	if err != nil {
		return err
	}

	err = json.NewDecoder(response.Body).Decode(responseData)
	if err != nil {
		fmt.Println("Decode error", err)
		return err
	}

	return nil
}

func (c *httpClient) Put(url string, config *PutConfig, responseData any) error {
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(config.Body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("do request error", err)
		return err
	}
	defer response.Body.Close()

	fmt.Println("Before decode")
	err = config.OnResponse(response)
	if err != nil {
		return err
	}

	err = json.NewDecoder(response.Body).Decode(responseData)
	if err != nil {
		fmt.Println("Decode error", err)
		return err
	}

	return nil
}
