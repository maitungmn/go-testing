package locations_provider

import (
	"encoding/json"
	"fmt"
	"github.com/maitungmn/go-testing/src/api/clients/restclient"
	"github.com/maitungmn/go-testing/src/api/domain/locations"
	"github.com/maitungmn/go-testing/src/api/utils/errors"
	"io/ioutil"
	"net/http"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%v"
)

var (
	CountryProvider countryServiceInterface = &countryProvider{}
)

type countryServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}
type countryProvider struct{}

func (p *countryProvider) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	url := fmt.Sprintf(urlGetCountry, countryId)
	response, err := restclient.ClientStruct.Get(url, countryId)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(bytes, &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
			}
		}
		return nil, &apiErr
	}

	var result locations.Country
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}
	return &result, nil
}
