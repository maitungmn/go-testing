package app

import "github.com/maitungmn/go-testing/src/api/controllers"

func mapUrls() {
	router.GET("/location/countries/:country_id", controllers.GetCountry)
}
