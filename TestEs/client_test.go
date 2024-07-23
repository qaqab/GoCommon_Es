package GoCommon_Es

import (
	"fmt"
	"testing"

	"github.com/qaqab/GoCommon_DbManager"

	"github.com/qaqab/GoCommon_Es"
)

func TestEsClient(t *testing.T) {

	clientAll := GoCommon_DbManager.ClientAll{ConfigSetting: struct {
		ConfigPath string
		ConfigName string
	}{ConfigPath: "/root/go-project/EasyCommon/YamlFile/", ConfigName: "test"}}
	clientAll.DbManagerClient("es.data")
	es_Client := clientAll.EsClient

	es_Setting := GoCommon_Es.EsSetting{Addresse: clientAll.EsSettingData.Addresse, Username: clientAll.EsSettingData.Username, Password: clientAll.EsSettingData.Password}
	es_Setting.EsClient = es_Client

	fmt.Println(es_Client)

}
