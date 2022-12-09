package nospaceleftondevice

import (
	"aoc-2022-go/utils"
	"strconv"
	"strings"
)

const INPUT_PATH = "07_nospaceleftondevice/input.txt"

func GetSumOfDirSizeAtMost() int {
	_, dirs := parseTreeAndDirs()

	size := 0
	treshold := 100000
	for _, dir := range dirs {
		dirSize := dir.getSize()
		if dirSize <= treshold {
			size += dirSize
		}
	}

	return size
}

func GetSmallestDirToDelete() int {
	tree, dirs := parseTreeAndDirs()

	fileSystemTotalSpace := 70000000
	requiredSpace := 30000000
	spaceToBeFreed := requiredSpace - (fileSystemTotalSpace - tree.getSize())

	smallestDirSize := fileSystemTotalSpace
	for _, dir := range dirs {
		dirSize := dir.getSize()
		if dirSize >= spaceToBeFreed &&
			dirSize < smallestDirSize {
			smallestDirSize = dirSize
		}
	}

	return smallestDirSize
}

func parseTreeAndDirs() (*Dir, []*Dir) {
	dirs := []*Dir{}
	tree := newDir(nil)
	tree.subDirs["/"] = newDir(tree)
	currentDir := tree

	for line := range utils.GetInputLines(INPUT_PATH) {
		if strings.HasPrefix(line, "$ ls") {
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			currentDir = getCurrentDir(line, currentDir)
		} else if strings.HasPrefix(line, "dir") {
			dirname := strings.Split(line, "dir ")[1]
			dir := newDir(currentDir)
			dirs = append(dirs, dir)
			currentDir.subDirs[dirname] = dir
		} else {
			split := strings.Split(line, " ")
			size, _ := strconv.Atoi(split[0])
			currentDir.files = append(currentDir.files, &File{
				name: split[1],
				size: size,
			})
		}
	}

	return tree, dirs
}

func getCurrentDir(line string, currentDir *Dir) *Dir {
	targetDir := strings.Split(line, "$ cd ")[1]
	if targetDir == ".." {
		return currentDir.parent
	} else {
		return currentDir.subDirs[targetDir]
	}
}

func getFile(line string) *File {
	split := strings.Split(line, " ")
	size, _ := strconv.Atoi(split[0])
	return &File{
		name: split[1],
		size: size,
	}
}
