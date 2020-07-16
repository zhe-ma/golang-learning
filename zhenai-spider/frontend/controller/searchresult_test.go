package controller

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestElasticQuery(t *testing.T) {
	client, err := elastic.NewClient(elastic.SetURL("http://10.196.102.145:9200"), elastic.SetSniff(false))
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Search("zhenai").Query(elastic.NewQueryStringQuery("Age:30")).Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log("TotalHints", resp.TotalHits())

	resp, err = client.Search("zhenai").Type("profiles").Query(elastic.NewQueryStringQuery("Age:34 Height:166")).Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log("TotalHints", resp.TotalHits())

	resp, err = client.Search("zhenai").Query(elastic.NewQueryStringQuery("Height:>177")).Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log("TotalHints", resp.TotalHits())

	resp, err = client.Search("zhenai").Query(elastic.NewQueryStringQuery("Age:(<25) Sex:male Height:([166 TO 168])")).Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Log("TotalHints", resp.TotalHits())
}
