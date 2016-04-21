package marathon

import (
	"fmt"
	"net/http"
	"net/url"

	"gms/apihandler/commons"
	"gms/apihandler/marathon/domain"

	"github.com/bitly/go-simplejson"
	"github.com/parnurzeal/gorequest"
)

const (
	MARATHON_HOST = "http://10.53.15.219:18080/v2/"
	GET_GROUPS    = "/v2/groups"
)

/**
 * get apps of marathon
 */
func getApps() *domain.MarathonApp {

	return nil
}

func getSubGroups() *domain.MarathonGroup {
	return nil
}

func parseJsonArray(obj *simplejson.Json) (*simplejson.Json, int, error) {
	objArray, err := obj.Array()
	objLen := len(objArray)

	return obj, objLen, err
}

/**
 * get groups of marathon
 */
func GetGroups(user string) map[string]domain.MarathonGroup {

	apiResult := commonHandler(commons.HTTP_GET, "groups")
	if apiResult.Status == 200 {

		jsonData := apiResult.Content

		var result = make(map[string]domain.MarathonGroup)

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
				}

				subGroupList, subGroupListLen, _ := parseJsonArray(groupData.Get("groups"))

				for j := 0; j < subGroupListLen; j++ {
					marathonSubGroup := new(domain.MarathonGroup)

					subGroupData := subGroupList.GetIndex(j)

					subGroupId, _ := subGroupData.Get("id").String()
					subGroupVersion, _ := subGroupData.Get("version").String()
					marathonSubGroup.Id = subGroupId
					marathonSubGroup.Version = subGroupVersion

					appList, appListLen, _ := parseJsonArray(subGroupData.Get("apps"))
					for m := 0; m < appListLen; m++ {
						appData := appList.GetIndex(m)

						appId, _ := appData.Get("id").String()
						appCmd, _ := appData.Get("cmd").String()
						appInstances, _ := appData.Get("instances").Int()
						appCpus, _ := appData.Get("cpus").Float64()
						appMem, _ := appData.Get("mem").Int()
						appDisk, _ := appData.Get("disk").Int()
						appServicePort, _ := appData.Get("ports").GetIndex(0).Int()
						appVersion, _ := appData.Get("version").String()

						appContainerObj := appData.Get("container")
						appContainerType, _ := appContainerObj.Get("type").String()

						appContainerDockerObj := appContainerObj.Get("docker")
						appContainerImage, _ := appContainerDockerObj.Get("image").String()
						appContainerNetwork, _ := appContainerDockerObj.Get("network").String()

						if appContainerNetwork == "bridge" {

						}

						appLabelsObj := appData.Get("labels")
						appBuildId, _ := appLabelsObj.Get("BUILDID").String()
						appTriggeredBy, _ := appLabelsObj.Get("TRIGGERED_BY").String()
						appCommit, _ := appLabelsObj.Get("COMMIT").String()
						appJobName, _ := appLabelsObj.Get("JOB_NAME").String()
						appProject, _ := appLabelsObj.Get("PROJECT").String()
						appUser, _ := appLabelsObj.Get("USER").String()
						appEnv, _ := appLabelsObj.Get("ENV").String()
						appDeployId, _ := appLabelsObj.Get("DEPLOYID").String()

						marathonApp := new(domain.MarathonApp)
						marathonApp.Id = appId
						marathonApp.Cmd = appCmd
						marathonApp.Instances = appInstances
						marathonApp.Cpus = appCpus
						marathonApp.Mem = appMem
						marathonApp.Disk = appDisk
					}
				}
			}
		}

		/*

			groupList := groupInfo.Groups

			var result = make(map[string]domain.MarathonGroup)

			if len(groupList) > 0 {
				for _, groupData := range groupList {
					groupId := groupData.ID

					var marathonGroup *domain.MarathonGroup
					if groupObj, isExist := result[groupId]; isExist {
						marathonGroup = &groupObj
					} else {
						marathonGroup = new(domain.MarathonGroup)
					}

					groupApps := marathonGroup.Apps
					append(groupApps)
					fmt.Println(marathonGroup)
				}
			}

			//fmt.Printf("groupList info aaaaa %#v", groupList)
		*/

		return result
	} else {

		return nil
	}
}

/**
 * destory groups
 */
func DestroyGroups(groupsId string) *commons.CommonApiResult {
	encodeUri := url.QueryEscape(groupsId)
	return commonHandler(commons.HTTP_DELETE, "groups/"+encodeUri)
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
