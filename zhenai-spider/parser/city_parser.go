package parser

import (
	"fmt"
	"regexp"
)

type CityParser struct {
}

func (p *CityParser) Parse(content []byte) {
	// <a href="http://album.zhenai.com/u/1509603866" target="_blank">如茶一般</a>
	reg := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	matches := reg.FindAllSubmatch(content, -1)
	for _, match := range matches {
		fmt.Println(string(match[2]), string(match[1]))
	}
}
