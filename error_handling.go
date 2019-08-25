package main

import (
	"fmt"
	"os"
)

func fileprocess(filename string)  {
	f,err:=os.Open(filename)

	if err!=nil{
		fmt.Println("File does not exist !!")
	}else {
		fmt.Println("The name is: ",f.Name())
	}
}

func main() {
	fileprocess("/Users/rajnigam/workspace/go-app/first.go");
}