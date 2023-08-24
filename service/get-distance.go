package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func IsTravelModeValid(mode string) bool {
	ValidTravelModes := []string{"driving", "walking", "public transit"}

	for _, validMode := range ValidTravelModes {
		if validMode == mode {
			return true
		}
	}

	return false
}

type Coordinate struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type BingAPIResponse struct {
	AuthenticationResultCode string `json:"authenticationResultCode"`
	BrandLogoURI             string `json:"brandLogoUri"`
	Copyright                string `json:"copyright"`
	ResourceSets             []struct {
		EstimatedTotal int `json:"estimatedTotal"`
		Resources      []struct {
			Type         string `json:"__type"`
			Destinations []struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"destinations"`
			Origins []struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"origins"`
			Results []struct {
				DestinationIndex  int     `json:"destinationIndex"`
				OriginIndex       int     `json:"originIndex"`
				TotalWalkDuration int     `json:"totalWalkDuration"`
				TravelDistance    float64 `json:"travelDistance"`
				TravelDuration    float64 `json:"travelDuration"`
			} `json:"results"`
		} `json:"resources"`
	} `json:"resourceSets"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	TraceID           string `json:"traceId"`
}

func GetDistance(origin *Coordinate, destination *Coordinate, travelMode string) (*fiber.Map, error) {
	// things to get started
	apiKey := os.Getenv("BING_KEY")
	originCoords := FormatCoordinates(origin)
	destinationCoords := FormatCoordinates(destination)
	url := os.Getenv("BING_URL")

	// output variable
	var result BingAPIResponse

	uri := fmt.Sprintf("%s?origins=%s&destinations=%s&travelMode=%s&key=%s", url, originCoords, destinationCoords, travelMode, apiKey)

	httpResponse, httpError := http.Get(uri)

	if httpError != nil {
		return nil, httpError
	}

	body, bodyError := ioutil.ReadAll(httpResponse.Body)

	if bodyError != nil {
		return nil, bodyError
	}

	jsonError := json.Unmarshal(body, &result)

	if jsonError != nil {
		return nil, jsonError
	}

	return &fiber.Map{
		"travel":      result.ResourceSets[0].Resources[0].Results,
		"origin":      result.ResourceSets[0].Resources[0].Origins,
		"destination": result.ResourceSets[0].Resources[0].Destinations,
	}, nil

}

func FormatCoordinates(coordinate *Coordinate) string {
	return fmt.Sprintf("%s,%s", coordinate.Lat, coordinate.Long)
}
