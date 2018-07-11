package logic

import (
	"os"
	"path/filepath"
)

type FileGroup struct {
	URL      string
	FileName string
}

func ToList() ([]FileGroup, error) {
	var groups []FileGroup
	if err := filepath.Walk("groups/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			_, filename := filepath.Split(path)
			group := FileGroup{
				URL:      path,
				FileName: filename,
			}
			groups = append(groups, group)
		}

		return nil
	}); err != nil {
		return nil, nil
	}

	return groups, nil
}
