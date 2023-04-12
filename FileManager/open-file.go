package FileManager

import (
	"fmt"
	"os"
)

func (fm *FileManager) open() {
  file, err := os.OpenFile(fm.File.Name, os.O_RDWR, os.ModeAppend)

	if err != nil {
		fmt.Printf("Occurred and error while trying to open the file %q", fm.File.Name)
		os.Exit(1)
	}

	fm.File.Data = file
}
