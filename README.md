# ldLoggerKit [![Build Status](https://travis-ci.org/hkloudou/ldLoggerKit.svg?branch=master)](https://travis-ci.org/hkloudou/ldLoggerKit) [![Build Status](https://godoc.org/hkloudou/ldLoggerKit.svg?status.svg)](https://godoc.org/github.com/hkloudou/ldLoggerKit)
- ldLoggerKit is a loggerkit writen by golang
### depend
- github.com/astaxie/beego/config/env
- gopkg.in/mgo.v2/bson
- github.com/hkloudou/ldDataBaseKit
### download and update
`go get -u -v github.com/hkloudou/ldLoggerKit`

### enviroment
``` sh

##database part (the same as github.com/hkloudou/ldDataBaseKit)
ENV DB_CONFIGFILE #default conf/database.ini(the config file path)
ENV DB_CONFIGTYPE #default ini (format of config file ex.https://github.com/astaxie/beego/tree/master/config)

ENV DB_MONGO_ADDRS  #default "" example localhost,8.1.213.1
ENV DB_MONGO_DATABASE   #deault "" example testdatabase
ENV DB_MONGO_USERNAME   #default ""
ENV DB_MONGO_PASSWORD   #default ""



##log part
ENV LOG_TABLENAME   #default log_dev
```

``` go
if DB_MONGO_ADDRS == "" || DB_MONGO_DATABASE == "" {
  // read from DB_CONFIGFILE
} else {
  // use parame from env define
}
//read LOG_TABLENAME from env define

```



### example
``` go
//SYSTEM System
var SYSTEM = (&ldLoggerKit.LogBuilder{}).InitLogerFactory("system")
SYSTEM.Debugln("test")
SYSTEM.Debugf("%s","some string")
```
