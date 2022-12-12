package golearning_test

import (
	"strings"
	"testing"
)

func TestParseMailAddress(t *testing.T) {
	mail := "coach@neuralgalaxy.com"
	t.Logf("origin mail address: %s \n", mail)
	userName := GetUserNameFromMailAddress(mail, "JW")
	t.Logf("username: %s \n", userName)
}

func TestMobile(t *testing.T) {
	t.Logf("mobile: %s", strings.Replace("+8613331077291", "+86", "", 1))
}

func GetUserNameFromMailAddress(address string, currentUserName string) string {
	addresses := strings.Split(strings.TrimSpace(address), "@")
	if len(addresses) == 1 {
		return currentUserName
	}
	userName := addresses[0]
	if len(userName) == 0 {
		return currentUserName
	}
	return userName
}
