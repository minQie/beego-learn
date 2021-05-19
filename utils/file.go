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
	var (
		tempFile *os.File
		err error
	)

	if err = os.MkdirAll(saveDir, os.ModeDir); err != nil {
		log.Error(fmt.Sprintf("文件的具体目录 【%s】 创建失败：%s", saveDir, err))
		return nil, err
	}
	if tempFile, err = os.OpenFile(saveDir+"/"+saveName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666); err != nil {
		log.Error(fmt.Sprintf("文件创建失败：%s", err))
		return nil, err
	}
	if _, err = io.Copy(tempFile, file); err != nil {
		log.Error(fmt.Sprintf("文件保存失败：%s", err))
		return nil, err
	}
	if err = tempFile.Close(); err != nil {
		log.Error(fmt.Sprintf("文件关闭失败：%s", err))
		return nil, err
	}

	return tempFile, nil
}

// 是否存在
func Exists(fileOrDir string) (bool, error) {
	var err error

	if _, err = os.Lstat(fileOrDir); os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, LogError("判断文件是否存在失败", err)
	}
	return true, nil
}

// 如果目标目录不存在就创建
func MkDirIfNotExists(dir, createTip string, createFailTip ...string) {
	if ok, _ := Exists(dir); !ok {
		return
	}

	logs.Info("%s %s", createTip, dir)
	tip := append(createFailTip, "文件夹创建失败")[0]

	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		panic(fmt.Sprintf("%s：%s", tip, err))
	}
}
