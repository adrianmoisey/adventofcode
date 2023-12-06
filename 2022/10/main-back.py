from collections import deque

#with open("input.txt") as file:
with open("sample.txt") as file:
    input = file.read().strip()

stack = deque()
X = 1

cycle = 0

for index, i in enumerate(input.split('\n')):
    if cycle == 20 or (cycle-20) % 40 == 0:
        print(f"Cycle: {cycle}")
        print(X)
    cycle += 1

    if i.startswith('addx'):
        _, value = i.split()
        stack.append(int(value))
    else:
        continue

    if len(stack) == 1:
        if cycle == 20 or (cycle-20) % 40 == 0:
            print(f"Cycle: {cycle}")
            print(X)
        value = int(stack.popleft())
        X += value
        cycle += 1

