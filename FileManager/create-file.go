package FileManager

import (
	"fmt"
	"os"
)

func (fm *FileManager) create() {
	file, err := os.Create(fm.File.Name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}

	fm.File.Data = file
}
