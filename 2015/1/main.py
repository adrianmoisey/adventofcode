with open('input.txt') as input_file:
    input = input_file.read()

floor = 0
count = 0
for i in input:
    count += 1
    if i == '(':
        floor += 1
    if i == ')':
        floor -= 1

    if floor == -1:
        print(count)
        break



