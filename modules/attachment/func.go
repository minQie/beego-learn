package attachment

import (
	"beego-learn/base/config"
	"beego-learn/utils"
	"fmt"
	"os"
	"path"
	"strings"
)

// 根据文件拓展名获取文件类型，如何没有匹配到全局配置的文件类型，则返回 nil
func GetAttachmentType(fileExtension string) (t *Type) {
	lowerExtension := strings.ToLower(fileExtension)

	for _, extension := range config.C.Attachment.DocSupportType {
		if lowerExtension == extension {
			t = &Type{
				Name: extension,
				E:    DOC,
			}
		}
	}

	for _, extension := range config.C.Attachment.PicSupportType {
		if lowerExtension == extension {
			t = &Type{
				Name: extension,
				E:    PIC,
			}
		}
	}

	return
}

// 获取文件名
func GetFileNameWithExtension(a *Attachment) string {
	return a.SaveName + "." + a.FileExtension
}

// 获取文件存储目录
func GetExactSaveDir(a *Attachment) string {
	return path.Join(
		config.C.Attachment.SaveDir,
		a.FileExtension,
		utils.DateTimeToDateString(a.CreateTime),
	)
}

// 获取完整文件名（文件存储目录 + 文件名）
func GetFilePath(a *Attachment) string {
	return path.Join(
		GetExactSaveDir(a),
		GetFileNameWithExtension(a),
	)
}

// 数据从数据库查询出来，需要对一些额外的字段进行赋值
func Load(a *Attachment) {
	if a.Type == nil {
		a.Type = GetAttachmentType(a.FileExtension)
	}
	if a.File == nil {
		filePath := a.GetFilePath()
		file, err := os.Open(filePath)
		if err != nil {
			panic(fmt.Sprintf("加载文件：%s, 失败：%s", filePath, err))
		}
		if file == nil {
			panic(fmt.Sprintf("文件不存在：%s", filePath))
		}
		a.File = file
	}
}
