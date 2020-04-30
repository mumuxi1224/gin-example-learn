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
	db.Preload("Tag").Where(maps).Limit(pageSize).Offset(pageNum).Find(&articles)
	return
}

func GetArticle(id int)(article Article){
	db.Where("id=?",id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func EditArticle(id int,data interface{})bool{
	db.Model(&Article{}).Where("id=?",id).Update(data)
	return  true
}

func AddArticle(data map[string]interface{})bool{
	db.Create(&Article{
		TagID: data["tag_id"].(int),
		Title: data["title"].(string),
		Desc: data["desc"].(string),
		Content: data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State: data["state"].(int),
	})

	return true
}

func DeleteArticle(id int)bool{
	db.Where("id=?",id).Delete(&Article{})
	return true
}
