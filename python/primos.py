import math


def primos(numero: int) -> list:
    if numero < 2:
        return []

    primos = []
    for i in range(2, numero+1):
        ehPrimo = True
        limite = int(math.sqrt(i))
        for j in range(2, limite+1):
            if i % j == 0:
                ehPrimo = False
                break

        if ehPrimo:
            primos.append(i)

    return primos
