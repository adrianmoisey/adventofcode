from collections import Counter
from statistics import median

with open('input-real.txt') as input_file:
#with open('input.txt') as input_file:
    raw_input = input_file.read()

crabs =  raw_input.strip().split(',')
crabs = [int(crab) for crab in crabs]
crabs.sort()

positions = list(range(1, crabs[-1] + 1, 1))

moves = []
for high in range(1, crabs[-1] + 1, 1):
    total = 0
    for crab in crabs:
        movement = high - crab
        movement = abs(movement)
        total += sum(list(range(1, movement + 1, 1)))
    moves.append(total)

print(moves)
print(min(moves))

# 86398084 too high
