from collections import defaultdict

limit = 4000000; filename = 'input.txt'
#limit = 20; filename = 'sample.txt'
#limit = 20; filename = 'sample.txt.1'


def make_diamond(s, b, row):
    #print('make_diamond called')
    positions = set()
    s_x, s_y = s
    b_x, b_y = b

    b_x = abs(b_x - s_x)
    b_y = abs(b_y - s_y)
    size = b_x + b_y
    og_size = size

    x = 0
    y = 1

    size = size + 1
    #if row == 11:
    #    print(f'Our triangle is size: {size}, s_x: {s_x}, s_y: {s_y}, row: {row}')
    # answer is 11

    highest_point = s_y + size
    lowest_point = s_y - size
    #print(f"{s_y=}, {row=}")
    if lowest_point <= row <= highest_point:
        #print('were in the cube')

        if row == s_y:
            size = size
        elif row > s_y:
            difference = abs(highest_point - row)
            size = difference
        else:
            difference = abs(row - lowest_point)
            size = difference

        left = s_x - size+1
        right = s_x + size-1

        if left < 0:
            left = 0
        if right > limit:
            right = limit
        if right < 0:
            right = 0
        if left > limit:
            left = limit
        if left > right:
            return
        #print(f"{left=} {right=}: {s_x=} {s_y=}, {og_size=}")
        return left, right

with open(filename) as file:
    input = file.read().strip().split('\n')

smallest_x = 0
smallest_y = 0
biggest_x = 0
biggest_y = 0

for row in range(0, limit):
    print(row, ' ---------')
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

        leftright = make_diamond((sensor_x, sensor_y), (beacon_x, beacon_y), row)

        if leftright != None:
            grid.add(leftright)

    print(grid)

    #count = 0

    a = list(zip(*grid))

    lefts = sorted(a[0])
    rights = sorted(a[1])

    length = [lefts[0], rights[0]]

    for a, b in zip(lefts, rights):
        if a <= length[1] or a == length[1] + 1:
            length[1] = b

    #print(f"{length=}")

    if length[0] == 0 and length[1] != limit:
        print('we got our row!')
        print(row)
        print(length[1] + 1)
        x = length[1] + 1
        y = row
        print((x * 4000000) + y)
        break
