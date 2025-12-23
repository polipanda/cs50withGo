## Week 1 Problems
https://cs50.harvard.edu/x/psets/1/

Complete the following exercises using Go.

### Hello, World
Implement a function that returns `hello, world\n`.

### Hello, It's Me
Implement a function that returns `hello, <name>`, where name is provided to the function. For instance, if the name 
is Adele, it should return `hello, Adele\n`.

### Mario
#### Mario Stair
Implement a function that creates a pyramid using hashes(#) for bricks, following this pattern:
```
       #
      ##
     ###
    ####
   #####
  ######
 #######
########
```
But accept an int for the pyramid’s actual height, so that it can also create shorter pyramids like the below:
```
  #
 ##
###
```
#### Mario Stairs With Gap
Implement a function that creates a pyramid using hashes(#) for bricks, following this pattern:
```
   #  #
  ##  ##
 ###  ###
####  ####
```
Accept an int for the height of the pyramid.
### Cash and Credit
#### Cash
Suppose you work at a store and a customer gives you \$1.00 (100 cents) for candy that costs $0.50 (50 cents). 
You’ll need to pay them their “change,” the amount leftover after paying for the cost of the candy. When making change, 
odds are you want to minimize the number of coins you’re dispensing for each customer, lest you run out.

Implement a function that returns the minimum coins needed to make the given amount of change, in cents, like so:
```
Change owed: 26
2
```
The function should only accept an int greater than 0, and works for any amount of change:
```
Change owed: 70
4
```
#### Credit
Implement a function that reports whether a given credit card number it is a valid American Express, MasterCard, or 
Visa card number. Implementation details:
1. Use the following card validation rules
   1. A VISA card starts with the number 4 and can be 13 or 16 digits long
   2. A Mastercard starts with 51, 52, 53, 54, 55 and can be 16 digits long
   3. American Express starts with 34 or 37 and can be 15 digits long
2. The function should return AMEX, MASTERCARD, VISA, or INVALID as confirmation strings
3. The input will be a string but without hyphens and without leading zeros
4. A card must pass the Luhn's algorithm checksum (see below)

#### Luhn's Algorithm
Most cards use an algorithm invented by Hans Peter Luhn of IBM. According to Luhn’s algorithm, you can determine 
if a credit card number is (syntactically) valid as follows:

    1. Multiply every other digit by 2, starting with the number’s second-to-last digit, and then add those 
    products’ digits together. (The product can be derived by subtracting the number by 9 if the result was >= 10)
        ex. 9*2=18, this is >=10, so subtract by 9 to get 9. 3*2 = 6, so keep the 6
    2. Add the sum to the sum of the digits that weren’t multiplied by 2.
    3. If the total’s last digit is 0 (or, put more formally, if the total modulo 10 is congruent to 0), 
    the number is valid!