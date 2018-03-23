package ldLoggerKit

import (
	"log"

	"github.com/astaxie/beego/config/env"
	"github.com/hkloudou/ldDataBaseKit"
)

var database = ""
var tableName = ""
var runmode = ""

func init() {
	if err := ldDataBaseKit.Err(); err != nil {
		log.Println("ldLoggerKit can't reach database")
		//panic(err)
	}
	database = ldDataBaseKit.GetDataBaseName()
	tableName = env.Get("LOG_TABLENAME", "log_dev")
}
