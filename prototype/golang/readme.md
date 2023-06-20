# Go **原型**模式讲解和代码示例

**原型**是一种创建型设计模式， 使你能够复制对象， 甚至是复杂对象， 而又无需使代码依赖它们所属的类。

所有的原型类都必须有一个通用的接口， 使得即使在对象所属的具体类未知的情况下也能复制对象。 原型对象可以生成自身的完整副本， 因为相同类的对象可以相互访问对方的私有成员变量。

[ 进一步了解原型模式 ](https://refactoringguru.cn/design-patterns/prototype)



## 概念示例

让我们尝试通过基于操作系统文件系统的示例来理解原型模式。 操作系统的文件系统是递归的： 文件夹中包含文件和文件夹， 其中又包含文件和文件夹， 以此类推。

每个文件和文件夹都可用一个 `inode`接口来表示。  `inode`接口中同样也有 `clone`克隆功能。

`file`文件和 `folder`文件夹结构体都实现了 `print`打印和 `clone`方法， 因为它们都是 `inode`类型。 同时， 注意 `file`和 `folder`中的 `clone`方法。 这两者的 `clone`方法都会返回相应文件或文件夹的副本。 同时在克隆过程中， 我们会在其名称后面添加 “_clone” 字样。

####  **inode.go:** 原型接口

```go
package main

type Inode interface {
	printer(string)
	clone() Inode
}


```

####  **file.go:** 具体原型

```go
package main

import "fmt"

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

```

####  **folder.go:** 具体原型

```go
package main

import "fmt"

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, v := range f.children {
		v.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, v := range f.children {
		copy := v.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

```

####  **main.go:** 客户端代码

```go
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

```

####  **output.txt:** 执行结果

```sh
Print hierarchy for folder2
 folder2
  folder1
    file1
  file2
  file3

Printing hierarchy for clone Folder
  folder2_clone
    folder1_clone
        file1_clone
    file2_clone
    file3_clone
```