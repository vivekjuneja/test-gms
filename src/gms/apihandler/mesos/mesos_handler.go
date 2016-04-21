package mesos

import (
    "fmt"
    "gms/apihandler/commons"
    "github.com/parnurzeal/gorequest"
    "net/http"
    "github.com/bitly/go-simplejson"
)

const (
    MESOS_HOST = "http://10.53.15.219:15050"
)

/**
 * get groups
 */
func GetMasterState() map[string]interface{} {
    commonHandler(commons.HTTP_GET, "/master/state.json")
    return nil
}

func commonHandler(method string, uri string) *commons.CommonApiResult {
    callback := func(resp *http.Response, bodyBytes []byte) *commons.CommonApiResult {
        jsonData, _ := simplejson.NewFromReader(resp.Body)

        return &commons.CommonApiResult{
            Status:  resp.StatusCode,
            Message: resp.Status,
            Content: *jsonData,
        }
    }

    fmt.Println("request url to marathon : " + uri)

    request := gorequest.New()
    var resp *http.Response
    var bodyBytes []byte
    switch method {
    case commons.HTTP_GET:
        resp, bodyBytes, _ = request.Get(MESOS_HOST + uri).EndBytes()
    case commons.HTTP_POST:
        resp, bodyBytes, _ = request.Post(MESOS_HOST + uri).EndBytes()
    case commons.HTTP_PUT:
        resp, bodyBytes, _ = request.Put(MESOS_HOST + uri).EndBytes()
    case commons.HTTP_DELETE:
        resp, bodyBytes, _ = request.Delete(MESOS_HOST + uri).EndBytes()
    }

    callbackData := callback(resp, bodyBytes)

    return callbackData
}
