package main

import (
	"fmt"
	"strings"
)

var ()

func LUDecomposition(matrix1 [][]float64) (UMatrix, LMatrix [][]float64) {
	fmt.Println(strings.Repeat("#", 15))
	fmt.Println("Starting LU decomposition.")

	numberOfMisToCreate := len(matrix1) - 1 //Number of Mi matrices we will need to find.
	matrixSize := len(matrix1)              //Size of the Mi matrix

	//Initializa LMatrix
	LMatrix = InitializeMatrixWithZeros(matrixSize, matrixSize)

	//We wont initialize UMatrix since it will be stored in matrix1.

	//Populate UMatrix and LMatrix values
	for i := 0; i < numberOfMisToCreate; i++ {
		MMatrix := createPivotMatrixM(matrix1, i)
		matrix1, _ = MultiplyMatrices(MMatrix, matrix1)
		if i == 0 {
			LMatrix = generateLiMatrixFromUiMatrix(MMatrix)
		} else {
			MInverse := generateLiMatrixFromUiMatrix(MMatrix)
			LMatrix, _ = MultiplyMatrices(LMatrix, MInverse)
		}
		fmt.Printf("\n")
	}

	fmt.Println("Finished LU decomposition")
	//We return matrix1 in place of UMatrix since matrix1 actually stores UMatrix. Also, we keep UMatrix on the signature to keep the function easy to understand and use.
	return matrix1, LMatrix
}

func solutionViaLUDecomposition(c configuration) (res [][]float64) {
	U, L := LUDecomposition(c.matrixA)
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
	//Escrevendo resultado final
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Resultado final:\n%s\n", res2String))
	return res2
}
