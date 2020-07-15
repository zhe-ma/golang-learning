package view

import (
	"io"
	"text/template"
	"zhenai-spider/frontend/model"
)

type SearchResultView struct {
	t *template.Template
}

func NewSearchResultView(templateFile string) SearchResultView {
	return SearchResultView{t: template.Must(template.ParseFiles(templateFile))}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.t.Execute(w, data)
}
