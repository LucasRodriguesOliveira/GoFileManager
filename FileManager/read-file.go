package FileManager

import (
	"fmt"
	"os"
)

func (fm *FileManager) Read() string {
  if !fm.Check() {
		fmt.Printf("File %q does not exists.", fm.File.Name)
	}

	data, err := os.ReadFile(fm.File.Name)

	if err != nil {
		fmt.Println("Error while reading the file:", err)
		os.Exit(1)
	}

	return string(data)
}
