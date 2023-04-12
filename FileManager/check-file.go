package FileManager

import (
	"fmt"
	"os"
)

func (fm *FileManager) Check() bool {
	info, err := os.Stat(fm.File.Name);

	if err == nil && info.IsDir() {
		fmt.Printf("%q is a directory.\n", info.Name())
		os.Exit(1)
	}

	return !os.IsNotExist(err)
}
