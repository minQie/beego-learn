package controllers

import (
	"beego-learn/modules/attachment"
	"beego-learn/modules/config"
	"beego-learn/services"
	"beego-learn/utils"
	"errors"
	"github.com/alecthomas/units"
	"mime/multipart"
	"net/http"
	"strings"
)

type AttachmentController struct {
	BaseController
	s services.AttachmentService
}

func (c *AttachmentController) Prepare() {
	c.s = services.NewAttachment()
}

// @Title 附件上传
// @Description 文件上传管理之新增文件
// @Param files formData binary true 文件列表
func (c *AttachmentController) Upload() {
	// TODO 预期外异常的拦截，保证接口调用者获取定义好的数据格式
	// defer utils.HttpPanic(func(msg string) {
	//     c.ResponseJson(nil, errors.New(http.StatusInternalServerError, msg))
	// })

	var (
		userId      = c.GetLoginUserId()
		fileHeaders []*multipart.FileHeader
		file        multipart.File
		attachments []*attachment.Attachment
		err         error
	)

	if fileHeaders, err = c.GetFiles("files"); err == http.ErrMissingFile {
		c.ResponseJson(nil, utils.LogError("请指定要上传的文件"))
		return
	}
	if err != nil {
		c.ResponseJson(nil, utils.LogError("获取上传文件信息失败", err))
		return
	}
	if len(fileHeaders) == 0 {
		c.ResponseJson(nil, utils.LogError("获取上传文件信息失败", errors.New("没有获取到相关的文件头信息")))
		return
	}

	// 通过校验后待保存的附件数据
	attachments = make([]*attachment.Attachment, len(fileHeaders))
	for index, fileHeader := range fileHeaders {
		filename := fileHeader.Filename

		// 文件名长度校验
		fileNameLength := config.C.Attachment.MaxFileNameLength
		if len(filename) > fileNameLength {
			c.ResponseJson(nil, utils.LogError("文件名过长，不允许超过 %d", fileNameLength))
			return
		}
		// 文件类型校验
		pointIndex := strings.LastIndex(filename, ".")
		if pointIndex == -1 {
			c.ResponseJson(nil, utils.LogError("不支持的文件类型", "文件名："+filename))
			return
		}
		attachmentType := attachment.GetType(filename[pointIndex:])
		if attachmentType == nil {
			c.ResponseJson(nil, utils.LogError("不支持的文件类型", "文件类型："+filename[pointIndex:]))
			return
		}
		// 文件大小校验
		maxSingleFileSize := config.C.Attachment.MaxFileSizeMB * int64(units.MiB)
		if fileHeader.Size > maxSingleFileSize {
			c.ResponseJson(nil, utils.LogError("文件大小超过限制", maxSingleFileSize))
			return
		}

		if file, err = fileHeader.Open(); file == nil || err != nil {
			c.ResponseJson(nil, utils.LogError("获取上传文件失败", err))
			return
		}

		attachments[index] = &attachment.Attachment{
			SaveDir:       config.C.Attachment.SaveDir,
			UploadName:    fileHeader.Filename,
			SaveName:      utils.Get32BitUUID(),
			FileExtension: attachmentType.Name,
			FileSize:      fileHeader.Size,
			CreatedBy:     userId,

			File: file,
			Type: attachmentType,
		}
	}

	err = c.s.SaveFiles(attachments)
	c.ResponseJson(nil, err)
}

// 分页获取当前用户上传的文件信息列表（用户模块怎么样也要先做好）
func (c *AttachmentController) PageData() {
	// TODO 同上还是得有用户系统，用户可以查看自己上传的文件，管理员筛选用户查看上传的文件

}

// 文件下载
// 这个接口需要结合实际调用的场景来考虑，结论是根据 id 作为参数，来指定想要下载的文件（用户查看附件列表，点击某个进行下载）
// 目前的文件管理，是用户上传的静态文件，能下载的也就是这些文件，没有动态生成的文件，所以用户下载的文件名，就按照用户上传的文件名，否则就实时生成一个随机的名字
func (c *AttachmentController) Download() {
	// TODO 需要完成模块（响应头可参见具体排名项目）
}
