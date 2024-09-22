package main

import (
	"fmt"
)

var eps float64 = 0.001

func Solution(objCoeff [][]float64, constCoeff, vectorCoeff []float64) {
	fmt.Println("Here will be solution")
	fmt.Println(objCoeff)
	fmt.Println(constCoeff)
	fmt.Println(vectorCoeff)
}

func main() {
	var n, m int
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		fmt.Println("Wrong input")
		return
	}

	vectorCoeff := make([]float64, m)
	constCoeff := make([]float64, m)
	objCoeff := make([][]float64, n)
	for i := range objCoeff {
		objCoeff[i] = make([]float64, n)
	}

	for i := 0; i < n; i++{
		for j := 0; j < n; j++{
			var temp int
			_, err := fmt.Scan(&temp)
			if err != nil {
				fmt.Println("Wrong input for obj. function")
				return
			}
			objCoeff[i][j] = float64(temp)
		}	
	}

	for i := 0; i < m; i++{
		var temp int
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Wrong input for const. function")
			return
		}
		constCoeff[i] = float64(temp)
	}

	for i := 0; i < m; i++{
		var temp int
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Wrong input for vector")
			return
		}
		vectorCoeff[i] = float64(temp)
	}

	Solution(objCoeff, constCoeff, vectorCoeff)
}