package main

import (
	"fmt"
	"sync"
)

func main(){
	//Anonymous function

	wg := &sync.WaitGroup{}

	//mutexes (critical section)

	wg.Add(2)
	go func(){
		fmt.Println("Hello")
		wg.Done()
	}()

	go func(){
		fmt.Println("World")
		wg.Done()
	}()
	
	fmt.Println("fin")
	wg.Wait()
	fmt.Println("exit")
	//select{}
}