package real

import (
	"errors"
	"fmt"
	model "github.com/ptsypyshev/go_level2_final/cli/internal/model/filesystem"
	"os"
	"path/filepath"
)

const FileSystemKind = "real"

type FileSystem struct {
	fileTree map[string][]*model.File
}

// NewFileSystem is a constructor for FileSystem object
func NewFileSystem() *FileSystem {
	return &FileSystem{
		fileTree: make(map[string][]*model.File, 0),
	}
}

// ListFiles walks throw the directory and returns structure with list of all children files
func (fs *FileSystem) ListFiles(dirPath string) (*model.FileStats, error) {
	if _, err := os.Stat(dirPath); err != nil {
		return nil, model.ErrDirNotFound{DirPath: dirPath}
	}

	var files []*model.File
	fmt.Printf("Try to scan folder: %s\n", dirPath)
	if err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		// Used to print warning if WalkFunc returns an error
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, &model.File{
				Name:      info.Name(),
				ParentDir: filepath.Dir(path),
				SizeBytes: int(info.Size()),
				Content:   "",
			})
		}
		return nil
	}); err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	return &model.FileStats{List: files}, nil
}

// DeleteFile tries to delete "name" file from "dirPath" and returns an error if something went wrong
func (fs *FileSystem) DeleteFile(dirPath, name string) error {
	if _, err := os.Stat(dirPath); err != nil {
		return model.ErrDirNotFound{DirPath: dirPath}
	}

	fullPath := filepath.Join(dirPath, name)
	if err := os.Remove(fullPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return model.ErrFileNotFound{}
		}
		return err
	}
	return nil
}
