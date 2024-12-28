package main

import (
	"fmt"
	"math"
)

func gaussianElimination(matrix [][]float64, b []float64) ([]float64, error) {
	n := len(matrix)

	// Прямой ход (приведение к верхнетреугольной форме)
	for k := 0; k < n; k++ {
		// Поиск ведущего элемента
		maxRow := k
		for i := k + 1; i < n; i++ {
			if math.Abs(matrix[i][k]) > math.Abs(matrix[maxRow][k]) {
				maxRow = i
			}
		}

		// Перестановка строк
		matrix[k], matrix[maxRow] = matrix[maxRow], matrix[k]
		b[k], b[maxRow] = b[maxRow], b[k]

		// Проверка на вырожденность
		if math.Abs(matrix[k][k]) < 1e-9 {
			return nil, fmt.Errorf("система уравнений вырождена")
		}

		// Преобразование строк
		for i := k + 1; i < n; i++ {
			factor := matrix[i][k] / matrix[k][k]
			for j := k; j < n; j++ {
				matrix[i][j] -= factor * matrix[k][j]
			}
			b[i] -= factor * b[k]
		}
	}

	// Обратный ход (вычисление решения)
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += matrix[i][j] * x[j]
		}
		x[i] = (b[i] - sum) / matrix[i][i]
	}

	return x, nil
}

func main() {
	matrix := [][]float64{
		{2, 1, -1},
		{-3, -1, 2},
		{-2, 1, 2},
	}
	b := []float64{8, -11, -3}

	solution, err := gaussianElimination(matrix, b)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Решение:")
	for i, x := range solution {
		fmt.Printf("x%d = %.4f\n", i+1, x)
	}
}
