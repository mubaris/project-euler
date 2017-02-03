sum = 0

def fibonacci(n):
    if n == 1:
        return 1
    elif n == 2:
        return 2
    else:
        return fibonacci(n - 1) + fibonacci(n - 2)


i = 2

while fibonacci(i) <= 4000000:
    sum += fibonacci(i)
    i = i + 3

print(sum)
