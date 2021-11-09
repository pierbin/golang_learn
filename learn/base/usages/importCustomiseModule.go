package main

/*
	It should off GO111MODULE.
	If GO111MODULE=on, it will always "base/usages/importCustomiseModule.go:7:8: package apps/greetings is not in GOROOT"
	After set GO111MODULE=off, the result is "Hi, welcome to local package."
	Turn off GO111MODULE, using "export GO111MODULE=off"
	Turn on GO111MODULE, using "export GO111MODULE=on"
*/

//below is used for local greetings module work.
//import "apps/greetings"

import (
	"github.com/ramseyjiang/golang_learn/greetings"
)

func main() {
	greetings.Hi()
}
