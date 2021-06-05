// @APIVersion 1.0
// @Title beego-learn api doc
// @Description 没有描述
// @Contact 1459049487@qq.com
package routers

import (
	"beego-learn/controllers"
	"beego-learn/controllers/trial_param"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

const (
	HelloPathPre = "/hello"
	MainPathPre  = "/param"
)

/*
关于 beego 路由的概念说明：
当前是路由的配置，决定了 请求 到 接口方法 的映射关系
但是，是间接的，beego 启动初始化时，会根据当前文件的路由配置生成一个直接生效的路由配置文件
也就是执行里边的 init 方法的逻辑，看一下就可以明白

上面的描述说明了一个问题，就是说本次 beego 运行生成的文件，将会在 beego 下一次运行生效
*/
func init() {
	// web.Router("/", &controllers.DefaultController{})

	// 方式一：想要实现什么请求方式，就是实现对应名称的方法 - 注意大小写
	// web.Router("/hello", &HelloController{})
	// 方式二：请求方式:方法名、请求方式如果有多个，就用逗号隔开、如果配置了多个规则，就用分号隔开
	// web.Router("/hello",&HelloController{},"get:Hello")
	// 方式三：通过 /hello/hello 访问 HelloController 的 hello 方法
	// web.AutoRouter(&HelloController{})
	// 方式四：开始失败 - 没有像文档中说的那样生成什么 commonRoute.go 文件（注意一定要给方法的所属结构体声明的其类型前加上 *）
	// 配置的路由是 @router /hello/:key [get] ，就得通过以 get 的请求方式，以 /hello/xxx 的请求路径请求该接口
	// web.Include(&HelloController{})
	// 方式五：开始失败 - 同上，一样需要注意结构体声明的其类型前加上 *
	// 像下面这样配置，以及同 5 中的注解路由配置，就得通过以 get 的请求方式，以 /hello/hello/xxx 的请求路径请求该接口
	v1Routers := web.NewNamespace(HelloPathPre,
		web.NSInclude(
			&controllers.HelloController{},
		),
	)

	v2Routers := web.NewNamespace(MainPathPre,
		web.NSInclude(
			&trial_param.BasicController{},
			&trial_param.ConstructController{},
			&trial_param.PathParamController{},
			&trial_param.TController{},
		),
		web.NSNamespace("/multi",
			web.NSInclude(&trial_param.MultiParamController{}),
		),
		web.NSNamespace("/cookie",
			web.NSInclude(&trial_param.CookieController{}),
		),
		web.NSNamespace("/test",
			web.NSInclude(&controllers.TestController{}),
		),
	)

	web.AddNamespace(v1Routers, v2Routers)

	// 权限
	// web.InsertFilter(MainPathPre+"/*", web.BeforeRouter, filter.Authenticate)
	// 缓存
	// web.InsertFilter(MainPathPre+"/*", web.BeforeRouter, filter.ResponseCache)
	// 跨域
	// TODO 修改时机可行么？
	web.InsertFilter("*", web.AfterExec, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "Content-Type", "Origin"},
		ExposeHeaders:    []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Content-Length", "Content-Type"},
	}))
}
