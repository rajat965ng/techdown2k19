package main

import (
	"fmt"
	"sync"
	"time"
)

var count=0
var mu sync.Mutex

func process(n int)  {
	mu.Lock()
	for i:=0;i<5 ;i++  {
		count++
	}
	fmt.Println("The value at ",n," iteration is ", count)
	mu.Unlock()
}

func main() {
	for i:=1;i<=5 ;i++  {
		go process(i)
	}
	time.Sleep(2*time.Second)
	fmt.Println("The final value is ",count)
}