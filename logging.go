package ldLoggerKit

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/hkloudou/ldDataBaseKit"
	"github.com/hkloudou/websoketkit"
)

//LogBuilder LogBuilder
type LogBuilder struct {
	Source string
	WS     *websoketkit.WebsocketHandler
}

//InitLogerFactory InitLogerFactory
func (m *LogBuilder) InitLogerFactory(s string) *LogBuilder {
	m.Source = s
	//m.WS =
	return m
}

//Println Println
func (m *LogBuilder) Println(v ...interface{}) {
	text := fmt.Sprint(v...)
	log.Println(v...)
	m.Log(LTDEBUG, text)
}

//Printf Printf
func (m *LogBuilder) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
	text := fmt.Sprintf(format, v...)
	log.Printf(format, v...)
	m.Log(LTDEBUG, text)
}

//Infoln Info Log
func (m *LogBuilder) Infoln(v ...interface{}) {
	text := fmt.Sprint(v...)
	m.Log(LTINFO, text)
}

//Infof Info Log
func (m *LogBuilder) Infof(format string, v ...interface{}) {
	text := fmt.Sprintf(format, v...)
	m.Log(LTINFO, text)
}

//Debugln Debug
func (m *LogBuilder) Debugln(v ...interface{}) {
	text := fmt.Sprint(v...)
	m.Log(LTDEBUG, text)
}

//Debugf Debug
func (m *LogBuilder) Debugf(format string, v ...interface{}) {
	text := fmt.Sprintf(format, v...)
	m.Log(LTDEBUG, text)
}

//Warnln Warn
func (m *LogBuilder) Warnln(v ...interface{}) {
	text := fmt.Sprint(v...)
	m.Log(LTWARN, text)
}

//Warnf Warn
func (m *LogBuilder) Warnf(format string, v ...interface{}) {
	text := fmt.Sprintf(format, v...)
	m.Log(LTWARN, text)
}

//Errorln Error
func (m *LogBuilder) Errorln(v ...interface{}) {
	text := fmt.Sprint(v...)
	m.Log(LTERROR, text)

}

//Errorf Error
func (m *LogBuilder) Errorf(format string, v ...interface{}) {
	text := fmt.Sprintf(format, v...)
	m.Log(LTERROR, text)
}

//Log log
func (m *LogBuilder) Log(t LogType, data string) {
	if m.Source == "" {
		m.Source = "system"
	}
	dbse, errDB := ldDataBaseKit.GetMongoDBSession()
	if errDB != nil {
		//log to text
		//log.Println("hcLogger.Info", "errDB:", errDB.Error())
		return
	}
	defer dbse.Close()
	item := &LogItem{
		ID:       bson.NewObjectId(),
		Source:   m.Source,
		Type:     t,
		Timenano: time.Now().UnixNano(),
		Data:     data,
	}
	db := dbse.DB(database).C(GetTableName())
	if err := db.Insert(item); err != nil {
	}
	if m.WS != nil {
		m.WS.WriteMsgByChannelName("log", item)
	}
}

//GetTableName GetTableName
func GetTableName() string {
	return tableName
}
