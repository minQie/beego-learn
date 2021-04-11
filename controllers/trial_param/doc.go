package trial_param

/* 测试的接口方法与起对应的路由方法明明规则：
在需要对比请求方式的情况下：/方法名 [请求方式] ←→ 方法名请求方式
不需要对比请求方式的情况再：/方法名 [请求方式] ←→ 方法名
*/

// ------------------------------------------------------

/* 遵从 HTTP 协议，当请求方式为 GET 时，beego 会忽略请求体中的参数 */

/* POST 提交请求体参数的 form-data 和 x-www-form-urlencoded 方式的区别

共同点：
1.浏览器原生支持

一、x-www-form-urlencoded
对应请求头：application/x-www-form-urlencoded
简介：将 name、value 中的空格替换为加号；将非 ascii 字符做百分号编码；将 input 的 name、value 用 = 连接，不同的input之间用 & 连接

二、form-data
对应请求头：multipart/form-data（键值对都是通过 & 间隔分开的）
简介：将表单的数据处理为一条消息，以标签为单元，用分隔符分开。既可以上传键值对，也可以上传文件。当上传的字段是文件时，
会有 Content-Type 来表明文件类型；content-disposition，用来说明字段的一些信息；
由于有 boundary 隔离，所以 multipart/form-data 既可以上传文件，也可以上传键值对，它采用了键值对的方式，所以可以上传多个文件

每个 input 转为了一个由 boundary 分割的小格式，没有转码，直接将 utf8 字节拼接到请求体中，在本地有多少字节实际就发送多少字节，极大提高了效率，适合传输长字节
如：Content-Type: multipart/form-data; boundary=----WebKitFormBoundarymNhhHqUh0p0gfFa8（在这里指明作为分割线的字符是什么）
*/

// ------------------------------------------------------

/* beego 源码 - 参数模块

c.Ctx.Request.Form
 包含 URL 参数 和 表单参数（不包含 json 参数）
c.Ctx.RequestBody
 包含 请求体中参数
*/

// ------------------------------------------------------

/* Cookie

由于使用 cookie 可能带来安全性问题，所以我们最好在设置 cookie 时，指定两个有利于安全的选项

https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Cookies#%E9%99%90%E5%88%B6%E8%AE%BF%E9%97%AE_cookie

	1、设置 Cookie
	参数顺序对应的含义：key, value, max-age（有效期：秒）, path, domain, secure（只在 https 协议下传送 cookie）, httponly（让 js 无法访问 cookie）
	httponly：参数保证了 Cookie 的安全性，但是带来了额外的麻烦，就是退出登录（清除 Cookie），只能调用接口，通过后端来清除了
	c.Ctx.SetCookie("token, "123", 60, "", "", true, true)

	2、清除 Cookie
	c.Ctx.SetCookie("token", "", -1)

	3、获取 Cookie
	c.Ctx.GetCookie("token")
*/
