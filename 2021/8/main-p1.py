with open('input.txt') as input_file:
    INPUT_S = input_file.read()

segments = {
    0: 6,
    1: 2,
    2: 5,
    3: 5,
    4: 4,
    5: 5,
    6: 6,
    7: 3,
    8: 7,
    9: 6,
}

input_lines = INPUT_S.strip().split('\n')

count = 0
for line in input_lines:
    left, right = line.split('|')

    right_l = right.split()
    for digit in right_l:
        if len(digit) in [segments[1], segments[4], segments[7], segments[8]]:
            count += 1

print(count)






