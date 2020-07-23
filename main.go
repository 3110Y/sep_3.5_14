package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		panic(fmt.Sprintf("Unable to open file:: %v\n", err))
	}
	defer file.Close()
	i := 1
	reader := bufio.NewReader(file)
	for number, err := reader.ReadString(';'); err != io.EOF && number != "0;"; number, err = reader.ReadString(';') {
		i++
		if err != nil {
			panic(fmt.Sprintf("Unable to read: %v\n", err))
		}
	}
	fmt.Println(i)
}
