package utils

import (
	"encoding/hex"
	"github.com/satori/go.uuid"
)

// 返回去除 36位带有 4个- 的UUID string
func Get32BitUUID() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}
