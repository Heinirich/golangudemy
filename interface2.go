package main

import (
	"fmt"
)

func Anything(anything interface{}){
	fmt.Println(anything)
}

func main(){
	Anything(2.44)
	mymap := make(map[string]interface{})
	mymap["name"] = "Spartan";
	fmt.Println(mymap)
}