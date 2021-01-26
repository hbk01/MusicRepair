package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2"
)

// Revert 删除元信息
func Revert(path string) error {
	tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		return err
	}
	defer tag.Close()

	fmt.Println(path + "：")
	fmt.Println("\t作者：", tag.Artist(), " -> ")
	fmt.Println("\t标题：", tag.Title(), " -> ")
	fmt.Println("\t专辑：", tag.Album(), " -> ")

	tag.SetDefaultEncoding(id3v2.EncodingUTF16)
	tag.SetTitle("")
	tag.SetArtist("")
	tag.SetAlbum("")

	if err = tag.Save(); err != nil {
		return err
	}

	return nil
}

// Repair 修复一个音频文件
func Repair(path string) error {
	tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		return err
	}
	defer tag.Close()

	_, filename := filepath.Split(path)
	metadata := GetMetadata(filename[0 : len(filename)-4])

	fmt.Println(path + "：")
	fmt.Println("\t作者：", tag.Artist(), " -> ", strings.Join(metadata.Artists, "、"))
	fmt.Println("\t标题：", tag.Title(), " -> ", metadata.Title)
	fmt.Println("\t专辑：", tag.Album(), " -> ", metadata.Album)

	tag.SetDefaultEncoding(id3v2.EncodingUTF16)
	tag.SetTitle(metadata.Title)
	tag.SetAlbum(metadata.Album)
	tag.SetArtist(strings.Join(metadata.Artists, "、"))

	if err = tag.Save(); err != nil {
		return err
	}

	return nil
}

// GetMetadata 获取
func GetMetadata(name string) (metadata Metadata) {
	//Title   string   // 标题
	//Artists []string // 艺术家
	//Album   string   // 专辑

	// for example:
	//     aaa - bbb.mp3    -> aaa is artists, bbb is title
	//     aaa、bbb - ccc.mp3  -> aaa and bbb is artists, ccc is title
	s := strings.Split(name, "-")

	title := strings.TrimSpace(s[1])
	artists := strings.Split(s[0], "、")

	for i, v := range artists {
		artists[i] = strings.TrimSpace(v)
	}

	metadata = Metadata{
		Title:   title,
		Artists: artists,
		Album:   title,
	}
	return metadata
}
