package radar

import "encoding/base64"

// Base64 encoded routes to hide domain & endpoint sources.
var domainName = "aHR0cHM6Ly93d3cucmFkYXJib3guY29t"
var apiDomainName = "aHR0cHM6Ly9kYXRhLnJiMjQuY29t"
var scannerRoute = "bGl2ZQ=="
var dataRoute = "bGl2ZS1mbGlnaHQtaW5mbw=="

// Decode the encoded domain and route names such that they
// could be used for getting aircraft info.
func init() {
	decodeAPIs := []*string{&domainName, &apiDomainName, &scannerRoute, &dataRoute}

	for _, str := range decodeAPIs {
		decodedStr, _ := base64.StdEncoding.DecodeString(*str)
		*str = string(decodedStr)
	}
}
