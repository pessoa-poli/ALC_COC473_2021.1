package main

import (
	"fmt"
	"math"
)

func regressaoLinear(c configuration) {
	//Monta matriz A e c
	var a11, a12, a21, a22, y11, y21 float64
	for _, v := range c.PointsList {
		a11 += 1
		a12 += v[0]
		a21 += v[0]
		a22 += math.Pow(v[0], 2)
		y11 += v[1]
		y21 += v[0] * v[1]
	}
	mA := [][]float64{{a11, a12}, {a21, a22}}
	mC := [][]float64{{y11}, {y21}}
	newC := configuration{
		matrixA: mA,
		vectorB: mC,
	}
	melhoresB := solutionViaLUDecomposition(newC)
	CoefAngular := melhoresB[1][0]
	CoefLinear := melhoresB[0][0]
	sol := CoefLinear + CoefAngular*c.TargetX
	valoresDaEquacao := fmt.Sprintf("Coeficiente Angular:%v\nCoeficienteLinear:%v\n", CoefAngular, CoefLinear)
	Pw(OUTPUT_FILE_PATH, valoresDaEquacao)
	fim := fmt.Sprintf("Dado x: %v temos y=%v\n", c.TargetX, sol)
	Pw(OUTPUT_FILE_PATH, fim)
}
