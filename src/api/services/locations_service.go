package services

import (
	"fmt"
	"github.com/maitungmn/go-testing/src/api/domain/locations"
	"github.com/maitungmn/go-testing/src/api/providers/locations_provider"
	"github.com/maitungmn/go-testing/src/api/utils/errors"
)

type locationsService struct {}

type locationServiceInterface interface {
	GetCountry(string)(*locations.Country, *errors.ApiError)
}

var (
	LocationsService locationServiceInterface
)

func init() {
	fmt.Println("Init Locations Service")
	LocationsService = &locationsService{}
}

func (s *locationsService)GetCountry(countryId string)(*locations.Country, *errors.ApiError) {
	fmt.Println("inside service")
	return locations_provider.CountryProvider.GetCountry(countryId)
}
