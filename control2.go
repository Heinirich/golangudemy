package main


import (
	"fmt"
)

func main()  {
	day := "Fri"

	switch day{
	case "Fris":
		fmt.Println("Friday")
	case "Sat","Sund":
		fmt.Println("weekend")
	default :
		fmt.Println("Not found")
	}

	
}