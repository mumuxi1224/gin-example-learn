package jwt

import (
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT()gin.HandlerFunc{
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code=e.SUCCESS
		token:=context.Query("token")
		if token=="" {
			code=e.INVALID_PARAMS
		}else{
			cliams,err:=util.ParseToken(token)
			if err!=nil {
				code=e.ERROR_AUTH_CHECK_AUTH_FAIL
			}else if time.Now().Unix()>cliams.ExpiresAt {
				code=e.ERROR_AUTH_CHECK_AUTH_TIMEOUT
			}
		}

		if code!=e.SUCCESS{
			context.JSON(http.StatusUnauthorized,gin.H{
				"code":code,
				"msg":e.GetMsg(code),
				"data":data,
			})

			context.Abort()  //防止调用挂起的处理程序 主要用在鉴权中
			return
		}

		context.Next()
	}
}
