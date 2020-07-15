package fetcher

import (
	"encoding/json"
	"regexp"
	"time"
	"zhenai-spider/model"
	"zhenai-spider/util"
)

type ProfileFetcher struct {
	URL string
}

var profileRateLimiter = time.Tick(time.Second)

func (f *ProfileFetcher) Run() (result Result) {
	<-profileRateLimiter

	content, err := util.HttpRequestGet(f.URL)
	if err != nil {
		util.WarnLog.Println(err)
		return
	}

	parseProfiles(content, &result)
	parseSubFetchers(content, &result)

	return
}

func parseProfiles(content []byte, result *Result) {
	reg := regexp.MustCompile(`<script>window\.__INITIAL_STATE__=(.+);\(function\(\)`)
	matches := reg.FindAllSubmatch(content, -1)
	if len(matches) == 0 {
		return
	}

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(matches[0][1], &jsonMap)
	if err != nil {
		util.WarnLog.Println("Failed to parse profile json:", err)
		return
	}

	memberListData := jsonMap["memberListData"].(map[string]interface{})
	// !这里的类型是[]interface{}，不是[]map[string]interface{}.
	memberList := memberListData["memberList"].([]interface{})

	profiles := []model.Profile{}

	for _, member := range memberList {
		memberData := member.(map[string]interface{})
		var profile model.Profile
		profile.MemberId = memberData["memberID"].(float64)
		profile.NickName = memberData["nickName"].(string)
		profile.Age = int32(memberData["age"].(float64))
		profile.Height = int32(memberData["height"].(float64))
		profile.Education = memberData["education"].(string)
		if int32(memberData["sex"].(float64)) == 0 {
			profile.Sex = "Male"
		} else {
			profile.Sex = "Female"
		}
		profile.WorkCity = memberData["workCity"].(string)
		profile.Occutation = memberData["occupation"].(string)
		profile.Salary = memberData["salary"].(string)
		profiles = append(profiles, profile)
	}

	result.Items = append(result.Items, &profiles)
}

func parseSubFetchers(content []byte, result *Result) {
	reg := regexp.MustCompile(`(http://www.zhenai.com/zhenghun/[a-zA-Z]*/[\d]+)`)
	matches := reg.FindAllSubmatch(content, -1)

	for _, match := range matches {
		fetcher := &ProfileFetcher{URL: string(match[1])}
		result.SubFetchers = append(result.SubFetchers, fetcher)
	}
}
