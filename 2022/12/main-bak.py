import string

import sys

letters = string.ascii_lowercase

#with open("input.txt") as file:
with open("sample.txt") as file:
    input = file.read().strip().split('\n')

height = len(input) - 1
width = len(input[0]) - 1

grid = {}

# X is up or down
# Y is left or right

for x, line in enumerate(input):
    for y, character in enumerate(line):
        if character == "S":
            start = (x, y)

        if character == "E":
            end = (x, y)

        grid[(x, y)] = [character, sys.maxsize]

unvisited_set = set(grid.keys())


def current_node(position):
    x, y = position
    neighbors = [(x+1, y), (x, y+1), (x-1, y), (x, y-1)]

    if grid[position][0] == "S":
        our_index = letters.index('a')
    elif grid[position][0] == "E":
        our_index = letters.index('z')
    else:
        our_index = letters.index(grid[position][0])

    for n in neighbors:
        if grid.get(n, None) != None:

            if grid[n][0] == "E":
                their_index = letters.index('z')
            elif grid[n][0] == "S":
                their_index = letters.index('a')
            else:
                their_index = letters.index(grid[n][0])

            if our_index+1 == their_index or their_index <= our_index:
                # our position + 1
                distance = grid[position][1] + 1
                # if they have a value already, and ours is smaller
                if grid[n][1] > distance:
                    grid[n][1] = distance


    unvisited_set.discard(position)

grid[start][1] = 0

while grid[end][1] == sys.maxsize:
    smallest = sys.maxsize
    smallest_tuple = None

    for node in tuple(unvisited_set):
        if grid[node][1] < smallest:
            smallest = grid[node][1]
            smallest_tuple = node

    current_node(smallest_tuple)


#print(start, end)
#print(grid)
print(grid[end])

# 186 too low
