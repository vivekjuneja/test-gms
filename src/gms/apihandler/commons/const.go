package commons

import (
	"github.com/bitly/go-simplejson"
)

const (
	// http status GET constant
	HTTP_GET    = "GET"
	// http statu POST constant
	HTTP_POST   = "POST"
	// http status PUT constant
	HTTP_PUT    = "PUT"
	// http status DELETE constant
	HTTP_DELETE = "DELETE"
)

// common api result Struct in GMS
type CommonApiResult struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Content simplejson.Json `json:"content"`
}

// common marathon response Struct in GMS
type CommonMarathonResponseStruct struct {
	Version string `json:"version"`
}
