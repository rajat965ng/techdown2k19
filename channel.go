package main

import (
	"fmt"
	"time"
)

func displayChannel(ch chan int){
	fmt.Println("Inside Display")
	ch <- 1234
}

func consumer(ch chan int)  {
	for {
		x, flag:=<-ch
		if flag {
			fmt.Println(x)
		}else {
			fmt.Println("Empty channel")
			break
		}
	}
}

func producer(ch chan int)  {
	for i:=1;i<=5;i++{
		fmt.Println("Produce:",i)
		ch <- i
	}
	close(ch)
}

func prog()  {
	fmt.Println("inside prog")
	ch:=make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(5*time.Second)

}

func main() {
	ch := make(chan int)
	go displayChannel(ch)
	x:= <- ch
	fmt.Println("The main method is called !!")
	fmt.Println("The value of x is:",x)
	prog()
}