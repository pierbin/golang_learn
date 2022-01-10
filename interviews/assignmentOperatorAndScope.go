package main

import (
	"fmt"
	"os"
)

func main() {
	example()
	solution()
}

/** Expected outputs.
there
are
no
strings
on
me
*/
func example() {
	var data []string

	killswitch := os.Getenv("KILLSWITCH")

	if killswitch == "" {
		fmt.Println("kill switch is off")
		data, err := getData()

		if err != nil {
			panic("ERROR!")
		}

		fmt.Printf("Data was fetched! %d\n\n", len(data))
	}

	for _, item := range data {
		fmt.Println(item)
	}
}

func getData() ([]string, error) {
	return []string{"there", "are", "no", "strings", "on", "me"}, nil
}

func solution() {
	var data []string
	var err error // Declaring err to make sure we can use = instead of :=

	killswitch := os.Getenv("KILLSWITCH")

	if killswitch == "" {
		fmt.Println("kill switch is off")
		data, err = getData() // Here replace ":=" by "=", then the data scope is changed.

		if err != nil {
			panic("ERROR!")
		}

		fmt.Printf("Data was fetched! %d\n", len(data))
	}

	for _, item := range data {
		fmt.Println(item)
	}
}

// In the firstone(),  it defines the data at the beginning, and the date is also defined in the if condition,
// so the scope of data, outside the if condition is the global data, inside the if condition is the partial data.
