package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func compareTriplets(a []int32, b []int32) []int32 {
	var aliceScore int32 = 0
	var bobScore int32 = 0

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			aliceScore++
		} else if a[i] < b[i] {
			bobScore++
		}
	}

	result := []int32{aliceScore, bobScore}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var stdout *os.File
	outputPath := os.Getenv("OUTPUT_PATH")
	if outputPath != "" {
		var err error
		stdout, err = os.Create(outputPath)
		checkError(err)
	} else {
		stdout = os.Stdout
	}

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < 3; i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var b []int32

	for i := 0; i < 3; i++ {
		bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	result := compareTriplets(a, b)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
