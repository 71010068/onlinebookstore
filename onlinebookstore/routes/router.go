package routes

import (
	"github.com/gin-gonic/gin"
	v1 "onlinebookstore/api/v1"
	"github.com/gin-contrib/multitemplate"
	"onlinebookstore/middleware"
	"onlinebookstore/utils"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "static/admin/index.html")
	p.AddFromFiles("front", "static/front/index.html")
	return p
}

// 路由的入口文件
func InitRouter()  {
	gin.SetMode(utils.AppMode)
	//r := gin.Default() //gin.New 默认没有中间件
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())	// 引入日志库
	r.Use(middleware.Cors())	// 跨域
	r.Use(gin.Recovery())

	//// 托管admin
	////r.LoadHTMLGlob("static/admin/index.html")
	//r.Static("admin/static","static/admin/static")
	//r.GET("admin", func(c *gin.Context) {
	//	c.HTML(200,"index.html",nil)
	//})
	//// 托管front
	//r.LoadHTMLGlob("static/front/index.html")	// 加载的页面
	//r.Static("front/static","static/front/static")
	//r.GET("front", func(c *gin.Context) {		// 托管
	//	c.HTML(200,"index.html",nil)
	//})
	r.Static("/admin", "./static/admin")
	r.Static("/front", "./static/front")
	//r.Static("/static", "./static/front/static")
	//r.StaticFile("/favicon.ico", "static/front/favicon.ico")

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})
	r.GET("/front", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})




	// 初始化路由地址
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken()) 	// 需要中间件的路由
	{	//下面这些路由接口需要权限才能访问
		// User用户模块的路由接口
		//router.GET("user",v1.GetUser)		//查询单个用户
		auth.PUT("user/:id",v1.EditUser)	//编辑用户
		auth.PUT("changepw/:id", v1.ChangeUserPassword) // 修改密码
		auth.DELETE("user/:id",v1.DeleteUser)	//删除用户
		// Category分类模块的路由接口
		auth.POST("category/add",v1.AddCategory)	//添加分类
		auth.PUT("category/:id",v1.EditCate )	//编辑分类
		auth.DELETE("category/:id",v1.DeleteCate)	//删除分类
		// Book图书模块的路由接口
		auth.POST("book/add",v1.AddBook)	//添加分类
		auth.PUT("book/:id",v1.EditBook )	//编辑分类
		auth.DELETE("book/:id",v1.DeleteBook)	//删除分类
		// 上传文件
		//auth.POST("upload",v1.UpLoad)
		// 更新个人信息
		//auth.PUT("personal/:id",v1.UpdateProfile)

		//购物车订单系统
		//auth.GET("cart/:id",v1.GetBookByUID)
		//auth.POST("cart/add",v1.AddBookInUID)
		//auth.DELETE("cart/:id",v1.DeleteCart)


	}

	// 初始化路由地址
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add",v1.AddUser)	// 添加用户
		router.GET("user/:id",v1.GetUserInfo)  // 查询单个用户
		router.GET("users",v1.GetUsers)	 // 查询用户列表
		// 图书分类
		router.GET("category",v1.GetCates)	 // 查询分类列表
		router.GET("category/:id", v1.GetCateInfo)  // 获取单个分类下的信息
		// 图书模块
		router.GET("book",v1.GetBooks)	 // 查询分类列表
		router.GET("book/list/:id",v1.GetCateBook)	// 查询单个图书列表
		router.GET("book/info/:id",v1.GetBook)	// 查询单个图书
		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)  // 前台登陆
		// 获取个人设置信息
		router.GET("personal/:id",v1.GetPersonal) // 获取个人信息设置

		// 更新个人信息
		router.PUT("personal/:id",v1.UpdateProfile)

		// 上传文件
		router.POST("upload",v1.UpLoad)

		//购物车订单系统
		router.GET("cart/:id",v1.GetBookByUID)
		router.POST("cart/add",v1.AddBookInUID)
		router.DELETE("cart/:id",v1.DeleteCart)



	}


	//启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(utils.HttpPort)
}