package main

import (
	"fmt"
	"math"
)

var ()

func LUViaCholeskyDecomposition(c configuration) (L, U [][]float64) {
	littleLMatrix := initializeMatrixWithZeros(c.systemOrder, c.systemOrder)
	L = initializeMatrixWithZeros(c.systemOrder, c.systemOrder)
	U = initializeMatrixWithZeros(c.systemOrder, c.systemOrder)
	for i := 0; i < c.systemOrder; i++ {
		var sum float64
		for k := 0; k < i; k++ {
			sum = sum + math.Pow(littleLMatrix[i][k], 2)
		}
		fmt.Printf("For i:%v sum is %v\n", i, sum)
		littleLMatrix[i][i] = math.Pow((c.matrixA[i][i] - sum), (-1 / 2))
		if littleLMatrix[i][i] <= 0 {
			panic("L[%v][%v] has value %v -> The matrix is not positive-definite.")
		}
		//fmt.Printf("LilMatrix%v%v is:%v\n", i, i, littleLMatrix[i][i])
		for j := i + 1; j < c.systemOrder; j++ {
			var sum float64
			for k := 0; k < i; k++ {
				sum = sum + littleLMatrix[i][k]*littleLMatrix[j][k]
			}
			//fmt.Printf("sum for i:%v j:%v is %v\n", i, j, sum)
			littleLMatrix[j][i] = (1 / littleLMatrix[i][i]) * (c.matrixA[i][j] - sum)
			fmt.Printf("LilMatrix%v%v is %v\n", j, i, littleLMatrix[j][i])
		}
	}

	//build L matrix
	for i := 0; i < c.systemOrder; i++ {
		for j := 0; j < c.systemOrder; j++ {
			if i < j {
				continue
			} else {
				L[i][j] = littleLMatrix[i][j]
			}
		}
	}

	//build U matrix
	for i := 0; i < c.systemOrder; i++ {
		for j := 0; j < c.systemOrder; j++ {
			U[i][j] = L[j][i]
		}
	}
	return L, U
}

func LUViaCholeskyDecomposition2(c configuration) (L, U [][]float64) {
	//Inicializando as matrizes necessárias
	L = initializeMatrixWithZeros(c.systemOrder, c.systemOrder)
	U = initializeMatrixWithZeros(c.systemOrder, c.systemOrder)
	n := c.systemOrder
	// Decomposing a matrix into Lower Triangular
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			var sum float64
			if j == i { // summation for diagonals
				for k := 0; k < j; k++ {
					sum += math.Pow(L[j][k], 2)
				}
				L[j][j] = math.Sqrt(c.matrixA[j][j] - sum)
			} else {

				// Evaluating L(i, j) using L(j, j)
				for k := 0; k < j; k++ {
					sum += (L[i][k] * L[j][k])
				}
				L[i][j] = (c.matrixA[i][j] - sum) / L[j][j]
			}
		}
	}

	//Transpose L to find U
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			U[i][j] = L[j][i]
		}
	}
	return L, U
}

func solutionViaCholeskyDecomposition(c configuration) (res [][]float64) {
	L, U := LUViaCholeskyDecomposition(c)
	res1 := forwardSubstitution(L, c.vectorB)
	res2 := backwardsSubstitution(U, res1)
	return res2
}
