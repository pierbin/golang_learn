package main

import (
	"fmt"
	"reflect"
)

//How to declare a pointer type.
//"var myIntPointer *int" or "myFloatPointer := &myFloat"
func declareHoldPointer() {
	var myInt int
	var myIntPointer *int
	myInt = 4
	myIntPointer = &myInt
	fmt.Println(myIntPointer)  //myIntPointer prints the pointer itself, the pointer address is 0xc000014098.
	fmt.Println(*myIntPointer) //\*myIntPointer prints the value at the pointer, the pointer value is 4.

	var myFloat float64
	myFloat = 1.2
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)  //myFloatPointer prints the pointer itself, the pointer address is 0xc0000140a0
	fmt.Println(*myFloatPointer) //\*myFloatPointer prints the value at the pointer, the pointer value is 1.2

	var myBool bool
	myBool = true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)  //myBoolPointer prints the pointer itself, the pointer address is 0xc0000140a8
	fmt.Println(*myBoolPointer) //\*myBoolPointer prints the value at the pointer, the pointer value is true
}

//Update the pointer value, it will also update the variable directly.
func updatePointerValue() {
	myInt := 4
	fmt.Println(myInt) //4
	myIntPointer := &myInt
	*myIntPointer = 8          //Assign a new value to the variable at the pointer.
	fmt.Println(*myIntPointer) //Print the value of the variable at the pointer, the value is 8.
	fmt.Println(myInt)         //8
}

/*
	\*float64 is used to declare that the function returns a float64 pointer.
	&myFloat is used to return a pointer of the specified type.
*/
func returnPointerFunc() *float64 {
	var myFloat = 24.81
	return &myFloat
}

/*
	\*bool is used to utilise a pointer type for this parameter.
	\*myBoolPointer is used to print the value at the pointer that gets passed in.
*/
func passPointerFunc(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer) //true.
}

//A pointer variable can only hold pointers to one type of value, so a variable might only hold *int pointers, only *float64 pointers, and so on.
func main() {
	amount := 6

	fmt.Println(amount)  //6 is the amount value.
	fmt.Println(&amount) //0xc000014060 is the amount address.

	var myInt int
	fmt.Println(reflect.TypeOf(&myInt)) //Output is *int, it get a pointer to myInt and print the pointer's type.

	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat)) //Output is *float64, it get a pointer to myFloat and print the pointer's type.

	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool)) //Output is *bool, it get a pointer to myBool and print the pointer's type.

	declareHoldPointer()
	updatePointerValue()

	//Assign the returned pointer to a variable.
	var myFloatPointer *float64 = returnPointerFunc()

	//Print the value at the pointer.
	fmt.Println(*myFloatPointer) //24.81

	myBool = true

	//&myBool is used to pass a pointer to the function.
	passPointerFunc(&myBool)
}
