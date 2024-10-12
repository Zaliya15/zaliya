def atbash(text):
  alphabet = "abcdefghijklmnopqrstuvwxyz"
  result = ""
  for char in text:
    if char.isalpha():
      if char.isupper():
        result += alphabet[25 - (ord(char) - ord('A'))]
      else:
        result += alphabet[25 - (ord(char) - ord('a'))]
    else:
      result += char
  return result

print(atbash("QWERTY"))
