package main

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/gin-gonic/gin"

    "gms/apihandler/marathon"
)

const V1  = "/v1/"

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello Gravity Management System")
    })

    router.GET("/settings", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello Gravity Management System Settings")
    })

    router.GET(V1 + "apps", func(c *gin.Context) {

        user := c.Query("user")

        fmt.Println("user parameter : " + user)
        groupData := marathon.GetGroups(user)

        fmt.Printf("is result %v", groupData)

        // if using jsonp callback.
        callback := c.Query("callback")
        if len(callback) > 0 {
            groupDataJson, _ := json.Marshal(groupData)
            c.String(200, callback + "(" + string(groupDataJson) + ");")
        } else {
            c.JSON(200, groupData)
        }
    })

    router.DELETE("/apps/:groupsId", func(c *gin.Context) {
        groupsId := c.Param("groupsId")
        resultData := marathon.DestroyGroups("/" + groupsId)
        c.JSON(200, resultData)
    })

    router.Run(":4000") // listen and server on 0.0.0.0:4000
}
