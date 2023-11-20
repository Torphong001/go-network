package main

import "fmt"

// var x int =5
// var y int

// func Init(){
// 	fmt.Println("x =",x)
// }
func main(){
	// var t bool =true
	// var f bool 


	// names := "Torphong Sriboonruang"

	// fmt.Println("t is ",t)
	// fmt.Println("f is ",f)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(names)

	var age int =40
	var favNum float64 = 1.6180339

	var complex128 complex128 = 5+5i

	var r rune = 10

	fmt.Println("age is ",age, "favNum is ", favNum)
	fmt.Println("complex is ", complex128)
	fmt.Println("rune is ", r)

	var myName string = "Torgon"

	fmt.Println(myName + " is a god")
	fmt.Println(myName[3])
	fmt.Println("ความยาวของ myName",len(myName))

	var arry[5] float64
	arry[0] = 98.7
	arry[1] = 93.2
	arry[2] = 77.2
	arry[3] = 66.2
	arry[4] = 51.2

	fmt.Println(arry)
	fmt.Println("ความยาวของ arry",len(arry))
	fmt.Println("กำหนด arry",arry[2])

	arry3 := [3] float64 {68, 63, 77}
	fmt.Println(arry3)

	var s [] float64 = arry[2:5]
	fmt.Println(s)

	type Person struct{
		name string
		age int
	}

	var p Person
	p.name = "Torphong"
	p.age = 40
	fmt.Println(p)

	var x int = 5
	var xPtr *int = &x
	fmt.Println(xPtr)
}

