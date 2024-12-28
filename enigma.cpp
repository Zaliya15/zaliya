#include <iostream>
#include <string>
#include <vector>
#include <cctype>

using namespace std;

// Класс для реализации машины "Энигма"
class Enigma {
private:
    vector<vector<int>> rotors;
    vector<int> reflector;
    vector<int> ringSettings;
    vector<int> positions;

    void rotateRotors() {
        // Поворот первого ротора
        positions[0] = (positions[0] + 1) % 26;
        if (positions[0] == 0) { // Двойной шаг
            positions[1] = (positions[1] + 1) % 26;
            if (positions[1] == 0) {
                positions[2] = (positions[2] + 1) % 26;
            }
        }
    }

    char encodeChar(char ch) {
        if (!isalpha(ch)) return ch; // Пропускаем неалфавитные символы

        ch = toupper(ch);
        int index = ch - 'A';

        // Поворот роторов перед шифрованием символа
        rotateRotors();

        // Проход через роторы в прямом направлении
        for (int i = 0; i < rotors.size(); ++i) {
            index = (rotors[i][(index + positions[i] - ringSettings[i] + 26) % 26] - positions[i] + ringSettings[i] + 26) % 26;
        }

        // Проход через отражатель
        index = reflector[index];

        // Проход через роторы в обратном направлении
        for (int i = rotors.size() - 1; i >= 0; --i) {
            index = (find(rotors[i].begin(), rotors[i].end(), (index + positions[i]) % 26) - rotors[i].begin() - positions[i] + ringSettings[i] + 26) % 26;
        }

        return static_cast<char>(index + 'A');
    }

public:
    Enigma(vector<vector<int>> rotors, vector<int> reflector, vector<int> ringSettings, vector<int> initialPositions)
        : rotors(rotors), reflector(reflector), ringSettings(ringSettings), positions(initialPositions) {}

    string encrypt(const string& text) {
        string result;
        for (char ch : text) {
            result += encodeChar(ch);
        }
        return result;
    }
};

int main() {
    // Пример настроек
    vector<vector<int>> rotors = {
        {4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9},
        {0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4},
        {1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14}
    };
    vector<int> reflector = { 24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19 };
    vector<int> ringSettings = { 0, 0, 0 };
    vector<int> initialPositions = { 0, 0, 0 };

    // Инициализация машины
    Enigma enigma(rotors, reflector, ringSettings, initialPositions);

    // Шифрование текста
    string plaintext = "HELLO WORLD";
    string ciphertext = enigma.encrypt(plaintext);

    cout << "Зашифрованный текст: " << ciphertext << endl;

    return 0;
}
