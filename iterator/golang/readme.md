# Go **迭代器**模式讲解和代码示例

**迭代器**是一种行为设计模式， 让你能在不暴露复杂数据结构内部细节的情况下遍历其中所有的元素。

在迭代器的帮助下， 客户端可以用一个迭代器接口以相似的方式遍历不同集合中的元素。

## 示例

迭代器模式的主要思想是将集合背后的迭代逻辑提取至不同的、 名为迭代器的对象中。 此迭代器提供了一种泛型方法， 可用于在集合上进行迭代， 而又不受其类型影响。

####  **collection.go:** 集合

```go
package main

type Collection interface {
	createIterator() Iterator
}

```

####  **userCollection.go:** 具体集合

```go
package main

type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

```

####  **iterator.go:** 迭代器

```go
package main

type Iterator interface {
	hasNext() bool
	getNext() *User
}

```

####  **userIterator.go:** 具体迭代器

```go
package main

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

```

####  **user.go:** 客户端代码

```go
package main

type User struct {
	name string
	age  int
}

```

####  **main.go:** 客户端代码

```go
package main

import "fmt"

func main() {
	user1 := &User{
		name: "a",
		age:  30,
	}

	user2 := &User{
		name: "b",
		age:  20,
	}

	UserCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := UserCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v \n", user)
	}
}

```

####  **output.txt:** 执行结果

```html
User is &{name:a age:30} 
User is &{name:b age:20} 
```