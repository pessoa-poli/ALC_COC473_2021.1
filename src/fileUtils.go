package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func WriteToFile(filePath, stuffToWrite string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	_, err = f.Write([]byte(stuffToWrite))
	if err != nil {
		panic(err.Error())
	}
}

func Pw(filePath, stuffToWriteAndPrint string) {
	WriteToFile(filePath, stuffToWriteAndPrint)
	fmt.Print(stuffToWriteAndPrint)
}

func DeleteFile(filePath string) {
	_, err := os.Open(filePath)
	if errors.Is(os.ErrNotExist, err) {
		fmt.Printf("No file found to delete on path: %v\n", filePath)
	}

	_, err = os.Stat(filePath)
	if err == nil {
		err = os.Remove(filePath)

		if err != nil && err != os.ErrNotExist {
			panic(err.Error())
		}
		return
	}
	fmt.Println("No output file found to delete")
}

func CreateMatrixString(matrix [][]float64) (matrixASString string) {
	strB := strings.Builder{}
	//Find out the longest number to print
	numLen := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			str := fmt.Sprintf("%v", matrix[i][j])
			if len(str) > numLen {
				numLen = len(str)
			}
		}
	}
	//Print each number and fill in gaps to match numLen
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			var numAsStringLen int
			num := matrix[i][j]
			numAsStringLen = len(fmt.Sprintf("%v", num))
			strB.Write([]byte(fmt.Sprintf("%v", num)))
			if numAsStringLen < numLen {
				strB.Write([]byte(strings.Repeat(" ", numLen-numAsStringLen)))
			}
			if j+1 == len(matrix[i]) {
				strB.Write([]byte("\n"))
			} else {
				strB.Write([]byte("  "))
			}
		}
	}
	return strB.String()
}
