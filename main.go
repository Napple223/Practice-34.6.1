package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var outputFileName string
	if len(os.Args) == 1 {
		fmt.Println("Введите хотябы 1 аргумент (имя входного файла) в параметр командной строки")
		return
	}
	if len(os.Args) == 3 {
		outputFileName = os.Args[2]
	} else {
		fmt.Println("имя выходного файла не введено, будет создан файл outputData.txt")
		outputFileName = "./outputData.txt"
	}
	_ = os.Remove(outputFileName)
	outputFile, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()
	inputFileName := os.Args[1]
	inputFile, err := os.OpenFile(inputFileName, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()
	fileReader := bufio.NewReader(inputFile)
	fileWriter := bufio.NewWriter(outputFile)
	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		re := regexp.MustCompile(`([0-9]+)+([\+-])+([0-9+])=`)
		sub := re.FindAllStringSubmatch(string(line), -1)
		if len(sub) == 0 {
			continue
		}
		switch {
		case sub[0][2] == "+":
			f, _ := strconv.Atoi(sub[0][1])
			s, _ := strconv.Atoi(sub[0][3])
			res := f + s
			output := sub[0][1] + sub[0][2] + sub[0][3] + "=" + strconv.Itoa(res) + "\n"
			_, _ = fileWriter.WriteString(output)
		case sub[0][2] == "-":
			f, _ := strconv.Atoi(sub[0][1])
			s, _ := strconv.Atoi(sub[0][3])
			res := f - s
			output := sub[0][1] + sub[0][2] + sub[0][3] + "=" + strconv.Itoa(res) + "\n"
			_, _ = fileWriter.WriteString(output)
		}
	}
	_ = fileWriter.Flush()
	fmt.Printf("Данные записаны в файл %v", outputFileName)
}
