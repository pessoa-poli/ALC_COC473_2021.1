package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestReadMatrixPairToMemory(t *testing.T) {
	fmt.Println(strings.Repeat("#", 15))
	fmt.Println("Started Loading matrices.")
	start := time.Now()

	matrix1, matrix2 := readMatrixPairToMemory(MATRIX_FILE_PATH)

	fmt.Printf("First Matrix: %v\nSecond Matrix: %v\n", matrix1, matrix2)
	fmt.Println("Finished loading matrices.")
	timeElapsed := time.Since(start)
	fmt.Printf("This operation took %v\n", timeElapsed)
	fmt.Println(strings.Repeat("#", 15))
}

func TestCheckIfMatricesCanMultiply(t *testing.T) {
	matrix1, matrix2 := [][]float64{{2, 2}, {2, 2}}, [][]float64{{2, 2}, {2, 2}}
	canMultiply := CheckIfMatricesCanMultiply(matrix1, matrix2)
	fmt.Printf("Can matrices multiply? %v\n", canMultiply)
	if !canMultiply {
		log.Fatal("We should be able to multiply these matrices.")
	}
}

func TestMultiplyMatrices(t *testing.T) {
	matrix1, matrix2 := [][]float64{{1, 2, 2}, {4, 4, 2}, {4, 6, 4}}, [][]float64{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	result, _ := MultiplyMatrices(matrix1, matrix2)
	fmt.Printf("Result is: %v\n", result)
	resultShouldBe := [][]float64{{5, 10, 15}, {10, 20, 30}, {14, 28, 42}}
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			if resultShouldBe[i][j] != result[i][j] {
				fmt.Printf("Entry %v,%v is %v and should be %v\n", i, j, result[i][j], resultShouldBe[i][j])
				log.Fatal("Mutiplication went wrong.")
			}
		}
	}
}

func TestCreatePivotMatrixM(t *testing.T) {
	matrix1 := [][]float64{{1, 2, 2}, {4, 4, 2}, {4, 6, 4}}
	pivotMatrix := createPivotMatrixM(matrix1, 0)
	pivotMatrixShouldBe := [][]float64{{1, 0, 0}, {-4, 1, 0}, {-4, 0, 1}}
	fmt.Printf("PivotMatrixResult: %v", pivotMatrix)
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1); j++ {
			if pivotMatrixShouldBe[i][j] != pivotMatrix[i][j] {
				fmt.Printf("Entry %v,%v is %v and should be %v\n", i, j, pivotMatrix[i][j], pivotMatrixShouldBe[i][j])
				log.Fatal("Creating pivot Matrix went wrong")
			}
		}
	}
}

func TestGenerateLiMatrixFromUiMatrix(t *testing.T) {
	matrix1 := [][]float64{{1, 0, 0}, {4, 1, 0}, {4, 0, 1}}
	resultShouldBe := [][]float64{{1, 0, 0}, {-4, 1, 0}, {-4, 0, 1}}
	result := generateLiMatrixFromUiMatrix(matrix1)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			if resultShouldBe[i][j] != result[i][j] {
				fmt.Printf("Entry %v,%v is %v and should be %v\n", i, j, result[i][j], resultShouldBe[i][j])
				log.Fatal("Generating LU from LI went wrong.")
			}
		}
	}
	fmt.Println("Generating LU from LI succeeded.")
}

func TestFindUAndLForGivenMatrix(t *testing.T) {
	matrix1 := [][]float64{{1, 2, 2}, {4, 4, 2}, {4, 6, 4}}
	resultShouldBeL, resultShouldBeU := [][]float64{{1, 0, 0}, {4, 1, 0}, {4, 0.5, 1}}, [][]float64{{1, 2, 2}, {0, -4, -6}, {0, 0, -1}}
	resultU, resultL := LUDecomposition(matrix1)
	fmt.Printf("Results are:\nU:%v\nL:%v\n", resultU, resultL)

	for i := 0; i < len(resultShouldBeL); i++ {
		for j := 0; j < len(resultShouldBeL); j++ {
			if resultShouldBeL[i][j] != resultL[i][j] {
				fmt.Printf("Entry %v,%v is %v and should be %v\n", i, j, resultL[i][j], resultShouldBeL[i][j])
				log.Fatal("Generating L from matrix went wrong.")
			}
		}
	}

	for i := 0; i < len(resultShouldBeU); i++ {
		for j := 0; j < len(resultShouldBeU); j++ {
			if resultShouldBeU[i][j] != resultU[i][j] {
				fmt.Printf("Entry %v,%v is %v and should be %v\n", i, j, resultU[i][j], resultShouldBeU[i][j])
				log.Fatal("Generating U from matrix went wrong.")
			}
		}
	}
	fmt.Println("Generating LU from LI succeeded.")
}

func TestLoadRunConfiguration(t *testing.T) {
	c := loadRunConfiguration(CONF_DAT_PATH)
	fmt.Printf("%+v\n", c)
}

func TestForwardSubstitution(t *testing.T) {
	matrixA := [][]float64{{1, 0, 0}, {4, 1, 0}, {4, 0.5, 1}}
	vectorB := [][]float64{{3}, {6}, {10}}
	resExpected := [][]float64{{3}, {-6}, {1}}
	res := forwardSubstitution(matrixA, vectorB)
	for i := range res {
		if resExpected[i][0] != res[i][0] {
			fmt.Printf("Result was %v. It should be %v\n", res, resExpected)
			panic("Result does not match the expected.")
		}
	}

	fmt.Println(res)
}

func TestBackwardsSubstitution(t *testing.T) {
	matrixA := [][]float64{{1, 2, 2}, {0, -4, -6}, {0, 0, -1}}
	vectorB := [][]float64{{3}, {-6}, {1}}
	resExpected := [][]float64{{-1}, {3}, {-1}}
	res := backwardsSubstitution(matrixA, vectorB)
	for i := range res {
		if resExpected[i][0] != res[i][0] {
			fmt.Printf("Result was %v. It should be %v\n", res, resExpected)
			panic("Result does not match the expected.")
		}
	}

	fmt.Println(res)
}

func TestSolutionViaLUDecomposition(t *testing.T) {
	matrixA := [][]float64{{1, 2, 2}, {4, 4, 2}, {4, 6, 4}}
	vectorB := [][]float64{{3}, {6}, {10}}
	resExpected := [][]float64{{-1}, {3}, {-1}}
	c := configuration{
		systemOrder: 3,
		ICOD:        1,
		IDET:        1,
		matrixA:     matrixA,
		vectorB:     vectorB,
		TOLm:        1,
	}
	res := solutionViaLUDecomposition(c)
	for i := range res {
		if resExpected[i][0] != res[i][0] {
			fmt.Printf("Result was %v. It should be %v\n", res, resExpected)
			panic("Result does not match the expected.")
		}
	}
	fmt.Println(res)
}

func TestLUViaCholeskyDecomposition(t *testing.T) {
	matrix1 := [][]float64{{1, 0.2, 0.4}, {0.2, 1, 0.5}, {0.4, 0.5, 1}}
	resultShouldBeL, resultShouldBeU := [][]float64{{1, 0, 0}, {0.2, 0.98, 0}, {0.4, 0.43, 0.81}}, [][]float64{{1, 0.2, 0.4}, {0, 0.98, 0.43}, {0, 0, 0.81}}
	c := configuration{
		systemOrder: 3,
		ICOD:        0,
		IDET:        0,
		matrixA:     matrix1,
		vectorB:     [][]float64{},
		TOLm:        0,
	}
	resultL, resultU := LUViaCholeskyDecomposition(c)
	fmt.Printf("Results are:\nL:%.2f\nU:%.2f\n", resultL, resultU)

	for i := 0; i < len(resultShouldBeL); i++ {
		for j := 0; j < len(resultShouldBeL); j++ {
			num1 := fmt.Sprintf("%.2f", resultShouldBeL[i][j])
			num2 := fmt.Sprintf("%.2f", resultL[i][j])
			if num1 != num2 {
				fmt.Printf("Entry %v,%v is %v and should be %v\n", i, j, resultL[i][j], resultShouldBeL[i][j])
				log.Fatal("Generating L from matrix went wrong.")
			}
		}
	}

	for i := 0; i < len(resultShouldBeU); i++ {
		for j := 0; j < len(resultShouldBeU); j++ {
			num1 := fmt.Sprintf("%.2f", resultShouldBeU[i][j])
			num2 := fmt.Sprintf("%.2f", resultU[i][j])
			if num1 != num2 {
				fmt.Printf("Entry %v,%v is %v and should be %v\n", i, j, resultU[i][j], resultShouldBeU[i][j])
				log.Fatal("Generating U from matrix went wrong.")
			}
		}
	}
	fmt.Println("Generating LU from LI succeeded.")
}

func TestSolutionViaCholeskyDecomposition(t *testing.T) {
	matrixA := [][]float64{{1, 0.2, 0.4}, {0.2, 1, 0.5}, {0.4, 0.5, 1}}
	vectorB := [][]float64{{0.6}, {-0.3}, {-0.6}}
	resExpected := [][]float64{{1}, {-0}, {-1}}
	c := configuration{
		systemOrder: 3,
		ICOD:        1,
		IDET:        1,
		matrixA:     matrixA,
		vectorB:     vectorB,
		TOLm:        1,
	}
	res := SolutionViaCholeskyDecomposition(c)
	for i := range res {
		if resExpected[i][0] != res[i][0] {
			fmt.Printf("Result was %v. It should be %v\n", res, resExpected)
			panic("Result does not match the expected.")
		}
	}
	fmt.Println(res)
}

func TestWriteToFile(t *testing.T) {
	someStuff := "asdqwdqwd"
	WriteToFile("../thatfile.deleteme", someStuff)
}

func TestDeleteFile(t *testing.T) {
	DeleteFile(OUTPUT_FILE_PATH)
}

func TestCreateMatrixString(t *testing.T) {
	matrixA := [][]float64{{1, 0.2, 0.4}, {0.2, 1, 0.5}, {0.4, 0.5, 1}}
	str := CreateMatrixString(matrixA)
	fmt.Println(str)
}

func TestSolucaoPeloProcedimentoIterativoDeJacobi(t *testing.T) {
	c := configuration{
		systemOrder: 3,
		ICOD:        0,
		IDET:        0,
		matrixA:     [][]float64{{3, -1, -1}, {-1, +3, -1}, {-1, -1, +3}},
		vectorB:     [][]float64{{1}, {2}, {1}},
		TOLm:        0.001,
	}
	res := SolucaoPeloProcedimentoIterativoDeJacobi(c)
	fmt.Printf("A solucao encontrada foi %v\n", res)
}

func TestSolucaoPeloProcedimentoIterativoDeGaussSeidel(t *testing.T) {
	c := configuration{
		systemOrder: 3,
		ICOD:        0,
		IDET:        0,
		matrixA:     [][]float64{{3, -1, -1}, {-1, +3, -1}, {-1, -1, +3}},
		vectorB:     [][]float64{{1}, {2}, {1}},
		TOLm:        0.001,
	}
	res := SolucaoPeloProcedimentoIterativoDeGaussSeidel(c)
	fmt.Printf("A solucao encontrada foi %v\n", res)
}

func TestSolucaoViaMetodoDaPotencia(t *testing.T) {
	c := configuration{
		systemOrder: 3,
		ICOD:        0,
		IDET:        0,
		matrixA:     [][]float64{{1, 0.2, 0}, {0.2, 1, 0.5}, {0, 0.5, 1}},
		vectorB:     [][]float64{},
		TOLm:        0.001,
	}
	autovalor, autovetor := SolucaoViaMetodoDaPotencia(c)
	fmt.Printf("A solucao encontrada foi\nautovalor:%v\nautovetor:%v\n", autovalor, autovetor)
}

func TestAchaAutovaloresEAutovetoresViaMetodoDeJacobi(t *testing.T) {
	c := configuration{
		systemOrder: 3,
		ICOD:        0,
		IDET:        0,
		matrixA:     [][]float64{{1, 0.2, 0}, {0.2, 1, 0.5}, {0, 0.5, 1}},
		vectorB:     [][]float64{},
		TOLm:        0.01,
	}
	autovalores, autovetores := AchaAutovaloresEAutovetoresViaMetodoDeJacobi(c)
	autovaloresString := CreateMatrixString(autovalores)
	autovetoresString := CreateMatrixString(autovetores)
	fmt.Printf("A solucao encontrada foi\nautovalor:\n%s\nautovetor:\n%s\n", autovaloresString, autovetoresString)
}

func TestSolucaoViaMetodoDeJacobi(t *testing.T) {
	c := configuration{
		systemOrder: 3,
		ICOD:        0,
		IDET:        0,
		matrixA:     [][]float64{{1, 0.2, 0}, {0.2, 1, 0.5}, {0, 0.5, 1}},
		vectorB:     [][]float64{{1.2}, {1.7}, {1.5}},
		TOLm:        0.01,
	}
	sol := SolucaoViaMetodoDeJacobi(c)
	fmt.Printf("SolucaoEncontrada:\n%s\n", CreateMatrixString(sol))
}

func TestPw(t *testing.T) {
	Pw(OUTPUT_FILE_PATH, "asdad")
}
