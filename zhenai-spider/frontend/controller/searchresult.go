package controller

import (
	"context"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"zhenai-spider/frontend/model"
	"zhenai-spider/frontend/view"
	profile "zhenai-spider/model"

	"github.com/olivere/elastic"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.FormValue("query"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}

	result, err := s.QuerySearchResult(query, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = s.view.Render(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s SearchResultHandler) QuerySearchResult(query string, from int) (model.SearchResult, error) {
	result := model.SearchResult{}

	result.QueryString = query
	resp, err := s.client.Search("zhenai").Type("profiles").Query(elastic.NewQueryStringQuery(query)).From(from).Do(context.Background())
	if err != nil {
		return result, nil
	}

	result.Hints = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(profile.Profile{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func NewSearchResultHandler(templateFile string, elasticHost string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetURL(elasticHost), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.NewSearchResultView(templateFile),
		client: client,
	}
}
