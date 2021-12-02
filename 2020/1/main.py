with open('input.txt') as input_file:
    input = input_file.read()

input = input.split()

for index, value1 in enumerate(input):
    for value2 in input[index:]:
        if int(value1) + int(value2) == 2020:
            print("Part 1:")
            print(int(value1) * int(value2))
        for value3 in input[index:]:
            if int(value1) + int(value2) + int(value3) == 2020:
                print("Part 2:")
                print(int(value1) * int(value2) * int(value3))

