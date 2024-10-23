def linearSearch(array, value):
    for i in range(len(array)):
        if array[i] == value:
            return 1
    return -1


print(linearSearch([20, 30, 50, 80, 90], 90))
