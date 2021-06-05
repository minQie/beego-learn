package services

import (
	"beego-learn/modules/attachment"
	"beego-learn/utils"
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/prometheus/common/log"
	"os"
)

type AttachmentService interface {
	SaveFiles([]*attachment.Attachment) error
}

func NewAttachment() AttachmentService {
	return new(attachmentService)
}

type attachmentService struct {
}

func (s *attachmentService) SaveFiles(attachments []*attachment.Attachment) error {
	var (
		tempFiles = make([]*os.File, len(attachments))
		file      *os.File
		err       error
	)

	o := orm.NewOrm()
	_ = o.Begin()
	utils.HandleTransaction(o, err)

	// 事务回滚（回滚失败，不删除临时文件）
	for index, a := range attachments {
		if file, err = s.saveFile(a, o); err != nil {
			return nil
		}
		tempFiles[index] = file
	}

	// 临时文件删除（某个删除失败，不影响其他临时文件的删除）
	for _, file = range tempFiles {
		if file == nil {
			continue
		}
		if err = os.Remove((*file).Name()); err != nil {
			log.Error("文件保存失败且事务回滚成功，文件 【%s】 删除失败：%s", file.Name(), err)
		}
	}
	return nil
}

// 附件信息入库 以及 物理保存文件
func (s *attachmentService) saveFile(a *attachment.Attachment, attachmentDao orm.Ormer) (*os.File, error) {
	if _, err := attachmentDao.Insert(a); err != nil {
		return nil, err
	}

	exactSaveDir := a.GetExactSaveDir()
	nameWithExtension := a.GetFileNameWithExtension()

	return utils.SaveFile(exactSaveDir, nameWithExtension, a.File)
}
