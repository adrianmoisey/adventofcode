from collections import defaultdict, deque

with open("input-real.txt") as file:
    input = file.read().splitlines()

paths = defaultdict(list)

for i in input:
    left, right = i.split('-')
    if right != 'start':
        paths[left].append(right)
    if left != 'start':
        paths[right].append(left)

print(paths)

stack = deque([('start', {'start'}, False)])

total = 0
while stack:
    us, seen, double = stack.pop()

    if us == 'end':
        total += 1
        continue
    for node in paths[us]:
        if node not in seen or node.isupper():
            stack.append((node, seen | {node}, double))
            continue
        if double:
            continue

        stack.append((node, seen | {node}, True ))

print(total)
