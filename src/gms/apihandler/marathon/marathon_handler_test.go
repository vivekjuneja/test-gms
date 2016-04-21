package marathon_test

import (
	"gms/apihandler/marathon"
	"testing"
)

func TestGetGroups(t *testing.T) {
	result := marathon.GetGroups("jace")

	if result == nil {
		t.Error("Wrong Result.")
	}
}
