package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`  //	创建具有或不带名称的索引, 如果多个索引同名则创建复合索引
	Tag Tag	`json:"tag"`	//belongs to 一对一关系

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`

}

func (a *Article)BeforeCreate(scoep *gorm.Scope)error{
	scoep.SetColumn("CreatedOn",time.Now().Unix())
	return nil
}

func (a *Article)BeforeUpdate(scoep *gorm.Scope)error{
	scoep.SetColumn("ModifiedOn",time.Now().Unix())
	return nil
}


func ExistArticleById(id int)bool{
	var article Article
	db.Select("id").Where("id=?",id).First(&article)
	if article.ID>0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{})(count int){
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum,pageSize int,maps interface{})(articles []Article){
	db.Limit(pageSize).Offset(pageNum).Where(maps).Find(&articles)

	return
}