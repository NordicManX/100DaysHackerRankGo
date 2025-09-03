package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dayOfProgrammer(year int32) string {
	// Write your code here
	if year == 1918 {
		return "26.09.1918"
	}

	isLeap := false
	if year < 1918 {
		isLeap = year%4 == 0
	} else {
		isLeap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	}

	if isLeap {
		return fmt.Sprintf("12.09.%d", year)
	}
	return fmt.Sprintf("13.09.%d", year)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
