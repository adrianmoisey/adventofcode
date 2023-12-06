from collections import deque

with open("input.txt") as file:
    input = file.read()

stack = deque(maxlen=14)

for pos, letter in enumerate(input.strip()):
    stack.append(letter)

    if len(set(stack)) == 14:
        print(pos + 1)
        break
