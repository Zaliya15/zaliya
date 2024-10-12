#include <iostream>
#include <vector>

using namespace std;

int binarSearch(vector<int> a, int b) {
    int first = 0;
    int last = a.size() - 1;

    while (first <= last) {
        int mid = (first + last) / 2;


        if (a[mid] == b) {
            return mid;
        }
        else if (a[mid] < b) {
            first = mid + 1;
        }
        else {
            last = mid - 1;
        }
    }

    return -1;
}

int main() {
    setlocale(LC_ALL, "rus");
    vector<int> a = { 1, 2, 3, 4, 5, 6, 7, 8, 9 };
    int b = 5;

    int index = binarSearch(a, b);

    if (index != -1) {
        cout << "Элемент найден на индексе: " << index << endl;
    }
    else {
        cout << "Элемент не найден" << endl;
    }

    return 0;
}
