package GoCommon_Es

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type EsSetting struct {
	Addresse string `json:"addresse"`
	Username string `json:"username"`
	Password string `json:"password"`

	EsClient *elasticsearch.Client `json:"client"`
}

func (esSetting EsSetting) IdGetSource(index string, id string) (*esapi.Response, error) {
	data, err := esSetting.EsClient.Get(index, id)
	return data, err
}
