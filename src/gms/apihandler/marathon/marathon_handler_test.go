package marathon_test

import (
    "gms/apihandler/marathon"
    "testing"
    "fmt"
)

func TestGetGroups(t *testing.T) {
    result := marathon.GetGroups("jace")

    if result == nil {
        t.Error("Wrong Result.")
    } else {
        fmt.Print(result)
    }
}
