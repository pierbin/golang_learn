package main

import (
	"fmt"
	"testing"
)

//Execute go test base/tests/example_test.go in command line. This file will be run.
//You can use Errorf to include additional information in your test’s failure messages, such as the
//arguments you passed to a function, the return value you got, and the value you were expecting.

//Function name should begin with "Test".
//Name after "Test" can be whatever you want.
//Test functions must accept a single parameter: a pointer to a testing.T value.
func TestTwoElements(t *testing.T) { //Function will be passed a  pointer to a testing.T value.
	list := []string{"apple", "orange"}
	want := "apple and orange"  //want is the return value we want.
	got := JoinWithCommas(list) //git is the return value we actually got.
	if got != want {
		//Instead of calling t.Errorf(), call errorStirng(), it's an error helper function.
		t.Error(errorString(list, got, want))
	}
	t.Error("no test written yet") //If remove t.Error(), the test will be passed.
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple and orange"  //want is the return value we want.
	got := JoinWithCommas(list) //git is the return value we actually got.

	if got != want {
		//Instead of calling t.Errorf(), call errorStirng(), it's an error helper function.
		t.Error(errorString(list, got, want))
	}
	t.Error("no test here either") //If remove t.Error(), the test will be passed.
}

//Functions within a _test.go file whose names do not begin with Test are not run by go test.
//They can be used by tests as “helper” functions.
//errorString is an error helper function.
func errorString(list []string, got string, want string) string {
	//Errorf() works similarly to Error, but it accepts a formatting string just like the fmt.Printf function.
	//t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
