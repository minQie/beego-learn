package attachment

import (
	"beego-learn/models/const/attachment"
	"mime/multipart"
	"time"
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
	Id            int64      `orm:"column(id)"             json:"id"`
	SaveDir       string     `orm:"column(save_dir)"       json:"save_dir"`       // 存储文件的文件夹绝对路径（仅做记录使用，实际根据项目的配置来）
	UploadName    string     `orm:"column(upload_name)"    json:"upload_name"`    // 上传时的文件名
	SaveName      string     `orm:"column(file_name)"      json:"file_name"`      // 存储时的文件名
	FileExtension string     `orm:"column(file_extension)" json:"file_extension"` // 文件拓展名
	FileSize      int64      `orm:"column(file_size)"      json:"file_size"`      // 文件大小（单位：字节）
	CreateTime    time.Time  `orm:"column(create_time)"    json:"-"`              // 创建时间
	CreatedBy     int64      `orm:"column(created_by)"     json:"createdBy"`      // 创建人
	UpdateTime    time.Time  `orm:"column(update_time)"    json:"-"`              // 更新时间
	UpdatedBy     int64      `orm:"column(updated_by)"     json:"updatedBy"`      // 更新人
	DeleteTime    *time.Time `orm:"column(is_deleted)"     json:"-"`              // 删除标识
	DeleteBy      int64      `orm:"column(delete_by)"      json:"-"`              // 删除人

	File multipart.File `json:"-"`
	Type *Type          `json:"-"`
}

// 描述文件类型的枚举结构体
type Type struct {
	Name string
	E    attachment.TypeEnum
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
