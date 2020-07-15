package view

import (
	"os"
	"testing"
	"zhenai-spider/frontend/model"
	profile "zhenai-spider/model"
)

func TestSearchResultView(t *testing.T) {
	view := NewSearchResultView("index.html")

	result := model.SearchResult{}
	result.Hints = 100
	result.Start = 30
	result.PrevFrom = 100
	result.NextFrom = 200
	result.QueryString = "age=10"

	item := profile.Profile{
		NickName:   "Nick",
		Age:        34,
		Height:     166,
		Education:  "大学",
		Sex:        "male",
		WorkCity:   "上海",
		Occutation: "厨师",
		Salary:     "1000元",
	}

	for i := 0; i < 10; i++ {
		result.Items = append(result.Items, item)
	}

	out, err := os.Create("index_test.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = view.Render(out, result)
	if err != nil {
		t.Error(err)
	}
}
