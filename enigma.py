class Enigma:
    def __init__(self, rotors, reflector, ring_settings, initial_positions):
        self.rotors = rotors
        self.reflector = reflector
        self.ring_settings = ring_settings
        self.positions = initial_positions

    def rotate_rotors(self):
        # Поворот первого ротора, возможно движение других
        self.positions[0] = (self.positions[0] + 1) % 26
        if self.positions[0] == 0:  # Двойной шаг
            self.positions[1] = (self.positions[1] + 1) % 26
            if self.positions[1] == 0:
                self.positions[2] = (self.positions[2] + 1) % 26

    def encode_char(self, char):
        if not char.isalpha():
            return char  # Пропускаем неалфавитные символы

        char = char.upper()
        index = (ord(char) - ord('A'))

        # Поворот роторов перед шифрованием символа
        self.rotate_rotors()

        # Проход через роторы в прямом направлении
        for i in range(len(self.rotors)):
            index = (self.rotors[i][(index + self.positions[i] - self.ring_settings[i]) % 26] - self.positions[i] + self.ring_settings[i]) % 26

        # Проход через отражатель
        index = self.reflector[index]

        # Проход через роторы в обратном направлении
        for i in range(len(self.rotors) - 1, -1, -1):
            index = (self.rotors[i].index((index + self.positions[i]) % 26) - self.positions[i] + self.ring_settings[i]) % 26

        return chr(index + ord('A'))

    def encrypt(self, text):
        return ''.join(self.encode_char(char) for char in text)


# Пример настроек
rotor1 = [4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9]
rotor2 = [0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4]
rotor3 = [1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14]
reflector = [24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19]

ring_settings = [0, 0, 0]  # Настройки кольца
initial_positions = [0, 0, 0]  # Начальные позиции

# Инициализация машины
enigma = Enigma([rotor1, rotor2, rotor3], reflector, ring_settings, initial_positions)

# Шифрование сообщения
plaintext = "HELLO WORLD"
ciphertext = enigma.encrypt(plaintext)
print("Зашифрованный текст:", ciphertext)
