package utils

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/prometheus/common/log"
	"io"
	"mime/multipart"
	"os"
)

// 文件数据 file 保存到指定的目录 saveDir
func SaveFile(saveDir, saveName string, file multipart.File) (*os.File, error) {
	err := os.MkdirAll(saveDir, os.ModeDir)
	if err != nil {
		log.Error(fmt.Sprintf("文件的具体目录 【%s】 创建失败：%s", saveDir, err))
		return nil, err
	}

	tempFile, err := os.OpenFile(saveDir+"/"+saveName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		log.Error(fmt.Sprintf("文件创建失败：%s", err))
		return nil, err
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		log.Error(fmt.Sprintf("文件保存失败：%s", err))
		return nil, err
	}

	// TODO 文件 close 一定要用 defer么
	err = tempFile.Close()
	if err != nil {
		log.Error(fmt.Sprintf("文件关闭失败：%s", err))
		return nil, err
	}

	return tempFile, nil
}

// 是否存在
func Exists(fileOrDir string) bool {
	_, err := os.Lstat(fileOrDir)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		logs.Error("读取文件报错：%s", err)
		return false
	}
	return true
}

// 如果目标目录不存在就创建
func MkDirIfNotExists(dir, createTip string, createFailTip ...string) {
	if Exists(dir) {
		return
	}

	logs.Info("%s %s", createTip, dir)

	tip := "文件夹创建失败"
	if len(createFailTip) != 0 {
		tip = createFailTip[0]
	}
	// 创建
	err := os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		panic(fmt.Sprintf("%s：%s", tip, err))
	}
}
