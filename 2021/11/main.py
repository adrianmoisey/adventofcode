from collections import defaultdict
import sys

with open("input-real.txt") as file:
    input = file.readlines()

octopuses = defaultdict(lambda: sys.maxsize)

for x, line in enumerate(input):
    for y, c in enumerate(line.strip()):
        octopuses[(x, y)] = int(c)

for a in range(0,11):
    for b in range(0,11):
        if octopuses[(a, b)] < sys.maxsize:
            print(octopuses[(a, b)], end='')
    print()

def contains_nine(octs):
    for value in octs.values():
        if value > 9 and value < sys.maxsize:
            return False
    return True

def contains_zero(octs):
    for value in octs.values():
        if value != 0 and value < sys.maxsize:
            return False
    return True

count = 0

for i in range(0,99999):
    # Increase by 1
    for pos in octopuses:
        octopuses[pos] += 1
    flashed = set()
    while contains_nine(octopuses) == False:
        for (x, y), c in tuple(octopuses.items()):
            if octopuses[(x, y)] >= 10 and (x, y) not in flashed and octopuses[(x, y)] < sys.maxsize:
                flashed.add((x, y))
                count += 1
                octopuses[(x + 1, y + 1)] += 1
                octopuses[(x + 1, y)] += 1
                octopuses[(x + 1, y - 1)] += 1
                octopuses[(x, y + 1)] += 1
                octopuses[(x, y - 1)] += 1
                octopuses[(x - 1, y + 1)] += 1
                octopuses[(x - 1, y)] += 1
                octopuses[(x - 1, y - 1)] += 1
        for (x, y) in flashed:
            octopuses[(x, y)] = 0
    if contains_zero(octopuses) is True:
        print('hello')
        print(i+1)
        break

for a in range(0,11):
    for b in range(0,11):
        if octopuses[(a, b)] < sys.maxsize:
            print(octopuses[(a, b)], end='')
    print()

print(count)
