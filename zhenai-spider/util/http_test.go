package util

import (
	"io/ioutil"
	"testing"
)

// go test -run TestHttpRequestGet1 -v
func TestHttpRequestGet1(t *testing.T) {
	content, err := HttpRequestGet("http://www.zhenai.com/zhenghun")
	if err != nil {
		t.Error(err)
	}

	ioutil.WriteFile("citylist.html", content, 0666)
}

func TestHttpRequestGet2(t *testing.T) {
	content, err := HttpRequestGet("http://www.zhenai.com/zhenghun/huaibei")
	if err != nil {
		t.Error(err)
	}

	ioutil.WriteFile("city.html", content, 0666)
}
