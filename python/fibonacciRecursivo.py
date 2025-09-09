def fibonacci(numero: int) -> int:
    if numero < 0:
        return -1

    if numero <= 1:
        return numero

    return fibonacci(numero - 1) + fibonacci(numero - 2)


if __name__ == "__main__":
    print(fibonacci(6))
