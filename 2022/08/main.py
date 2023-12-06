from collections import OrderedDict

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.readlines()

visible = set()

width = len(input[0].strip())
height = len(input)

for x in range(width):
    # West -> East
    previous_height = -1
    for y in range(height):
        tree_height = int(input[x][y])
        if tree_height > previous_height:
            previous_height = tree_height
            visible.add((x, y))
        else:
            continue

    # East -> West
    previous_height = -1
    for y in range(height-1, 0-1, -1):
        tree_height = int(input[x][y])
        if tree_height > previous_height:
            previous_height = tree_height
            visible.add((x, y))
        else:
            continue

for y in range(height):
    # North -> South
    previous_height = -1
    for x in range(width):

        tree_height = int(input[x][y])
        if tree_height > previous_height:
            previous_height = tree_height
            visible.add((x, y))
        else:
            continue

    # South -> North
    previous_height = -1
    for x in range(width-1, 0-1, -1):
        tree_height = int(input[x][y])
        if tree_height > previous_height:
            previous_height = tree_height
            visible.add((x, y))
        else:
            continue


print(len(sorted(visible)))
