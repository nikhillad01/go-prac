package main

import ("fmt"
		"reflect")// formatting

func main(){
	fmt.Println("Hello world !")	// prints in new line
	fmt.Print("Hey")
	fmt.Print("you ! \n")

	fmt.Printf("Hello %s Age is %d \n","Nikhil",24)		// Uses verbs to fill the dynamic values 

	// %s string
	// %d for integer ch
	// fmt.Println(true && true)
	// fmt.Println(true && false)
	// fmt.Println(false && false)

	const a = 1
	// var b int= 2
	// var c string="333333"

	fmt.Println(reflect.TypeOf(a))
	// var c =23
	// sum := 0
	// for i:=0;i<10;i++{
		
	// 	sum+=i
		
	// }
	// fmt.Println("--",sum)


}