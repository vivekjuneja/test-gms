package mesos

import (
	"encoding/json"
	"fmt"
	"gms/apihandler/commons"
	"gms/apihandler/mesos/domain"

	"github.com/parnurzeal/gorequest"
)

const (
	MESOS_HOST = "http://10.53.15.219:15050"
)

/**
 * get groups
 */
func GetMasterState() commons.CommonApiResult {
	return commonHandler(commons.HTTP_GET, "/master/state.json", domain.MasterState{})
}

func commonHandler(method string, uri string, jsonmappingobj interface{}) commons.CommonApiResult {
	result := commons.CommonApiResult{}

	callback := func(resp gorequest.Response, body string, errs []error) {
		result.Status = 200
		result.Message = "OK"

		json.Unmarshal([]byte(body), &jsonmappingobj)
		result.Content = &jsonmappingobj
	}

	fmt.Println("request url to marathon : " + uri)

	request := gorequest.New()
	switch method {
	case commons.HTTP_GET:
		request.Get(MESOS_HOST + uri).End(callback)
	case commons.HTTP_POST:
		request.Post(MESOS_HOST + uri).End(callback)
	case commons.HTTP_PUT:
		request.Put(MESOS_HOST + uri).End(callback)
	case commons.HTTP_DELETE:
		request.Delete(MESOS_HOST + uri).End(callback)
	}

	return result
}
