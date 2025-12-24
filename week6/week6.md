## Week 6 Problems
https://cs50.harvard.edu/x/psets/6/

Complete the following exercises using Python.

### Mario Stairs
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

### Mario Stairs With Gap
Implement a function that creates a pyramid using hashes(#) for bricks, following this pattern:
```
   #  #
  ##  ##
 ###  ###
####  ####
```
Accept an int for the height of the pyramid.

### Cash
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

### Credit

Implement a function that reports whether a given credit card number it is a valid American Express, MasterCard, or
Visa card number. Implementation details:
1. Use the following card validation rules
    1. A VISA card starts with the number 4 and can be 13 or 16 digits long
    2. A Mastercard starts with 51, 52, 53, 54, 55 and can be 16 digits long
    3. American Express starts with 34 or 37 and can be 15 digits long
2. The function should return AMEX, MASTERCARD, VISA, or INVALID as confirmation strings
3. The input will be a string but without hyphens and without leading zeros
4. A card must pass the Luhn's algorithm checksum (see below)
5. 
#### Luhn's Algorithm
Most cards use an algorithm invented by Hans Peter Luhn of IBM. According to Luhn’s algorithm, you can determine
if a credit card number is (syntactically) valid as follows:

    1. Multiply every other digit by 2, starting with the number’s second-to-last digit, and then add those 
    products’ digits together. (The product can be derived by subtracting the number by 9 if the result was >= 10)
        ex. 9*2=18, this is >=10, so subtract by 9 to get 9. 3*2 = 6, so keep the 6
    2. Add the sum to the sum of the digits that weren’t multiplied by 2.
    3. If the total’s last digit is 0 (or, put more formally, if the total modulo 10 is congruent to 0), 
    the number is valid!

### Readability
There are algorithms that allow programmers to estimate the reading level of any given text. One such algorithm is the
Coleman-Liau index

> index = 0.0588 * L - 0.296 * S - 15.8

where L is the average number of letters per 100 words in the text, and S is the average number of sentences per
100 words in the text.

Implement a function that calculates the approximate grade level to comprehend some text.
1. Return 'Grade X' where X is the grade calculated.
2. Return 'Grade 16+' if the calculation is 16 or higher
3. Return 'Before Grade 1' if the calculation is less than 1

For the purpose of this problem, any sequence of characters separated by a space to be a word (so a
hyphenated word like “sister-in-law” should be considered one word, not three). You may assume that a sentence:

    will contain at least one word;
    will not start or end with a space; and
    will not have multiple spaces in a row.

### DNA
Individual can be identified through DNA profiling. Using a DNA database and a DNA sequence, people can be identified
by finding the number of repeat occurrences of Short Tandem Repeats (STRs). A DNA database in a CSV file may look 
like this:

>name,AGAT,AATG,TATC
Alice,28,42,14
Bob,17,22,19
Charlie,36,18,25

The data above suggests that Alice has the sequence `AGAT` repeated 28 times consecutively, `AATG` 42 times, and `TATC`
14 times. Using a single sequence of DNA, and finding the longest consecutive sequence of repeated `AGAT`, `AATG`, and
`TATC`, we can search the database for any matches to all 3 counts.

Write a function that will take a sequence of DNA and a CSV file containing STR counts for a list of individuals and 
then output to whom the DNA (most likely) belongs.

#### Specification
1. The function should accept the name of a CSV file containing the STR counts for a list of individuals, and accept
the name of a text file containing the DNA sequence to identify
2. Open the CSV file and read the contents into memory
   1. You may assume the first row of the CSV is the column names. The first column will be the word `name` and the
   remaining columns will be the STR sequences
3. Open the DNA sequence file and read the contents into memory
4. For each STR (from the CSV column names), the function should compute the longest run of consecutive repeats of the
STR in the DNA sequence (a helper function `longest_match` will do this for you)
5. If the STR counts match with any individual in the CSV file, print their name
   1. You may assume the STR counts are unique to only one name
6. If the STR counts do not match any individual, print `No match`