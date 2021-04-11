package attachment

import (
	"beego-learn/models"
	"mime/multipart"
)

/*
-- DROP TABLE IF EXISTS `attachment`;

CREATE TABLE `attachment` (
    `id`              BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT KEY         COMMENT '主键id',
    `save_dir`        VARCHAR(255)     NOT NULL                            COMMENT '存储文件的文件夹绝对路径',
    `upload_name`     VARCHAR(64)      NOT NULL                            COMMENT '文件上传时文件名',
    `save_name`       VARCHAR(64)      NOT NULL                            COMMENT '文件存储时文件名',
    `file_extension`  VARCHAR(4)       NOT NULL                            COMMENT '文件拓展名',
    `file_size`       BIGINT UNSIGNED  NOT NULL                            COMMENT '文件大小',
    `uploaded_by`     BIGINT UNSIGNED  NOT NULL                            COMMENT '文件上传者',
    `create_time`     DATETIME         NOT NULL DEFAULT c_TIMESTAMP  COMMENT '创建时间',
    `update_time`     DATETIME         NOT NULL DEFAULT c_TIMESTAMP ON UPDATE c_TIMESTAMP  COMMENT '更新时间'
) COMMENT = '文件表';

CREATE UNIQUE INDEX `uk_save_name` ON `attachment` (save_name);
*/

// 文件存储形式 FileDir|FileExtension|CreateDate|SaveName.FileExtension
type Attachment struct {
	models.BaseEntity
	SaveDir       string `orm:"column(save_dir)"        json:"save_dir"`       // 存储文件的文件夹绝对路径（仅做记录使用，实际根据项目的配置来）
	UploadName    string `orm:"column(upload_name)"     json:"upload_name"`    // 上传时的文件名
	SaveName      string `orm:"column(file_name)"       json:"file_name"`      // 存储时的文件名
	FileExtension string `orm:"column(file_extension)"  json:"file_extension"` // 文件拓展名
	FileSize      int64  `orm:"column(file_size)"       json:"file_size"`      // 文件大小（单位：字节）
	UploadedBy    int64  `orm:"column(uploaded_by)"     json:"uploaded_by"`    // 上传者

	File multipart.File `json:"-"`
	Type *Type          `json:"-"`
}

// 描述文件类型的枚举结构体
type Type struct {
	Name string
	E    TypeEnum
}

func (a *Attachment) GetExactSaveDir() string {
	return GetExactSaveDir(a)
}

func (a *Attachment) GetFileNameWithExtension() string {
	return GetExactSaveDir(a)
}

func (a *Attachment) GetFilePath() string {
	return GetFilePath(a)
}

func (a *Attachment) Load() {
	Load(a)
}
