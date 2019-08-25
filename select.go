package main

import "fmt"

func data1(ch chan string)  {
	ch<-"This is data1"
}

func data2(ch chan string)  {
	ch<-"This is data2"
}

func main() {

	chan1:=make(chan string)
	chan2:=make(chan string)

	go data1(chan1)
	go data2(chan2)

	select {
	case x:=<-chan1:
		fmt.Println(x)
	case y:=<-chan2:
		fmt.Println(y)
	default:
		fmt.Println("Default is executed")
	}

}