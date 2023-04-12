package FileManager

import (
	"fmt"
	"io"
	"os"

	"FileIO/log"
)

func (fm *FileManager) Write(data string) {
  fm.OpenOrCreate()
	defer fm.File.Data.Close()

	_, err := fm.File.Data.Seek(0, io.SeekEnd)

	if err != nil {
		log.Error("Seeking end of the file")
		os.Exit(1)
	}

	_, err = fm.File.Data.Write([]byte(data))

	if err != nil {
		log.Error(fmt.Sprintf("Error writing file: %s", err))
		os.Exit(1)
	}

	log.Info(fmt.Sprintf("The content %q was written in %q file successfully.", data, fm.File.Name))
}
