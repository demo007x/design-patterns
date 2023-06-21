package main

// 形状结构体
type Shape interface {
	getType() string
	accept(Visitor)
}
