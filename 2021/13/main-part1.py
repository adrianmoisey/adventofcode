from collections import defaultdict

#with open("input.txt") as file:
with open("input-real.txt") as file:
    input = file.read().splitlines()


max_x = 0
max_y = 0
coords = defaultdict(list)
folds = []

for line in input:
    if line.startswith('fold along'):
        _, _, i = line.split(' ')
        folds.append(i)
        continue
    if len(line) > 0:
        a, b = line.split(',')
        a, b = int(a), int(b)
        if a > max_x:
            max_x = a
        if b > max_y:
            max_y = b
        coords[(a, b)] = '#'
print("MAX:", max_x, max_y)
#print(coords)
print(folds)
folds = [folds[0]]
print(folds)
print("MaxX:", max_x)
print("MaxY:", max_y)
max_x = max_x+1
max_y = max_y+1

for fold in folds:
    if fold.startswith('y'):
        print('HERE: Y')
        _, fold_pos = fold.split('=')
        fold_pos = int(fold_pos)
        print(fold_pos)

        for pos in tuple(coords):
            if coords[pos] == '#':
                x, y = pos
                if y >= fold_pos:
                    coords[(x, y)] = '.'
                    y = max_y - y
                    coords[(x,y)] = '#'


    if fold.startswith('x'):
        print('HERE: X')
        _, fold_pos = fold.split('=')
        fold_pos = int(fold_pos)
        print(fold_pos)

        for pos in tuple(coords):
            if coords[pos] == '#':
                x, y = pos
                if x >= fold_pos:
                    coords[(x, y)] = '.'
                    x = max_x - x
                    print(x,y)
                    coords[(x, y)] = '#'

count = 0
for i in coords:
    if coords[i] == '#':
        count += 1
print(count)

# 865 is too high
# 865 too high
# 650 is too low
