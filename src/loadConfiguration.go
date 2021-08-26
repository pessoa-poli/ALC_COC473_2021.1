package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	MATRIX_LINE_NUM_SEPARATOR = " "
	CONFIGURATION             configuration
)

func initLoadConfigurations() {
	CONFIGURATION = loadRunConfiguration(CONF_DAT_PATH)
}

type configuration struct {
	systemOrder int
	ICOD        int
	IDET        int
	matrixA     [][]float64
	vectorB     [][]float64
	TOLm        float64
}

func readLineIntoFloat64Array(line string, vec *[]float64) {
	//Split the line we read so we can work on each number.
	stringArray := strings.Split(line, MATRIX_LINE_NUM_SEPARATOR)
	numSlice := []float64{}
	for i := range stringArray {
		num, err := strconv.ParseFloat(stringArray[i], 64)
		if err != nil {
			fmt.Printf("Error parsing line %v\n", line)
			panic(err.Error())
		}
		numSlice = append(numSlice, float64(num))
	}
	*vec = numSlice
}

func loadRunConfiguration(configurationFilePath string) (c configuration) {
	file, err := os.Open(configurationFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	matrixLine := 0
	for scanner.Scan() {
		//Grab a line
		line := scanner.Text()
		//Check for any error during line scan
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if string(line[0]) == "#" {
			continue
		}
		switch lineNum {
		case 0:
			c.systemOrder, err = strconv.Atoi(line)
			if err != nil {
				panic("conf.dat file probably has a mistake on the systemOrder field")
			}
			c.matrixA = InitializeMatrixWithZeros(c.systemOrder, c.systemOrder)
			c.vectorB = InitializeMatrixWithZeros(c.systemOrder, 1)
		case 1:
			c.ICOD, err = strconv.Atoi(line)
			if err != nil {
				panic("conf.dat file probably has a mistake on the ICOD field")
			}
		case 2:
			c.IDET, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println("conf.dat file probably has a mistake on the IDET field")
				panic(err.Error())
			}
		case 3:
			readLineIntoFloat64Array(line, &c.matrixA[matrixLine])
			matrixLine++
			if matrixLine < c.systemOrder {
				continue
			}
			matrixLine = 0
		case 4:
			readLineIntoFloat64Array(line, &c.vectorB[matrixLine])
			matrixLine++
			if matrixLine < c.systemOrder {
				continue
			}
			matrixLine = 0
		case 5:
			c.TOLm, err = strconv.ParseFloat(line, 64)
			if err != nil {
				fmt.Println("There might be an error on TOLm field.")
				panic(err.Error())
			}
		}

		lineNum++
	}
	return c
}
