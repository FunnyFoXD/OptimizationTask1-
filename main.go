package main

import (
	"fmt"
	"math"
)

func Solution(table [][]float64, n, m, m2 int, approx float64, solVars []float64, maxOrMin int) {
	var pivotElement float64
	var minimumRatio float64 = 1e10
	var minimum float64 = 1e10
	var minimumIndex, minimumRatioIndex int
	var flag bool = true
	var prevSol float64 = 1e10
	var eps int = 3

	for flag {
		// Finding minimum element by z row
		minimum = 1e10
		for i := 1; i < n+m2+1; i++ {
			if table[m][i] < minimum {
				minimum = table[m][i]
				minimumIndex = i
			}
		}

		// Checking minimum value, if it's great or equal to 0,
		// then there are no negative values in z row, we can finish
		if minimum >= 0 {
			break
		}

		var flag bool = false
		// Finding column and row with minimum value
		minimumRatio = 1e10
		for i := 0; i < m; i++ {
			if table[i][minimumIndex] > 0 { // We cav divide only by positive values
				ratio := table[i][n+m2+1] / table[i][minimumIndex]
				if ratio < minimumRatio {
					minimumRatio = ratio
					minimumRatioIndex = i
				}
				flag = true
			}
		}

		if !flag {
			fmt.Println("Problem is unbounded")
			return
		}
		// Finding pivot element
		pivotElement = table[minimumRatioIndex][minimumIndex]

		// Changing indexes while moving variables
		if minimumIndex <= n {
			table[minimumRatioIndex][0] = float64(minimumIndex)
		} else {
			table[minimumRatioIndex][0] = float64(-(minimumIndex - n))
		}

		// Change table

		// Change all rows, except pivot row
		for i := 0; i < m+1; i++ {
			if i != minimumRatioIndex {
				factor := table[i][minimumIndex] // element in the pivot column
				for j := 1; j < n+m2+2; j++ {
					table[i][j] -= factor * table[minimumRatioIndex][j] / pivotElement
				}
			}
		}

		// Divide pivot row by pivot element
		for i := 1; i < n+m2+2; i++ {
			table[minimumRatioIndex][i] /= pivotElement
		}

		// If solution changes by smaller value than approximation
		if (math.Abs(prevSol - table[m][n+m2+1])) < approx {
			break
		}
		prevSol = table[m][n+m2+1]
	}

	for i := 0; i < m; i++ {
		if table[i][0] > 0 {
			solVars[int(table[i][0]-1)] = table[i][n+m2+1]
		}
	}
	fmt.Print("Decision variables: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%.*f ", eps, solVars[i])
	}

	fmt.Println()

	if maxOrMin == 0 {
		fmt.Printf("Maximum value of the objective function: %.*f ", eps, table[m][n+m2+1])
	} else {
		fmt.Printf("Minimum value of the objective function: %.*f ", eps, (-1)*table[m][n+m2+1])
	}
}

func main() {
	var maxOrMin int // 0- maximization, 1- minimization
	fmt.Println("Choose problem: 0 - maximization, 1 - minimization:")
	_, err := fmt.Scan(&maxOrMin)
	if err != nil || (maxOrMin != 0 && maxOrMin != 1) {
		fmt.Println("Wrong input for choosing problem!")
		return
	}

	var n, m int
	fmt.Println("Enter number of variables and constraints:")
	_, err = fmt.Scan(&n, &m)
	if err != nil {
		fmt.Println("Wrong input for number of variables or constraints!")
		return
	}

	objCoeff := make([]float64, n)
	rhsCoeff := make([]float64, m)
	constCoeff := make([][]float64, m)
	table := make([][]float64, m+1)

	for i := range constCoeff {
		constCoeff[i] = make([]float64, n)
	}

	fmt.Println("Enter objective function coefficients:")
	for i := 0; i < n; i++ {
		var temp int
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Wrong input for objective function!")
			return
		}
		objCoeff[i] = float64(temp)
	}

	fmt.Println("Enter constraints coefficients:")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var temp int
			_, err := fmt.Scan(&temp)
			if err != nil {
				fmt.Println("Wrong input for constraint function!")
				return
			}
			constCoeff[i][j] = float64(temp)
		}
	}

	fmt.Println("Enter right hand side coefficients:")
	for i := 0; i < m; i++ {
		var temp int
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Println("Wrong input for vector!")
			return
		}
		if temp < 0 {
			fmt.Println("The method is not applicable!")
			return
		}
		rhsCoeff[i] = float64(temp)
	}

	var approx float64
	fmt.Println("Enter approximation(eps):")
	_, err = fmt.Scan(&approx)
	if err != nil {
		fmt.Println("Wrong input for approximation!")
		return
	}

	extras := make([]int, m)
	m2 := m
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if objCoeff[j] == 0 && constCoeff[i][j] == 1 {
				extras[i] = j
				m2 -= 1
			}
		}
	}

	for i := range table {
		table[i] = make([]float64, n+m2+3)
	}

	// Fullfiling indexes of the table
	for i := 0; i < m; i++ {
		if extras[i] == 0 {
			table[i][0] = float64(-(i + 1))
		} else {
			table[i][0] = float64(extras[i] + 1)
		}
	}

	table[m][0] = float64(0)

	// Fulfilling table
	for i := 0; i < m+1; i++ {
		for j := 1; j < n+m2+3; j++ {
			if i < m && j <= n { // work with constraints
				table[i][j] = constCoeff[i][j-1]
			} else if i == m && j <= n {
				if maxOrMin == 0 {
					table[i][j] = float64((-1) * objCoeff[j-1])
				} else {
					table[i][j] = float64(objCoeff[j-1])
				}
			} else if i < m && j == n+m2+1 { // work with rhsCoeff
				table[i][j] = rhsCoeff[i]
			} else if j == n+m2+2 { // work with ratio
				table[i][j] = float64(0)
			} else if i < m && j > n && j <= n+m2 && m2 != 0 { //work with s1, s2 ...
				if j == n+i+1 {
					table[i][j] = float64(1)
				} else {
					table[i][j] = float64(0)
				}
			} else if i == m && j > n && m2 != 0 { // work with s1, s2 .. and rhs of z
				table[i][j] = float64(0)
			}
		}
	}

	solVars := make([]float64, n)
	for i := 0; i < n; i++ {
		solVars[i] = 0
	}

	Solution(table, n, m, m2, approx, solVars, maxOrMin)
}
