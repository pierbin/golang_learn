package main

import (
	"fmt"
)

type IFootwear interface {
	GetPrice(discount float32) float32
}

type Shoe struct {
	Size     int64
	Price    float32
	Category string
}

func (shoe *Shoe) GetPrice(discount float32) float32 {
	return shoe.Price * (1 - discount)
}

func newShoe(size int64, category string, price float32) IFootwear {
	return &Shoe{
		Size:     size,
		Category: category,
		Price:    price,
	}
}

type Slipper struct {
	Size     int64
	Price    float32
	Category string
}

func (slipper *Slipper) GetPrice(discount float32) float32 {
	return slipper.Price * (1 - discount)
}

func newSlipper(size int64, category string, price float32) IFootwear {
	return &Slipper{
		Size:     size,
		Category: category,
		Price:    price,
	}
}

type Sandal struct {
	Size     int64
	Price    float32
	Category string
}

func (sandal *Sandal) GetPrice(discount float32) float32 {
	return sandal.Price * (1 - discount)
}

func newSandal(size int64, category string, price float32) IFootwear {
	return &Sandal{
		Size:     size,
		Category: category,
		Price:    price,
	}
}

func GetFootWear(footwearType string, size int64, category string, price float32) (IFootwear, error) {
	switch footwearType {
	case "Shoe":
		return newShoe(size, category, price), nil
	case "Slipper":
		return newSlipper(size, category, price), nil
	case "Sandal":
		return newSandal(size, category, price), nil
	default:
		return nil, fmt.Errorf("invalid footwear type")
	}
}

func main() {
	footwear1, _ := GetFootWear("Slipper", 42, "Gents", 200)
	fmt.Println(footwear1.GetPrice(0.1))
	footwear2, _ := GetFootWear("Sandal", 42, "Gents", 1000)
	fmt.Println(footwear2.GetPrice(0.15))
	footwear3, err := GetFootWear("test", 42, "Gents", 100)
	fmt.Println(footwear3, err)
}
