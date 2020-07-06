package main

import (
	"zhenai-spider/engine"
	"zhenai-spider/parser"
)

func main() {
	seed := engine.NewFetcher("http://www.zhenai.com/zhenghun", new(parser.CitylistParser))
	engine.Run(seed)
	// engine.Run()
	// content, _ := util.Fetch("http://www.zhenai.com/zhenghun")

	// fmt.Println(content)
	// resp, err := http.Get("http://www.zhenai.com/zhenghun")
	// if err != nil {
	// 	log.Fatalf("Failed to request URL: %v.", err)
	// }
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// // err = ioutil.WriteFile("cityList.html", body, 0666)
	// // if err != nil {
	// // 	log.Fatalf("Failed to write file cityList.html: %v.", err)
	// // }

	// // E.g. <a href="http://www.zhena	i.com/zhenghun/huaibei" data-v-2cb5b6a2>淮北</a>
	// reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	// matches := reg.FindAllSubmatch(body, -1)
	// // for _, match := range matches {
	// // 	fmt.Println(string(match[2]), string(match[1]))
	// // }

	// resp1, err := http.Get(string(matches[0][1]))
	// if err != nil {
	// 	log.Fatalf("Failed to request URL: %v.", err)
	// }
	// defer resp1.Body.Close()

	// body1, _ := ioutil.ReadAll(resp1.Body)
	// err = ioutil.WriteFile("city.html", body1, 0666)
	// if err != nil {
	// 	log.Fatalf("Failed to write file city.html: %v.", err)
	// }

	// <a href="http://album.zhenai.com/u/1509603866" target="_blank">如茶一般</a>
	// reg2 := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	// matches2 := reg2.FindAllSubmatch(body1, -1)
	// for _, match := range matches2 {
	// 	fmt.Println(string(match[2]), string(match[1]))
	// }

	// resp2, err := http.Get(string(matches2[0][1]))
	// if err != nil {
	// 	log.Fatalf("Failed to request URL: %v.", err)
	// }
	// defer resp2.Body.Close()

	// body2, _ := ioutil.ReadAll(resp2.Body)
	// err = ioutil.WriteFile("profile.html", body2, 0666)
	// if err != nil {
	// 	log.Fatalf("Failed to write file profile.html: %v.", err)
	// }

	// request, err := http.NewRequest("GET", `http://album.zhenai.com/u/1427091188`, nil)
	// // if err != nil {
	// // 	return err
	// // }
	// // 设置请求投
	// request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	// request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	// request.Header.Add("Connection", "keep-alive")
	// request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	// client := http.Client{}
	// // Do sends an HTTP request and returns an HTTP response
	// // 发起一个HTTP请求，返回一个HTTP响应
	// resp2, err := client.Do(request)
	// defer resp2.Body.Close()

	// body2, _ := ioutil.ReadAll(resp2.Body)
	// err = ioutil.WriteFile("profile.html", body2, 0666)
	// if err != nil {
	// 	log.Fatalf("Failed to write file profile.html: %v.", err)
	// }
}
