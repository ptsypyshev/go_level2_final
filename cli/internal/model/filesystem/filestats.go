package filesystem

import (
	"bytes"
	"fmt"
	"strconv"
)

type FileStats struct {
	List []*File
}

func (fs *FileStats) String() string {
	lastFileInd := len(fs.List) - 1

	var buff bytes.Buffer
	for fileInd := range fs.List {
		buff.WriteString(strconv.Itoa(fileInd + 1))
		buff.WriteString(" ")
		buff.WriteString("File: ")
		buff.WriteString(fs.List[fileInd].Name)
		buff.WriteString("\n")
		buff.WriteString("  Size: ")
		buff.WriteString(strconv.Itoa(fs.List[fileInd].SizeBytes))
		buff.WriteString(" bytes\n")
		buff.WriteString("  Directory: ")
		buff.WriteString(fs.List[fileInd].ParentDir)
		buff.WriteString("\n")
		buff.WriteString("  Content: ")
		buff.WriteString(fs.List[fileInd].Content)
		buff.WriteString("\n")
		if fileInd != lastFileInd {
			buff.WriteString("\n")
		}
	}
	return buff.String()
}

// смотрит только на имена файлов надо допилить
func (fs *FileStats) FindDuplicates() *FileStats {
	duplicatesFilter := make(map[string][]*File)
	for _, file := range fs.List {
		listByName, ok := duplicatesFilter[file.Name]
		if !ok {
			duplicatesFilter[file.Name] = []*File{file}
		} else {
			listByName = append(listByName, file)
			duplicatesFilter[file.Name] = listByName
		}
	}

	list := make([]*File, 0)
	for _, listByName := range duplicatesFilter {
		if len(listByName) > 1 {
			list = append(list, listByName...)
		}
	}

	return &FileStats{List: list}
}

type File struct {
	Name      string
	ParentDir string
	SizeBytes int
	Content   string
}

func (f *File) FullPath() string {
	return fmt.Sprintf("%s/%s", f.ParentDir, f.Name)
}
