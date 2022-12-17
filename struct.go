package main

import (
	"fmt"
)

//Abstract data type that combines two logically related data
type Car struct{
	Name string
	Age int
	ModelNo int
}

//Method

func (c Car) Print(){
	fmt.Println(c)
}

func (c Car) Drive(){
	fmt.Println("Driving")
}

func (c Car) GetName() string{
	return c.Name
}
func main(){
	c := Car{
		Name : "Chevy",
		Age :  12,
		ModelNo :  242323,
	}
	//var c1 Car
	fmt.Println(c)
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}