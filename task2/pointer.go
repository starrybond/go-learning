package main

import (
	"fmt"
)

/*
1.编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
*/

// addTen 接收一个 *int 类型的指针参数，将其指向的值增加 10
func addTen(p *int) {
	*p += 10
}

/*
2.实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
*/

func doubleSlice(nums *[]int) {
	if nums == nil {
		return
	}
	//for i := 0; i < len(*nums); i++ {
	//	(*nums)[i] *= 2
	//}
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func main() {
	// 1.指针
	x := 42
	fmt.Println("调用前 x 的值:", x)

	addTen(&x) // 传入 x 的地址

	fmt.Println("调用后 x 的值:", x)

	s := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", s)

	doubleSlice(&s)
	fmt.Println("乘以2后的切片:", s)

}
