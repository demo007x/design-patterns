package main

// 具体工厂
type Adidas struct{}

func (a Adidas) makeShoe() IShoe {
	return &AdidasShirt{Shirt{
		logo: "adidas",
		size: 14,
	}}
}

func (a Adidas) makeShirt() IShirt {
	return AdidasShirt{Shirt{
		logo: "shirt",
		size: 29,
	}}
}
