package fileSystem

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func GetFiles(path string) ([]os.DirEntry, []os.DirEntry) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []os.DirEntry{}, []os.DirEntry{}
	}

	var dirs []os.DirEntry
	var files []os.DirEntry
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e)
		} else {
			files = append(files, e)
		}
	}

	return dirs, files
}

type File struct {
	Name      string
	Path      string
	Type      string
	Extension string
}

type Dir struct {
	Name string
	Path string
}

// GetAllFiles 取得指定目錄下所以目錄與檔案
func GetAllFiles(path string) ([]Dir, []File) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, nil
	}
	var dirs []Dir
	var files []File
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, getDirMetaData(path, e))
		} else {
			files = append(files, getFileMetadata(path, e))
		}
	}

	return dirs, files
}

func fileType(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(buffer), nil
}

func getDirMetaData(path string, e os.DirEntry) Dir {
	return Dir{
		Name: e.Name(),
		Path: fmt.Sprintf("%s/%s", path, e.Name()),
	}
}

// 取得檔案類型
func getFileMetadata(rootPath string, e os.DirEntry) File {
	p := fmt.Sprintf("%s/%s", rootPath, e.Name())
	t, _ := fileType(p)
	return File{
		Name:      e.Name(),
		Path:      p,
		Type:      t,
		Extension: fileExtension(p),
	}
}

// 取得檔案副檔名
func fileExtension(filePath string) string {
	extension := filepath.Ext(filePath)

	// 去除掉句點
	if len(extension) <= 1 {
		return extension[0:]
	}

	return extension[1:]
}
