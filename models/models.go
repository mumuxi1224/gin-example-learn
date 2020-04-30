package models

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init(){
	var (
		err error
		dbType,dbName,user,password,host,tablePrefix string
	)

	sec,err:=setting.Cfg.GetSection("database")
	if err!=nil {
		log.Fatal(2,"Fail to get section 'database':%v",err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	//mysql的连接 "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	db,err=gorm.Open( dbType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName) )

	if err!=nil{
		log.Println(err)
	}

	//默认表名  加上前缀
	gorm.DefaultTableNameHandler= func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix+defaultTableName
	}

	// 全局禁用表名复数    不开启的话 自动创建的表名后面会加上s
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	//详细日志 set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
	db.LogMode(true)

	//连接池
	db.DB().SetMaxIdleConns(10)  //设置闲置的连接数
	db.DB().SetMaxOpenConns(100)	//设置最大打开的连接数
}

func CloseDB(){
	defer db.Close()
}
