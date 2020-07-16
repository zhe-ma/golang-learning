package fetcher

import (
	"regexp"
	"time"
	"zhenai-spider/util"
)

type CityFetcher struct {
	URL string
}

var cityRateLimiter = time.Tick(time.Second)

func (f *CityFetcher) Run() (result Result) {
	<-cityRateLimiter

	content, err := util.HttpRequestGet(f.URL)
	if err != nil {
		util.WarnLog.Println(err)
		return
	}

	// E.g. <a href="http://www.zhenai.com/zhenghun/huaibei" data-v-2cb5b6a2>淮北</a>
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := reg.FindAllSubmatch(content, -1)
	for _, match := range matches {
		fetcher := &ProfileFetcher{URL: string(match[1])}
		result.SubFetchers = append(result.SubFetchers, fetcher)
		result.Items = append(result.Items, string(match[2]))
		// if len(result.SubFetchers) > 1 {
		// 	util.TraceLog.Println("Finish fetch")
		// 	return
		// }
	}

	return
}
