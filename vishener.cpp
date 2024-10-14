#include <iostream>
#include <string>

using namespace std;
string vishenerEncrypt(const string& text, const string& key) {
    string result;
    int keyIndex = 0;
    int keyLength = key.length();

    for (char c : text) {
        if (isalpha(c)) {
            char base = islower(c) ? 'a' : 'A';
            char encryptedChar = (c - base + (tolower(key[keyIndex % keyLength]) - 'a')) % 26 + base;
            result += encryptedChar;
            keyIndex++;
        }
        else {
            result += c;
        }
    }
    return result;
}

string vishenerDecrypt(const string& text, const string& key) {
    string result;
    int keyIndex = 0;
    int keyLength = key.length();

    for (char c : text) {
        if (isalpha(c)) {
            char base = islower(c) ? 'a' : 'A';
            char decryptedChar = (c - base - (tolower(key[keyIndex % keyLength]) - 'a') + 26) % 26 + base;
            result += decryptedChar;
            keyIndex++;
        }
        else {
            result += c;
        }
    }
    return result;
}

int main() {
    setlocale(LC_ALL, "RU");
    string text, key;

    cout << "Введите текст: ";
    getline(cin, text);

    cout << "Введите ключ: ";
    getline(cin, key);

    string encryptedText = vishenerEncrypt(text, key);
    cout << "Зашифрованный текст: " << encryptedText << endl;

    string decryptedText = vishenerDecrypt(encryptedText, key);
    cout << "Расшифрованный текст: " << decryptedText << endl;

    return 0;
}
