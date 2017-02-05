def largestPrimeFactor(n):
    lpf = 2
    while n > lpf:
        if n % lpf == 0:
            n = n // lpf
            lpf = 2
        else:
            lpf += 1
    return lpf

print(largestPrimeFactor(600851475143))
