package fetcher

import (
	"encoding/json"
	"regexp"
	"zhenai-spider/model"
	"zhenai-spider/util"
)

type ProfileFetcher struct {
	URL string
}

func (f *ProfileFetcher) Run() (result Result) {
	content, err := util.HttpRequestGet(f.URL)
	if err != nil {
		util.WarnLog.Println(err)
		return
	}

	reg := regexp.MustCompile(`<script>window\.__INITIAL_STATE__=(.+);\(function\(\)`)
	matches := reg.FindAllSubmatch(content, -1)
	if len(matches) == 0 {
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(matches[0][1], &jsonMap)
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

	result.Items = append(result.Items, profiles)

	return
}
