package v1

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context){
	name:=c.Query("name")

	maps:=make(map[string]interface{})
	data:=make(map[string]interface{})

	if name!="" {
		maps["name"]=name
	}
	var state=-1
	if arg:=c.Query("state");arg!=""{
		state=com.StrTo( arg ).MustInt()
		maps["state"] = state
	}

	code:=e.SUCCESS
	data["lists"] = models.GetTags( util.GetPage(c),setting.PageSize,maps )
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}
//添加文章标签
func AddTag(c *gin.Context){
	name:=c.Query("name")
	state:=com.StrTo( c.DefaultQuery("state","0") ).MustInt()
	createdBy:=c.Query("created_by")

	valid:=validation.Validation{}
	valid.Required(name,"name").Message("名称不能为空")
	valid.MaxSize(name,100,"name").Message("名称不能超过100个字符")
	valid.Required(createdBy,"created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy,100,"created_by").Message("创建人不能超过100个字符")
	valid.Range(state,0,1,"state").Message("状态准许为0或1")

	fmt.Println("name:",name,",state:",state,",createdby:",createdBy)
	code:=e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code=e.SUCCESS
			models.AddTag(name,state,createdBy)
		}else{
			code=e.ERROR_EXIST_TAG
		}
	}else{
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})

}
//编辑文章标签
func EditTag(c *gin.Context){
	id:=com.StrTo( c.Param("id") ).MustInt()
	name:=c.Query("name")
	modifiedBy:=c.Query("modified_by")
	state:=-1
	valid:=validation.Validation{}
	if arg:=c.Query("state");arg!="" {
		state=com.StrTo( arg).MustInt()
		valid.Range(state,0,1,"state").Message("状态只允许0或1")
	}

	valid.Required(id,"id").Message("ID不能为空")
	valid.Required(modifiedBy,"modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy,100,"修改人不能超过100个字符")
	valid.MaxSize(name,100,"名称不能超过100个字符")

	code:=e.INVALID_PARAMS

	if !valid.HasErrors() {
		code=e.SUCCESS
		if models.ExistTagById(id){
			data:=make(map[string]interface{})
			data["modified_by"] =modifiedBy
			if name!=""{
				data["name"]=name
			}
			if state!=-1{
				data["state"]=state
			}
			models.EditTag(id,data)
		}
	}else{
		code=e.ERROR_NOT_EXIST_TAG
	}

	c.JSON(code,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})

}
//删除文章标签
func DeleteTag(c *gin.Context){
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")


	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})

}