from collections import deque

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.read().strip()

stack = deque()
X = 1

for i in input.split('\n'):
    stack.append(i)

cycle = 1
total = []

horizontal = 0

# Look Jonathan, a list comprehension!
display = [['.'] * 40 for _ in range(6)]

while True:
    sprite = [X + (40 * horizontal), X+1 + (40 * horizontal), X+2 + (40 * horizontal)]

    if cycle in sprite:
        tmp_cycle = cycle - (40 * horizontal)
        display[horizontal][tmp_cycle-1] = "#"

    if len(stack) == 0:
        break

    if (cycle) % 40 == 0:
        horizontal += 1
    instruction = stack.popleft()
    cycle += 1

    if instruction.startswith('addx'):
        _, value = instruction.split()
        stack.appendleft(value)

    elif instruction.startswith('noop'):
        pass

    else:
        X += int(instruction)

for x in display:
    for y in x:
        print(y, end = '')
    print()
