package parser

import (
	"encoding/json"
	"fmt"
	"regexp"

	"zhenai-spider/model"
)

type ProfileParser struct {
}

func (p *ProfileParser) Parse(content []byte) {
	reg := regexp.MustCompile(`<script>window\.__INITIAL_STATE__ = (.+); \(function \(\)`)
	matches := reg.FindAllSubmatch(content, -1)
	if len(matches) == 0 {
		return
	}

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(matches[0][1], &jsonMap)
	if err != nil {
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
		profile.Age = memberData["age"].(int32)
		profile.Height = memberData["height"].(int32)
		profile.Education = memberData["education"].(string)
		if memberData["Sex"].(int32) == 0 {
			profile.Sex = "Male"
		} else {
			profile.Sex = "Female"
		}
		profile.WorkCity = memberData["workCity"].(string)
		profile.Occutation = memberData["occupation"].(string)
		profile.Salary = memberData["salary"].(string)
		profiles := append(profiles, profile)
	}

	fmt.Println(profiles)
}

// type Profile struct {
// 	NickName   string
// 	Age        int32
// 	Height     int32
// 	Education  string
// 	Sex        string
// 	LiveCity   string
// 	WorkCity   string
// 	Occutation string
// 	Salary     string
// }

// {
// 	"memberListData":{
// 		 "currentPage":1,
// 		 "memberList":[
// 				{
// 					 "age":64,
// 					 "education":"大专",
// 					 "height":172,
// 					 "marriage":"离异",
// 					 "nickName":"云卷云舒",
// 					 "occupation":"政府机构",
// 					 "salary":"3001-5000元",
// 					 "sex":0,
// 					 "workCity":"安徽淮北"
// 				}
//		 ]
// 	}
