package entity

type Student struct {
	Name string
	Age  int
}

// type 声明别名   var声明变量     type可以用于类型转换，如setI()
//// 当前变量是i ， 类型也是i
type i int

func SetI(value int) i {
	return i(value)
}
