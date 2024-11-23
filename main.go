package main

import (
	"fmt"
	"os"

	row "github.com/OJarrisonn/gols/pkg"
)

func main() {
	path := "."

	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		r := row.NewFileRow(file)
		fmt.Println(r)
	}
}