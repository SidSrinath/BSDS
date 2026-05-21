package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func writeToFile() {
	//open the file
	file, err := os.OpenFile("file-access1.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file file-access1.txt")
	}
	//defer close
	defer file.Close()

	//record start time
	startTime := time.Now()

	//loop 100k times
	for i := 0; i < 100000; i++ {
		fmt.Fprintf(file, "%d\n", i)
	}

	//record end time
	endTime := time.Now()

	fmt.Println("Time taken in regular write: ", endTime.Sub(startTime))
}

func writeToFileBuffered() {
	//open file
	file, err := os.OpenFile("file-access2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error opening file file-access2.txt")
	}

	//defer close
	defer file.Close()

	//wrap file in bufio.NewWriter
	writer := bufio.NewWriter(file)

	//record start time
	startTime := time.Now()

	//run loop and write to file
	for i := 0; i < 100000; i++ {
		fmt.Fprintf(writer, "%d\n", i)
	}

	//record end time
	endTime := time.Now()

	fmt.Println("Time taken in buffered approach: ", endTime.Sub(startTime))
}

func main() {
	writeToFile()

	writeToFileBuffered()
}
