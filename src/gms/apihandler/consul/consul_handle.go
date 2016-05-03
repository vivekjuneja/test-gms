package consul

import (
    "strings"
    "github.com/parnurzeal/gorequest"
)

const (
    CONSUL_HOST = "http://10.53.15.219:8500"
)

// Get Container ( Mesos Slave ) Ip Address
func GetContainerIpAddr(flag string) string {
    request := gorequest.New()
    resp, body, _ := request.Get(CONSUL_HOST + "/v1/kv/haproxy_" + flag + "?raw").EndBytes()
    if resp.StatusCode == 200 {
        return strings.TrimSpace(string(body))
    }
    return ""
}