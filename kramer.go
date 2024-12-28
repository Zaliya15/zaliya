package main

import (
	"errors"
	"fmt"
	"math"
)

// Функция для вычисления определителя матрицы
func determinant(matrix [][]float64) float64 {
	n := len(matrix)
	temp := make([][]float64, n)
	for i := range matrix {
		temp[i] = append([]float64(nil), matrix[i]...)
	}

	det := 1.0

	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if math.Abs(temp[j][i]) > math.Abs(temp[pivot][i]) {
				pivot = j
			}
		}

		if math.Abs(temp[pivot][i]) < 1e-9 {
			return 0
		}

		temp[i], temp[pivot] = temp[pivot], temp[i]
		if i != pivot {
			det = -det
		}

		det *= temp[i][i]
		for j := i + 1; j < n; j++ {
			temp[i][j] /= temp[i][i]
		}

		for j := i + 1; j < n; j++ {
			for k := i + 1; k < n; k++ {
				temp[j][k] -= temp[j][i] * temp[i][k]
			}
		}
	}

	return det
}

// Функция для решения СЛАУ методом Крамера
func solveCramer(coefficients [][]float64, constants []float64) ([]float64, error) {
	n := len(coefficients)
	if len(constants) != n {
		return nil, errors.New("размер матрицы коэффициентов и вектора свободных членов не совпадает")
	}

	detMain := determinant(coefficients)
	if math.Abs(detMain) < 1e-9 {
		return nil, errors.New("система не имеет единственного решения (det = 0)")
	}

	solutions := make([]float64, n)
	for i := 0; i < n; i++ {
		modifiedMatrix := make([][]float64, n)
		for j := 0; j < n; j++ {
			modifiedMatrix[j] = append([]float64(nil), coefficients[j]...)
			modifiedMatrix[j][i] = constants[j]
		}
		solutions[i] = determinant(modifiedMatrix) / detMain
	}

	return solutions, nil
}

func main() {
	coefficients := [][]float64{
		{2, -1, 3},
		{1, 3, 2},
		{3, 1, -1},
	}

	constants := []float64{5, 14, -1}

	solution, err := solveCramer(coefficients, constants)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Решение:")
	for i, x := range solution {
		fmt.Printf("x%d = %.4f\n", i+1, x)
	}
}
