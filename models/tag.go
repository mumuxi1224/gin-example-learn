package models

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy	string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int,pagesize int,maps interface{} )(tags []Tag){
	db.Where(maps).Limit(pagesize ).Offset(pageNum).Find(tags )
	return
}

func GetTagTotal(maps interface{})(count int){
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
