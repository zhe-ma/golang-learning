package parser

import (
	"fmt"
	"regexp"
)

type CityParser struct {
}

func (p *CityParser) Parse(content []byte) {
	// E.g. <a href="http://www.zhenai.com/zhenghun/huaibei" data-v-2cb5b6a2>淮北</a>
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := reg.FindAllSubmatch(content, -1)
	for _, match := range matches {
		fmt.Println(string(match[2]), string(match[1]))
	}
}
