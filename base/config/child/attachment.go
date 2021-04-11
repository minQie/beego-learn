package child

type AttachmentConfig struct {
	MaxFileSizeMB     int64    `yaml:"max_file_size_mb"`
	MaxFileNameLength int      `yaml:"max_file_name_length"`
	DocSupportType    []string `yaml:"doc_support_type,flow"`
	PicSupportType    []string `yaml:"pic_support_type,flow"`
	SaveDir           string   `yaml:"save_dir"`
	DocDirName        string   `yaml:"doc_dir_name"`
	PicDirName        string   `yaml:"pic_dir_name"`
}
