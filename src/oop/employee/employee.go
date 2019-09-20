package employee

import (
	"fmt"
)

/**
我们把 Employee struct 的开头字母 e 改成了小写，也就是说我们把类型 Employee
struct 改成了 type employee struct。
这样做我们成功地取消了 employee 结构的导出，并阻止了来自其他包的访问。
将未导出结构的所有字段也设置为未导出，这是一个很好的实践，除非需要导出它们。
由于我们不需要employee 结构的字段在包之外的任何地方，所以我们也取消了所有字段的导出。
我们在 LeavesRemaining() 方法中相应地更改了字段名。
*/
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

/**
由于 employee 没有导出，所以不可能从其他包创建 Employee 类型的值。
因此，我们在第一行中提供了一个导出的 New 函数。
它接受所需的参数作为输入，并返回一个新创建的 employee。
*/
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}

/**
在上面的程序中，第一行指定该文件属于 employee package。第7行声明 Employee 结构。第14行的 Employee 结构中添加了一个名为 LeavesRemaining 的方法。
这将计算并显示员工剩余的休假数。现在我们有了一个结构体和一个方法，该方法对捆绑在一起的结构体进行操作，类似于类。
*/
