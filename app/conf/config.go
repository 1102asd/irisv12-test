package conf

import (
	"github.com/json-iterator/go"
	"io/ioutil"
)

type mysql struct {
	Port       string `json:"port"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Hostname   string `json:"hostname"`
	Database   string `json:"database"`
	ServerPort string `json:"serverport"`
}

var (
	CONFIGFILE  = "config.json"
	MysqlConfig = &mysql{}
)

func init() {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile("app/conf/config.json")
	if err != nil {
		panic("Sys config read err")
	}
	err = jsoniter.Unmarshal(b, MysqlConfig)
	if err != nil {
		panic(err)
	}
}
