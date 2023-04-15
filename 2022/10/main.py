from collections import deque

with open("input.txt") as file:
    input = file.read().strip()

stack = deque()
X = 1

for i in input.split('\n'):
    stack.append(i)

# The first thing we do in the loop is to increment this by one... effectively making the first cycle "2"
# This makes no sense to me, but it was the only way that my code would work.
# I can't figure it out.
# If cycle were to start at 0, the loop works in a way that makes sense to me, but doesn't give the correct answer
cycle = 1
total = []

while len(stack) != 0:
    cycle += 1
    instruction = stack.popleft()

    if instruction.startswith('addx'):
        _, value = instruction.split()
        stack.appendleft(value)

    elif instruction.startswith('noop'):
        pass

    else:
        X += int(instruction)

    if cycle == 20 or (cycle-20) % 40 == 0:
        total.append(X * cycle)

print(total)
print(sum(total))
