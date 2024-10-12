def binarSearch(a,b):
    first, last = 0, len(a) - 1

    while first <= last:
        mid = (first + last) // 2
        if a[mid] == b:
            return mid
        elif a[mid] < b:
            first = mid + 1
        else:
            last = mid - 1

    return -1

a = [1, 2, 3, 4, 5, 6, 7, 8, 9]
b = 1
index = binarSearch(a, b)

if index != -1:
    print(f"Элемент найден на индексе: {index}")
else:
    print("Элемент не найден")
