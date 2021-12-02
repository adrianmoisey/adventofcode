def read_input():
    with open('input.txt') as input_file:
        input = input_file.read()

    return input.splitlines()

# Part 1
input = read_input()

index = 3
multiplier = 3
length = len(input[0]) - 1
tree = '#'
tree_count = 0

for i in input[1:]:
    if i[index] is tree:
        tree_count+=1
    index += multiplier
    if index > length:
        index = index - length-1

print(f'Part 1: {tree_count}')

# Part 2

input = read_input()

length = len(input[0]) - 1
tree = '#'
tree_count = 0

variants = [
    (1, 1),
    (3, 1),
    (5, 1),
    (7, 1),
    (1, 2),
]

totals = []

for h, v in variants:
    index = 0
    multiplier = h
    tree_count = 0

    for count, i in enumerate(input):
        if count == 0:
            continue
        if count % v == 0:
            if i[index] is tree:
                tree_count+=1
            index += multiplier
            if index > length:
                index = index - length-1

    totals.append(tree_count)

print(tree_count)
print(totals)
