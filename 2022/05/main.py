with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.read()

stacks, instructions = input.split('\n\n')

flipped = list(zip(*reversed(stacks.split('\n'))))

stacks = {}

for i in flipped:
    if i[0] == ' ':
        continue

    boxes = []

    for z in i[1:]:
        if z == ' ':
            continue
        else:
            boxes.append(z)

    stacks[i[0]] = boxes

print(stacks)

for i in instructions.strip().split('\n'):
    i = i.split()
    count = i[1]
    source = i[3]
    destination = i[5]

    print(count, source, destination)

    for z in range(int(count)):
        grab = stacks[source].pop()
        stacks[destination].append(grab)

print(stacks)

for i in stacks.values():
    print(i[-1])
