package main

import (
	"fmt"
)

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

/*
第15行中的 EmployeeOperations 接口是通过嵌入 SalaryCalculator 和LeaveCalculator 接口创建的。
如果它为 SalaryCalculator 和 LeaveCalculator 接口中提供的方法提供方法定义，
则称任何类型都实现 EmployeeOperations 接口。
Employee 结构实现了 EmployeeOperations 接口，
因为它分别为第29行和第33行中的 DisplaySalary 和 CalculateLeavesLeft 方法提供了定义。
在第46行，类型为 Employee 的 e 被分配给 EmployeeOperations 类型的 empOp 。
在接下来的两行中，在 empOp 上调用 DisplaySalary() 和 CalculateLeavesLeft()方法。
*/
func main() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}
