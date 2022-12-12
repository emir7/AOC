package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type folder struct {
	name         string
	files        []file
	innerFolders []*folder
	parent       *folder
	size         int
}

func sumDirectories(root *folder) int {
	sum := 0

	if len(root.innerFolders) == 0 {
		if root.size < 100000 {
			return root.size
		}

		return 0
	}

	for _, childFolder := range root.innerFolders {
		sum += sumDirectories(childFolder)
	}

	if root.size < 100000 {
		sum += root.size
	}

	return sum

}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	root := &folder{name: "/", files: []file{}, innerFolders: []*folder{}, parent: nil, size: 0}
	var currentFolder *folder

	for fileScanner.Scan() {
		var line string = strings.TrimSpace(fileScanner.Text())

		if len(line) == 0 {
			break
		}

		if err != nil {
			return
		}

		splittedCmd := strings.Fields(line)

		if splittedCmd[0] == "$" {
			if splittedCmd[1] == "cd" {
				if splittedCmd[2] == "/" {
					currentFolder = root
				} else if splittedCmd[2] == ".." {
					currentFolder.parent.size += currentFolder.size
					currentFolder = currentFolder.parent
				} else {
					for _, folder := range currentFolder.innerFolders {
						if folder.name == splittedCmd[2] {
							currentFolder = folder
						}
					}
				}
			}
			if splittedCmd[1] == "ls" {
				continue
			}
		} else {
			if splittedCmd[0] == "dir" {
				folderName := splittedCmd[1]
				newFolder := folder{name: folderName, files: []file{}, innerFolders: []*folder{}, parent: currentFolder, size: 0}
				currentFolder.innerFolders = append(currentFolder.innerFolders, &newFolder)
			} else {
				fileSize, _ := strconv.Atoi(splittedCmd[0])
				fileName := splittedCmd[1]
				currentFolder.files = append(currentFolder.files, file{name: fileName, size: fileSize})
				currentFolder.size += fileSize
			}
		}

	}

	rootSize := 0

	for _, innerFolder := range root.innerFolders {
		rootSize += innerFolder.size
	}

	for _, file := range root.files {
		rootSize += file.size
	}
	root.size = rootSize

	s := sumDirectories(root)

	fmt.Println(s)
	readFile.Close()

}
