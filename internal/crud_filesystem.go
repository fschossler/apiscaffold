package internal

import (
	"fmt"
	"log"
	"os"
)

func CreateFolder(folderName string) {
	os.Mkdir(folderName, 0775)
	fmt.Printf("The folder %s was created.\n", folderName)
}

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Printf("File %s was created.\n", fileName)
}

func AllFoldersCreation(folderNames ...string) {
	for _, folder := range folderNames {
		CreateFolder(folder)
	}
}

func AllFilesCreation(fileNames ...string) {
	for _, file := range fileNames {
		CreateFile(file)
	}
}
