package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func mdc(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmc(a, b int32) int32 {
	if a == 0 || b == 0 {
		return 0
	}
	res := (a * b) / mdc(a, b)
	if res < 0 {
		return -res
	}
	return res
}

func getTotalX(a []int32, b []int32) int32 {

	mmcDeA := a[0]
	for i := 1; i < len(a); i++ {
		mmcDeA = mmc(mmcDeA, a[i])
	}

	mdcDeB := b[0]
	for i := 1; i < len(b); i++ {
		mdcDeB = mdc(mdcDeB, b[i])
	}

	contador := int32(0)
	multiplo := mmcDeA
	for multiplo <= mdcDeB {
		if mdcDeB%multiplo == 0 {
			contador++
		}
		multiplo += mmcDeA
	}

	return contador
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

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
