from collections import defaultdict

answer = 2000000
#answer = 10

def make_diamond(s, b):
    #print('make_diamond called')
    positions = []
    s_x, s_y = s
    b_x, b_y = b

    b_x = abs(b_x - s_x)
    b_y = abs(b_y - s_y)
    size = b_x + b_y

    x = 0
    y = 1
    #print('---')
    #print(s_x, s_y)

    size = size + 1
    for i in range(size):
        y_size = size - i
        if s_y + i == answer or s_y - i == answer:
            for j in range(y_size):
                positions.append((s_x + j, s_y + i))
                positions.append((s_x - j, s_y + i))
                positions.append((s_x - j, s_y - i))
                positions.append((s_x + j, s_y - i))

    return  positions

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.read().strip().split('\n')

grid = {}

smallest_x = 0
smallest_y = 0
biggest_x = 0
biggest_y = 0

for i in input:
    i = i.split()

    sensor_x, sensor_y = i[2], i[3]
    beacon_x, beacon_y = i[8], i[9]

    sensor_x = int(sensor_x.split('=')[1].strip(','))
    sensor_y = int(sensor_y.split('=')[1].strip(':'))
    beacon_x = int(beacon_x.split('=')[1].strip(','))
    beacon_y = int(beacon_y.split('=')[1])

    smallest_x = min(smallest_x, sensor_x, beacon_x)
    smallest_y = min(smallest_y, sensor_y, beacon_y)

    biggest_x = max(biggest_x, sensor_x, beacon_x)
    biggest_y = max(biggest_y, sensor_y, beacon_y)

    for i in make_diamond((sensor_x, sensor_y), (beacon_x, beacon_y)):
        #print('make_diamond ended')
        if i[1] == answer:
            #print(i)
            if grid.get(i) != "B" or grid.get(i) != "S":
                grid[i] = '#'

    grid[sensor_x, sensor_y] = 'S'
    grid[beacon_x, beacon_y] = 'B'

#print(smallest_x, smallest_y)
#print(biggest_x, biggest_y)
#print(grid)

#for y in range(smallest_y, biggest_y+1):
#    for x in range(smallest_x, biggest_x+1):
#        print(grid.get((x, y), '.'), end='')
#    print()




print("----")
#a = make_diamond((8, 7), (2, 10))
#for i in a:
#    grid[i] = '#'
#

grid[(0,0)] = 'X'

#for y in range(smallest_y-10, biggest_y+1+10):
#    for x in range(smallest_x-10, biggest_x+1+10):
#        print(grid.get((x, y), '.'), end='')
#    print()

count = 0
a = []

for i in grid:
    if i[1] == answer:
        if grid[i] == '#':
            a.append(i)
            count += 1

#print(sorted(a))
print(count)
