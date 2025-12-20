def main():
    n = 0
    while n <= 0:
        try:
            n = int(float(input('Please enter a dollar amount: ')) * 100)
        except ValueError:
            print("invalid input")

    count = 0
    for coin in [25, 10, 5, 1]:
        while n >= coin:
            count += 1
            n -= coin

    print(count)


main()