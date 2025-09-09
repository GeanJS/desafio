def fibonacci(numero: int) -> int:
    if numero < 0:
        return -1
    if numero <= 1:
        return numero

    a, b = 0, 1
    for i in range(numero):
        a, b = b, a+b
    return a


if __name__ == "__main__":
    print(fibonacci(6))
