with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.read()

stack = []

for pos, letter in enumerate(input.strip()):
    stack += letter

    if len(stack) == 15:
        stack = stack[1:]

    if len(stack) == 14:
        if len(set(stack)) == 14:
            print(pos + 1)
            break
