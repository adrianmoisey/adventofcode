def read_input():
    with open('input-sample.txt') as input_file:
        input = input_file.read()

    return input.splitlines()


input = read_input()
input2 = read_input()

flip = list(zip(*reversed(input)))

gamma = ''
epsilon = ''

for i in flip:
    zero = i.count('0')
    one = i.count('1')

    if zero > one:
        gamma += '0'
        epsilon += '1'

    elif one > zero:
        gamma += '1'
        epsilon += '0'

#print(gamma, epsilon)
#print(int(gamma, 2) * int(epsilon, 2))

# Part 2
oxygen = ''
co2 = ''

current_bit = 0

def get_all_bits(input, bit):
    all_bits = []
    for line in input:
        all_bits.append(line[bit])
    return(all_bits)

answer = []

for bit in range(0, len(input[0])):
    print(input)
    all_bits = get_all_bits(input, bit)
    count_zero = all_bits.count('0')
    count_one = all_bits.count('1')
    if count_one > count_zero:
        oxygen += '1'
        input = [x for x in input if x.startswith(oxygen) ]
    elif count_zero > count_one:
        oxygen += '0'
        input = [x for x in input if x.startswith(oxygen) ]
    elif count_zero == count_one:
        oxygen += '1'
        input = [x for x in input if x.startswith(oxygen) ]
    if len(input) == 1:
        answer.append(input[0])

print('===')

for bit in range(0, len(input2[0])):
    if len(input2) == 1:
        answer.append(input2[0])
    print(input2)
    all_bits = get_all_bits(input2, bit)
    count_zero = all_bits.count('0')
    count_one = all_bits.count('1')
    if count_one > count_zero:
        co2 += '0'
        input2 = [x for x in input2 if x.startswith(co2) ]
    if count_zero > count_one:
        co2 += '1'
        input2 = [x for x in input2 if x.startswith(co2) ]
    if count_zero == count_one:
        co2 += '0'
        input2 = [x for x in input2 if x.startswith(co2) ]

print(answer)

print(int(answer[0], 2) * int(answer[1], 2))
