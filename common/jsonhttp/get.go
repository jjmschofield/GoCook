package jsonhttp

import (
	"encoding/json"
	"github.com/jjmschofield/GoCook/common/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string, model interface{}) error {
	response, responseError := makeGetRequest(url)

	if responseError != nil {
		return responseError
	}

	body, bodyError := readResponseBody(response)

	if bodyError != nil {
		return bodyError
	}

	bindError := bindResponseBody(body, model)

	if bindError != nil {
		return bindError
	}

	return nil
}

func makeGetRequest(url string) (response *http.Response, error error) {
	var httpClient = &http.Client{ // The default http client does not set a sensible timeout - see https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
		Timeout: time.Second * 10,
	}

	response, responseError := httpClient.Get(url)

	if responseError != nil {
		logger.Error("HTTP Request failed", zap.Error(responseError))
	}

	return response, responseError
}

func readResponseBody(response *http.Response) ([]byte, error) {
	body, bodyError := ioutil.ReadAll(response.Body)

	if bodyError != nil {
		logger.Error("Reading body of HTTP response failed", zap.Error(bodyError))
	}

	return body, bodyError
}

func bindResponseBody(bodyBytes []byte, model interface{}) error {
	jsonErr := json.Unmarshal(bodyBytes, model)

	if jsonErr != nil {
		logger.Error("Binding JSON in body of HTTP response failed", zap.Error(jsonErr))
	}

	return jsonErr
}
