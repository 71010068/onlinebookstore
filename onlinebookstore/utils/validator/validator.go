package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"onlinebookstore/utils/errmsg"
	"reflect"
)

// 数据验证
func Validate(data interface{}) (string,int){
	// 实例化
	validate := validator.New()
	// 翻译实例化
	uni := unTrans.New(zh_Hans_CN.New())  // 转换成汉语中文
	trans , _ := uni.GetTranslator("zh_Hans_CN")

	// 对翻译方法进行注册
	err := zhTrans.RegisterDefaultTranslations(validate,trans)
	if err != nil{
		fmt.Println("err：",err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	// 对传进来打data进行验证
	err = validate.Struct(data)
	if err != nil{
		for _,v :=range err.(validator.ValidationErrors){
			return v.Translate(trans),errmsg.ERROR
		}
	}
	// 成功
	return "",errmsg.SUCCESS
}