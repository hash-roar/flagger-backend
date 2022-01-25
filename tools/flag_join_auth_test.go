package tools_test

import (
	"flagger-backend/tools"
	"testing"
)

func TestGetAuthNum(t *testing.T) {
	var testAuthArray = [...]int{1, 6, 3}
	result := tools.GetAuthNum(testAuthArray[:])
	t.Error(result)

}

func TestAuthArray(t *testing.T) {
	var authNum uint64 = 74
	if tools.IsAuthorized(6, authNum) {
		t.Error("error")
	}
}
