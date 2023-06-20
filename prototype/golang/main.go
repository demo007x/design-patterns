package main

import "fmt"

func main() {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "folder2",
	}

	fmt.Println("Print hierarchy for folder2")
	folder2.print(" ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
