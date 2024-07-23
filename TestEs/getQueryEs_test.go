package GoCommon_Es

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/qaqab/GoCommon_DbManager"

	"github.com/qaqab/GoCommon_Es"
)

type GlobalBrand struct {
	Took int `json:"took"`
	Hits struct {
		Hits []struct {
			Source struct {
				Brand      string `json:"brand"`
				SerialId   int    `json:"serialId"`
				StatusCode int    `json:"statusCode"`
				Lokatien   string `json:"lokatien"`
			} `json:"_source"`
			Id string `json:"_id"`
		} `json:"hits"`
	} `json:"hits"`
}

func TestGetQuery(t *testing.T) {
	clientAll := GoCommon_DbManager.ClientAll{ConfigSetting: struct {
		ConfigPath string
		ConfigName string
	}{ConfigPath: "/root/go-project/EasyCommon/YamlFile/", ConfigName: "test"}}
	clientAll.DbManagerClient("es.data")
	es_Setting := GoCommon_Es.EsSetting{Addresse: clientAll.EsSettingData.Addresse, Username: clientAll.EsSettingData.Username, Password: clientAll.EsSettingData.Password}
	es_Setting.EsClient = clientAll.EsClient
	docChan := make(chan []byte)
	go es_Setting.GetSource("global_brand_basic_info", `{ "query": { "match": {"brand":"TORNADO"} } }`, []string{"brand", "serialId", "statusCode", "lokatien"}, docChan)
	for {
		doc, ok := <-docChan
		if !ok {
			break
		}
		var globalBrandData GlobalBrand

		json.Unmarshal(doc, &globalBrandData)
		for _, k := range globalBrandData.Hits.Hits {
			log.Println(k)
			bytes, _ := json.Marshal(k)

			// 打印JSON字符串
			log.Println(string(bytes))

		}
	}

}
