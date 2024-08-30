package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func main() {

}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func getCoordinates(address string) (float64, float64) {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the API key from the environment variable
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		log.Fatalf("API key not found in environment variables")
	}
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.GeocodingRequest{
		Address:  "16 bunyip walkway",
		Language: "en",
		Region:   "au",
	}

	latlng := "-33.732,150.914"

	l := strings.Split(latlng, ",")
	lat, err := strconv.ParseFloat(l[0], 64)
	if err != nil {
		log.Fatalf("Couldn't parse latlng: %#v", err)
	}
	lng, err := strconv.ParseFloat(l[1], 64)
	if err != nil {
		log.Fatalf("Couldn't parse latlng: %#v", err)
	}
	r.LatLng = &maps.LatLng{
		Lat: lat,
		Lng: lng,
	}

	resp, err := c.Geocode(context.Background(), r)
	check(err)

	pretty.Println(resp)
	pretty.Println(resp[0].Geometry.Location.Lat)
	pretty.Println(resp[0].Geometry.Location.Lng)
	return resp[0].Geometry.Location.Lat, resp[0].Geometry.Location.Lng
}
