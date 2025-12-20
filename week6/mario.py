def main():
    n = 0
    while n <= 0 or n > 8:
        try:
            n = int(input('Please enter a number between 1 and 8: '))
        except ValueError:
            print("invalid input")

    for i in range(1, n+1):
        print(' ' * (n-i) + '#' * i)

    for i in range(1, n+1):
        print(' ' * (n-i)  + '#' * i + ' ' + '#' * i)

main()