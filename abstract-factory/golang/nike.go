package main

// Nike 具体工厂
type Nike struct{}

func (a Nike) makeShoe() IShoe {
	return NikeShirt{
		Shirt{
			logo: "nike",
			size: 10,
		},
	}
}

func (a Nike) makeShirt() IShirt {
	return NikeShirt{Shirt{
		logo: "nike",
		size: 10,
	}}
}
