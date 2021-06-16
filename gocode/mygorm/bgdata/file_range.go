package bgdata

import (
	"os"
	"path/filepath"
)

type fileRange struct {
}

func NewFileRange() *fileRange {
	return &fileRange{}
}

func (s *fileRange) Walk(root string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		logger.Printf("root=%s, path=%s, err=%+v\n", root, path, err)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
