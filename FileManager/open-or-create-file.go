package FileManager

func (fm *FileManager) OpenOrCreate() {
	if fm.Check() {
		fm.open()
	} else {
		fm.create()
	}
}
