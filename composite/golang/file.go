package main

import "fmt"

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}
