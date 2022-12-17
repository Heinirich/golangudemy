package main

import (
	"fmt"
)

func main()  {
	fmt.Println("This is Control Structure")
	f := true
	flag := &f
	if flag == nil{
		fmt.Println("Flag is empty")
	} else if *flag{
		fmt.Println("true")
	} else{
		fmt.Println("false")
	}


	for i:=0;i<=10;i++{
		fmt.Println(i)
	}

	x := 0
	for ;x<5;{
		fmt.Println(x)
		x++
	}
}
