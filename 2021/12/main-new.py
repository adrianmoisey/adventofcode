from collections import defaultdict

with open("input.txt") as file:
    input = file.read().splitlines()

paths = defaultdict(list)

for i in input:
    left, right = i.split('-')
    if right != 'start':
        paths[left].append(right)
    if left != 'start':
        paths[right].append(left)
print(paths)


# (current, [seen])

array = [('start', [])]
path = 0

while array:
    current, seen = array.pop()
    for next in paths[current]:
        if current.islower() and current in seen:
            continue
        if next == 'end':
            path += 1
            continue

        v = seen[:] + [current]
        t = (next, v)
        array.append(t)

print(path)
