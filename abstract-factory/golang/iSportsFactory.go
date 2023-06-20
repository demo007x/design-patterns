package main

import "errors"

// 抽象工厂接口

type ISportsFactory interface {
	makeShoe() IShoe   // 制造鞋子
	makeShirt() IShirt // 制造 shirt
}

// GetSportFactory 根据品牌来获取一个工厂
func GetSportFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, errors.New("wrong brand type passed")
	}
}
