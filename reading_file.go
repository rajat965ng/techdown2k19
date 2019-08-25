package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readfile(name string){
	data,err:=ioutil.ReadFile(name)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func writefile(name string,data string)  {
	f,err:=os.OpenFile(name,os.O_APPEND|os.O_WRONLY,0644)
	if err!=nil {
		fmt.Println(err)
		return
	}
	n,err:=f.WriteString(data)
	fmt.Println("The value of n is:",n)
	f.Close()
}

func main() {
	readfile("simple.txt")
	writefile("simple.txt","\nOnce more write here...")
	readfile("simple.txt")
}
