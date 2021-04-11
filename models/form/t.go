package form

import (
	"github.com/beego/beego/v2/core/validation"
	"strings"
)

// 源码不难，要是忘了，就用当前方法进行调试（beego validation 的里面的函数式编程是非常值得学习和借鉴的，很优秀）
// 点：beego validation 没有采用开发者自定义校验标签提示的设计

// 源码：
// 1.从 c.ParseForm → c.Input（决定后边能处理的所有参数） → ParseForm → parseFormToStruct 中可以看得到对 form 和 default 标签的处理（如 form 支持 - 符号，代表忽略）
// 2.beego\v2@v2.0.1\core\validation\util.go 中可以看到对 valid 和 label 标签属性的定义

// 紧密相关的拓展 - beego 的参数解析到结构体的注意点
// 1.url 的 path 参数肯定会解析，只有请求类型为 POST PUT PATCH 才从请求体中读取参数
// 2.结构体字段名一定需要大写（parseFormToStruct.go 282 行）
// 3.直接将请求体参数解析到实体中，并不支持小驼峰参数名到大驼峰变量名的自动转换识别
//  可以通过 form 标签手动指定；如果参数格式是 json，那么就可以通过 go 标准库的特性进行默认的小驼峰到大驼峰的转换了

// 4.Input 方法会将参数都转化进 c.Ctx.Request.Form，ParseForm 方法本质就是解析这个的
//  而进行 json 到结构体的转化，使用的是 c.Ctx.Input.RequestBody（这个对应着请求，不需要进行什么特殊处理，即可以直接拿到请求体数据）
type TForm struct {
	DevelopName  string `form:"developName"  valid:"Required"`                   // 直接使用 beego 默认提供的校验标签
	BusinessName string `form:"businessName" valid:"CustomRequired" label:"用户名"` // 如果希望自定义校验提示的 方式一
}

// 假如业务就是需要自定义校验信息，那就通过自定义校验的方式
// 方式一：实现 beego validation 的接口方法
// 注意：beego validation 会在标签校验通过后，进行回调
// （推荐这种，这是 beego 对自身校验规则不够全面的预留用户拓展）
func (f *TForm) Valid(valid *validation.Validation) {
	if len(f.DevelopName) == 0 {
		_ = valid.SetError("Name", "用户名不能为空")
	}
}

// 方式二：自定义标签
// 目前不是很推荐这种方式：
// 1、在一定程度有些违反 beego validation 模块的设计了，违背了使用初衷，作者希望自定义标签及其对应的方法，是希望给出泛用的校验规则
// 当然其他原因
// 2、技术负责人并不喜欢 beego validation 模块，不建议使用（自己觉得好，但是做不了主）
// 3、像当前这样写法的目的是拓展 beego validation，所以 beego 自身支持才是最好的解法
func init() {
	_ = validation.AddCustomFunc("CustomRequired", CustomRequired)
}

// v：  如果校验出现错误，应该将自定义错误设置进去
// obj：标签修饰的 字段值
// key：标签修饰的 字段名.当前方法名
func CustomRequired(v *validation.Validation, obj interface{}, compoundKey string) {
	if r := v.Required(obj, compoundKey); !r.Ok {
		// 删除调用 Required 方法设置进去的校验错误
		var index = -1
		for i, err := range v.Errors {
			if err == r.Error {
				index = i
			}
		}
		if index != -1 {
			v.Errors = append(v.Errors[:index], v.Errors[index+1:]...)
		}
		// 设置自定义的错误
		var (
			params = strings.Split(compoundKey, ".")
			key    = params[0]
			value  = params[len(params)-1]
		)
		_ = v.SetError(key, value+"不能为空")
	}
}
