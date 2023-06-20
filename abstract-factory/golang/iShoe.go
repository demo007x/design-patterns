package main

// 抽象产品

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

// Shoe 鞋子实例
type Shoe struct {
	logo string
	size int
}

func (s Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s Shoe) setSize(size int) {
	s.size = size
}

func (s Shoe) getLogo() string {
	return s.logo
}

func (s Shoe) getSize() int {
	return s.size
}
