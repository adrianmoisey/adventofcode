from collections import OrderedDict

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.read()

input = input.strip().split('\n')

visible = set()

width = len(input[0].strip())
height = len(input)

my_trees = []

for w in range(width):
    for h in range(height):
        my_trees.append((w, h))

totals = []

for my_tree in my_trees:
    w, h = my_tree
    my_height = input[w][h]

    left = []
    right = []
    up = []
    down = []

    # Left
    for i in range(w, 0-1 ,-1):
        if i == w:
            continue
        if input[i][h] < my_height:
            left += input[i][h]
        else:
            left += input[i][h]
            break

    # Right
    for i in range(w, width):
        if i == w:
            continue
        #w = i
        if input[i][h] < my_height:
            right += input[i][h]
        else:
            right += input[i][h]
            break

    # Up
    for i in range(h, 0-1 ,-1):
        if i == h:
            continue
        # y = i
        if input[w][i] < my_height:
            up += input[w][i]
        else:
            up += input[w][i]
            break

    # Down
    for i in range(h, height):
        if i == h:
            continue
        # y = i
        if input[w][i] < my_height:
            down += input[w][i]
        else:
            down += input[w][i]
            break
    total = len(left) * len(right) * len(up) * len(down)
    totals.append(total)

print(max(totals))
