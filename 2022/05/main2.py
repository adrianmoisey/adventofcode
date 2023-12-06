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
    print(i)
    i = i.split()
    count = i[1]
    source = i[3]
    destination = i[5]

    pickup = []

    for z in range(int(count)):
        grab = stacks[source].pop()
        pickup.append(grab[0])
        print(pickup)

    stacks[destination] = stacks[destination] + pickup[::-1]
    print(stacks)

print(stacks)

for i in stacks.values():
    print(i[-1])
