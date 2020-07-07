package util

import (
	"bufio"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("wrong http reponse status code: %d of %s", resp.StatusCode, url)
	// }

	r := bufio.NewReader(resp.Body)
	e := DetermineEncoding(r)
	reader := transform.NewReader(r, e.NewDecoder())
	return ioutil.ReadAll(reader)
}

// // 根据URL提取
// func Fetch(url string) ([]byte, error) {
// 	resp, err := httpRequest(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("wrong status code: %d of %s", resp.StatusCode, url)
// 	}
// 	r := bufio.NewReader(resp.Body)
// 	e := determineEncoding(r)
// 	utf8Reader := transform.NewReader(r,  e.NewDecoder())
// 	return ioutil.ReadAll(utf8Reader)
// }

// // 编码识别
// func determineEncoding(
// 	r *bufio.Reader) encoding.Encoding{
// 	bytes, err:=r.Peek(1024)
// 	if err!=nil {
// 		log.Printf("fetch error : %v\n", err )
// 		// 如果没有识别到，返回一个UTF-8(默认)
// 		return unicode.UTF8
// 	}
// 	e, _, _ := charset.DetermineEncoding(
// 		bytes, "")
// 	return e
// }
