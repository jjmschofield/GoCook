package jsonHttp

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

func Get(url string, model interface{}) (error) {

	response, responseError := makeGetRequest(url)

	if(responseError != nil){
		return responseError
	}

	body, bodyError := readResponseBody(response)

	if bodyError != nil {
		return bodyError
	}

	bindError := bindResponseBody(body, model)

	if(bindError != nil){
		return bindError
	}

	return nil
}

func makeGetRequest(url string) (response *http.Response, error error){
	response, responseError := http.Get(url)

	if(responseError != nil){
		log.Fatal(responseError)
	}

	return response, responseError
}

func readResponseBody(response *http.Response) ([]byte, error){
	body, bodyError := ioutil.ReadAll(response.Body)

	if bodyError != nil {
		log.Fatal(bodyError)
	}

	return body, bodyError
}

func bindResponseBody(bodyBytes []byte, model interface{}) (error){
	jsonErr := json.Unmarshal(bodyBytes, model)

	if(jsonErr != nil){
		log.Fatal(jsonErr)
	}

	return jsonErr;
}