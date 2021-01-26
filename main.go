package main

import (
	"flag"
	"fmt"
)

var (
	root        = flag.String("dir", "./", "Specifies the directory where the music files are located")
	isRecursive = flag.Bool("recursive", false, "If set, Musicrepair will run recursively in the given directory")
	isRevert    = flag.Bool("revert", false, "If set, Musicrepair will revert the files")
	list        = flag.Bool("list", false, "Show the output by list")
)

func main() {
	flag.Parse()

	fileList := WalkDir(*root) // List of all files to work on

	for _, file := range fileList {
		if *isRevert {
			err := Revert(file)
			if err != nil {
				fmt.Println("error: ", err.Error())
			}
		} else {
			err := Repair(file)
			if err != nil {
				fmt.Println("error: ", err.Error())
			}
		}
	}
}
