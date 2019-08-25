package main

import "fmt"

func display(num int){
	fmt.Println(num)
}

func (e emp) display()  {
	fmt.Println(e.name)
}

type emp struct {
	id int
	name string
	isMale bool
}

func main() {
	fmt.Println("This is 2nd go !!")

	defer display(1)
	defer display(2)
	display(3)

	a:=20
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*b+1)

	var employee emp
	employee.id,employee.name,employee.isMale = 8,"Rajat",true;

	emp2 := emp{9,"Abc",false}

	fmt.Println(employee)
	fmt.Println(emp2)

	employee.display()
}