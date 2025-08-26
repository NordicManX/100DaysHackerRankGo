package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func timeConversion(s string) string {

	periodo := s[len(s)-2:]
	horarioSemPeriodo := s[:len(s)-2]
	partesHorario := strings.Split(horarioSemPeriodo, ":")

	hora, _ := strconv.Atoi(partesHorario[0])

	if periodo == "PM" && hora != 12 {
		hora += 12
	} else if periodo == "AM" && hora == 12 {
		hora = 0
	}

	novaHoraStr := fmt.Sprintf("%02d", hora)

	return novaHoraStr + ":" + partesHorario[1] + ":" + partesHorario[2]

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
