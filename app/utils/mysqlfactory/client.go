package mysqlfactory

import (
	"database/sql"
	"fmt"
	"go-admin/app/global/errors"
	"go-admin/app/utils/helper"
	"log"
	"time"
)

// InitSqlDriver 初始化数据库驱动
func InitSqlDriver() *sql.DB {
	configFac := helper.CreateYamlFactory()
	dbType := configFac.GetString("dbType")
	host := configFac.GetString("mysql.host")
	port := configFac.GetString("mysql.port")
	user := configFac.GetString("mysql.user")
	password := configFac.GetString("mysql.password")
	dataBase := configFac.GetString("mysql.dataBase")
	maxIdleConns := configFac.GetInt("mysql.maxIdleConns")
	maxOpenConns := configFac.GetInt("mysql.maxOpenConns")
	connMaxLifetime := configFac.GetDuration("mysql.connMaxLifetime")
	// 连接数据库
	db, err := sql.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dataBase))
	if err != nil {
		log.Fatal(errors.Errors_Db_SqlDriverInitFail, err)
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime + time.Second)
	return db
}

// 从连接池获取一个连接
func GetOneEffectivePing() *sql.DB {
	configFac := helper.CreateYamlFactory()
	maxRertyTimes := configFac.GetInt("mysql.pingFailRetryTimes")
	// ping 失败运行重试
	vDbDriver := InitSqlDriver()
	for i := 1; i <= maxRertyTimes; i++ {
		if err := vDbDriver.Ping(); err != nil { // 获取一个连接失败，进行重试
			vDbDriver = InitSqlDriver()
			time.Sleep(time.Second * 1)
			if i == maxRertyTimes {
				log.Fatal(errors.Errors_Db_GetConnFail, err)
			}
		} else {
			break
		}
	}
	return vDbDriver
}
