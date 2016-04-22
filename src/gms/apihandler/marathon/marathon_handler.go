package marathon

import (
    "fmt"
    "net/http"
    "net/url"
    "strings"
    "gms/apihandler/commons"
    "gms/apihandler/marathon/domain"

    "github.com/bitly/go-simplejson"
    "github.com/parnurzeal/gorequest"
)

const (
    MARATHON_HOST = "http://10.53.15.219:18080/v2/"
    GET_GROUPS = "/v2/groups"
)

// get apps in marathon
func getApps() *domain.MarathonApp {

    return nil
}

// get sub groups of groups in marathon
func getSubGroups() *domain.MarathonGroup {
    return nil
}

func parseJsonArray(obj *simplejson.Json) (*simplejson.Json, int, error) {
    objArray, err := obj.Array()
    objLen := len(objArray)

    return obj, objLen, err
}

// get groups in marathon
func GetGroups(user string) []domain.MarathonGroup {

    apiResult := commonHandler(commons.HTTP_GET, "groups")
    if apiResult.Status == 200 {

        jsonData := apiResult.Content

        groupList, groupListLen, _ := parseJsonArray(jsonData.Get("groups"))

        if groupListLen > 0 {

            var result = make(map[string]domain.MarathonGroup)

            for i := 0; i < groupListLen; i++ {
                groupData := groupList.GetIndex(i)

                groupId, _ := groupData.Get("id").String()
                groupVersion, _ := groupData.Get("version").String()

                var marathonGroup *domain.MarathonGroup
                if groupObj, isExist := result[groupId]; isExist {
                    marathonGroup = &groupObj
                } else {
                    marathonGroup = new(domain.MarathonGroup)

                    marathonGroup.Id = groupId
                    marathonGroup.Version = groupVersion

                    marathonSubGroups := make([]domain.MarathonGroup, 0)

                    marathonGroup.Groups = marathonSubGroups
                }

                subGroupList, subGroupListLen, _ := parseJsonArray(groupData.Get("groups"))

                for j := 0; j < subGroupListLen; j++ {
                    marathonSubGroup := new(domain.MarathonGroup)

                    subGroupData := subGroupList.GetIndex(j)

                    subGroupId, _ := subGroupData.Get("id").String()

                    subGroupId = strings.Replace(subGroupId, groupId + "/", "", -1)

                    subGroupVersion, _ := subGroupData.Get("version").String()
                    marathonSubGroup.Id = subGroupId
                    marathonSubGroup.Version = subGroupVersion

                    marathonApps := make([]domain.MarathonApp, 0)

                    appList, appListLen, _ := parseJsonArray(subGroupData.Get("apps"))
                    for m := 0; m < appListLen; m++ {
                        appData := appList.GetIndex(m)

                        appLabelsObj := appData.Get("labels")
                        appUser, _ := appLabelsObj.Get("USER").String()

                        if len(user) > 0 && appUser != user {
                            fmt.Println("is contine!!!!!!!")
                            continue
                        }

                        appId, _ := appData.Get("id").String()
                        appCmd, _ := appData.Get("cmd").String()
                        appInstances, _ := appData.Get("instances").Int()
                        appCpus, _ := appData.Get("cpus").Float64()
                        appMem, _ := appData.Get("mem").Int()
                        appDisk, _ := appData.Get("disk").Int()
                        appVersion, _ := appData.Get("version").String()

                        appContainerObj := appData.Get("container")
                        appContainerType, _ := appContainerObj.Get("type").String()

                        appContainerDockerObj := appContainerObj.Get("docker")
                        appContainerImage, _ := appContainerDockerObj.Get("image").String()
                        appContainerNetwork, _ := appContainerDockerObj.Get("network").String()

                        appServicePorts := make([]int, 0)
                        appContainerPorts := make([]int, 0)

                        if strings.ToLower(appContainerNetwork) == "bridge" {
                            portMappings, portMappingLen, _ := parseJsonArray(appContainerDockerObj.Get("portMappings"))

                            for pi := 0; pi < portMappingLen; pi++ {
                                portMappingData := portMappings.GetIndex(pi)

                                tempServicePort, _ := portMappingData.Get("servicePort").Int()
                                tempContainerPort, _ := portMappingData.Get("containerPort").Int()

                                appServicePorts = append(appServicePorts, tempServicePort)
                                appContainerPorts = append(appContainerPorts, tempContainerPort)
                            }
                        } else {
                            portList, portListLen, _ := parseJsonArray(appData.Get("ports"))
                            for pi := 0; pi < portListLen; pi++ {
                                portMappingData := portList.GetIndex(pi)
                                tempContainerPort, _ := portMappingData.Get("containerPort").Int()
                                appContainerPorts = append(appContainerPorts, tempContainerPort)
                            }
                        }

                        appBuildId, _ := appLabelsObj.Get("BUILDID").String()
                        appTriggeredBy, _ := appLabelsObj.Get("TRIGGERED_BY").String()
                        appCommit, _ := appLabelsObj.Get("COMMIT").String()
                        appJobName, _ := appLabelsObj.Get("JOB_NAME").String()
                        appProject, _ := appLabelsObj.Get("PROJECT").String()
                        appEnv, _ := appLabelsObj.Get("ENV").String()
                        appDeployId, _ := appLabelsObj.Get("DEPLOYID").String()

                        marathonApp := new(domain.MarathonApp)
                        marathonApp.Id = appId
                        marathonApp.Cmd = appCmd
                        marathonApp.Instances = appInstances
                        marathonApp.Cpus = appCpus
                        marathonApp.Mem = appMem
                        marathonApp.Disk = appDisk
                        marathonApp.ServicePorts = appServicePorts
                        marathonApp.ContainerPorts = appContainerPorts
                        marathonApp.ContainerType = appContainerType
                        marathonApp.ContainerNetwork = appContainerNetwork
                        marathonApp.ContainerImage = appContainerImage
                        marathonApp.BuildId = appBuildId
                        marathonApp.TriggeredBy = appTriggeredBy
                        marathonApp.Commit = appCommit
                        marathonApp.JobName = appJobName
                        marathonApp.Project = appProject
                        marathonApp.User = appUser
                        marathonApp.Env = appEnv
                        marathonApp.DeployId = appDeployId
                        marathonApp.Version = appVersion

                        marathonApps = append(marathonApps, *marathonApp)
                    } // end of apps loop

                    // filter된 apps 데이터가 없는 경우 해당 sub group은 버림
                    if len(marathonApps) > 0 {
                        marathonSubGroup.Apps = marathonApps
                        marathonGroup.Groups = append(marathonGroup.Groups, *marathonSubGroup)
                    }
                }

                if len(marathonGroup.Groups) > 0 {
                    result[groupId] = *marathonGroup
                }
            }

            resultList := []domain.MarathonGroup{}
            for _, val := range result {
                resultList = append(resultList, val)
            }

            return resultList
        }
    }

    return nil
}

// destory groups
func DestroyGroups(groupsId string) *commons.CommonApiResult {
    encodeUri := url.QueryEscape(groupsId)
    return commonHandler(commons.HTTP_DELETE, "groups/" + encodeUri)
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
        resp, bodyBytes, _ = request.Get(MARATHON_HOST + uri).EndBytes()
    case commons.HTTP_POST:
        resp, bodyBytes, _ = request.Post(MARATHON_HOST + uri).EndBytes()
    case commons.HTTP_PUT:
        resp, bodyBytes, _ = request.Put(MARATHON_HOST + uri).EndBytes()
    case commons.HTTP_DELETE:
        resp, bodyBytes, _ = request.Delete(MARATHON_HOST + uri).EndBytes()
    }

    callbackData := callback(resp, bodyBytes)

    return callbackData
}
