//func main() {
//	e := employee.Employee {
//		FirstName: "Sam",
//		LastName: "Adolf",
//		TotalLeaves: 30,
//		LeavesTaken: 20,
//	}
//	e.LeavesRemaining()
//}

//import "oop/employee"
//
//func main() {
//	var e employee.Employee
//	e.LeavesRemaining()
//}

package main

import "oop/employee"

/**
现在你可以理解虽然 Go 不支持类，但是可以有效地使用 struct 来代替类，
并且可以使用 New(parameters) 的方法来代替构造函数。
*/
func main() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
