from collections import defaultdict

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.read().strip().split('\n')

grid = defaultdict()

max_x = 0
max_y = 0

min_x = 500
min_y = 500

for line in input:
    split_line = line.split(' -> ')
    for index, _ in enumerate(split_line):
        if index+1 == len(split_line):
            break
        current_point = split_line[index].split(',')
        next_point = split_line[index+1].split(',')

        current_point = (int(current_point[0]) ,int(current_point[1]))
        next_point = (int(next_point[0]) ,int(next_point[1]))

        # Find our bounderies
        max_x = max(max_x, current_point[0])
        max_y = max(max_y, current_point[1])
        min_x = min(min_x, current_point[0])
        min_y = min(min_y, current_point[1])

        max_x = max(max_x, next_point[0])
        max_y = max(max_y, next_point[1])
        min_x = min(min_x, next_point[0])
        min_y = min(min_y, next_point[1])

        if current_point[0] == next_point[0]:
            # yuck
            for i in range(min(current_point[1], next_point[1]), max(current_point[1], next_point[1])+1):
                grid[(current_point[0], i)] = '#'

        if current_point[1] ==  next_point[1]:
            for i in range(min(current_point[0], next_point[0]), max(current_point[0], next_point[0])+1):
                grid[(i, current_point[1])] = '#'

#grid[(500,0)] = 'o'
print(grid)

if min_y > 0:
    min_y = 0

print(min_x, max_x, min_y, max_y)

def print_grid():
    for y in range(min_y, max_y+1):
        for x in range(min_x, max_x+1):
            print(grid.get((x, y), '.'), end ='')
        print()





bail = False
count = 0

floor = max_y+2

while bail == False:
#for _ in range(0,100):
    count += 1
    sand_x, sand_y = (500,0)

    if grid.get((500, 0), '.') == 'o':
        bail = True
        print(count-1)
        break

    while True:
        if sand_y+1 == floor:
            grid[(sand_x, sand_y)] = 'o'
            break

        elif grid.get((sand_x, sand_y+1), '.') == '.':
            sand_y += 1
        else:
            if grid.get((sand_x-1, sand_y+1), '.') != '.':
                if grid.get((sand_x+1, sand_y+1), '.') != '.':
                    grid[(sand_x, sand_y)] = 'o'
                    break
                else:
                    sand_y += 1
                    sand_x += 1
            else:
                sand_y += 1
                sand_x -= 1

        if grid.get((500, 0), '.') == 'o':
            bail = True
            print(count)
            break


for y in range(min_y-50, max_y+2):
    for x in range(min_x-154, max_x+140):
        character = grid.get((x, y), '.')
        if character == 'o':
            print('🟨', end='')
        elif character == '#':
            print('🟫', end='')
        elif character == '.':
            print('⬛️', end='')
        else:
            print(character, end='')
    print()
