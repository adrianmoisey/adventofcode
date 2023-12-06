import string

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.readlines()

index = string.ascii_lowercase + string.ascii_uppercase
total = 0

for i in input:
    print(i)
    length = len(i.strip())
    half = length/2
    half = int(half)
    first = set(i[:half])
    second = set(i[half:-1])

    common = first & second
    common = common.pop()
    print(common)
    value = index.index(common) + 1
    total = total + value

print(total)
