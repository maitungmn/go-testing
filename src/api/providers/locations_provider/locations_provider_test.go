package locations_provider

import (
	"github.com/maitungmn/go-testing/src/api/clients/restclient"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

var (
	getRequestFunc func(url string) (*http.Response, error)
)

type getClientMock struct{}

func (cm *getClientMock) Get(url string) (*http.Response, error) {
	return getRequestFunc(url)
}

func TestGetCountryNotFound(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(strings.NewReader(LocationsApiNotFound)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	country, err := CountryProvider.GetCountry("AR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	invalidCloser, _ := os.Open("-asf3")
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       invalidCloser,
		}, nil
	}
	restclient.ClientStruct = &getClientMock{}

	country, err := CountryProvider.GetCountry("AR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country AR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(LocationsApiInvalidJsonResponse)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	country, err := CountryProvider.GetCountry("AR")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(LocationsApiResponse)),
		}, nil
	}
	restclient.ClientStruct = &getClientMock{} //without this line, the real api is fired

	country, err := CountryProvider.GetCountry("AR")

	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.States))
}