package main

import "fmt"

func main() {
	arrayDefine()
	floatArrayCalculate()
	simpleForArrayLoop() //Highly recommend using this way
}

/*
	A single underscore (_) is used to tell the compiler that we don’t need this. Because in go, it doesn't allow useless variable exists.
	value is the same as x[i].
	The following is similar as php foreach $array as $key => $value.
	for _, value := range x {
		total += value
	}
*/
func simpleForArrayLoop() {
	x := [5]float64{98, 93, 77, 82, 83}
	var total float64
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))
}

func arrayDefine() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func floatArrayCalculate() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}
