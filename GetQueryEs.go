package GoCommon_Es

import (
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func (esSetting EsSetting) GetSource(index string, body string, source []string, docChan chan []byte) {
	// docChan := make(chan []byte)
	req := esapi.SearchRequest{
		Index:  []string{index},
		Body:   strings.NewReader(body),
		Scroll: 120 * time.Second,
		Source: source}
	Scroll(esSetting.EsClient, req, docChan)

}

type scrollResponse struct {
	ScrollId string `json:"_scroll_id"`
	Hits     struct {
		Hits []any `json:"hits"`
	} `json:"hits"`
}

func Scroll(es *elasticsearch.Client, req esapi.SearchRequest, docChan chan []byte) {
	body := ParseResponse(req.Do(context.Background(), es))
	var response scrollResponse
	json.Unmarshal(body, &response)
	docChan <- body
	defer close(docChan)
	if response.ScrollId != "" && len(response.Hits.Hits) > 0 {
		scroll_req := esapi.ScrollRequest{
			ScrollID: response.ScrollId,
		}
		for {
			body = ParseResponse(scroll_req.Do(context.Background(), es))
			var new_response scrollResponse
			json.Unmarshal(body, &new_response)
			docChan <- body
			if len(new_response.Hits.Hits) == 0 {
				break
			}
		}
	}
}
func ParseResponse(res *esapi.Response, err error) []byte {
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body
}
