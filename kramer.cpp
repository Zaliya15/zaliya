#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

// Функция для вычисления определителя матрицы
double determinant(const vector<vector<double>>& matrix) {
    int n = matrix.size();
    vector<vector<double>> temp = matrix;
    double det = 1.0;

    for (int i = 0; i < n; ++i) {
        int pivot = i;
        for (int j = i + 1; j < n; ++j) {
            if (fabs(temp[j][i]) > fabs(temp[pivot][i])) {
                pivot = j;
            }
        }

        if (fabs(temp[pivot][i]) < 1e-9) {
            return 0;
        }

        swap(temp[i], temp[pivot]);
        if (i != pivot) {
            det = -det;
        }

        det *= temp[i][i];
        for (int j = i + 1; j < n; ++j) {
            temp[i][j] /= temp[i][i];
        }

        for (int j = i + 1; j < n; ++j) {
            for (int k = i + 1; k < n; ++k) {
                temp[j][k] -= temp[j][i] * temp[i][k];
            }
        }
    }

    return det;
}

// Функция для решения СЛАУ методом Крамера
vector<double> solveCramer(const vector<vector<double>>& coefficients, const vector<double>& constants) {
    int n = coefficients.size();
    double detMain = determinant(coefficients);

    if (fabs(detMain) < 1e-9) {
        throw runtime_error("Система не имеет единственного решения (det = 0).");
    }

    vector<double> solutions(n);
    for (int i = 0; i < n; ++i) {
        vector<vector<double>> modifiedMatrix = coefficients;
        for (int j = 0; j < n; ++j) {
            modifiedMatrix[j][i] = constants[j];
        }
        solutions[i] = determinant(modifiedMatrix) / detMain;
    }

    return solutions;
}

int main() {
    vector<vector<double>> coefficients = {
        {2, -1, 3},
        {1, 3, 2},
        {3, 1, -1}
    };

    vector<double> constants = { 5, 14, -1 };

    try {
        vector<double> solution = solveCramer(coefficients, constants);
        cout << "Решение:" << endl;
        for (int i = 0; i < solution.size(); ++i) {
            cout << "x" << i + 1 << " = " << solution[i] << endl;
        }
    }
    catch (const runtime_error& e) {
        cout << "Ошибка: " << e.what() << endl;
    }

    return 0;
}
