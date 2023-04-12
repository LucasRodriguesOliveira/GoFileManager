package main

import (
	fm "FileIO/FileManager"
	"FileIO/log"
)

func main() {
	log.Info("Creating fileManager.")
	var fileManager = fm.NewFileManager("example.txt", ".\\")

	log.Info("Checking if the file exists.")
	exists := fileManager.Check()
	
	if exists {
		log.Info("Exist.")
	} else {
		log.Info("Do not exist.")
	}

	log.Info("Creating file or opening to write.")
	fileManager.Write("Hello")
	log.Info("Content writen inside the file.")

	log.Info("Reading the content of the file.")
	content := fileManager.Read()
	log.Info("The content is")
	log.Info(content)

	fileManager.Write(", World!")
	log.Info("Trying to update the file.")
	log.Info("Reading the content again.")
	content = fileManager.Read()
	log.Info("The content is")
	log.Info(content)
}
