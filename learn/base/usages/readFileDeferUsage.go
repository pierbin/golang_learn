package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	defer CloseFile(file)	//Close the file even if it has an error.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)	//convert string to float64
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

/**
	Using this command to run correct result.
    go run base/usages/readFileDeferUsage.go base/usages/data.txt
	The correct result is as below:

	/Users/jiangdawei/go/src/apps/learn
	Opening base/usages/data.txt
	Closing file
	Sum: 252.80


	Using this command to run incorrect result.
	go run base/usages/readFileDeferUsage.go base/usages/badData.txt
	The incorrect result is as below:

	/Users/jiangdawei/go/src/apps/learn
	Opening base/usages/badData.txt
	Closing file
	2021/11/19 12:50:21 strconv.ParseFloat: parsing "hello": invalid syntax
	exit status 1

	The file is closed by both of them, because the incorrect one is used defer to trigger it.
*/
func main() {
	//Getwd() returns a rooted path name corresponding to the current directory
	p, _ := os.Getwd()
	println(p)

	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %0.2f\n", sum)
}
