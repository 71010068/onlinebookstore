// 错误处理信息
package errmsg


const (
	SUCCESS = 200
	ERROR = 500
	// code = 1000... 用户模块错误
	ERROR_USERNAME_USED 	= 1001		// 用户名已使用
	ERROR_PASSWORD_WRONG 	= 1002		// 用户密码错误
	ERROR_USER_NOT_EXIST 	= 1003  	// 用户不存在
	ERROR_TOKEN_EXIST	 	= 1004		// 用户携带token不存在
	ERROR_TOKEN_RUNTIME 	= 1005		// 用户携带token超时
	ERROR_TOKEN_WRONG 		= 1006		// 用户携带token错误
	ERROR_TOKEN_TYPE_WRONG 	= 1007		// token格式不正确
	ERROR_USER_NO_RIGHT		= 1008		// 用户没有管理权限
	// code = 2000... 图书模块错误
	ERROR_BOOK_NOT_EXIST	= 2001		//查询的图书不存在
	// code = 3000... 分类模块错误
	ERROR_CATENAME_USERD	= 3001		//创建分类名存在
	ERROR_CATE_NOT_EXIST 	= 3002		//创建分类不存在
)
// map：int是错误码，string是错误描述
var codeMsg = map[int]string{
	SUCCESS : "OK",
	ERROR 	: "FAIL",
	ERROR_USERNAME_USED 	: 	"用户名已存在！",
	ERROR_PASSWORD_WRONG 	: 	"密码错误！",
	ERROR_USER_NOT_EXIST 	: 	"用户不存在！",
	ERROR_TOKEN_EXIST 		: 	"TOKEN不存在。",
	ERROR_TOKEN_RUNTIME		: 	"TOKEN已过期。",
	ERROR_TOKEN_WRONG 		: 	"TOKEN不正确。",
	ERROR_TOKEN_TYPE_WRONG	:	"TOKEN格式错误。",
	ERROR_USER_NO_RIGHT		:	"该用户无权限",

	ERROR_BOOK_NOT_EXIST	:	"该图书不存在。",

	ERROR_CATENAME_USERD	:	"该分类已存在。",
	ERROR_CATE_NOT_EXIST	:	"该分类不存在。",

}

// 输出错误信息的函数
func GetErrMsg(code int) string{
	return codeMsg[code]
}