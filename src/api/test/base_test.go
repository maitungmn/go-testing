package test

import (
	"github.com/maitungmn/go-testing/src/api/app"
	"net/http"
	"os"
	"testing"
)

var (
	getRequestFunc func(url string) (*http.Response, error)
)

type getClientMock struct{}

func (cm *getClientMock) Get(url string) (*http.Response, error) {
	return getRequestFunc(url)
}

func TestMain(m *testing.M) {
	go app.StartApp()
	os.Exit(m.Run())
}