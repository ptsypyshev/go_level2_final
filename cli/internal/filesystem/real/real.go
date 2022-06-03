package real

import (
	model "github.com/ptsypyshev/go_level2_final/cli/internal/model/filesystem"
)

const FileSystemKind = "real"

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

type FileSystem struct{}

func (fs *FileSystem) ListFiles(dirPath string) (*model.FileStats, error) {
	return nil, model.ErrDirNotFound{DirPath: dirPath}
}

func (fs *FileSystem) DeleteFile(dirPath, name string) error {
	return model.ErrFileNotFound{DirPath: dirPath, FileName: name}
}
