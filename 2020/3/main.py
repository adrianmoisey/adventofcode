from functools import reduce

def read_input():
    with open('input.txt') as input_file:
        input = input_file.read()

    return input.splitlines()

input = read_input()

length = len(input[0])
height = len(input)
tree = '#'
tree_count = 0

variants = [
    (1, 1),
    (3, 1),
    (5, 1),
    (7, 1),
    (1, 2),  # ?
]

total = []

for right, down in variants:
    tree_count = 0
    row, column = 0, 0
    while True:
        row += right
        column += down
        if row > (length - 1):
            row = row - length
        if column >= height:
            break
        if input[column][row] is tree:
            tree_count += 1

    print(tree_count)
    total.append(tree_count)

print(reduce(lambda x, y: x*y, total))
