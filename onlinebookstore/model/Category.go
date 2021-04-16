package model

import (
	"gorm.io/gorm"
	"onlinebookstore/utils/errmsg"
)

// 分类
type Category struct {
	ID uint `gorm:"primary_key;auto_increment"  json:"id"`
	Name string  `gorm:"type:varchar(20);not null"  json:"name"`  //分类名称
}

// -------------------编写数据库对应的DAO操作的方法--------------------
// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?",name).First(&cate)
	if cate.ID > 0{	// 说明用户名已存在
		return errmsg.ERROR_CATENAME_USERD	// 2001
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCate(data *Category)(code int){

	// gorm中添加的方法：create
	err := db.Create(&data).Error
	if err != nil{ //出错
		return errmsg.ERROR	// 500
	}
	return errmsg.SUCCESS	//200

}

// 查询单个分类信息
func GetCateInfo(id int) (Category,int){
	var cate  Category
	db.Where("id = ?",id).First(&cate)
	return cate,errmsg.SUCCESS
}

// 查询分类列表
//pageSize：页数， pageNum：当前多少页     通过前端传
func GetCate(pageSize int,pageNum int) ([]Category,int64) {
	var cates []Category
	var total int64
	// 数据库查找
	// db.Scopes()
	err = db.Find(&cates).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cates).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil,0
	}
	return cates,total
}

// todo 查询分类下的所有图书

// 编辑分类信息
func EditCate(id int,data *Category) int{
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	//db.Model(&User{}).Where("id = ?",id).Updates(maps)
	err = db.Model(&cate).Where("id = ?",id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS // 200
}



// 删除分类
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?",id).Delete(&cate).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}
