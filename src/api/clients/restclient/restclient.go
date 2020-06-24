package restclient

import (
	"fmt"
	"net/http"
)

var (
	ClientStruct ClientInterface = &clientStruct{}
)

type clientStruct struct{}

type ClientInterface interface {
	Get(string, string) (*http.Response, error)
}

func (ci *clientStruct) Get(url string, countryId string) (*http.Response, error) {
	compiledUrl := fmt.Sprintf(url, countryId)
	request, err := http.NewRequest(http.MethodGet, compiledUrl, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}

	return client.Do(request)
}
