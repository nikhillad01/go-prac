package main

import ("fmt"
		"reflect"
	)

func main(){
	var name string = "Nikhil"
	var name1,name2 = "first","name"
	fmt.Println(name,name1,name2)


	newVar := "variabl declared without declaring type or with keyword" //Dont use underscores in go variables use Camel-case

	fmt.Println(newVar,reflect.TypeOf(newVar))

	age := 20

	if age < 20{
		fmt.Println("You are minor")
	} else if age > 20 {
		fmt.Println("Greater than 20")
	} else {
		fmt.Println("Exat 20 ha")
	}
	

}