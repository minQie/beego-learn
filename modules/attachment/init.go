package attachment

import (
	"beego-learn/modules/config"
	"beego-learn/utils"
)

// 初始化文件大类别的存储目录
func Init() {
	saveDir := config.C.Attachment.SaveDir
	docDirName := config.C.Attachment.DocDirName
	picDirName := config.C.Attachment.PicDirName

	utils.MkDirIfNotExists(saveDir, "创建文件存储主目录", "创建文件存储主目录失败")
	utils.MkDirIfNotExists(saveDir+"/"+docDirName, "创建文档存储目录", "创建文档存储目录失败")
	utils.MkDirIfNotExists(saveDir+"/"+picDirName, "创建图片存储目录", "创建图片存储目录失败")
}
