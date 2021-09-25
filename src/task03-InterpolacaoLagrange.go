package main

import (
	"fmt"
	"strings"
)

func interpolacaoLagrange(c *configuration) {
	listaDePontos := c.PointsList
	listaDePontosString := CreateMatrixString(listaDePontos)
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Received points list:\n%s\n", listaDePontosString))
	Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Received TargetX: %v\n", c.TargetX))
	xTarget := c.TargetX
	var res float64 = 0
	//Para cada ponto v, calculamos um fi
	for k, v := range listaDePontos {
		yi := v[1]
		var fiNumerador, fiDenominador float64
		firstIter := true
		//Montagem do fi_i para ki
		for k2, v2 := range listaDePontos {
			if k2 == k {
				//Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Skipping k2=k\n")) //Pula o incremento do fiNumerador e fiDenominador se k=i
				continue
			}
			if firstIter {
				firstIter = false
				fiNumerador = xTarget - v2[0]
				fiDenominador = v[0] - v2[0]
				//fmt.Printf("k2: %v; fiNumerador: %v; fiDenominador: %v\n", k2, fiNumerador, fiDenominador)
				continue
			}
			fiNumerador = fiNumerador * (xTarget - v2[0])
			fiDenominador = fiDenominador * (v[0] - v2[0])
			//fmt.Printf("k2: %v; fiNumerador: %v; fiDenominador: %v\n", k2, fiNumerador, fiDenominador)
		}
		fii := fiNumerador / fiDenominador
		res += yi * fii
		//Pw(OUTPUT_FILE_PATH, fmt.Sprintf("Res parcial: %v para k=%v yi=%v fii=%v\n", res, k, yi, fii))
	}
	fmt.Println(strings.Repeat("#", 16))
	resString := fmt.Sprintf("Resultado final foi %v\n", res)
	Pw(OUTPUT_FILE_PATH, resString)
}
