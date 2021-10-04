package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ochom/focus/domain"
)

//MakeHttpRequest makes the http request
func MakeHttpRequest(params domain.HttpRequestParam) ([]byte, error) {
	//parse payload to bytes
	reqBody, err := json.Marshal(params.Payload)
	if err != nil {
		return nil, err
	}

	//create a new request
	req, err := http.NewRequest(params.Method, params.URL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	//add request headers
	for k, v := range params.Headers {
		req.Header.Add(k, v)
	}

	//init http client
	client := http.DefaultClient

	//make request
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	//check if response status is failure
	if res.StatusCode > 300 {
		return nil, fmt.Errorf("request failed: %v", res.Status)
	}

	//parse response body to bytes array
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetEnv(envName string) string {
	env := os.Getenv(envName)
	if env == "" {
		log.Fatalf("environment variables: %v not declared", envName)
		os.Exit(1)
	}

	return env
}
