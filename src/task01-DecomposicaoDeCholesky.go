package main

import (
	"fmt"
	"math"
)

var ()

func LUViaCholeskyDecomposition(c configuration) (L, U [][]float64) {
	//Inicializando as matrizes necessárias
	L = InitializeMatrixWithZeros(c.systemOrder, c.systemOrder)
	U = InitializeMatrixWithZeros(c.systemOrder, c.systemOrder)
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

func SolutionViaCholeskyDecomposition(c configuration) (res [][]float64) {
	L, U := LUViaCholeskyDecomposition(c)
	Lstring := CreateMatrixString(L)
	Ustring := CreateMatrixString(U)
	//Escrevendo em arquivo
	Pw(OUTPUT_FILE_PATH, "Matriz L encontrada\n")
	Pw(OUTPUT_FILE_PATH, Lstring)
	Pw(OUTPUT_FILE_PATH, "Matriz U encontrada\n")
	Pw(OUTPUT_FILE_PATH, Ustring)
	res1 := forwardSubstitution(L, c.vectorB)
	res2 := backwardsSubstitution(U, res1)
	res2String := CreateMatrixString(res2)
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Resultado final\n%s\n", res2String))
	return res2
}
