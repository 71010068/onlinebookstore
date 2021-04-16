package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"onlinebookstore/utils/errmsg"
)

type User struct {

	gorm.Model
	//gorm 提供的模型参数
	Username string	 `gorm:"type:varchar(20);not null"  json:"username" validate:"required,min=4,max=12" label:"用户名"`		// 用户名
	Password string  `gorm:"type:varchar(200);not null"  json:"password" validate:"required,min=6,max=20" label:"密码"`		// 密码
	Phone int		 `gorm:"type:int;not null"  json:"phone" validate:"required" label:"手机号"`		//手机号
	Email string   	 `gorm:"type:varchar(50);not null"  json:"email" validate:"required" label:"邮箱"`		//邮箱
	Role int		 `gorm:"type:int;DEFAULT:2"  json:"role" validate:required,get=2 label:"角色码"`		//角色

	Name   string `gorm:"type:varchar(20)" json:"name"`
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	Qqchat string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat string `gorm:"type:varchar(100)" json:"wechat"`
	Avatar string `gorm:"type:varchar(200)" json:"avatar"`
}

// -------------------编写数据库对应的DAO操作的方法--------------------
// 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?",name).First(&users)
	if users.ID > 0{	// 说明用户名已存在
		return errmsg.ERROR_USERNAME_USED	// 1001
	}
	return errmsg.SUCCESS
}

// 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

// 添加用户
func CreateUser(data *User)(code int){
	// 密码进入数据库前进行hash加盐加密
	//data.Password = ScryptPw(data.Password)

	// gorm中添加的方法：create
	err := db.Create(&data).Error
	if err != nil{ //出错
		return errmsg.ERROR	// 500
	}
	return errmsg.SUCCESS	//200

}

// 查询单个用户
func GetUser(id int) (User,int){
	var user User
	err = db.Where("ID = ?",id).First(&user).Error
	if err != nil{
		return user,errmsg.ERROR
	}
	return user,errmsg.SUCCESS
}

// 查询用户列表
//pageSize：页数， pageNum：当前多少页     通过前端传
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id,username,phone,email,role").Where("username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)

		db.Model(&users).Where("username LIKE ?", username+"%", ).Count(&total)
		return users, total
	}
	db.Select("id,username,phone,email,role").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	if err != nil {
		return users, 0
	}
	return users, total

//func GetUsers(pageSize int,pageNum int) ([]User,int64) {
//	var users []User
//	var total int64
//	// 数据库查找
//	db.Scopes()
//	err = db.Limit(pageSize).Offset((pageNum-1) * pageSize).Find(&users).Count(&total).Error
//	if err != nil && err != gorm.ErrRecordNotFound{
//		return nil,0
//	}
//	return users,total
//}
}




// 编辑用户信息
func EditUser(id int,data *User) int{
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["phone"] = data.Phone
	maps["email"] = data.Email
	maps["role"] = 	data.Role
	//maps["name"] = data.Name
	//maps["desc"] = data .Desc
	//maps["qq_chat"] = data.Qqchat
	//maps["wechat"] = data.Wechat
	//maps["avatar"] = data.Avatar
	//db.Model(&User{}).Where("id = ?",id).Updates(maps)
	err = db.Model(&user).Where("id = ?",id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS // 200
}

// 修改密码
func ChangePassword(id int, data *User) int {
	//var user User
	//var maps = make(map[string]interface{})
	//maps["password"] = data.Password
	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?",id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}

// 密码加密
// 钩子函数完成入库前hash加盐加密, 不需要手动调用，框架自动调用。方法名确定：BeforeCreate()
func (u *User)BeforeCreate(tx *gorm.DB)(err error){
	u.Password = ScryptPw(u.Password)
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}
//func (u *User) BeforeSave(_ *gorm.DB) (err error) {
//	u.Password = ScryptPw(u.Password)
//	return nil
//}
//
//func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
//	u.Password = ScryptPw(u.Password)
//	return nil
//}
// 生成密码
func ScryptPw(password string)string{
	//const KeyLen = 10
	//salt := make([]byte,8)
	//salt = []byte{1,3,1,4,5,2,0,2}
	//
	//HashPw , err := scrypt.Key([]byte(password) , salt,16384,8,1,KeyLen)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fpw := base64.StdEncoding.EncodeToString(HashPw)
	//return fpw
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

//----------------------------编写控制登陆的的函数---------------------------------------
// 登陆验证
//func CheckLogin(username , password string) int{
//	var user User
//
//	db.Where("username = ?", username).First(&user)
//	if user.ID == 0{
//		return errmsg.ERROR_USER_NOT_EXIST
//	}
//	if ScryptPw(password) != user.Password{
//		return errmsg.ERROR_PASSWORD_WRONG
//	}
//	if user.Role != 1 {
//		return errmsg.ERROR_USER_NO_RIGHT
//	}
//	return errmsg.SUCCESS
//}
// 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCESS
}
// 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}


//----------------------------编写个人信息设置的函数---------------------------------------
//获取个人信息设置
func GetPersonal(id int) (User, int) {
	var user User
	err = db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// 更新个人信息设置
func UpdatePersonal(id int, data *User) int {
	var user User
	err = db.Model(&user).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//----------------------------编写购物车的函数---------------------------------------
// 查询添加图书列表