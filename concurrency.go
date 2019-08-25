package main

import (
	"fmt"
	"time"
)

func displayloop()  {
	for i:=1;i<=5 ;i++  {
		fmt.Println(i)
		time.Sleep(1*time.Second)
	}
}

func main() {

	go displayloop()
	for i:=1;i<5 ;i++  {
		fmt.Println("In main !!")
		time.Sleep(2*time.Second)
	}
}