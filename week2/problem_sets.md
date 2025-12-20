## Week2 Problem Sets
### Scrabble
In the game of Scrabble, players create words to score points, and the number of points is the sum of the point values 
of each letter in the word.

| A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q  | R | S | T | U | V | W | X | Y | Z  |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|----|---|---|---|---|---|---|---|---|----|
| 1 | 3 | 3 | 2 | 1 | 4 | 2 | 4 | 1 | 8 | 5 | 1 | 3 | 1 | 1 | 3 | 10 | 1 | 1 | 1 | 1 | 4 | 4 | 8 | 4 | 10 |

For example, if we wanted to score the word “CODE”, we would note that the ‘C’ is worth 3 points, the ‘O’ is worth 1 
point, the ‘D’ is worth 2 points, and the ‘E’ is worth 1 point. Summing these, we get that “CODE” is worth 7 points.

Implement a program that determines the winner of a short Scrabble-like game. Your program should prompt for input 
twice: once for “Player 1” to input their word and once for “Player 2” to input their word. Then, depending on which 
player scores the most points, your program should either print “Player 1 wins!”, “Player 2 wins!”, or “Tie!” 
(in the event the two players score equal points).

### Readability
According to Scholastic, E.B. White’s Charlotte’s Web is between a second- and fourth-grade reading level, and Lois 
Lowry’s The Giver is between an eighth- and twelfth-grade reading level. What does it mean, though, for a book to be
at a particular reading level?

Well, in many cases, a human expert might read a book and make a decision on the grade (i.e., year in school) for 
which they think the book is most appropriate. But an algorithm could likely figure that out too!

Implement a program that calculates the approximate grade level needed to comprehend some text. Your program should 
print as output “Grade X” where “X” is the grade level computed, rounded to the nearest integer. If the grade level is 
16 or higher (equivalent to or greater than a senior undergraduate reading level), your program should output 
“Grade 16+” instead of giving the exact index number. If the grade level is less than 1, your program should output 
“Before Grade 1”.

For the purpose of this problem, we’ll consider any sequence of characters separated by a space to be a word (so a 
hyphenated word like “sister-in-law” should be considered one word, not three). You may assume that a sentence:

    will contain at least one word;
    will not start or end with a space; and
    will not have multiple spaces in a row.

Under those assumptions, you might already be able to find a mathematical relationship between the number of words and 
the number of spaces.

You are, of course, welcome to attempt a more sophisticated solution that can tolerate multiple spaces between words!

### Caesar
Supposedly, Caesar (yes, that Caesar) used to “encrypt” (i.e., conceal in a reversible way) confidential messages by 
shifting each letter therein by some number of places. For instance, he might write A as B, B as C, C as D, ..., and, 
wrapping around alphabetically, Z as A. And so, to say HELLO to someone, Caesar might write IFMMP instead. Upon 
receiving such messages from Caesar, recipients would have to “decrypt” them by shifting letters in the opposite 
direction by the same number of places.

The secrecy of this “cryptosystem” relied on only Caesar and the recipients knowing a secret, the number of places by 
which Caesar had shifted his letters (e.g., 1). Not particularly secure by modern standards, but, hey, if you’re perhaps
the first in the world to do it, pretty secure!

Unencrypted text is generally called `plaintext`. Encrypted text is generally called `ciphertext`. And the secret used 
is called a key.

To be clear, then, here’s how encrypting `HELLO` with a key of 1 yields `IFMMP`:

| plaintext    | `H` | `E` | `L` | `L` | `O` |
|--------------|-----|-----|-----|-----|-----|
| key          | 1   | 1   | 1   | 1   | 1   |
| = ciphertext | `I` | `F` | `M` | `M` | `P` |

More formally, Caesar’s algorithm (i.e., cipher) encrypts messages by “rotating” each letter by 𝑘 positions. More 
formally, if 𝑝 is some plaintext (i.e., an unencrypted message), 𝑝𝑖 is the 𝑖𝑡⁢ℎ character in 𝑝, and 𝑘 is a secret key 
(i.e., a non-negative integer), then each letter, 𝑐𝑖, in the ciphertext, 𝑐, is computed as 

> 𝑐𝑖=(𝑝𝑖+𝑘) % 26

wherein % 26 here means “remainder when dividing by 26.” This formula perhaps makes the cipher seem more complicated 
than it is, but it’s really just a concise way of expressing the algorithm precisely. Indeed, for the sake of 
discussion, think of A(or a) as 0 , B (or b) as 1 , …, H (or h) as 7 , I (or i) as 8 , …, and Z (or z) as 25 . Suppose 
that Caesar just wants to say Hi to someone confidentially using, this time, a key, 𝑘 , of 3. And so his plaintext, 𝑝 , 
is Hi, in which case his plaintext’s first character, 𝑝0 , is H (aka 7), and his plaintext’s second character, 𝑝1 , is 
i (aka 8). His ciphertext’s first character, 𝑐0 , is thus K, and his ciphertext’s second character, 𝑐1 , is thus L. 
Make sense?

Write a program that enables you to encrypt messages using Caesar’s cipher. At the time the user executes the program, 
they should decide, by providing a command-line argument, what the key should be in the secret message they’ll provide 
at runtime. You may assume that it will be a positive integer.

#### Specification
1. Your program must accept a non-negative integer. Let’s call it 𝑘 for the sake of discussion.
2. Do not assume that 𝑘 will be less than or equal to 26. Your program should work for all non-negative integral values 
of 𝑘 less than 2^31 −26. In other words, you don’t need to worry if your program eventually breaks if the user chooses a 
value for 𝑘 that’s too big or almost too big to fit in an `int`. (Recall that an `int` can overflow.) But, even if 𝑘 is 
greater than 26, alphabetical characters in your program’s input should remain alphabetical characters in your program’s 
output. For instance, if 𝑘 is 27 , `A` should not become `\` even though `\` is 27 positions away from `A` in ASCII, per
[asciitable.com](https://www.asciitable.com/); `A` should become `B`, since `B` is 27 positions away from `A`, provided 
you wrap around from `Z` to `A`.
3. Your program must accept a `plaintext` to be ciphered.
4. Your program must return the `ciphertext`.
5. Your program must preserve case: capitalized letters, though rotated, must remain capitalized letters; lowercase 
letters, though rotated, must remain lowercase letters.
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

Create a program that enables you to encrypt messages using a substitution cipher. At the time the user executes the 
program, they should decide on what the key should be in the secret message they’ll provide at runtime.

#### Specification
1. Your program must accept the key to use for the substitution. The key should be case-insensitive so whether any 
character in the key is uppercase or lowercase should not affect the behavior of your program.
2. If the key is invalid (as by not containing 26 characters, containing any character that is not an alphabetic 
character, or not containing each letter exactly once), your program should return an error message of your choice.
3. Your program must accept the `plaintext` to be ciphered.
4. Your program must return the `ciphertext` with each alphabetical character in the plaintext substituted for the 
corresponding character in the ciphertext; non-alphabetical characters should be outputted unchanged.
5. Your program must preserve case: capitalized letters must remain capitalized letters; lowercase letters must remain 
lowercase letters.
