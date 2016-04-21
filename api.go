package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "gms/apihandler/marathon"
    "gms/apihandler/mesos"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello Gravity Management System")
    })

    router.GET("/settings", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello Gravity Management System Settings")
    })

    router.GET("/marathon/groups", func(c *gin.Context) {

        user := c.Query("user")
        fmt.Println("user parameter : " + user)
        groupData := marathon.GetGroups(user)

        fmt.Printf("is result %v", groupData)

        c.JSON(200, groupData)
    })

    router.DELETE("/marathon/groups/:groupsId", func(c *gin.Context) {
        groupsId := c.Param("groupsId")
        resultData := marathon.DestroyGroups("/" + groupsId)
        c.JSON(200, resultData)
    })

    router.GET("/mesos/master_state", func(c *gin.Context) {
        resultData := mesos.GetMasterState()
        c.JSON(200, resultData)
    })

    router.Run(":4000") // listen and server on 0.0.0.0:5000
}
