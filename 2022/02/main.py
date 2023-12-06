#A for Rock, B for Paper, and C for Scissors. The second column--"
# X for Rock, Y for Paper, and Z for Scissors

# The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).


# PART 2:
#X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.readlines()

score = 0

for i in input:
    opponent, me = i.split()

    if me == "X":
        score += 1
        me = "A"
    elif me == "Y":
        score += 2
        me = "B"
    elif me == "Z":
        score += 3
        me = "C"

    if opponent == me:
        score += 3
    elif opponent == "A" and me == "B": # I win
        score += 6
    elif opponent == "B" and me == "C": # I win
        score += 6
    elif opponent == "C" and me == "A": # I win
        score += 6
    elif opponent == "A" and me == "C": #  win
        score += 0
    elif opponent == "B" and me == "A": # I win
        score += 0
    elif opponent == "C" and me == "B": # I win
        score += 0




print(score)


