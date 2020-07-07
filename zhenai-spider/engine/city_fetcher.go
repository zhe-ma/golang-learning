package engine

import (
	"regexp"
	"zhenai-spider/util"
)

type CityFetcher struct {
	URL string
}

func (f *CityFetcher) Run() (fetchers []Fetcher) {
	content, err := util.HttpRequestGet(f.URL)
	if err != nil {
		util.WarnLog.Println(err)
		return
	}

	// E.g. <a href="http://www.zhenai.com/zhenghun/huaibei" data-v-2cb5b6a2>淮北</a>
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := reg.FindAllSubmatch(content, -1)
	for _, match := range matches {
		fetcher := &ProfileFetcher{string(match[1])}
		fetchers = append(fetchers, fetcher)

		if len(fetchers) > 1 {
			return
		}
	}

	return
}
