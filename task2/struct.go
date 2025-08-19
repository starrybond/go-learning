package main

import "fmt"

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
