package engine

type Fetcher interface {
	Run() []Fetcher
}
