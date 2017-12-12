/**********************************************
** @Des: This file ...
** @Author: yaoxf
** @Date:   2017-12-01 00:16:00
** @Last Modified by:   yaoxf
** @Last Modified time: 2017-12-01 17:26:48
***********************************************/

package models

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// mysql配置
type DbConf struct {
	DbHost     string
	DbUser     string
	DbPasswd   string
	DbPort     int
	DbName     string
	DbPrefix   string
	DbTimezone string
}

var dbConf DbConf

// 初始化MySQL数据库
func Init() (err error) {
	// 初始化配置
	err = InitConfig()
	if err != nil {
		return errors.New("db conf error, plase check.")
	}

	// 设置默认DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbConf.DbUser, dbConf.DbPasswd, dbConf.DbHost, dbConf.DbPort, dbConf.DbName)
	if dbConf.DbTimezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(dbConf.DbTimezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn, 30, 30)

	// 注册Model
	orm.RegisterModel(new(Admin), new(Auth), new(Role))

	// debug开关
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	return nil
}

// 初始化DB配置
func InitConfig() (err error) {

	// 数据库配置
	dbHost := beego.AppConfig.String("db.host")
	dbUser := beego.AppConfig.String("db.user")
	dbPasswd := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	dbPrefix := beego.AppConfig.String("db.prefix")
	dbTimezone := beego.AppConfig.String("db.timezone")

	// conf check
	dbPort, err := beego.AppConfig.Int("db.port")
	if err != nil {
		return errors.New("db.port conf error.")
	}

	if len(dbHost) == 0 || len(dbUser) == 0 || len(dbPasswd) == 0 || len(dbName) == 0 {
		return errors.New("db conf error, plase check.")
	}
	dbConf.DbHost = dbHost
	dbConf.DbUser = dbUser
	dbConf.DbPasswd = dbPasswd
	dbConf.DbPort = dbPort
	dbConf.DbName = dbName
	dbConf.DbPrefix = dbPrefix
	dbConf.DbTimezone = dbTimezone

	return nil
}

// 拼接完整表名称
func TableName(name string) string {
	return dbConf.DbPrefix + name
}

// 计算分页offset
func LimitOffset(page, pageSize int) int {
	return (page - 1) * pageSize
}
