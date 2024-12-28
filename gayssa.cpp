#include <iostream>
#include <vector>
#include <omp.h>

using namespace std;
void gaussian_elimination(vector<vector<double>>& A, vector<double>& B) {
    int n = A.size();

    for (int i = 0; i < n; i++) {
#pragma omp parallel for
        for (int j = i + 1; j < n; j++) {
            double factor = A[j][i] / A[i][i];
            for (int k = i; k < n; k++) {
                A[j][k] -= factor * A[i][k];
            }
            B[j] -= factor * B[i];
        }
    }

    vector<double> X(n);
    for (int i = n - 1; i >= 0; i--) {
        X[i] = B[i];
        for (int j = i + 1; j < n; j++) {
            X[i] -= A[i][j] * X[j];
        }
        X[i] /= A[i][i];
    }

    cout << "Решение: ";
    for (const auto& x : X) {
        cout << x << " ";
    }
    cout << endl;
}

int main() {
    setlocale(LC_ALL, "rus");
    vector<vector<double>> A = {
        {3, 2, -4},
        {2, 3, 3},
        {5, -3, 1}
    };
    vector<double> B = { 3, 15, 14 };

    gaussian_elimination(A, B);

    return 0;
}