package filesystem

import (
	"github.com/ptsypyshev/go_level2_final/cli/internal/filesystem/mock"
	"github.com/ptsypyshev/go_level2_final/cli/internal/filesystem/real"
	model "github.com/ptsypyshev/go_level2_final/cli/internal/model/filesystem"
)

type Filesystem interface {
	ListFiles(dirPath string) (*model.FileStats, error)
	DeleteFile(dirPath, name string) error
}

func NewFileSystem(kind string) (Filesystem, error) {
	switch kind {
	case mock.FileSystemKind:
		return mock.NewFileSystem(), nil
	case real.FileSystemKind:
		return real.NewFileSystem(), nil
	}
	return nil, model.ErrInvalidFilesystem{FilesystemKind: kind}
}
