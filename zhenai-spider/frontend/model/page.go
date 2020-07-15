package model

type SearchResult struct {
	Hints       int64
	Start       int
	PrevFrom    int
	NextFrom    int
	QueryString string
	Items       []interface{}
}
