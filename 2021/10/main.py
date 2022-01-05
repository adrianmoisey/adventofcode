with open("input-real.txt") as file:
    input = file.read().split()

paren = {
    '(': ')',
    '[': ']',
    '{': '}',
    '<': '>',
}

total = []

for line in input:
    points = 0
    next_expected = []
    corrupted = False
    for c in line:
        if c in paren.keys():
            next_expected.append(paren[c])
        else:
            if c == next_expected.pop():
                continue
            else:
                corrupted = True
                break
    if corrupted == False:
        for e in next_expected[::-1]:
            points = points * 5
            if e == ')':
                points += 1
            if e == ']':
                points += 2
            if e == '}':
                points += 3
            if e == '>':
                points += 4
        print(points)
        total.append(points)

total.sort()
print('part 2')
print(total[int(len(total)/2)])

