package util

import (
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//分页

func GetPage(c *gin.Context)int{
	result:=0
	//现将page参数转化为 com下的StrTo类型    Int()转化为10机制什么的
	page,_:=com.StrTo( c.Query("page") ).Int()
	if page>0 {
		result=(page-1)*setting.PageSize
	}

	return result
}
