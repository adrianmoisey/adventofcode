from collections import OrderedDict

#with open("input.txt") as file:
with open("sample.txt") as file:
    input = file.readlines()

visible = set()
all_trees = {}

width = len(input[0].strip())
height = len(input)

for height in range(len(input)):
    for width in range(len(input[0]) - 1):
        visible.add((0, height))
        visible.add((width , 0))

        visible.add((len(input) - 1, height))
        visible.add((width , len(input[0].strip()) - 1))

print(sorted(visible))
print(len(visible))

for height in range(len(input)):
    for width in range(len(input[0]) - 1):
        height = int(height)
        width = int(width)
        all_trees[(height, width)] = input[height][width]

print(all_trees)

# East West
for x in range(0, height + 1):
    row = OrderedDict()
    counter = []
    for y in range(0, width + 1):
        #print(all_trees[(x, y)])
        row[(x, y)] = all_trees[(x, y)]

    tree_previous = 0
    tree_index = 0
    for i, tree in enumerate(row.values()):
        tree = int(tree)
        if tree > tree_previous:
            tree_previous = tree
            tree_index = i
        else:
            break
    for index, tree in enumerate(row):
        if index > tree_index:
            break
        else:
            visible.add(tree)

    row = OrderedDict(sorted(row.items(), reverse=True))

    tree_previous = 0
    tree_index = 0
    for i, tree in enumerate(row.values()):
        tree = int(tree)
        if tree > tree_previous:
            tree_previous = tree
            tree_index = i
        else:
            break
    print("row")
    print(row)
    for index, tree in enumerate(row):
        if index > tree_index:
            break
        else:
            print(tree)
            visible.add(tree)

# North South
for x in range(0, height + 1):
    row = OrderedDict()
    counter = []
    for y in range(0, width + 1):
        row[(x, y)] = all_trees[(x, y)]

    print(row)

    tree_previous = 0
    tree_index = 0
    for i, tree in enumerate(row.values()):
        tree = int(tree)
        if tree > tree_previous:
            tree_previous = tree
            tree_index = i
        else:
            break
    for index, tree in enumerate(row):
        if index > tree_index:
            break
        else:
            visible.add(tree)

    row = OrderedDict(sorted(row.items(), reverse=True))

    tree_previous = 0
    tree_index = 0
    for i, tree in enumerate(row.values()):
        tree = int(tree)
        if tree > tree_previous:
            tree_previous = tree
            tree_index = i
        else:
            break
    for index, tree in enumerate(row):
        if index > tree_index:
            break
        else:
            visible.add(tree)
print()
print(sorted(visible))
print(len(visible))
