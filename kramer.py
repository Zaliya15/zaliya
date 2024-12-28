import numpy as np


def solve_cramer(coefficients, constants):
    # Определяем размер матрицы
    n = len(coefficients)
    # Вычисляем главный определитель
    det_main = np.linalg.det(coefficients)
    if abs(det_main) < 1e-9:
        raise ValueError("Система не имеет единственного решения (det = 0).")

    # Вычисляем решения для каждого неизвестного
    solutions = []
    for i in range(n):
        # Создаём копию матрицы и заменяем i-й столбец на столбец свободных членов
        modified_matrix = coefficients.copy()
        modified_matrix[:, i] = constants
        # Вычисляем определитель модифицированной матрицы
        det_modified = np.linalg.det(modified_matrix)
        # Находим значение неизвестного
        solutions.append(det_modified / det_main)

    return solutions


# Пример использования
coefficients = np.array([
    [2, -1, 3],
    [1, 3, 2],
    [3, 1, -1]
], dtype=float)

constants = np.array([5, 14, -1], dtype=float)

try:
    solution = solve_cramer(coefficients, constants)
    print("Решение:")
    for i, x in enumerate(solution, start=1):
        print(f"x{i} = {x:.4f}")
except ValueError as e:
    print(e)
