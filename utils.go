package main

import (
	"os"
	"path/filepath"
)

//Metadata : Structure for one track's metadata
type Metadata struct {
	Title   string   // 标题
	Artists []string // 艺术家
	Album   string   // 专辑
}

// WalkDir 寻找目录下的所有 .mp3 文件
func WalkDir(root string) (fileList []string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !*isRecursive && filepath.Dir(path) != filepath.Dir(root) {
			return filepath.SkipDir
		}
		if filepath.Ext(path) == ".mp3" {
			fileList = append(fileList, path)
		}
		return nil
	})
	return
}
