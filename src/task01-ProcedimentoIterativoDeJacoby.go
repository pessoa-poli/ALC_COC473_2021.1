package main

import (
	"fmt"
	"math"
)

func inicializarVetorSolucaoZero(c configuration) (vetSol [][]float64) {
	for i := 0; i < c.systemOrder; i++ {
		line := []float64{1}
		vetSol = append(vetSol, line)
	}
	return vetSol
}

func achaProxVetSolDadoVetSolAtualEConfiguracoes(vetSolVelho [][]float64, c configuration) (vetSolNovo [][]float64) {
	for i := 0; i < c.systemOrder; i++ {
		var xi []float64
		var soma float64 = 0
		for j := 0; j < c.systemOrder; j++ {
			if i == j {
				continue
			}
			soma += c.matrixA[i][j] * vetSolVelho[j][0]
		}
		xiNum := (c.vectorB[i][0] - soma) / c.matrixA[i][i]
		xi = append(xi, xiNum)
		vetSolNovo = append(vetSolNovo, xi)
	}
	return vetSolNovo
}

func MediaEuclidiana(vet [][]float64) (res float64) {
	var soma float64 = 0
	for i := 0; i < len(vet); i++ {
		soma += math.Pow(vet[i][0], 2)
	}
	res = math.Sqrt(soma)
	return res
}

//subtraiVetores ... retorna o resultado de a-b
func subtraiVetores(a, b [][]float64) (res [][]float64) {
	for i := 0; i < len(a); i++ {
		num := a[i][0] - b[i][0]
		res = append(res, []float64{num})
	}
	return res
}

func CalcResiduo(vetSolNovo, vetSolVelho [][]float64) (residuo float64) {
	resSub := subtraiVetores(vetSolNovo, vetSolVelho)
	dividendo := MediaEuclidiana(resSub)
	divisor := MediaEuclidiana(vetSolNovo)
	residuo = dividendo / divisor
	return residuo
}

func chechaSeMatrizEDiagonalPrincipal(c configuration) (podeAplicar bool) {
	for i := 0; i < c.systemOrder; i++ {
		aii := c.matrixA[i][i]
		var soma float64 = 0
		for j := 0; j < c.systemOrder; j++ {
			if i == j {
				continue
			}
			soma += math.Abs(c.matrixA[i][j])
		}
		if aii < soma {
			return false
		}
	}
	return true
}

func SolucaoPeloProcedimentoIterativoDeJacobi(c configuration) (vetSol [][]float64) {
	fmt.Println("Iniciando solução pelo Procedimento Iterativo de Jacobi")
	podeAplicar := chechaSeMatrizEDiagonalPrincipal(c)
	if !podeAplicar {
		panic("O método Iterativo de Jacobi não pode ser aplicado a matriza dada pois esta não é diagonal dominante.")
	}
	vetSolAnterior := inicializarVetorSolucaoZero(c)
	vetSol = achaProxVetSolDadoVetSolAtualEConfiguracoes(vetSolAnterior, c)
	residuo := CalcResiduo(vetSol, vetSolAnterior)

	//Printando output
	vetSolAnteriorString := CreateMatrixString(vetSolAnterior)
	stringDepuracao := fmt.Sprintf("VetorSolucao:\n%s\n", vetSolAnteriorString)
	stringDepuracao2 := fmt.Sprintf("Residuo:\n%v\n", residuo)
	//Pw(OUTPUT_FILE_PATH, stringDepuracao)
	Pw(OUTPUT_FILE_PATH, stringDepuracao2)
	//Pw(OUTPUT_FILE_PATH, SEPARADOR)
	iterations := 1
	for residuo > c.TOLm {
		vetSolAnterior = vetSol
		vetSol = achaProxVetSolDadoVetSolAtualEConfiguracoes(vetSolAnterior, c)
		residuo = CalcResiduo(vetSol, vetSolAnterior)

		//Printando output loop
		//stringDepuracao = fmt.Sprintf("VetorSolucao:\n%s\n", vetSolAnteriorString)
		stringDepuracao2 = fmt.Sprintf("Residuo:\n%v\n", residuo)
		//Pw(OUTPUT_FILE_PATH, stringDepuracao)
		Pw(OUTPUT_FILE_PATH, stringDepuracao2)
		//Pw(OUTPUT_FILE_PATH, SEPARADOR)
		iterations++
	}
	fmt.Printf("Iterações total: %v\n", iterations)
	//Printando output final
	vetSolString := CreateMatrixString(vetSol)
	stringDepuracao = fmt.Sprintf("VetorSolucaoFinal:\n%s\n", vetSolString)
	stringDepuracao2 = fmt.Sprintf("Residuo Final:\n%v\n", residuo)
	Pw(OUTPUT_FILE_PATH, stringDepuracao)
	Pw(OUTPUT_FILE_PATH, stringDepuracao2)
	Pw(OUTPUT_FILE_PATH, SEPARADOR)
	return vetSol
}
