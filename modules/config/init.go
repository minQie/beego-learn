package config

import (
	"fmt"
	_ "github.com/beego/beego/v2/core/config/yaml"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Read(configs ...string) {
	conf := append(configs, "conf/app.yml")[0]

	// 就只想留下一个配置文件 app.yaml
	//  本来想着，无视默认的 beego/v2/config.init 中，没有读取到 conf/app.conf 的 Warn 日志
	//  然后，手动加载 conf/app/yaml
	// 	但是实际发现你手动加载的配置文件，并不会放到 Beego 配置文件对应配置实例中
	//  真正的配置实例实在 server/web.init 中，而这里边没有拓展只读取 conf/app.conf
	//    相关变量有：web.BConfig 和 web.AppConfig
	//    可以看到这样一句：// now only support ini, next will support json.

	// TODO bConfig.InitGlobalInstance 到底有什么用，web.BConfig 在使用 conf/app.conf 配置时，是会受到影响的
	//   这里的表现却没有，依旧是默认值
	//   最终发现：在 web.init 最后一条语句 parsConfig 中 → newAppConfig 没有报错，就执行 assignConfig → parseConfigForV1 中
	//    → 将 newAppConfig 得到的 AppConfig 的值赋值到 BConfig, &BConfig.Listen, &BConfig.WebConfig, &BConfig.Log, &BConfig.WebConfig.Session

	// 其实就不是说，反正像下面这样可以把 yml 的配置加载到 beego 内部的一个配置实例中（虽然这样弄启动时还有警告）
	// 最根本的是，这样配置对 Beego 内部的机制不起作用，也就是你配个端口什么的，不起作用，这个配置还要个毛

	// 结论1：结合上面的分析结果，就 web.init 方法中决定 AppConfig 数据源的配置文件就是写死的 app.conf，所以 Beego 默认的配置文件你根本没法自定义、拓展
	//    瞎搞的思路：名字是 app.conf 实际内容是 app.yaml，然后，也是将里边写死的 ini → IniConfig 解析，改成 ini → yaml.Config
	//    感觉可行，但是你没法在你希望的那个时间执行这个替换操作：config.init → web.init → 你的替换操作...
	// 结论2：bConfig.InitGlobalInstance 没任何卵用，实际查看赋值给的变量都没有地方使用到

	// beego 读取必要的服务配置
	if err := web.LoadAppConfig("yaml", conf); err != nil {
		logs.Warn("init global config instance failed. If you donot use this, just ignore it. ", err)
	}
	// 自定义实体配置
	YamlConfigLoad(conf)
}

// 配置文件名，conf 目录下找起
func YamlConfigLoad(config string) {
	yamlFile, err := ioutil.ReadFile(config)
	if err != nil {
		panic(fmt.Sprintf("读取自定义配置文件失败：%s", err))
	}
	if yamlFile == nil {
		panic("没有找到配置文件")
	}

	// 数据库配置文件加载
	// note: 不使用 beego 的 yaml 模块，是因为 beego 的 yaml 模块只支持，将某个 yaml 文件解析成一个 map
	if err = yaml.Unmarshal(yamlFile, &C); err != nil {
		panic(fmt.Sprintf("加载自定义配置失败：%s", err))
	}
}
