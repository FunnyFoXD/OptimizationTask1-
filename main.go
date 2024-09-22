package main

import (
	"fmt"
)

func main() {
	var n, m int
	var e float64
	fmt.Scan(&n, &m, &e)
	c := make([]float64, n)   //objective function
	a := make([][]float64, n) //coeficient matrix
	b := make([]float64, n)   //right side
	for i := 0; i < n; i++ {
		a[i] = make([]float64, m)
		fmt.Scan(&c[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	fmt.Scan(&e)

}
