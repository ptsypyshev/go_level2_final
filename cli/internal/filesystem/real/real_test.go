package real

import (
	"fmt"
	"github.com/ptsypyshev/go_level2_final/cli/internal/model/filesystem"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func ExampleFileSystem_ListFiles() {
	fs := NewFileSystem()
	fileStats, _ := fs.ListFiles("test")
	fmt.Println(fileStats)

	//Output:
	//Try to scan folder: test
	//1 File: exist.txt
	//   Size: 0 bytes
	//   Parent directory: test
	//   Content:

}

func ExampleFileSystem_DeleteFile() {
	fs := NewFileSystem()
	err := fs.DeleteFile("test", "example.txt")
	if err != nil {
		fmt.Printf("Failed to delete file, because of: %s", err)
	} else {
		fmt.Println("File is deleted")
	}

	//Output: Failed to delete file, because of: file [] not found in []
}

func TestFileSystem_DeleteFile(t *testing.T) {
	fs := NewFileSystem()
	if cwd, err := os.Getwd(); err != nil {
		t.Fatalf("cannot get current working directory, because of: %s\n", err)
	} else {
		testDir := filepath.Join(cwd, "test")
		testNotExistDir := filepath.Join(cwd, "testNotExist")
		testNoPermDir := filepath.Join(cwd, "testNoPerm")
		os.Create(filepath.Join(testDir, "exist.txt"))
		assert.ErrorIs(t, fs.DeleteFile(testNotExistDir, "example.txt"), filesystem.ErrDirNotFound{DirPath: testNotExistDir})
		assert.ErrorIs(t, fs.DeleteFile(testNoPermDir, "noPerm.txt"), os.ErrPermission)
		assert.ErrorIs(t, fs.DeleteFile(testDir, "noFile.txt"), filesystem.ErrFileNotFound{})
		assert.NoError(t, fs.DeleteFile(testDir, "exist.txt"))
	}
}

func TestFileSystem_ListFiles(t *testing.T) {
	fs := NewFileSystem()
	if cwd, err := os.Getwd(); err != nil {
		t.Fatalf("cannot get current working directory, because of: %s\n", err)
	} else {
		testDir := filepath.Join(cwd, "test")
		testNotExistDir := filepath.Join(cwd, "testNotExist")
		testNoPermReadDir := filepath.Join(cwd, "testNoPermRead")
		os.Create(filepath.Join(testDir, "exist.txt"))
		if _, errCase := fs.ListFiles(testNotExistDir); errCase != nil {
			assert.ErrorIs(t, errCase, filesystem.ErrDirNotFound{DirPath: testNotExistDir})
		}
		if _, errCase := fs.ListFiles(testNoPermReadDir); errCase != nil {
			assert.ErrorIs(t, errCase, os.ErrPermission)
		}

		_, errCase := fs.ListFiles(testDir)
		assert.NoError(t, errCase)
	}
}
