package filesystem

import (
	"bytes"
	"fmt"
	"path/filepath"
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
		buff.WriteString("  Parent directory: ")
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

func (fs *FileStats) FindDuplicates() *FileStats {
	duplicatesFilterByName := make(map[string][]*File)
	for _, file := range fs.List {
		listByName, ok := duplicatesFilterByName[file.Name]
		if !ok {
			duplicatesFilterByName[file.Name] = []*File{file}
		} else {
			listByName = append(listByName, file)
			duplicatesFilterByName[file.Name] = listByName
		}
	}

	list := make([]*File, 0)
	for _, listByName := range duplicatesFilterByName {
		if len(listByName) > 1 {
			duplicatesFilterBySize := make(map[int][]*File)
			for _, file := range listByName {
				listByName, ok := duplicatesFilterBySize[file.SizeBytes]
				if !ok {
					duplicatesFilterBySize[file.SizeBytes] = []*File{file}
				} else {
					listByName = append(listByName, file)
					duplicatesFilterBySize[file.SizeBytes] = listByName
				}
			}
			for _, listByNameAndSize := range duplicatesFilterBySize {
				if len(listByNameAndSize) > 1 {
					list = append(list, listByNameAndSize...)
				}
			}
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
	return fmt.Sprint(filepath.Join(f.ParentDir, f.Name))
}
