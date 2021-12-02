def read_input():
    with open('input.txt') as input_file:
        input = input_file.read()

    return input.splitlines()

# Part 1
h = 0
d = 0
a = 0
input = read_input()

for i in input:
    print(h, d, a)
    direction, distance = i.split()
    if direction == 'forward':
        h += int(distance)
        d += a * int(distance)

    if direction == 'down':
        a += int(distance)

    if direction == 'up':
        a -= int(distance)

print(h, d, a)
print('Part 1 ', h * d)

