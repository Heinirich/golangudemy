package main

import (
	"fmt"
)
type Car interface{
	Drive()
	Stop()
}

func NewModel(arg string) Car{
	return &Lambo{arg}
}

type Lambo struct{
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Stop(){
	fmt.Println("Lambo Stop",l.LamboModel)

}

func (l *Lambo) Drive(){
	fmt.Println("Lambo Drive ",l.LamboModel)

	
}
func (c *Chevy) Stop(){
	fmt.Println("Lambo Stop",c.l)

}

func (c *Chevy) Drive(){
	fmt.Println("Chevy Drive",c.ChevyModel)
}



func main(){
	l:= Lambo{"Gallardo"}
	c:=Chevy{"chevy"}
	l.Drive()
	c.Drive()
	
}