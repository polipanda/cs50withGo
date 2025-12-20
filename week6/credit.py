import re

def main():
    numbers = [
        '378282246310005',
        '371449635398431',
        '5555555555554444',
        '5105105105105100',
        '4111111111111111',
        '4012888888881881',
        '4003600000000014',
        '1234567890'
    ]
    for i in numbers:
        print(i)
        print(check_credit_card(i))

    # check_credit_card(input("Enter a credit card number (no hyphens): "))


def check_credit_card(number):
    try:
        if (re.search(r'^\d{13}$|^\d{15}$|^\d{16}$', number) is None
                or not is_valid_luhns(number)):
            return "INVALID"
    except ValueError:
        return "INVALID"

    if is_mastercard(number):
        return "MASTERCARD"
    elif is_amex(number):
        return "AMEX"
    elif is_visa(number):
        return "VISA"
    else:
        return "INVALID"

def is_valid_luhns(card):
    total = 0
    length = len(card)
    start = 0 if length % 2 == 0 else 1

    for i in range(start, length, 2):
        by_two = int(card[i]) * 2
        if by_two >= 10:
            by_two -= 9
        total += by_two

    start = 0 if start == 1 else 1
    for i in range(start, length, 2):
        total += int(card[i])

    if total % 10 == 0:
        return True

    return False

def is_mastercard(card):
    if card[0:2] not in ['51', '52', '53', '54', '55'] or len(card) != 16:
        return False

    return True

def is_amex(card):
    if card[0:2] not in ['34', '37'] or len(card) != 15:
        return False

    return True

def is_visa(card):
    if card[0] != "4" or len(card) not in [13, 16]:
        return False

    return True

main()