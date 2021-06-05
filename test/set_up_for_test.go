package test

import (
	"beego-learn/modules/config"
	_ "beego-learn/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func record() {
	// 直接运行下边的 test 运行会报错，原因是当前项目下的工具包 help 几乎在项目中的任何地方都有被引用
	// 而在 help 包下的某个文件中，有着读取 beego 配置的逻辑，且读取不到，就 panic
	// 实际就是 beego 没有读取到配置文件，控制台中打头的 warn 级别日志就可以看出来

	// 首先 beego 读取配置文件，可以追溯到底层的 os.Open(filename)，这个方法在 beego 实际调用时，参数值是 “conf/app.conf”
	// 也就是相对路径，也就是平时运行项目和这里实际相对的路径是不同的

	// 验证方式一：在当前目录下，放置一个 conf/app.conf（好处：单独维护一个测试环境的配置文件；坏处：这个配置文件得随着实际运行的代码文件变更目录）
	// 验证方式二：将 help 初始化配置的代码注释掉，然后在这里维护配置对象（不推荐）
	// 验证方式三：突然想到的，goland 的运行配置应该可以指定这个路径，对应修改了一下 Working directory，发现运行可以了（但感觉还不是一劳永逸的）
	// 验证方式四：通过 go 的基础 api，打印一下空的相对路径对应的绝对路径

	// 注意，通过查阅源码注释发现，你想获取当前目录的绝对路径，根据规范应该使用 "."，但是 go 会有将 "" → "." 的处理
	// 通过打印当前路径的方法，也再次验证了，goland 的 Working directory 就是设置这个基础路径的

	// 问题2：register db `default`, sql: unknown driver "mysql" (forgotten import?)（解决：添加导包，不能自己添加包下的代码，不然又去加载那个包，就重复了）
	// 问题1：must have one register DataBase alias named `default`（解决：添加下面这一段）
	// if err := orm.RegisterDataBase("default", config.C.DB.DbName, config.C.DB.DSN()); err != nil {
	// 	panic(err)
	// }

	// 问题3：这个配置有什么用（下面的配置在 beego 2.0 后不需要了，官方文档的 ORM 使用有说到默认支持了）
	// _ = orm.RegisterDriver("mysql", orm.DRMySQL)
}

// 上面是过程记录，下面是项目实战中如何设置好各项的环境
func init() {
	// 1、导入 mysql 驱动包
	// 2、加载自定义配置（不能指望修改工作路径就能起到执行 main.go 中 init 的效果，根路径中的方法不可能被其他包引用，也就是根据 go 的加载规则，不会加载到 main.go）
	config.YamlConfigLoad("./conf/app_dev.yml")
	// 3、注册默认连接配置
	if err := orm.RegisterDataBase("default", config.C.DB.DbName, config.C.DB.DSN()); err != nil {
		panic(err)
	}
}

func TestT(t *testing.T) {
	fmt.Println(123)
}
