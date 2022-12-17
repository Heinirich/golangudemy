package main

import (
	"fmt"
	//"time"
)

func main()  {
	// var c chan int

	c := make(chan int)

	//send

	go func ()  {
		c <- 1
	}()

	value := <-c
	fmt.Println(value)

	go func ()  {
		c <- 2
	}()

	//time.Sleep(time.Second *2)

	//sniff
	val := <-c
	fmt.Println(val)
	// fmt.Println(c)

}