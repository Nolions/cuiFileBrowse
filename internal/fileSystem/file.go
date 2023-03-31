package fileSystem

import "os"

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
