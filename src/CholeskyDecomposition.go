package main

import (
	"math"
)

var ()

func LUViaCholeskyDecomposition(c configuration) (L, U [][]float64) {
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
