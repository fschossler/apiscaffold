package internal

import (
	"fmt"
	"log"
	"os"
)

func CreateFolder(folderName string) {
	os.Mkdir(folderName, 0775)
	fmt.Printf("The folder %s was created.", folderName)
}

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Printf("File %s was created.", fileName)
}
