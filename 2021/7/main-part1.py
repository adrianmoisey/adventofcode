from collections import Counter
from statistics import median

with open('input-real.txt') as input_file:
#with open('input.txt') as input_file:
    raw_input = input_file.read()

crabs =  raw_input.strip().split(',')
crabs = [int(crab) for crab in crabs]
crabs.sort()

highest = crabs[-1]

moves = []

for high in crabs[::-1]:
    total = 0
    for crab in crabs:
        movement = high - int(crab)
        total += abs(movement)
    moves.append(total)

print(min(moves))
