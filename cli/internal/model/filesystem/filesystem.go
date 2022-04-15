package filesystem

import "fmt"

type ErrInvalidFilesystem struct {
	FilesystemKind string
}

func (e ErrInvalidFilesystem) Error() string {
	return fmt.Sprintf("[%s] is invalid filesystem", e.FilesystemKind)
}
