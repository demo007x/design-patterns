package main

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("searching recursvely for keyword %s in folder %s \n", keyword, f.name)
	for _, component := range f.components {
		component.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
