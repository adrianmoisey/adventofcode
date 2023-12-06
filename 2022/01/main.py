with open("input.txt") as file:
    input = file.readlines()

count = 0
elf = {}

for i in input:
    i = i.strip()

    if i == '\n' or i == '':
        count += 1
        continue

    current = elf.get(count, 0)
    current += int(i)
    elf[count] = current

print("Part 1")
print(sorted(elf.values())[-1])


print("Part 2")
print(sum(sorted(elf.values())[-3:]))

