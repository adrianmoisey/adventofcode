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


# (current, ([seen], visited))

array = [('start', ([], False))]
path = 0

while array:
    print(array)
    current, sv = array.pop()
    for next in paths[current]:
        seen, visited = sv
        if next.islower() and next in seen and visited == True:
            continue
        if next == 'end':
            print("END", ",".join(seen), f",{current}, {next}")
            path += 1
            continue
        if next.islower() and next in seen:
            visited = True
        v = seen[:] + [current]
        t = (next, (v, visited))
        array.append(t)

print(path)
