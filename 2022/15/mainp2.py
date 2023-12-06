from collections import defaultdict

limit = 4000000; filename = 'input.txt'
#limit = 20; filename = 'sample.txt'


def make_diamond(s, b, row):
    #print('make_diamond called')
    positions = set()
    s_x, s_y = s
    b_x, b_y = b

    b_x = abs(b_x - s_x)
    b_y = abs(b_y - s_y)
    size = b_x + b_y

    x = 0
    y = 1
    #print('---')
    #print(s_x, s_y)
    #print(size)

    size = size + 1
    if row == 11:
        print(f'Our triangle is size: {size}, s_x: {s_x}, s_y: {s_y}, row: {row}')
    # answer is 11

    #if (0 <= s_x + size <= limit) or \
    #   (0 <= s_x - size <= limit) or \
    #   (0 <= s_y + size <= limit) or \
    #   (0 <= s_y - size <= limit):

    #   #if (row >= s_y + size) or \
    #   #   (row <= s_y - size):

    #   for i in range(size):
    #       print('inside this loop')
    #       y_size = size - i
    #       if s_y + i == row or s_y - i == row:
    #           for j in range(y_size):
    #               if 0 <= s_x + j <= limit:
    #                   print('added a dot')
    #                   positions.add(s_x + j)
    #               if 0 <= s_x - j <= limit:
    #                   print('added a dot')
    #                   positions.add(s_x - j)
    #   #print('make_diamond finished')
    return  positions

with open(filename) as file:
    input = file.read().strip().split('\n')

smallest_x = 0
smallest_y = 0
biggest_x = 0
biggest_y = 0

for row in range(0, limit):
    #print(row, ' ---------')
    grid = set()
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

        for i in make_diamond((sensor_x, sensor_y), (beacon_x, beacon_y), row):
            grid.add(i)

    #if filename == 'sample.txt':
    #    for y in range(smallest_y-10, biggest_y+1+10):
    #        for x in range(smallest_x-10, biggest_x+1+10):
    #            print(grid.get((x, y), '.'), end='')
    #        print()

    count = 0
    #print(len(grid))

    if len(grid) == limit:
        print('we got our row!')
        print(row)
        for i,num in enumerate(grid):
            if i != num:
                print(num-1, row)
                print(((num-1) * 4000000) + row)
                break

        break
