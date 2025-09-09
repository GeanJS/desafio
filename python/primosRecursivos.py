import math


def primosRecursivos(numero: int) -> list[int]:
    if numero < 2:
        return []

    primos = primosRecursivos(numero - 1)
    ehPrimo = True
    limite = int(math.sqrt(numero))

    for i in range(2, limite + 1):
        if numero % i == 0:
            ehPrimo = False
            break

    if ehPrimo:
        primos.append(numero)

    return primos


if __name__ == "__main__":
    print(primosRecursivos(10))
