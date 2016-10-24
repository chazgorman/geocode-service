package service

import (
	"encoding/json"
	"errors"
	"github.com/grindhold/gominatim"
)

// GeocodeService provides operations on strings.
type GeocodeService interface {
	Geocode(string) (string, error)
	Count(string) int
}

type geocodeService struct{}

type Location struct {
	Placename string
	Latitude  string `json:float64`
	Longitude string `json:float64`
}

func (geocodeService) Geocode(placename string) (string, error) {
	respJsonString := ""
	if placename == "" {
		return respJsonString, ErrEmpty
	}
	gominatim.SetServer("http://nominatim.openstreetmap.org/")

	//Get by a Querystring
	qry := gominatim.SearchQuery{
		Q: placename,
	}
	resp, _ := qry.Get() // Returns []gominatim.Result

	location := Location{resp[0].DisplayName,resp[0].Lat,resp[0].Lon}
	respJson, err := json.Marshal(location)

	if(err != nil){
		respJsonString = "Unable to process geocode request"
	} else {
		respJsonString = string(respJson)
	}

	return respJsonString, nil
}

func (geocodeService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
