package neturl_test

import (
	"net/url"
	"testing"
)

func TestNetUrlFunc(t *testing.T) {
	rawurl := "https://www.baidu.com/aaa/bbb/ccc?a=1&a=2&b=3#11"
	parsedUrl, _ := url.Parse(rawurl)
	parsedUrl.Host = "localhost"

	t.Logf("url value: %#v", parsedUrl)
	t.Logf("url value: %#v", parsedUrl.String())
}
