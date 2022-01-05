from collections import Counter
from statistics import median

#with open('input-real.txt') as input_file:
with open('input.txt') as input_file:
    raw_input = input_file.read()

crabs =  raw_input.strip().split(',')
crabs = [int(crab) for crab in crabs]

crabs_count = Counter(crabs)

most_common = int(crabs_count.most_common()[0][0])
m = median(crabs)
print(most_common, m)
most_common = m

moves = 0

for crab in crabs:
    movement = int(crab) - most_common
    moves += abs(movement)

# 459589 is too high
print(moves)
