## Week 2 Problems
https://cs50.harvard.edu/x/psets/2/

Complete the following exercises using Go.

### Scrabble
In the game of Scrabble, players create words to score points, and the number of points is the sum of the point values 
of each letter in the word.

| A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q  | R | S | T | U | V | W | X | Y | Z  |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|----|---|---|---|---|---|---|---|---|----|
| 1 | 3 | 3 | 2 | 1 | 4 | 2 | 4 | 1 | 8 | 5 | 1 | 3 | 1 | 1 | 3 | 10 | 1 | 1 | 1 | 1 | 4 | 4 | 8 | 4 | 10 |

For example, if we wanted to score the word “CODE”, we would note that the ‘C’ is worth 3 points, the ‘O’ is worth 1 
point, the ‘D’ is worth 2 points, and the ‘E’ is worth 1 point. Summing these, we get that “CODE” is worth 7 points.

Implement a function that determines the winner of a Scrabble-like game. The function should accept any number of
players and their word, and use the above table to determine the winner. A tie is also possible.

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

### Caesar
One fairly basic cryptography cipher is Caesar's algorithm. It was supposedly used by Caesar himself. Caesar’s algorithm 
(i.e., cipher) encrypts messages by “rotating” each letter by 𝑘 positions. More formally, if 𝑝 is some plaintext 
(i.e., an unencrypted message), 𝑝𝑖 is the 𝑖𝑡⁢ℎ character in 𝑝, and 𝑘 is a secret key (i.e., a non-negative integer), 
then each letter, 𝑐𝑖, in the ciphertext, 𝑐, is computed as 

> 𝑐𝑖=(𝑝𝑖+𝑘) % 26

To be clear, then, here’s how encrypting `HELLO` with a key of 1 yields `IFMMP`:

| plaintext    | `H` | `E` | `L` | `L` | `O` |
|--------------|-----|-----|-----|-----|-----|
| key          | 1   | 1   | 1   | 1   | 1   |
| = ciphertext | `I` | `F` | `M` | `M` | `P` |

#### Specification
1. Your function must accept a non-negative integer as the key.
2. The function should work for all non-negative integers, even above 26.
   1. For clarity, a number above 26 just means there's been a rollover. For instance, the number 1 is A, but so is 27.
3. Accept a `plaintext` to be ciphered.
4. Return the `ciphertext`.
5. Preserve case: capitalized letters, though rotated, must remain capitalized. Same for lowercase.
6. You need only worry about rotating letters, all other characters can stay the same.

### Substitution
In a substitution cipher, we “encrypt” (i.e., conceal in a reversible way) a message by replacing every letter with 
another letter. To do so, we use a key: in this case, a mapping of each of the letters of the alphabet to the letter 
it should correspond to when we encrypt it. To “decrypt” the message, the receiver of the message would need to know the 
key, so that they can reverse the process: translating the encrypt text (generally called ciphertext) back into the 
original message (generally called plaintext).

A key, for example, might be the string `NQXPOMAFTRHLZGECYJIUWSKDVB`. This 26-character key means that `A` (the first 
letter of the alphabet) should be converted into `N` (the first character of the key), `B` (the second letter of the 
alphabet) should be converted into `Q` (the second character of the key), and so forth.

A message like `HELLO`, then, would be encrypted as `FOLLE`, replacing each of the letters according to the mapping 
determined by the key.

Write a function that encrypts a substitution cipher.

#### Specification
1. The function must accept the key to use for the substitution. The key should be case-insensitive so whether any 
character in the key is uppercase or lowercase should not affect the behavior
2. If the key is invalid (as by not containing 26 characters, containing any character that is not an alphabetic 
character, or not containing each letter exactly once), the function should return an error
3. Accept the `plaintext` to be ciphered
4. Return the `ciphertext` with each alphabetical character in the plaintext substituted for the corresponding 
5. character in the ciphertext; non-alphabetical characters must remain unchanged
6. Preserve case: capitalized letters must remain capitalized letters; lowercase letters must remain lowercase letters
