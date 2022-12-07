package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name   string
	files  []File
	dirs   []*Directory
	parent *Directory
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	root := Directory{"/", []File{}, []*Directory{}, nil}
	pwd := &root

	ls_buffer := []string{}

	scanner := bufio.NewScanner(file)
	defer file.Close()
	for scanner.Scan() {
		line := scanner.Text()
		// Process input
		if line[0] == '$' {
			// command
			paths := strings.Split(line, " ")
			if paths[1] == "cd" {
				ls(pwd, ls_buffer)
				ls_buffer = []string{}
				pwd = cd(pwd, paths[2])
			} else if paths[1] == "ls" {

			}
		} else {
			ls_buffer = append(ls_buffer, line)
		}
	}
	ls(pwd, ls_buffer)
	pwd = cd(pwd, "/")
	fmt.Println(part1(pwd))
	p2goal := size(pwd) - 40000000
	fmt.Println(part2(pwd, p2goal))
}

func cd(pwd *Directory, path string) *Directory {
	if path == ".." {
		// Go up one directory
		return pwd.parent
	} else if path == "." {
		// Do nothing
		return pwd
	} else if path == "/" {
		// Go to root
		for pwd.parent != nil {
			pwd = pwd.parent
		}
		return pwd
	} else {
		// Go down one directory
		for _, dir := range pwd.dirs {
			if dir.name == path {
				return dir
			}
		}
		// Create new directory
		new_dir := Directory{path, []File{}, []*Directory{}, pwd}
		pwd.dirs = append(pwd.dirs, &new_dir)
		return &new_dir
	}
}

func ls(pwd *Directory, files []string) *Directory {
	// Add files to pwd
	for _, line := range files {
		if line[:3] != "dir" {
			pwd.files = append(pwd.files, read_ls(line))
		}
	}
	return pwd
}

func read_ls(line string) File {
	paths := strings.Split(line, " ")
	if len(paths) < 2 {
		// error
	}
	size, _ := strconv.Atoi(paths[0])
	name := paths[1]
	return File{name, size}
}

func part1(root *Directory) int {
	// Get the sum of all file sizes in folders smaller than 100000
	sum := 0
	if size(root) <= 100000 {
		sum += size(root)
	}
	for _, dir := range root.dirs {
		//if size(dir) <= 100000 {
		//	sum += size(dir)
		//}
		sum += part1(dir)
	}
	return sum
}

func size(dir *Directory) int {
	// Get the size of a directory
	sum := 0
	for _, file := range dir.files {
		sum += file.size
	}
	for _, subdir := range dir.dirs {
		sum += size(subdir)
	}
	return sum
}

func part2(root *Directory, s int) int {
	// Find the smallest directory that has a size of at least s
	record := 1000000000
	for _, dir := range root.dirs {
		if size(dir) >= s {
			if size(dir) < record {
				record = size(dir)
			}
		}
		r := part2(dir, s)
		if r < record {
			record = r
		}
	}
	return record
}
