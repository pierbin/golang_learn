package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	message := []string{"D", "E", "F", "X", "S", "Y", "Z", "G", "H", "I", "J", "K", "V", "W", "L", "M", "N", "O", "P", "B", "C", "Q", "R", "A", "T", "U"}
	sort.Strings(message)
	fmt.Printf(strings.Join(message, "")) //ABCDEFGHIJKLMNOPQRSTUVWXYZ
}
