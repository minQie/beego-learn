package services

import (
	"beego-learn/modules/attachment"
	"beego-learn/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/prometheus/common/log"
	"os"
)

type AttachmentService interface {
	SaveFiles([]*attachment.Attachment) error
}

func NewAttachmentService() AttachmentService {
	return new(attachmentService)
}

type attachmentService struct {
}

// TODO 批处理
func (s attachmentService) SaveFiles(attachments []*attachment.Attachment) error {
	var (
		tx        orm.TxOrmer
		err       error
		file      *os.File
		tempFiles = make([]*os.File, len(attachments))
	)
	tx, _ = orm.NewOrm().Begin()
	for index, a := range attachments {
		file, err = s.saveFile(a, &tx)
		if err != nil {
			break
		}
		tempFiles[index] = file
	}
	if err == nil {
		return nil
	}

	// 事务回滚（回滚失败，不删除临时文件）
	if err = tx.Rollback(); err != nil {
		log.Error("文件保存失败且事务回滚失败：%s", err)
		return err
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
	_ = tx.Commit()
	return nil
}

// 附件信息入库 以及 物理保存文件
func (s attachmentService) saveFile(a *attachment.Attachment, attachmentDao *orm.TxOrmer) (*os.File, error) {
	if _, err := (*attachmentDao).Insert(a); err != nil {
		return nil, err
	}

	exactSaveDir := a.GetExactSaveDir()
	nameWithExtension := a.GetFileNameWithExtension()

	return utils.SaveFile(exactSaveDir, nameWithExtension, a.File)
}
