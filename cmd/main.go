package main

import "github.com/fschossler/apiscaffold/internal"

func main() {
	internal.CreateFile("test.txt")
	internal.CreateFolder("test")
}
