// Client for the Radarbox API. Get planes in the area
// and metadata about each aircraft.
package radar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Unique idenifier for each aircraft.
type FlightID struct {
	id uint
}

// Radar box from which aircraft should be tracked. Specifying the top-left
// and bottom right cord of this box, all aircraft that are inside of it will be
// observed and tracked.
type RadarBounds struct {
	LatTR, LongTR float64
	LatBR, LongBR float64
}

// Static headers that are required for the API to function properly.
var headers http.Header = http.Header(map[string][]string{
	"User-Agent":      {"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:88.0) Gecko/20100101 Firefox/88.0"},
	"Accept":          {"application/json, text/plain, */*"},
	"Accept-Language": {"en-US,en;q=0.5"},
	"Origin":          {domainName},
	"DNT":             {"1"},
	"Connection":      {"keep-alive"},
	"Referer":         {fmt.Sprintf("%s/", domainName)},
	"TE":              {"Trailers"},
})

// Captures all of the aircraft within the specified bounds. Returns a list of
// flight ids of all of the aircraft in the area.
func (rb RadarBounds) DetectFlights() (*[]FlightID, error) {
	queryArgs := rb.detectUrlQuery()
	endpoint := fmt.Sprintf("%s/%s?%s", apiDomainName, scannerRoute, queryArgs.Encode())
	req, err := http.NewRequest("GET", endpoint, nil)
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
	}

	flightIDs := rb.parseFlightIDs(&body)

	return flightIDs, nil
}

// Builds the url query for a aircraft detection requests. Sets the bounds and
// the filter metadata of the request.
func (rb RadarBounds) detectUrlQuery() url.Values {
	currentTime := int(time.Now().Unix())
	queryArgs := url.Values(map[string][]string{
		"ff":             {"false"},
		"designator":     {"iata"},
		"showLastTrails": {"true"},
		"timestamp":      {strconv.Itoa(currentTime)},
		"os":             {"web"},
		"adsb":           {"true"},
		"adsbsat":        {"true"},
		"asdi":           {"true"},
		"mlat":           {"true"},
		"sate":           {"true"},
		"uat":            {"true"},
		"hfdl":           {"true"},
		"sti":            {"true"},
		"asdex":          {"true"},
		"flarm":          {"true"},
		"onair":          {"true"},
		"class[]":        {"?", "A", "B", "C", "G", "H", "M"},
		"diverted":       {"false"},
		"delayed":        {"false"},
		"isga":           {"false"},
		"ground":         {"false"},
		"blocked":        {"false"},
	})

	bounds := fmt.Sprintf("%.3f,%.3f,%.3f,%.3f", rb.LatTR, rb.LongTR, rb.LatBR, rb.LongBR)
	queryArgs.Add("bounds", bounds)

	return queryArgs
}

// Parses the result of the detection event and returns the aircraft flight ids.
func (rb RadarBounds) parseFlightIDs(body *[]byte) *[]FlightID {
	flightIDs := []FlightID{}
	var radarData []map[uint][]interface{}
	_ = json.Unmarshal(*body, &radarData)

	for _, plane := range radarData {
		for fid := range plane {
			flightIDs = append(flightIDs, FlightID{fid})
		}
	}

	return &flightIDs
}
