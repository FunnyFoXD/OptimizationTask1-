package main

import (
	"fmt"
)

var eps float64 = 0.001

func Solution(table [][]float64, n, m int) [][]float64 {
	var pivotElement float64
	var minimumRatio float64 = 1e10
	var minimum float64 = 1e10
	var minimumIndex, minimumRatioIndex int
	var flag bool = true
	var count int = 0
	for flag {
		minimum = 1e10
		for i := 1; i < n+m+2; i++ {
			if table[m][i] < minimum {
				minimum = table[m][i]
				minimumIndex = i
			}

		}

		if minimum >= 0 {
			fmt.Println("Done")
			return table // Answer
		}

		for i := 0; i < m+1; i++ {
			table[i][n+m+2] = table[i][n+m+1] / table[i][minimumIndex]
			if table[i][n+m+2] < minimumRatio && table[i][n+m+2] > 0 {
				minimumRatio = table[i][n+m+2]
				minimumRatioIndex = i
			}
		}

		pivotElement = table[minimumRatioIndex][minimumIndex]

		for i := 0; i < m+1; i++ {
			for j := 1; j < m+n+2; j++ {
				if i != minimumRatioIndex {
					table[i][j] = table[minimumRatioIndex][j]/pivotElement*(-table[i][minimumIndex]) + table[i][j]
				}
			}
		}

		for i := 1; i < m+n+2; i++ {
			table[minimumRatioIndex][i] = table[minimumRatioIndex][i] / pivotElement
		}
		count += 1
		fmt.Println("Table change ", count)
		for i := 0; i < m + 1; i++ {
			for j := 1; j < n+m+3; j++ {
				fmt.Print(table[i][j], " ")
			}	
		}
		fmt.Println()
	}

	return table
}

func main() {
	var n, m int
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		fmt.Println("Wrong input")
		return
	}

	objCoeff := make([]float64, n)
	rhsCoeff := make([]float64, m)
	constCoeff := make([][]float64, m)
	table := make([][]float64, m+1)
	for i := range constCoeff {
		constCoeff[i] = make([]float64, n)
	}

	for i := range table {
		table[i] = make([]float64, n+m+3)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var temp int
			_, err := fmt.Scan(&temp)
			if err != nil {
				fmt.Println("Wrong input for obj. function")
				return
			}
			constCoeff[i][j] = float64(temp)
		}
	}

	for i := 0; i < m; i++ {
		var temp int
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Wrong input for const. function")
			return
		}
		rhsCoeff[i] = float64(temp)
	}

	for i := 0; i < n; i++ {
		var temp int
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Wrong input for vector")
			return
		}
		objCoeff[i] = float64(temp)
	}
	// fullfiling indexes of the table
	for i := 0; i < m+1; i++ {
		if i != m {
			table[i][0] = float64(-(i + 1))
		} else {
			table[i][0] = float64(0)
		}
	}
	//fulfilling table
	for i := 0; i < m+1; i++ {
		for j := 1; j < n+m+3; j++ {
			if i < m && j <= n { // work with constraints
				table[i][j] = constCoeff[i][j-1]
			} else if i == m && j <= n { // work with objCoeff
				table[i][j] = float64((-1) * objCoeff[j-1])
			} else if i < m && j == n+m+1 { // work with rhsCoeff
				table[i][j] = rhsCoeff[i]
			} else if j == n+m+2 { // work with ratio
				table[i][j] = float64(0)
			} else if i < m && j > n && j <= n+m { //work with s1, s2 ...
				if j == n+i+1 {
					table[i][j] = float64(1)
				} else {
					table[i][j] = float64(0)
				}
			} else if i == m && j > n { // work with s1, s2 .. and rhz of z
				table[i][j] = float64(0)
			}
		}
	}

	fmt.Println("Initial table:")
	for i := 0; i < m + 1; i++ {
		for j := 1; j < n+m+3; j++ {
			fmt.Print(table[i][j], " ")
		}
	}
	fmt.Println()
	table = Solution(table, n, m)

	fmt.Println("Solution: ", table[m][n+m+1])
	for i := 0; i < m; i++ {
		fmt.Println(table[i][n+m+1])
	}

	fmt.Println("Table after last change")
	for i := 0; i < m; i++ {
		for j := 1; j < n+m+3; j++ {
			fmt.Print(table[i][j], " ")
		}
	}
}