package main

import "fmt"

var ()

func LUDecomposition(matrix1 [][]float64) (UMatrix, LMatrix [][]float64) {
	numberOfMisToCreate := len(matrix1) - 1 //Number of Mi matrices we will need to find.
	matrixSize := len(matrix1)              //Size of the Mi matrix

	//Initializa LMatrix
	LMatrix = initializeMatrixWithZeros(matrixSize, matrixSize)

	//We wont initialize UMatrix since it will be stored in matrix1.

	//Populate UMatrix and LMatrix values
	for i := 0; i < numberOfMisToCreate; i++ {
		MMatrix := createPivotMatrixM(matrix1, i)
		matrix1, _ = multiplyMatrices(MMatrix, matrix1)
		if i == 0 {
			LMatrix = generateLiMatrixFromUiMatrix(MMatrix)
		} else {
			MInverse := generateLiMatrixFromUiMatrix(MMatrix)
			LMatrix, _ = multiplyMatrices(LMatrix, MInverse)
		}
		fmt.Printf("\n")
	}
	//We return matrix1 in place of UMatrix since matrix1 actually stores UMatrix. Also, we keep UMatrix on the signature to keep the function easy to understand and use.
	return matrix1, LMatrix
}

func solutionViaLUDecomposition(c configuration) (res [][]float64) {
	U, L := LUDecomposition(c.matrixA)
	res1 := forwardSubstitution(L, c.vectorB)
	res2 := backwardsSubstitution(U, res1)
	return res2
}
