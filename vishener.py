def VishenrEncrypt(text, key):
    alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
    shifrtext = ''
    key_index = 0
    key = key.upper()

    for char in text:
        if char.upper() in alphabet:
            p = alphabet.index(char.upper())
            k = alphabet.index(key[key_index % len(key)])
            c = (p + k) % 26
            if char.isupper():
                shifrtext += alphabet[c]
            else:
                shifrtext += alphabet[c].lower()
            key_index += 1
        else:
            shifrtext += char

    return shifrtext

def VishenrDecrypt(shifrtext, key):
    alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
    text = ''
    key_index = 0
    key = key.upper()

    for char in shifrtext:
        if char.upper() in alphabet:
            c = alphabet.index(char.upper())
            k = alphabet.index(key[key_index % len(key)])
            p = (c - k + 26) % 26
            if char.isupper():
                text += alphabet[p]
            else:
                text += alphabet[p].lower()
            key_index += 1
        else:
            text += char

    return text

if __name__ == "__main__":
    txt = "Hello, goodbye"
    key = "KEY"
    encrypted = VishenrEncrypt(txt, key)
    decrypted = VishenrDecrypt(encrypted, key)

    print("Original:", txt)
    print("Encrypted:", encrypted)
    print("Decrypted:", decrypted)
