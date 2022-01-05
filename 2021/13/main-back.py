with open("input-real.txt") as file:
    input = file.read().splitlines()


max_x = 0
max_y = 0
coords_list = []
folds = []

for line in input:
    if line.startswith('fold along'):
        _, _, i = line.split(' ')
        folds.append(i)
        continue
    if len(line) > 0:
        b, a = line.split(',')
        b, a = int(b), int(a)
        if a > max_x:
            max_x = a
        if b > max_y:
            max_y = b
        coords_list.append((a, b))

grid = ['.'] * (max_x+1)
for i in range(max_x+1):
    grid[i] = ['.'] * (max_y+1)

for i in coords_list:
    x, y = i
    grid[x][y] = '#'

for fold in folds:
    if fold.startswith('y'):
        print('heer')
        _, pos = fold.split('=')
        pos = int(pos)

        upper = grid[0:pos]
        lower = grid[pos+1:]

        lower_rev = lower[::-1]

        for x in range(len(upper)):
            for y in range(len(upper[0])):
                    if lower_rev[x][y] == "#":
                        upper[x][y] = "#"
        grid = upper[:]

    if fold.startswith('x'):
        print('hexer')
        _, pos = fold.split('=')
        pos = int(pos)

        left = []
        right = []
        right_rev = []

        for line in grid:
            left.append(line[0:pos])
            right.append(line[pos+1:])

        for line in right:
            right_rev.append(line[::-1])


        for x in range(len(left)):
            for y in range(len(left[0])):
                    if right_rev[x][y] == "#":
                        left[x][y] = "#"

        grid = left[:]



for i in left:
    print(i)
print()
for i in right:
    print(i)
print()
for i in right_rev:
    print(i)

count = 0
for i in grid:
    for c in i:
        if c == '#':
            count+=1
print(count)
