package main


import (
	"fmt"
	"time"
)

//A go routine is the way we call a function and it runs in background and thus
//dont stop what is below it
func heavy(){
	for{
		time.Sleep(time.Second*1)
		fmt.Println("Heavy")

	}
}
func superHeavy(){
	for {
		time.Sleep(time.Second*2)
		fmt.Println("Super Heavy")
	}
}

func main()  {
	go heavy()
	go superHeavy()
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println(123)
	select {}
}