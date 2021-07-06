// Client for the \R\a\d\a/r/b/o/x/ API. Get planes in the area
// and metadata about each aircraft.
package radar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Extracted flight metadata from the given live flight.
type AircraftData struct {
	La  float64 `json:"la"`  // current latitude
	Lo  float64 `json:"lo"`  // current longitude
	Alt uint    `json:"alt"` // current altitude
	GS  float64 `json:"gs"`  // ground speed
	Act string  `json:"act"` // model name shorthand(e.g. C172)
	Acd string  `json:"acd"` // model name full(e.g. Cessna 172N Skyhawk 100)
}

// Lookup a flight by the fid. Return various metadata about the live flight.
func (rb RadarBounds) FlightInfo(fid string) (*AircraftData, error) {
	queryArgs := url.Values{}
	queryArgs.Add("fid", fid)

	scannerEndpoint := fmt.Sprintf("%s/%s?%s", apiDomainName, dataRoute, queryArgs.Encode())
	req, err := http.NewRequest("GET", scannerEndpoint, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else if string(body) == "null" { // check if the fid is valid
		return nil, fmt.Errorf("no live flight exists for fid \"%s\"", fid)
	}

	var metadata AircraftData
	_ = json.Unmarshal(body, &metadata)

	return &metadata, nil
}
