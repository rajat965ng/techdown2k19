package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock sync.Mutex
)


func getOneToHundred(buff chan []int) {

	time.Sleep(time.Second)
	arr := make([]int, 100)
	for i := 1; i <= 100; i++ {
		arr[i-1] = i
	}
	buff <- arr

}

func getHundredToThousand(buff chan []int) {
	time.Sleep(time.Second)
	arr := make([]int, 1000)
	for i := 101; i <= 1000; i++ {
		arr[i-1] = i
	}
	buff <- arr

}

func getThousandTo10Thousand(buff chan []int) {
	time.Sleep(time.Second)
	arr := make([]int, 10000)
	for i := 1001; i <= 10000; i++ {
		arr[i-1] = i
	}
	buff <- arr
	close(buff)
}



func main() {
	buff := make(chan []int)
	go getOneToHundred(buff)
	go getHundredToThousand(buff)
	go getThousandTo10Thousand(buff)

	for v := range buff {
		fmt.Println(v)
	}

}
