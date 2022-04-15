package mock

import (
	"strings"

	model "github.com/ifireice/go_level2_final/cli/internal/model/filesystem"
)

const FileSystemKind = "mock"

func NewFileSystem() *FileSystem {
	fileTree := map[string][]*model.File{
		"/mock": []*model.File{
			{
				Name:      "test.mock",
				ParentDir: "/mock",
				SizeBytes: 10,
				Content:   "Geekbrains",
			},
			{
				Name:      "test.mock",
				ParentDir: "/mock/subdir",
				SizeBytes: 10,
				Content:   "Geekbrains",
			},
			{
				Name:      "test2.mock",
				ParentDir: "/mock",
				SizeBytes: 17,
				Content:   "Go Course level 2",
			},
		},
	}
	return &FileSystem{fileTree: fileTree}
}

type FileSystem struct {
	fileTree map[string][]*model.File
}

func (fs *FileSystem) ListFiles(dirPath string) (*model.FileStats, error) {
	list := make([]*model.File, 0)
	dir, ok := fs.fileTree[dirPath]
	if !ok {
		return nil, model.ErrDirNotFound{DirPath: dirPath}
	}
	for _, file := range dir {
		list = append(list, file)
	}
	return &model.FileStats{List: list}, nil
}

func (fs *FileSystem) DeleteFile(dirPath, name string) error {
	if strings.HasPrefix(dirPath, "/mock") {
		dirPath = "/mock"
	}
	dir, ok := fs.fileTree[dirPath]
	if !ok {
		return model.ErrDirNotFound{DirPath: dirPath}
	}
	for fileInd := range dir {
		if dir[fileInd].Name == name {
			copy(dir[fileInd:], dir[fileInd+1:])
			dir[len(dir)-1] = nil
			dir = dir[:len(dir)-1]
		}
		fs.fileTree[dirPath] = dir
		return nil
	}
	return model.ErrFileNotFound{DirPath: dirPath, FileName: name}
}
