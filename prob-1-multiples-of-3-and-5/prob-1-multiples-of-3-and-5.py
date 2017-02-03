total = 0

'''for i in range(1000):
    if i % 3 == 0 or i % 5 == 0:
        total += i'''
#Optimized Solution

def sumDivisibleBy(n):
    p = int(999 / n)
    return int(n * p * (p + 1) / 2)

total = sumDivisibleBy(3) + sumDivisibleBy(5) - sumDivisibleBy(15)

print(total)
