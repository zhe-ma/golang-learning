package engine

import "zhenai-spider/parser"
import "zhenai-spider/util"

type Fetcher struct {
	URL string
	Parser parser.IParser
}

func (f *Fetcher) Run() error {
	cotent, err := util.Fetch(f.URL)
	if err != nil {
		return err
	}

	f.Parser.Parse(cotent)

	return nil
}

func NewFetcher(url string, parser parser.IParser) *Fetcher {
	return &Fetcher{
		URL: url,
		Parser: parser,
	}
}