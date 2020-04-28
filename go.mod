module github.com/EDDYCJY/go-gin-example

go 1.13

require (
	github.com/astaxie/beego v1.12.1 // indirect  表单验证的表 中文文档 https://beego.me/docs/mvc/controller/validation.md
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ini/ini v1.55.0 // indirect   一个读取文本配置的包
	github.com/go-sql-driver/mysql v1.5.0 // indirect  mysql驱动
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect  Golang写的，开发人员友好的ORM库 http://gorm.book.jasperxu.com/
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => ./go-gin-example/pkg/conf

	github.com/EDDYCJY/go-gin-example/middleware => ./go-gin-example/middleware

	github.com/EDDYCJY/go-gin-example/models => ./go-gin-example/models

	github.com/EDDYCJY/go-gin-example/pkg/e => ./go-gin-example/pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => ./go-gin-example/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => ./go-gin-example/pkg/util

	github.com/EDDYCJY/go-gin-example/routers => ./go-gin-example/routers
)
