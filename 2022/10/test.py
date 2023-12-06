X, part1, part2 = 1, 0, '\n'
for cycle, value in enumerate(open('sample.txt').read().split(), 1):
    part1 += cycle * X  if cycle%40==20              else 0
    part2 += '#'        if abs((cycle-1)%40 - X) < 2 else ' '
    X     += int(value) if value[-1].isdigit()       else 0
print(part1, *part2)

