package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		log.Fatalf("Failed to request URL: %v.", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// err = ioutil.WriteFile("cityList.html", body, 0666)
	// if err != nil {
	// 	log.Fatalf("Failed to write file cityList.html: %v.", err)
	// }

	// E.g. <a href="http://www.zhenai.com/zhenghun/huaibei" data-v-2cb5b6a2>淮北</a>
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := reg.FindAllSubmatch(body, -1)
	// for _, match := range matches {
	// 	fmt.Println(string(match[2]), string(match[1]))
	// }

	resp1, err := http.Get(string(matches[0][1]))
	if err != nil {
		log.Fatalf("Failed to request URL: %v.", err)
	}
	defer resp1.Body.Close()

	body1, _ := ioutil.ReadAll(resp1.Body)
	err = ioutil.WriteFile("city.html", body1, 0666)
	if err != nil {
		log.Fatalf("Failed to write file city.html: %v.", err)
	}

}
