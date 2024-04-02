package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-admin/app/global/consts"
	"go-admin/app/global/variable"
	"go-admin/app/http/controller/captcha"
	"go-admin/app/http/middleware/authorization"
	"go-admin/app/http/middleware/cors"
	validatorFactory "go-admin/app/http/validator/core/factory"
	"go-admin/app/utils/gin_release"
	"go.uber.org/zap"
	"net/http"
)

// 该路由主要设置 后台管理系统等后端应用路由
func InitWebRouter() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		router = gin_release.ReleaseRouter()
	} else {
		router = gin.Default()
		pprof.Register(router)
	}

	// 设置可信任的代理服务器列表,gin (2021-11-24发布的v1.7.7版本之后出的新功能)
	if variable.ConfigYml.GetInt("HttpServer.TrustProxies.IsOpen") == 1 {
		if err := router.SetTrustedProxies(variable.ConfigYml.GetStringSlice("HttpServer.TrustProxies.ProxyServerList")); err != nil {
			variable.ZapLog.Error(consts.GinSetTrustProxyError, zap.Error(err))
		}
	} else {
		_ = router.SetTrustedProxies(nil)
	}

	// 根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld,这是后端模块")
	})

	// 处理静态资源
	router.Static("/public", "./public")             //  定义静态资源路由与实际目录映射关系
	router.StaticFS("/dir", http.Dir("./public"))    // 将public目录内的文件列举展示
	router.StaticFile("/abcd", "./public/readme.md") // 可以根据文件名绑定需要返回的文件名

	// 创建一个验证码路由
	verifyCode := router.Group("captcha")
	{
		// 验证码业务，该业务无需专门校验参数，所以可以直接调用控制器
		verifyCode.GET("/", (&captcha.Captcha{}).GenerateId)                          //  获取验证码ID
		verifyCode.GET("/:captcha_id", (&captcha.Captcha{}).GetImg)                   // 获取图像地址
		verifyCode.GET("/:captcha_id/:captcha_value", (&captcha.Captcha{}).CheckCode) // 校验验证码
	}

	//  创建一个后端接口路由组
	backend := router.Group("/admin/")
	{
		backend.GET("ws", validatorFactory.Create(consts.ValidatorPrefix+"WebsocketConnect"))

		// 不需要token中间件
		noAuth := backend.Group("users/")
		{
			noAuth.POST("register", validatorFactory.Create(consts.ValidatorPrefix+"UsersRegister"))
			noAuth.POST("login", validatorFactory.Create(consts.ValidatorPrefix+"UsersLogin"))
		}

		// 刷新token
		refreshToken := backend.Group("users/")
		{
			refreshToken.Use(authorization.RefreshTokenConditionCheck()).POST("refreshToken", validatorFactory.Create(consts.ValidatorPrefix+"RefreshToken"))
		}
		// 需要token的中间件
		backend.Use(authorization.CheckTokenAuh())
		{
			// 用户组路由
			users := backend.Group("users/")
			{
				// 查询 ，这里的验证器直接从容器获取，是因为程序启动时，将验证器注册在了容器，具体代码位置：App\Http\Validator\Web\Users\xxx
				users.GET("index", validatorFactory.Create(consts.ValidatorPrefix+"UsersShow"))
				// 新增
				users.POST("create", validatorFactory.Create(consts.ValidatorPrefix+"UsersStore"))
				// 更新
				users.POST("edit", validatorFactory.Create(consts.ValidatorPrefix+"UsersUpdate"))
				// 删除
				users.POST("delete", validatorFactory.Create(consts.ValidatorPrefix+"UsersDestroy"))
			}
			// 文件上传公共路由
			uploadFiles := backend.Group("upload/")
			{
				uploadFiles.POST("files", validatorFactory.Create(consts.ValidatorPrefix+"UploadFiles"))
			}
		}
		return router
	}

}
