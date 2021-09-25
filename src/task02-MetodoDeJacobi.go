package main

import (
	"fmt"
	"math"
)

func checaSePodeAplicarOMetodoDeJacobi(c configuration) (podeAplicar bool) {
	matrizSimetrica := checaSeMatrizESimetrica(c)
	//matrizDiagonalDominante := chechaSeMatrizEDiagonalPrincipal(c)
	return matrizSimetrica
}

//criaMatrizIdentidade ... Cria uma matriz identidade com as dimensoes de matrixA, dada nas configuracoes.
func criaMatrizIdentidade(c configuration) (I [][]float64) {
	numOfRows := len(c.matrixA)
	numOfCols := len(c.matrixA[0])
	I = InitializeMatrixWithZeros(numOfRows, numOfCols)
	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfCols; j++ {
			if i == j {
				I[i][j] = 1
			}
		}
	}
	return I
}

func achaMaiorElementoForaDaDiagPrincipal(matrix [][]float64) (elem float64, ir, jr int) {
	numOfRows := len(matrix)
	numOfCols := len(matrix[0])
	elem = 0
	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfCols; j++ {
			if i == j {
				continue
			}
			abs := math.Sqrt(math.Pow(matrix[i][j], 2))
			if abs > elem {
				elem = abs
				ir = i
				jr = j
			}
		}
	}
	return elem, ir, jr
}

func CalcularMatrizPk(matrix [][]float64, ir, jr int) (matrizPk [][]float64) {
	var teta float64
	if matrix[ir][ir] != matrix[jr][jr] {
		teta = 0.5 * math.Atan(2*matrix[ir][jr]/(matrix[ir][ir]-matrix[jr][jr]))
	} else {
		teta = math.Pi / 4
	}
	c := configuration{
		systemOrder: 0,
		ICOD:        0,
		IDET:        0,
		matrixA:     matrix,
		vectorB:     [][]float64{},
		TOLm:        0,
	}
	I := criaMatrizIdentidade(c)
	I[ir][ir] = math.Cos(teta)
	I[jr][jr] = math.Cos(teta)
	if ir > jr {
		I[ir][jr] = math.Sin(teta)
		I[jr][ir] = -1 * math.Sin(teta)
	} else {
		I[ir][jr] = -1 * math.Sin(teta)
		I[jr][ir] = math.Sin(teta)
	}
	return I
}

func achaMatrizTransposta(m [][]float64) (mT [][]float64) {
	numOfRows := len(m)
	numOfCols := len(m[0])
	mT = InitializeMatrixWithZeros(numOfRows, numOfCols)
	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfCols; j++ {
			mT[i][j] = m[j][i]
		}
	}
	return mT
}

func AchaAutovaloresEAutovetoresViaMetodoDeJacobi(c configuration) (autovalores, autovetores [][]float64) {
	fmt.Println("Starting solution via MetodoDeJacobi")
	podeAplicar := checaSePodeAplicarOMetodoDeJacobi(c)
	if !podeAplicar {
		fmt.Println("A matriz A não é simétrica, o método de Jacobi não pode ser aplicado")
		panic("Método escolhido não serve para matriz de input.")
	}
	//Passo1
	A1 := c.matrixA
	X1 := criaMatrizIdentidade(c)

	//passo2
	iteracao := 0

	//Passo2.1
	maiorElemento, im, jm := achaMaiorElementoForaDaDiagPrincipal(A1)

	//Passo2.2
	pk := CalcularMatrizPk(A1, im, jm)
	pkT := achaMatrizTransposta(pk)
	step0, _ := MultiplyMatrices(pkT, A1)
	Anovo, _ := MultiplyMatrices(step0, pk)
	Xnovo, _ := MultiplyMatrices(X1, pk)

	//Imprime valores da iteração
	/* 	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Iteracao %v\n", iteracao))
	   	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("A1:\n%s\nX1:\n%s\n", CreateMatrixString(A1), CreateMatrixString(X1)))
	   	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Maior elemento: A1(%v%v) %v\n", im, jm, maiorElemento))
	   	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Matriz pk:\n%s\n", CreateMatrixString(pk)))
	   	Pw(OUTPUT_FILE_PATH, fmt.Sprintln(SEPARADOR)) */

	//Atualiza valores
	A1 = Anovo
	X1 = Xnovo

	//Passo3
	for maiorElemento > c.TOLm {
		iteracao++

		maiorElemento, im, jm = achaMaiorElementoForaDaDiagPrincipal(A1)
		pk := CalcularMatrizPk(A1, im, jm)
		pkT := achaMatrizTransposta(pk)
		step0, _ := MultiplyMatrices(pkT, A1)
		Anovo, _ := MultiplyMatrices(step0, pk)
		Xnovo, _ := MultiplyMatrices(X1, pk)
		/*
			Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Iteracao %v\n", iteracao))
			Pw(OUTPUT_FILE_PATH, fmt.Sprintf("A1:\n%s\nX1:\n%s\n", CreateMatrixString(A1), CreateMatrixString(X1)))
			Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Maior elemento: A1(%v%v) %v\n", im, jm, maiorElemento))
			Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Matriz pk:\n%s\n", CreateMatrixString(pk)))
			Pw(OUTPUT_FILE_PATH, fmt.Sprintln(SEPARADOR)) */

		A1 = Anovo
		X1 = Xnovo
	}

	/* Pw(OUTPUT_FILE_PATH, fmt.Sprintf("----Resultado----\n"))
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("A1:\n%s\nX1:\n%s\n", CreateMatrixString(A1), CreateMatrixString(X1)))
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Maior elemento: A1(%v%v) %v\n", im, jm, maiorElemento))
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Matriz pk:\n%s\n", CreateMatrixString(pk)))
	Pw(OUTPUT_FILE_PATH, fmt.Sprintln(SEPARADOR)) */
	iterationsString := fmt.Sprintf("Foram necessários %v iteracoes.\n", iteracao)
	Pw(OUTPUT_FILE_PATH, iterationsString)
	return A1, X1
}

func achaInversaDeMatrizDiagonal(matrix [][]float64) (matrixInv [][]float64) {
	numOfRows := len(matrix)
	numOfColumns := len(matrix[0])
	matrixInv = InitializeMatrixWithZeros(numOfRows, numOfColumns)
	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfColumns; j++ {
			if i == j {
				matrixInv[i][j] = 1 / matrix[i][j]
			}
		}
	}
	return matrixInv
}

func formEigenValuesLineMatrix(matrix [][]float64) (res [][]float64) {
	numOfRows := len(matrix)
	numOfColumns := len(matrix[0])
	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfColumns; j++ {
			if i == j {
				res = append(res, []float64{matrix[i][j]})
			}
		}
	}
	return res
}

func achaProdutoDeAutovalores(matrizLinhaAutovalores [][]float64) (res float64) {
	res = matrizLinhaAutovalores[0][0]
	numOfRows := len(matrizLinhaAutovalores)
	for i := 1; i < numOfRows; i++ {
		res = res * matrizLinhaAutovalores[i][0]
	}
	return res
}

func SolucaoViaMetodoDeJacobi(c configuration) (sol [][]float64) {
	lambda, teta := AchaAutovaloresEAutovetoresViaMetodoDeJacobi(c)
	eigenvalues := formEigenValuesLineMatrix(lambda)
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Autovalores:\n%s\nAutovetores:\n%s\n", CreateMatrixString(eigenvalues), CreateMatrixString(teta)))
	tetaT := achaMatrizTransposta(teta)
	lambdaInv := achaInversaDeMatrizDiagonal(lambda)
	prodAutovalores := achaProdutoDeAutovalores(eigenvalues)
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("ProdAutovalores:%v\n", prodAutovalores))
	//Pw(OUTPUT_FILE_PATH, fmt.Sprintf("lambdaInversa:\n%s\n", CreateMatrixString(lambdaInv)))
	Ystep, _ := MultiplyMatrices(tetaT, c.vectorB)
	Y, _ := MultiplyMatrices(lambdaInv, Ystep)
	Ystring := CreateMatrixString(Y)
	//Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Y:\n%s\n", CreateMatrixString(Y)))
	X, _ := MultiplyMatrices(teta, Y)
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Valor final de Y:%s\n", Ystring))
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Resultado X:\n%s\n", CreateMatrixString(X)))
	return X
}
