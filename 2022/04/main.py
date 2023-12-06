import string

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.readlines()

total = 0

for i in input:
    print("START")
    i = i.strip()
    print(f"{i=}")
    a, b = i.split(',')
    a1, a2 = a.split('-')
    b1, b2 = b.split('-')

    a_range = range(int(a1), int(a2) + 1)
    b_range = range(int(b1), int(b2) + 1)

    a_set = list(a_range)
    b_set = list(b_range)

    print(a_set)
    print(b_set)

    #a_string = 'c'.join([str(x) for x in a_set])
    #b_string = 'c'.join([str(x) for x in b_set])


    if all(item in a_set for item in b_set) or all(item in b_set for item in a_set):
        print('match')
        total += 1
    print("END")
    print()

print(total)
# 888 too high
# 622 too high
# 611 too high
# 311 not right
