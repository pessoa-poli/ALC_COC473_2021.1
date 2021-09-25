package main

import (
	"fmt"
	"strings"
	"time"
)

var (
	OUTPUT_FILE_PATH = "result.dat"
	CONF_DAT_PATH    = "conf.dat"
	SEPARADOR        = strings.Repeat("#", 40)
	//CONF_DAT_PATH_ALT = "../conf.dat"
)

func init() {
	initLoadConfigurations()
	DeleteFile(OUTPUT_FILE_PATH)
}

func main() {
	fmt.Println("Starting Program")
	start := time.Now()
	switch CONFIGURATION.ICOD {
	case 1:
		Pw(OUTPUT_FILE_PATH, "Solving via LU decomposition.\n")
		solutionViaLUDecomposition(CONFIGURATION)
	case 2:
		Pw(OUTPUT_FILE_PATH, "Solving via Cholesky decomposition.\n")
		SolutionViaCholeskyDecomposition(CONFIGURATION)
	case 3:
		Pw(OUTPUT_FILE_PATH, "Procedimento iterativo de Jacobi\n")
		SolucaoPeloProcedimentoIterativoDeJacobi(CONFIGURATION)
	case 4:
		Pw(OUTPUT_FILE_PATH, "Procedimento iterativo Gauss-Seidel\n")
		SolucaoPeloProcedimentoIterativoDeGaussSeidel(CONFIGURATION)
	case 5:
		Pw(OUTPUT_FILE_PATH, "Método da potência\n")
		SolucaoViaMetodoDaPotencia(CONFIGURATION)
	case 6:
		Pw(OUTPUT_FILE_PATH, "Método de Jacobi\n")
		SolucaoViaMetodoDeJacobi(CONFIGURATION)
	case 7:
		Pw(OUTPUT_FILE_PATH, "Regressão linear\n")
		regressaoLinear(CONFIGURATION)
	case 8:
		Pw(OUTPUT_FILE_PATH, "Interpolacao de Lagrange\n")
		interpolacaoLagrange(&CONFIGURATION)
	}

	//Saídas relativas ao determinante para ICOD 1
	if CONFIGURATION.IDET > 0 && CONFIGURATION.ICOD == 1 {
		detA := achaDeterminante(CONFIGURATION.matrixA)
		U, L := LUDecomposition(CONFIGURATION.matrixA)
		detU := achaDeterminante(U)
		detL := achaDeterminante(L)
		fimString := fmt.Sprintf("Determinante de L:%v\nDeterminante de U:%v\nDeterminante de A:%v\n", detL, detU, detA)
		Pw(OUTPUT_FILE_PATH, fimString)
	}

	if CONFIGURATION.IDET > 0 && CONFIGURATION.ICOD == 2 {
		_, L := LUDecomposition(CONFIGURATION.matrixA)
		Ltrans := achaMatrizTransposta(L)
		detL := achaDeterminante(L)
		detLtrans := achaDeterminante(Ltrans)
		detA := achaDeterminante(CONFIGURATION.matrixA)
		fimString := fmt.Sprintf("Determinante de L:%v\nDet de Ltrans:%v\nDet de A:%v\n", detL, detLtrans, detA)
		Pw(OUTPUT_FILE_PATH, fimString)
	}

	if CONFIGURATION.IDET > 0 && CONFIGURATION.ICOD == 6 {
		detA := achaDeterminante(CONFIGURATION.matrixA)
		outdetString := fmt.Sprintf("Dete de A: %v\n", detA)
		Pw(OUTPUT_FILE_PATH, outdetString)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed since start: %v\n", elapsed)
}
