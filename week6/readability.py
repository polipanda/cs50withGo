import re

from sympy.strategies.core import switch


def main():
    for text in test_text():
        print(reading_level(text))

    # text = input("Text: ")
    # print(reading_level(text))

def reading_level(text):
    letters = count_letters(text)
    words = count_words(text)
    sentences = count_sentences(text)


    avg_letters = (letters / words) * 100
    avg_sentences = (sentences / words) * 100

    index = 0.0588 * avg_letters - 0.296 * avg_sentences - 15.8

    return grade(index)


def grade(index):
    rounded = round(index)
    rounded = 0 if rounded < 0 else rounded

    options = {
        0: "Before Grade 1",
        1: "Grade 1",
        2: "Grade 2",
        3: "Grade 3",
        4: "Grade 4",
        5: "Grade 5",
        6: "Grade 6",
        7: "Grade 7",
        8: "Grade 8",
        9: "Grade 9",
        10: "Grade 10",
        11: "Grade 11",
        12: "Grade 12",
        13: "Grade 13",
        14: "Grade 14",
        15: "Grade 15",
    }

    if rounded> 15:
        return "Grade 16+"

    return options[rounded]




def count_words(text):
    return len(re.split(' ', text))

def count_sentences(text):
    count = 0
    sentences = re.split(r'[.!?]', text)

    for sentence in sentences:
        if len(sentence) > 1:
            count += 1

    return count

def count_letters(text):
    count = 0
    for i in text:
        if i.isalpha():
            count += 1

    return count

def test_text():
    return [
        "One fish. Two fish. Red fish. Blue fish.",
        "Would you like them here or there? I would not like them here or there. I would not like them anywhere.",
        "Congratulations! Today is your day. You're off to Great Places! You're off and away!",
        "Harry Potter was a highly unusual boy in many ways. For one thing, he hated the summer holidays more than any other time of year. For another, he really wanted to do his homework, but was forced to do it in secret, in the dead of the night. And he also happened to be a wizard.",
        "In my younger and more vulnerable years my father gave me some advice that I've been turning over in my mind ever since.",
        "Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, 'and what is the use of a book,' thought Alice 'without pictures or conversation?'",
        "When he was nearly thirteen, my brother Jem got his arm badly broken at the elbow. When it healed, and Jem's fears of never being able to play football were assuaged, he was seldom self-conscious about his injury. His left arm was somewhat shorter than his right; when he stood or walked, the back of his hand was at right angles to his body, his thumb parallel to his thigh.",
        "There are more things in Heaven and Earth, Horatio, than are dreamt of in your philosophy.",
        "It was a bright cold day in April, and the clocks were striking thirteen. Winston Smith, his chin nuzzled into his breast in an effort to escape the vile wind, slipped quickly through the glass doors of Victory Mansions, though not quickly enough to prevent a swirl of gritty dust from entering along with him.",
        "A large class of computational problems involve the determination of properties of graphs, digraphs, integers, arrays of integers, finite families of finite sets, boolean formulas and elements of other countable domains."
    ]

main()