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
}

func main() {
	fmt.Println("Starting Program")
	start := time.Now()
	switch CONFIGURATION.ICOD {
	case 1:
		Pw(OUTPUT_FILE_PATH, "Solving via LU decomposition.")
		solutionViaLUDecomposition(CONFIGURATION)
	case 2:
		Pw(OUTPUT_FILE_PATH, "Solving via Cholesky decomposition.")
		SolutionViaCholeskyDecomposition(CONFIGURATION)
	case 3:
		Pw(OUTPUT_FILE_PATH, "Procedimento iterativo de Jacobi")
		SolucaoPeloProcedimentoIterativoDeJacobi(CONFIGURATION)
	case 4:
		Pw(OUTPUT_FILE_PATH, "Procedimento iterativo Gauss-Seidel")
		SolucaoPeloProcedimentoIterativoDeGaussSeidel(CONFIGURATION)
	case 5:
		Pw(OUTPUT_FILE_PATH, "Método da potência")
		SolucaoViaMetodoDaPotencia(CONFIGURATION)
	case 6:
		Pw(OUTPUT_FILE_PATH, "Método de Jacobi")
		SolucaoViaMetodoDeJacobi(CONFIGURATION)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed since start: %v\n", elapsed)
}
