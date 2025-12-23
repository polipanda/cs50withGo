## Week 3 Problems
https://cs50.harvard.edu/x/psets/3/

Complete the following exercises using Go.

### Sort
Some algorithms for sorting a sequence of numbers include: selection sort, bubble sort, and merge sort.

- Selection sort iterates through the unsorted portions of a list, selecting the smallest element each time and 
moving it to its correct location.
- Bubble sort compares pairs of adjacent values one at a time and swaps them if they are in the incorrect order. This 
continues until the list is sorted.
- Merge sort recursively divides the list into two repeatedly and then merges the smaller lists back into a larger one 
in the correct order.

In this problem, you’ll analyze three (compiled!) sorting programs to determine which algorithms they use.

For this problem, you’ll need some “distribution code”. Provided to you are three already-compiled C programs, 
sort1, sort2, and sort3, as well as several .txt files for input. Each of sort1, sort2, and sort3 implements a 
different sorting algorithm: selection sort, bubble sort, or merge sort (though not necessarily in that order!). Your 
task is to determine which sorting algorithm is used by each file. Write some benchmarking to timing code to figure it 
out. 

The files are found in sort_problems_files:
1. random5000.txt
2. random10000.txt
3. random50000.txt
4. reversed5000.txt
5. reversed10000.txt
6. reversed50000.txt
7. sorted5000.txt
8. sorted10000.txt
9. sorted50000.txt

### Plurality
Elections come in all shapes and sizes. Perhaps the simplest way to hold an election is via a method commonly known as 
the “plurality vote” (also known as “first-past-the-post” or “winner take all”). In the plurality vote, every voter 
gets to vote for one candidate. Whichever candidate has the greatest number of votes is declared the winner.

Write a function that determines the winner of an election using a plurality vote.

#### Specification
1. The function can accommodate a maximum of 9 candidates
2. No duplicate candidates should exist
3. Voting is done by name
4. Invalidate any votes that don't match a candidates name (throw the vote away)
5. Return the winner(s) of the election (can have multiple winners)

### Runoff
Let's explore the ranked-choice voting system known as runoff. 

In a ranked-choice system, voters can vote for more than one candidate. Instead of just voting for their top choice,
they can rank the candidates in order of preference. The resulting ballots might therefore look like the below.

|   | Ballot  | Ballot  | Ballot  | Ballot  | Ballot  |
|---|---------|---------|---------|---------|---------|
| 1 | Alice   | Alice   | Bob     | Bob     | Charlie |
| 2 | Bob     | Charlie | Alice   | Alice   | Alice   |
| 3 | Charlie | Bob     | Charlie | Charlie | Bob     |

In an instant runoff election, voters can rank as many candidates as they wish. If any candidate has a majority 
(more than 50%) of the first preference votes, that candidate is declared the winner of the election.

If no candidate has more than 50% of the vote, then an “instant runoff” occurs. The candidate who received the fewest 
number of votes is eliminated from the election, and anyone who originally chose that candidate as their first 
preference now has their second preference considered. The process repeats until someone earns >50% of the vote.

Write a function that determines the winner of an election using a runoff ranked-choice.

#### Specification
1. The function can accommodate a maximum of 9 candidates
2. No duplicate candidates should exist
3. Voting is done by name
4. Voting is done by ranks, with as many ranks as there are candidates. (ex. 5 candidates means each voter has 5 ranks)
5. Invalidate any votes that don't match a candidates name (throw the vote away)
6. Account for any runoffs
7. Return the winner(s) of the election (can have multiple winners)
8. The winner is chosen when they attain majority (>50%) or all candidates but one have been eliminated

### Tideman
Let's explore the ranked-choice voting system known as Tideman. The Tideman voting method (also known as “ranked pairs”)
is a ranked-choice voting method that’s guaranteed to produce the Condorcet winner of the election if one exists.

#### Background
Generally speaking, the Tideman method works by constructing a “graph” of candidates, where an arrow (i.e. edge) from 
candidate A to candidate B indicates that candidate A wins against candidate B in a head-to-head matchup. The graph for 
the above election, then, would look like the below.

![tideman_inconclusive_graph.png](markdown_images/tideman_inconclusive_graph.png)

The Tideman voting method consists of three parts:
1. Tally: Once all the voters have indicated their preferences, determine, for each pair of candidates, who 
the preferred candidate is and by what margin they are preferred.
2. Sort: Sort the pairs of candidates in decreasing order of strength of victory, where strength of victory is defined 
to be the number of voters who prefer the preferred candidate.
3. Lock: Starting with the strongest pair, go through the pairs of candidates in order and “lock in” each pair to the 
candidate graph, so long as locking in that pair does not create a cycle in the graph.

Once the graph is complete, the source of the graph (the one with no edges pointing towards it) is the winner.

Write a function that determines the winner of an election using a tideman ranked-choice system.

#### Specification
1. The function can accommodate a maximum of 9 candidates
2. No duplicate candidates should exist
3. Voting is done by name
4. Voting is done by ranks, with as many ranks as there are candidates. (eg. 5 candidates means each voter has 5 ranks)
5. Invalidate any votes that don't match a candidates name (throw the vote away)
6. Account for any runoffs
7. Return the winner(s) of the election
8. The winner is the candidate with the strongest pair and is the root of the graph (does not lose any pairings)