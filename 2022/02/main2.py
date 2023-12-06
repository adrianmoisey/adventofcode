#A for Rock, B for Paper, and C for Scissors. The second column--"

# The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).


# PART 2:
#X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.readlines()

score = 0

for i in input:
    opponent, outcome = i.split()

    if outcome == "X":
        score += 0
    elif outcome == "Y":
        score += 3
    elif outcome == "Z":
        score += 6

    if outcome == "Y": # draw
        me = opponent

    elif outcome == "X": # loss
        if opponent == "A":
            me = "C"
        elif opponent == "B":
            me = "A"
        elif opponent == "C":
            me = "B"
    elif outcome == "Z": # Win
        if opponent == "A":
            me = "B"
        elif opponent == "B":
            me = "C"
        elif opponent == "C":
            me = "A"

    if me == "A":
        score += 1
    elif me == "B":
        score += 2
    elif me == "C":
        score += 3



print(score)


