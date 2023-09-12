// this is for geocoding

package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GeoCode(adddress string) (*Coordinate, error) {
	// get Coordinates

	geocodeAPI := os.Getenv("GEOCODE_API")
	geocodeAPIKey := os.Getenv("GEOCODE_API_KEY")

	// turn address into the proper query.

	queryAddress := url.QueryEscape(adddress)

	uri := fmt.Sprintf("%stext=%s&apiKey=%s", geocodeAPI, queryAddress, geocodeAPIKey)

	httpResponse, httpError := http.Get(uri)

	if httpError != nil {
		return nil, httpError
	}

	body, bodyError := ioutil.ReadAll(httpResponse.Body)

	if bodyError != nil {
		return nil, bodyError
	}

	result := fiber.Map{}

	jsonError := json.Unmarshal(body, &result)

	if jsonError != nil {
		return nil, jsonError
	}

	return nil, nil
}
