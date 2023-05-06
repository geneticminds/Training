package main

import (
	"fmt"
)

type person struct {
	fname string `example:"Garo"`
	lname string
}
type employee struct {
	person
	empId int
}

func (p person) details() {
	fmt.Println(p, " "+" I am a person")
}
func (e employee) details() {
	fmt.Println(e, " "+"I am a employee")
}

func main() {
	p1 := person{"Raj", "Kumar"}
	//p1 := person{lname : "Raj",fname : "Kumar"}
	fmt.Println(p1)
	p1.details()

	e1 := employee{person{"John", "Ponting"}, 11}
	//e1 := employee{ person{"John", "Ponting"},  11}
	e1.details()
	//var p1 person
	//fmt.Println(p1.fname)
}
