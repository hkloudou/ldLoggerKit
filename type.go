package ldLoggerKit

import (
	"gopkg.in/mgo.v2/bson"
)

type LogType string

const (
	LTINFO  LogType = "info"
	LTDEBUG LogType = "debug"
	LTERROR LogType = "error"
	LTWARN  LogType = "warn"
)

//LogItem LogItem
type LogItem struct {
	ID       bson.ObjectId `bson:"_id" json:"i"`
	Source   string        `bson:"s" json:"s"`
	Type     LogType       `bson:"t" json:"t"`
	Data     string        `bson:"d" json:"d"`
	Timenano int64         `bson:"tn" json:"tn"`
}
