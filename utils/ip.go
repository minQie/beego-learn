package utils

import (
	"github.com/beego/beego/v2/server/web/context"
	"net"
	"strings"
)

/*
IsLoopback：是否是本地回环地址
127.0.0.1、::1

IsPrivate：是否是内网地址（Go 1.7 新增）
10.0.0.0 ~ 10.255.255.255
172.16.0.0 ~ 172.31.255.255
192.168.0.0 ~ 192.168.255.255
*/

func IsLoopback(ipStr string) bool {
	return net.ParseIP(ipStr).IsLoopback()
}

func IsInternalIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip.IsPrivate() || ip.IsLoopback()
}

// IP beego 没有封装相应的方法，但是 gin 有
func IP(c *context.Context) string {
	var req = c.Request

	ip := strings.TrimSpace(strings.Split(req.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(req.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	return c.Input.IP()
}