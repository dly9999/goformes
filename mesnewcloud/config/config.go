package config

import (
	"fmt"

	"github.com/Unknwon/goconfig"
	//"github.com/Unknwon/goconfig"
	"time"
)

type Db struct {
	Server string
	Port   int
	Usr    string
	Psd    string
	Dbname string
} //数据库
type MQconfig struct {
	Mqip     string
	Port     int
	Username string
	Password string
} //Rabitserver Mqserver
type Queconfig struct {
	Per      string
	Name     string
	Routkey  string
	Exchange string
	Constr   string
} //队列名称
type HttpCfg struct {
	Http_Port    int
	Read_Timeout time.Duration
	WriteTimeout time.Duration
}

var (
	dbconf  Db
	mqconf  MQconfig
	queplan Queconfig
	hptconf HttpCfg
	runmod  string
)
var cfg *goconfig.ConfigFile

func init() {
	cfg, err := goconfig.LoadConfigFile("config/app.ini")
	if err != nil {
		panic(err)
	}
	var er error
	if dbconf.Server, er = cfg.GetValue("DB", "ip"); er != nil {
		panic(er)
	}

	if dbconf.Port, er = cfg.Int("DB", "port"); er != nil {
		fmt.Println(er)
	}

	if dbconf.Usr, er = cfg.GetValue("DB", "username"); er != nil {
		fmt.Println(er)
	}

	if dbconf.Psd, er = cfg.GetValue("DB", "password"); er != nil {
		fmt.Println(er)
	}

	if dbconf.Dbname, er = cfg.GetValue("DB", "Dbname"); er != nil {
		fmt.Println(er)
	}

	fmt.Println(dbconf)
	if mqconf.Mqip, er = cfg.GetValue("Mqserver", "mqip"); er != nil {
		fmt.Println(er)
	}

	if mqconf.Port, er = cfg.Int("Mqserver", "port"); er != nil {
		fmt.Println(er)
	}

	if mqconf.Username, er = cfg.GetValue("Mqserver", "username"); er != nil {
		fmt.Println(er)
	}

	if mqconf.Password, er = cfg.GetValue("Mqserver", "password"); er != nil {
		fmt.Println(er)
	}

	fmt.Println(mqconf)
	if queplan.Per, er = cfg.GetValue("Que", "per"); er != nil {
		panic(er)
	}

	if queplan.Name, er = cfg.GetValue("Que", "name"); er != nil {
		panic(er)
	}
	fmt.Println(queplan)
	hptconf.Http_Port = cfg.MustInt("server", "HTTP_PORT", 8080)
	hptconf.Read_Timeout = time.Duration(cfg.MustInt("server", "READ_TIMEOUT", 60)) * time.Second
	hptconf.WriteTimeout = time.Duration(cfg.MustInt("server", "READ_TIMEOUT", 60)) * time.Second
	fmt.Println(hptconf)
	if runmod, er = cfg.GetValue("MODE", "RUN_MODE"); er != nil {
		fmt.Println(er)
	}

	fmt.Println(runmod)

}

/*
	********************
	返回DB MQ 配置及server
	端口
	********************

*/
func Getconfig() (Db, MQconfig, Queconfig, HttpCfg) {

	return dbconf, mqconf, queplan, hptconf
}

func Getmode() string {

	return runmod
}
