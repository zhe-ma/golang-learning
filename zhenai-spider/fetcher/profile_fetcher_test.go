package fetcher

import (
	"fmt"
	"testing"
	"zhenai-spider/model"
	"zhenai-spider/util"
)

func TestParseSubFetchers(t *testing.T) {
	content, err := util.HttpRequestGet(`http://www.zhenai.com/zhenghun/huaibei`)
	if err != nil {
		t.Error(err)
	}

	var result Result
	parseSubFetchers(content, &result)
	for _, f := range result.SubFetchers {
		profileFetcher := f.(*ProfileFetcher)
		fmt.Println(profileFetcher.URL)
	}
}

func TestParseProfiles(t *testing.T) {
	content, err := util.HttpRequestGet(`http://www.zhenai.com/zhenghun/huaibei`)
	if err != nil {
		t.Error(err)
	}

	var result Result
	parseProfiles(content, &result)

	for _, item := range result.Items {
		if profiles, ok := item.(*[]model.Profile); ok {
			for _, profile := range *profiles {
				fmt.Println(profile)
			}
		}
	}
}
