with open("input-real.txt") as file:
    input = file.read().split()

paren = {
    '(': ')',
    '[': ']',
    '{': '}',
    '<': '>',
}
points = 0

for line in input:
    next_expected = []
    print(line)
    for c in line:
        if c in paren.keys():
            next_expected.append(paren[c])
        else:
            if c == next_expected.pop():
                continue
            else:
                print('corrupted!')
                print(c)
                if c == ')':
                    points += 3
                if c == ']':
                    points += 57
                if c == '}':
                    points += 1197
                if c == '>':
                    points += 25137
print(points)
