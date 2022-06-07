package main

import (
	"flag"
	"fmt"
	"github.com/ptsypyshev/go_level2_final/cli/internal/filesystem"
	"os"
	"strconv"
	"strings"
	"sync"
)

// что бы работало с обычной fs нужно запустить c -fs real
var fsKind = flag.String("fs", "mock", "Specify file system, available 'mock' and 'real'")

var longHelpFlag = flag.Bool("help", false, "Print help information, just specify this flag")
var shortHelpFlag = flag.Bool("h", false, "Short alias for --help")
var dirPathFlag = flag.String("dir", "", "Specify full directory path to scan")
var longForceDeleteAllFlag = flag.Bool("force", false, "Force delete all duplicates (use carefully!)")
var shortForceDeleteAllFlag = flag.Bool("f", false, "Short alias for --force")

func main() {
	flag.Parse()

	if *longHelpFlag || *shortHelpFlag {
		fmt.Println("Usage: ./main [arguments]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	fs, err := filesystem.NewFileSystem(*fsKind)
	if err != nil {
		fmt.Printf("cannot continue because of: %s\n", err)
		fmt.Println("Usage: ./main [arguments]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var directory string
	if *dirPathFlag == "" {
		fmt.Printf("Enter directory path for %q filesystem:\n", *fsKind)
		fmt.Scanf("%s", &directory)
	} else {
		directory = *dirPathFlag
	}

	// ListFiles - возвращает структуру FileStats - список файлов
	// но у FileStats есть метод FindDuplicates, который тоже вернет список файлов-дубликатов
	fileStats, err := fs.ListFiles(directory)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(fileStats)
	}

	fmt.Printf("\nFound files:\n\n%s", fileStats)

	duplicatesStats := fileStats.FindDuplicates()
	if len(duplicatesStats.List) > 0 {
		fmt.Printf("\nFound duplicates:\n\n%s\n", duplicatesStats)
		var filesToDelete string
		if *shortForceDeleteAllFlag || *longForceDeleteAllFlag {
			indexes := make([]string, 0, len(duplicatesStats.List))
			for i, _ := range duplicatesStats.List {
				indexes = append(indexes, strconv.Itoa(i+1))
			}
			filesToDelete = strings.Join(indexes, ",")
		} else {
			fmt.Printf("Which files you would like to delete (enter comma-separated list):\n")
			fmt.Scanf("%s", &filesToDelete)
		}
		if filesToDelete != "" {
			fmt.Printf("\n")
			var deletedFilesCounter int
			files := strings.Split(filesToDelete, ",")
			var wg sync.WaitGroup
			for _, fileIndStr := range files {
				wg.Add(1)
				go func(fileIndStr string) {
					defer wg.Done()
					fileInd, _ := strconv.Atoi(fileIndStr)
					dirPath := duplicatesStats.List[fileInd-1].ParentDir
					fileName := duplicatesStats.List[fileInd-1].Name
					fmt.Printf("Deleting file [%s] from directory [%s]...\n", fileName, dirPath)
					err := fs.DeleteFile(dirPath, fileName)
					if err != nil {
						fmt.Printf("Failed to delete file [%s] from directory [%s], error [%s], skipping...\n", fileName, dirPath, err)
					} else {
						fmt.Printf("Deleted file [%s] from directory [%s]\n", fileName, dirPath)
						deletedFilesCounter++
					}
				}(fileIndStr)
			}
			wg.Wait()
			// отображаем что удалили и что осталось
			if deletedFilesCounter > 0 {
				fmt.Printf("Successfuly deleted %d duplicates\n\n", deletedFilesCounter)
				fileStats, _ = fs.ListFiles(directory)
				fmt.Printf("Current directory file list:\n\n%s", fileStats)
			}
		}
	}
}
