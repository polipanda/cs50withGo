from pathlib import Path
from typing import Dict, List, Tuple, Sequence
import csv

NAME_KEY = 'name'
NO_MATCH_RESULT = 'No match'

def main():
    for csv_path, sequence_dict in test_input().items():
        str_list, rows = unpack_db_csv(csv_path)

        for sequence_path, expected_person in sequence_dict.items():
            sequence = unpack_sequence_file(sequence_path)
            str_dict = make_dict_of_str_and_longest_occurrence(str_list, sequence)
            person = find_matching_person(rows, str_dict)

            print(person)


def find_matching_person(rows: List[Dict], str_dict: Dict[str, int]) -> str:
    for row in rows:
        if all(int(row[key]) == value for key, value in str_dict.items()):
            return row[NAME_KEY]

    return NO_MATCH_RESULT

def make_dict_of_str_and_longest_occurrence(str_list: Sequence[str], sequence: str) -> Dict[str, int]:
    str_dict = {}
    for l in str_list:
        longest = longest_match(sequence, l)
        str_dict[l] = longest

    return str_dict

def unpack_sequence_file(sequence_path: str) -> str:
    with open(join_path(sequence_path), "r") as file:
        return file.read()

def unpack_db_csv(csv_path: str) -> Tuple[Sequence[str], List[Dict]]:
    with open(join_path(csv_path), "r") as file:
        reader = csv.DictReader(file)
        str_list = reader.fieldnames[1:]
        rows = [row for row in reader]

    return str_list, rows

def longest_match(sequence, subsequence):
    """Returns length of longest run of subsequence in sequence."""

    # Initialize variables
    longest_run = 0
    subsequence_length = len(subsequence)
    sequence_length = len(sequence)

    # Check each character in sequence for most consecutive runs of subsequence
    for i in range(sequence_length):

        # Initialize count of consecutive runs
        count = 0

        # Check for a subsequence match in a "substring" (a subset of characters) within sequence
        # If a match, move substring to next potential match in sequence
        # Continue moving substring and checking for matches until out of consecutive matches
        while True:

            # Adjust substring start and end
            start = i + count * subsequence_length
            end = start + subsequence_length

            # If there is a match in the substring
            if sequence[start:end] == subsequence:
                count += 1

            # If there is no match in the substring
            else:
                break

        # Update most consecutive matches found
        longest_run = max(longest_run, count)

    # After checking for runs at each character in seqeuence, return longest run found
    return longest_run

def join_path(path: str) -> Path:
    script_dir = Path(__file__).parent
    return script_dir / path

def test_input() -> Dict[str, Dict[str, str]]:
    return {
        'databases/small.csv': {
            'sequences/1.txt': 'Bob',
            'sequences/2.txt': 'No match',
            'sequences/3.txt': 'No match',
            'sequences/4.txt': 'Alice',
        },
        'databases/large.csv': {
            'sequences/5.txt': 'Lavender',
            'sequences/6.txt': 'Luna',
            'sequences/7.txt': 'Ron',
            'sequences/8.txt': 'Ginny',
            'sequences/9.txt': 'Draco',
            'sequences/10.txt': 'Albus',
            'sequences/11.txt': 'Hermione',
            'sequences/12.txt': 'Lily',
            'sequences/13.txt': 'No match',
            'sequences/14.txt': 'Severus',
            'sequences/15.txt': 'Sirius',
            'sequences/16.txt': 'No match',
            'sequences/17.txt': 'Harry',
            'sequences/18.txt': 'No match',
            'sequences/19.txt': 'Fred',
            'sequences/20.txt': 'No match',
        },
    }

main()
