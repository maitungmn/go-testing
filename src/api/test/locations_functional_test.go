package test

import (
	"encoding/json"
	"fmt"
	"github.com/maitungmn/go-testing/src/api/clients/restclient"
	"github.com/maitungmn/go-testing/src/api/providers/locations_provider"
	"github.com/maitungmn/go-testing/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetCountriesNotFound(t *testing.T) {
	fmt.Println("about to get functional test get countries")

	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(strings.NewReader(locations_provider.LocationsApiNotFound)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	response, err := http.Get("http://localhost:8080/location/countries/AR")

	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountriesNoError(t *testing.T) {
	// TODO: Implement later
}