package main

import "fmt"

import "calculation"

func add(num1 int, num2 int) int {
	return num1*num2
}



func main()  {
	fmt.Println("Hello World !! this is my first go !!")

	var a int
	a = 3
	fmt.Println("a:",a)

	var b int = 10*2
	fmt.Println("b:",b)

	var c = 100
	fmt.Println("c:",c)

	var x,y,z = 3,"Hello World !!",true
	fmt.Println("x: y: z:",x,y,z)

	p:=30
	fmt.Println(p)

	for i:=0; i<=5; i++ {
		fmt.Println(i)
	}

	if x>3 {
		fmt.Println("X is large")
	}else {
		fmt.Println("X is small")
	}

	switch a {
	case b:
		fmt.Println("is equal to b")
	case 3:
		fmt.Println("is equal to 3")
	default:
		fmt.Println("N/A")
	}

	arr:=[5] int {3,5,7,8,0}
	fmt.Println(arr)
	fmt.Println(len(arr))

	slicearray := arr[1:3]
	fmt.Println(slicearray)
	slicebarr := arr[3:5]
	fmt.Println(slicebarr)

	slicearray= append(slicebarr, slicearray...)
	fmt.Println(slicearray)

	fmt.Println(add(5,5))

	sum :=calculation.Do_add(3,5)
	fmt.Println(sum)
}