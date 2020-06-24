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
	Get(string) (*http.Response, error)
}

func (ci *clientStruct) Get(url string) (*http.Response, error) {
	compiledUrl := fmt.Sprintf(url)
	request, err := http.NewRequest(http.MethodGet, compiledUrl, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}

	return client.Do(request)
}
