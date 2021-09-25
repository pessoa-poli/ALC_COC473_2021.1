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
	POINTS_LINE_NUM_SEPARATOR = ","
	CONFIGURATION             configuration
)

func initLoadConfigurations() {
	CONFIGURATION = loadRunConfiguration(CONF_DAT_PATH)
}

type configuration struct {
	systemOrder    int
	ICOD           int
	IDET           int
	matrixA        [][]float64
	vectorB        [][]float64
	TOLm           float64
	PointsList     [][]float64
	TargetX        float64
	QtdPontosDados int
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

func readLineIntoFloat64ArrayOfTuples(line string, vec *[][]float64) {
	//Split the line we read so we can work on each number.
	stringArray := strings.Split(line, POINTS_LINE_NUM_SEPARATOR)
	if len(stringArray) > 2 {
		errString := fmt.Sprintf("One of the lines in the points list, had more than 2 numbers: %s\n", line)
		panic(errString)
	}
	numSlice := []float64{}
	for i := range stringArray {
		num, err := strconv.ParseFloat(stringArray[i], 64)
		if err != nil {
			fmt.Printf("Error parsing line %v\n", line)
			panic(err.Error())
		}
		numSlice = append(numSlice, num)
	}
	*vec = append(*vec, numSlice)
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
	qtdPontosLidos := 0
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
		case 6:
			c.QtdPontosDados, err = strconv.Atoi(line)
			if err != nil {
				panic("conf.dat file probably has a mistake on the QtdPontosDados field")
			}
		case 7:
			readLineIntoFloat64ArrayOfTuples(line, &c.PointsList)
			qtdPontosLidos++
			if qtdPontosLidos == c.QtdPontosDados {
				lineNum++
			}
			continue
		case 8:
			fmt.Printf("Reading targetX: %s\n", line)
			c.TargetX, err = strconv.ParseFloat(line, 64)
			if err != nil {
				panic("conf.dat file probably has a mistake on the TargetX field")
			}
		}

		lineNum++
	}
	return c
}
