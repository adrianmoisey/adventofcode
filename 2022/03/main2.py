import string

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.readlines()

index = string.ascii_lowercase + string.ascii_uppercase
total = 0

grouping = []

for i in input:
    grouping.append(i)

    if len(grouping) == 3:
        first = set(grouping[0].strip())
        second = set(grouping[1].strip())
        third = set(grouping[2].strip())

        common = first & second & third
        common = common.pop()
        value = index.index(common) + 1
        total = total + value
        grouping = []

print(total)
