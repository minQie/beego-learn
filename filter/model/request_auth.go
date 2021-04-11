package model

import (
	"beego-learn/base/config"
	"encoding/base64"
	"errors"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ClientType int

const (
	AuthHeaderName  = "Authorization"
	AuthHeaderSplit = "#"
	POSTMAN         = 1
)

var (
	ClientKeyMap = map[string]ClientType{
		config.C.Client.PostmanClientKey: POSTMAN,
	}
)

// 请求头的信息结构
type AuthorizationHeader struct {
	ClientKey string
	Username  string
	Password  string
	Timestamp int64
}

// 唯一标识一次请求的相关信息
type RequestIdentity struct {
	Method string
	Path   string
	Params string
}

// 用于验证请求权限的相关信息
type RequestAuthorization struct {
	AuthHeader      *AuthorizationHeader
	RequestIdentity *RequestIdentity
	ClientReqTime   int64 // 客户端请求的时间戳（包含在请求头）
	ServerRecTime   int64 // 服务端接收到请求的时间戳（具体在解析请求参数时）
}

// 解析 AuthorizationHeader（base64加密的 客户端Key#用户名#密码#请求时间戳）
// 如果解析失败，将以没有权限的提示，终止请求
func ParseRequest(c *context.Context) *RequestAuthorization {
	// 将请求参数封装到对应的结构体中，r.Request.Form（所有的参数） 和 r.Request.PostForm（请求体中的参数） - 详见源码注释
	if c.Request.Form == nil {
		_ = c.Request.ParseForm()
	}

	// 解析 AuthorizationHeader
	authorizationHeader, err := ParseRequestAuthorizationHeader(c)
	if err != nil {
		logs.Error(err)
		c.Abort(http.StatusUnauthorized, strconv.Itoa(http.StatusUnauthorized))
	}

	// 将唯一标识一个请求的参数等信息 封装到自定义的结构体中
	return &RequestAuthorization{
		RequestIdentity: ParseRequestIdentity(c),
		AuthHeader:      authorizationHeader,
		ClientReqTime:   authorizationHeader.Timestamp,
		ServerRecTime:   time.Now().UnixNano() / 1e6,
	}
}

func ParseRequestIdentity(c *context.Context) *RequestIdentity {
	// RequestIdentity.Path（处理请求路径末尾的`/`）
	path := c.Request.RequestURI
	for strings.HasSuffix(path, "/") && len(path) > 1 {
		path = path[:len(path)-1]
	}

	// RequestIdentity.Params（排序，只要包含的参数相同就应该被认为是相同的请求）
	params := make([]string, 0)
	for key, values := range c.Request.Form {
		if len(values) == 0 {
			continue
		}
		sort.Strings(values)
		for _, value := range values {
			params = append(params, key+"="+value)
		}
	}
	sort.Strings(params)

	return &RequestIdentity{
		Method: c.Request.Method,
		Path:   path,
		Params: strings.Join(params, "&"),
	}
}

func ParseRequestAuthorizationHeader(c *context.Context) (*AuthorizationHeader, error) {
	// 判空
	authorizationHeaderString := c.Request.Header.Get(AuthHeaderName)
	authorizationHeaderString = strings.TrimSpace(authorizationHeaderString)
	if authorizationHeaderString == "" {
		logs.Error("认证信息解析失败，缺少 %s 请求头", AuthHeaderName)
		return nil, errors.New("请登录后重试，并联系网站管理员")
	}

	// base64 解密
	decodeBytes, err := base64.StdEncoding.DecodeString(authorizationHeaderString)
	if err != nil {
		logs.Error("认证信息解析失败，base64 %s 解密失败", authorizationHeaderString)
		return nil, errors.New("请登录后重试，并联系网站管理员")
	}
	decodeString := string(decodeBytes)

	// 获取关键信息
	info := strings.Split(decodeString, AuthHeaderSplit)
	clientKey := info[0]
	username := info[1]
	password := info[2]
	timestampString := info[3]

	// timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	timestamp, err := strconv.Atoi(timestampString)
	if err != nil {
		logs.Error("认证信息解析失败，非法的时间戳 %s", timestampString)
		return nil, errors.New("请登录后重试，并联系网站管理员")
	}

	return &AuthorizationHeader{
		ClientKey: clientKey,
		Username:  username,
		Password:  password,
		Timestamp: int64(timestamp),
	}, nil
}

func (auth *RequestAuthorization) Authenticate() error {
	// 关键信息校验
	if _, have := ClientKeyMap[auth.AuthHeader.ClientKey]; !have {
		logs.Error("认证信息解析失败，非法的用户端key %s", auth.AuthHeader.ClientKey)
		return errors.New("请登录后重试，并联系网站管理员")
	}

	// TODO 去数据库根据带有唯一索引的用户名查询
	// TODO 用户表需要额外添加一个 hash 字段

	// TODO 看下密码是否匹配

	// TODO 一个滞前 滞后的客户端请求时间和服务端收到时间 校验

	// TODO 核心签名校验（经过 SHA256 加密的数据）
	// rawStr := fmt.Sprintf("%v#%v#%v#%v", a.clientType, clientAuthKeys[a.clientType], a.ReqIdentity.AsString(), a.reqTimeMs)
	// hash := sha256.Sum256([]byte(rawStr))
	// return hex.EncodeToString(hash[:])

	return nil
}
