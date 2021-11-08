package main

import (
	"fmt"
	"time"
	"reflect"
)

/*
The time.Now function returns a new Time value for the current date and
time, which we store in the now variable.
*/
func main() {
	now := time.Now()
	fmt.Println(now)

	var year int = now.Year()
	var month int = int(now.Month()) //Here if use now.Month() directly, it will have a type error. So it uses int(now.Month())
	var day int = now.Day()
	var hour int = now.Hour()
	var minute int = now.Minute()
	var second int = now.Second()

	fmt.Println(reflect.TypeOf(now.Month()), now.Month())	//time.Month November
	fmt.Println(year, month, day, hour, minute, second)
}
