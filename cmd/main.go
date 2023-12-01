package main

import (
	"fmt"

	"github.com/fschossler/apiscaffold/internal"
)

func main() {
	allFolders := []string{
		"builds",
		"cmd",
		"configs",
		"db", "db/migrations",
		"internal", "internal/app", "internal/app/models", "internal/app/controllers", "internal/app/routes", "internal/app/business",
		"pkg",
		"scripts",
	}

	allFiles := []string{
		"builds/Containerfile", "builds/container-compose.yml",
		"cmd/main.go",
		"configs/config.go",
		"Makefile",
		"README.md",
	}

	internal.AllFoldersCreation(allFolders...)
	internal.AllFilesCreation(allFiles...)

	fmt.Println("########## The API Scaffold structure was created with success ##########")
}
