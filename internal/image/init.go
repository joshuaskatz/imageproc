package internal_image

func NewFile() *File {
	return &File{}
}

func (file *File) SetIsColor(isColor bool) {
	file.IsColor = isColor
}

func (file *File) SetFilePath(filePath string) {
	file.Path = filePath
}
