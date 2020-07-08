package fetcher

type Result struct {
	SubFetchers []Fetcher
	Items       []interface{}
}

type Fetcher interface {
	Run() Result
}
