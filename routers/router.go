package routers

import (
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine{
	r:=gin.New()
	r.Use( gin.Logger() )
	r.Use(gin.Recovery() )
	gin.SetMode( setting.RunMode )  //设置运行模式

	r.GET("/test", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"test",
		})
	})

	apiv1:=r.Group("/api/v1")
	{
		//获取tag
		apiv1.GET("/tags",v1.GetTags)
		//新增tag
		apiv1.POST("/tags",v1.AddTag)
		//修改tag
		apiv1.PUT("tags/:id",v1.EditTag)
		//删除tag
		apiv1.DELETE("tags/:id",v1.DeleteTag)

		//获取指定文章
		apiv1.GET("/articles/:id",v1.GetArticle)
		//获取文章列表
		apiv1.GET("/articles",v1.GetArticles)
		//新增article
		apiv1.POST("/articles",v1.AddArticle)
		//修改atricle
		apiv1.PUT("/articles/:id",v1.EditArticle)
		//删除article
		apiv1.DELETE("/articles/:id",v1.DeleteArticle)
	}

	return r
}
