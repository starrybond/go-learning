package main

import "fmt"

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息

// Person 基础结构体
type Person struct {
	Name string
	Age  int
}

// Employee 通过组合 Person 扩展而来
type Employee struct {
	Person     // 匿名组合（嵌入）
	EmployeeID string
}

// PrintInfo 输出员工完整信息
func (e Employee) PrintInfo() {
	fmt.Printf("EmployeeID: %s, Name: %s, Age: %d\n",
		e.EmployeeID, e.Name, e.Age)
}

func main() {
	emp := Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E1001",
	}
	emp.PrintInfo()
}
