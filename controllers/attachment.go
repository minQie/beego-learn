package controllers

import (
	"beego-learn/base/config"
	"beego-learn/modules/attachment"
	"beego-learn/services"
	"beego-learn/utils"
	"errors"
	"github.com/alecthomas/units"
	"github.com/beego/beego/v2/core/logs"
	"net/http"
	"strconv"
	"strings"
)

type AttachmentController struct {
	BaseController
	s services.AttachmentService
}

func (c *AttachmentController) Prepare() {
	// TODO 需要确认是否需要复用
	c.s = services.NewAttachmentService()
}

// 文件上传
// TODO 补上路由配置
func (c *AttachmentController) Upload() {
	// TODO 封装 log.error("XXX: %s", err) 有返回值，直接返回 err 就 OK 了
	// TODO 预期外异常的拦截，保证接口调用者获取好的数据格式
	// TODO 对 Spring 来说都具有非常好的拓展意义的 拦截向请求中添加参数
	// defer help.HttpPanic(func(msg string) {
	//     c.ResponseJson(nil, errors.New(http.StatusInternalServerError, msg))
	// })

	fileHeaders, err := c.GetFiles("files[]")
	if err == http.ErrMissingFile {
		fileHeaders, err = c.GetFiles("files")
	}
	if err != nil {
		logs.Error("上传失败，没有获取到上传的文件：", err)
		c.ResponseJson(nil, errors.New("上传失败，获取附件时发生错误"))
		return
	}

	// 请求头非空校验
	if len(fileHeaders) == 0 {
		logs.Error("上传失败，没有获取到请求头：", err)
		c.ResponseJson(nil, errors.New("上传失败，获取请求头信息时发生错误"))
		return
	}

	// 通过校验后待保存的附件数据
	var attachments = make([]*attachment.Attachment, len(fileHeaders))
	for index, fileHeader := range fileHeaders {
		filename := fileHeader.Filename
		// 文件名长度校验
		fileNameLength := config.C.Attachment.MaxFileNameLength
		if len(filename) > fileNameLength {
			logs.Error("上传失败，文件名长度超过上限：%d", fileNameLength)
			c.ResponseJson(nil, errors.New("上传失败，不支持的文件类型"))
			return
		}

		// 文件类型非空校验
		pointIndex := strings.LastIndex(filename, ".")
		if pointIndex == -1 {
			logs.Error("上传失败，非法文件，类型未知：%s", filename)
			c.ResponseJson(nil, errors.New("上传失败，不支持的文件类型"))
			return
		}

		// 文件类型校验
		fileExtension := filename[index:]
		attachmentType := attachment.GetAttachmentType(fileExtension)
		if attachmentType == nil {
			logs.Error("上传文件失败，非法文件类型：%s", fileExtension)
			c.ResponseJson(nil, errors.New("上传失败，不支持的文件类型"))
			return
		}

		// 文件大小限制校验
		maxSingleFileSize := config.C.Attachment.MaxFileSizeMB * int64(units.MiB)
		if fileHeader.Size > maxSingleFileSize {
			logs.Error("上传失败，文件大小超过限制：%d", maxSingleFileSize)
			c.ResponseJson(nil, errors.New("上传失败，文件大小超过限制："+strconv.FormatInt(maxSingleFileSize, 10)))
			return
		}

		// 文件校验
		file, err := fileHeader.Open()
		if file == nil || err != nil {
			logs.Error("上传失败，获取文件失败：%s", err)
			c.ResponseJson(nil, errors.New("上传失败，获取文件失败"))
			return
		}

		attachments[index] = &attachment.Attachment{
			SaveDir:       config.C.Attachment.SaveDir,
			UploadName:    fileHeader.Filename,
			SaveName:      utils.Get32BitUUID(),
			FileExtension: attachmentType.Name,
			FileSize:      fileHeader.Size,
			UploadedBy:    0, // TODO 上传人id

			File: file,
			Type: attachmentType,
		}
	}

	// TODO 类比 ThreadLocal 的 Context 机制实现，将用户信息绑定到当前 Goroutine 里边
	err = c.s.SaveFiles(attachments)
	if err != nil {
		c.ResponseJson(nil, errors.New(err.Error()))
	} else {
		c.ResponseJson("上传成功", nil)
	}
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
