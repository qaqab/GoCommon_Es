package GoCommon_Es

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/qaqab/GoCommon_DbManager"

	"github.com/qaqab/GoCommon_Es"
)

type GlobalBrandSource struct {
	Brand      string `json:"brand"`
	SerialId   int    `json:"serialId"`
	StatusCode int    `json:"statusCode"`
	Lokatien   string `json:"lokatien"`
}
type GlobalBrandGet struct {
	Index             string `json:"_index"`
	Id                string `json:"_id"`
	GlobalBrandSource `json:"_source"`
}
type GlobalBrandsss struct {
	Took int `json:"took"`
	Hits struct {
		Hits []struct {
			GlobalBrandSource `json:"_source"`
			Id                string `json:"_id"`
		} `json:"hits"`
	} `json:"hits"`
}

func TestIdGetSource(t *testing.T) {
	clientAll := GoCommon_DbManager.ClientAll{ConfigSetting: struct {
		ConfigPath string
		ConfigName string
	}{ConfigPath: "/root/go-project/EasyCommon/YamlFile/", ConfigName: "test"}}
	clientAll.DbManagerClient("es.data")
	es_Client := clientAll.EsClient

	es_Setting := GoCommon_Es.EsSetting{}
	es_Setting.EsClient = es_Client

	// 调用IdGetSource方法，获取指定索引和文档ID的源数据
	data, err := es_Setting.IdGetSource("global_brand_basic_info", "8440dcb5951275511720aea4f9ce5467")

	// 检查响应状态码是否为200
	if data.StatusCode != 200 {
		panic("请求失败")
	}

	// 检查是否发生错误
	if err != nil {
		panic(err)
	}

	// 读取响应体内容
	body, err := io.ReadAll(data.Body)
	if err != nil {
		panic(err)
	}

	// 将响应体内容反序列化为GlobalBrandGet结构体
	var globalBrandGet GlobalBrandGet
	json.Unmarshal(body, &globalBrandGet)

	// 打印反序列化后的结果
	fmt.Println(globalBrandGet)

	// 将GlobalBrandGet结构体序列化为JSON字符串
	bytes, _ := json.Marshal(globalBrandGet)

	// 打印JSON字符串
	fmt.Println(string(bytes))
}
