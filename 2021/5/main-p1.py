from dataclasses import dataclass

#with open('input-real.txt') as input_file:
with open('input.txt') as input_file:
    raw_input = input_file.read()

coordinates = raw_input.splitlines()

#An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
#An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
# x1,y1 -> x2,y2

@dataclass
class Coordinates:
    x1: int
    y1: int
    x2: int
    y2: int

data_co = []

for c in coordinates:
    x, _, y = c.split()
    C = Coordinates(
            int(x.split(',')[0]),
            int(x.split(',')[1]),
            int(y.split(',')[0]),
            int(y.split(',')[1]),
    )
    data_co.append(C)


max_x = 0
max_y = 0
for c in data_co:
    if c.x1 > max_x:
        max_x = c.x1
    if c.x2 > max_x:
        max_x = c.x2

    if c.y1 > max_y:
        max_x = c.x1
    if c.y2 > max_y:
        max_y = c.y2

max_x += 1
max_y += 1

grid = [0] * max_x
for i in range(max_x):
    grid[i] = [0] * max_y

for c in data_co:
    if c.x1 == c.x2:
        small = min(c.y1, c.y2)
        big = max(c.y1, c.y2)
        for p in range(small, big+1):
            grid[c.x1][p] += 1

    if c.y1 == c.y2:
        small = min(c.x1, c.x2)
        big = max(c.x1, c.x2)
        for p in range(small, big+1):
            grid[p][c.y1] += 1

for line in grid:
    print(line)

count = 0
for line in grid:
    for item in line:
        if item > 1:
            count += 1

print(count)
