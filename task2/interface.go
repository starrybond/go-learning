package main

import (
	"fmt"
	"math"
)

// 1. 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Rectangle 结构体
type Rectangle struct {
	Width, Height float64
}

// Rectangle 实现 Shape 接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 3. Circle 结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// 创建实例
	r := Rectangle{Width: 4, Height: 3}
	c := Circle{Radius: 5}

	// 调用方法
	shapes := []Shape{r, c}
	for _, s := range shapes {
		fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
	}
}
