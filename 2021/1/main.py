def read_input():
    with open('input.txt') as input_file:
        input = input_file.read()

    return input.splitlines()

# Part 1
count = 0
input = read_input()
previous = 0

for i in input:
    if int(i) > previous:
        count += 1
        previous = int(i)
    else:
        previous = int(i)

print('Part 1 ', count-1)

# Part 2

new_array = []
for i in range(0, len(input) - 2):

    sum = int(input[i]) + int(input[i+1]) + int(input[i+2])
    new_array.append(sum)

count = 0
previous = 0

for i in new_array:
    if int(i) > previous:
        count += 1
        previous = int(i)
    else:
        previous = int(i)

print('Part 2', count-1)

