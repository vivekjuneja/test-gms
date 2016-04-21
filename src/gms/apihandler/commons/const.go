package commons

import (
	"github.com/bitly/go-simplejson"
)

const (
	HTTP_GET    = "GET"
	HTTP_POST   = "POST"
	HTTP_PUT    = "PUT"
	HTTP_DELETE = "DELETE"
)

type CommonApiResult struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Content simplejson.Json `json:"content"`
}

type CommonMarathonResponseStruct struct {
	Version string `json:"version"`
}
