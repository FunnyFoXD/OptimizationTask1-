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
		// Step 1: Найти минимальный элемент в строке целевой функции (по Z)
		minimum = 1e10
		for i := 1; i < n+m+1; i++ {
			if table[m][i] < minimum {
				minimum = table[m][i]
				minimumIndex = i
			}
		}

		// Если минимальное значение больше или равно нулю, завершить
		if minimum >= 0 {
			fmt.Println("Done")
			return table
		}

		// Step 2: Найти строку с минимальным положительным отношением
		minimumRatio = 1e10
		for i := 0; i < m; i++ {
			if table[i][minimumIndex] > 0 { // Мы можем делить только на положительные значения
				ratio := table[i][n+m+1] / table[i][minimumIndex]
				if ratio < minimumRatio {
					minimumRatio = ratio
					minimumRatioIndex = i
				}
			}
		}

		// Step 3: Найти разрешающий элемент
		pivotElement = table[minimumRatioIndex][minimumIndex]

		// Step 4: Обновляем таблицу

		// 1. Нормализуем разрешающую строку (чтобы разрешающий элемент стал 1):
		for i := 1; i < n+m+2; i++ {
			table[minimumRatioIndex][i] /= pivotElement
		}

		// 2. Обновляем строки, кроме разрешающей:
		for i := 0; i < m+1; i++ {
			if i != minimumRatioIndex { // Не трогаем разрешающую строку
				factor := table[i][minimumIndex] // элемент, который должен стать 0
				for j := 1; j < n+m+2; j++ {
					table[i][j] -= factor * table[minimumRatioIndex][j]
				}
			}
		}


		// Вывести текущую таблицу для отладки
		count++
		fmt.Printf("Iteration %d\n", count)
		for i := 0; i < m+1; i++ {
			for j := 1; j < n+m+3; j++ {
				fmt.Printf("%f ", table[i][j])
			}
			fmt.Println()
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

	fmt.Println("Solution:", table[m][n+m+1])
	
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