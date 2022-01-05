import sys

with open("input.txt") as input:
    positions_s = input.read()

print("This is the grid")
print(positions_s)
print("This is the grid </end>")

grid = {}

max_x_orig = len(positions_s.splitlines() * 5) - 1
max_y_orig = len(positions_s.splitlines()[0] * 5) - 1

max_x = len(positions_s.splitlines() * 5) - 1
max_y = len(positions_s.splitlines()[0] * 5) - 1



for x, line in enumerate(positions_s.splitlines()):
    for y, c in enumerate(line):
                      # (value, tentative distance value)
        grid[(x, y)] = (int(c), sys.maxsize)

#print(grid)
print("Max X and Y")
print(max_x, max_y)
unvisited_set = set()
for x in range(0, max_x+1):
    for y in range(0, max_y+1):
        unvisited_set.add((x, y))
print(unvisited_set)
grid[(0,0)] = (1,0)

def current_node(position):
    x, y = position
        # calculate multplier
    if y >= 0 and y <= 9:
        y_multiplier = 1
        y_subtracion = 0
    elif y > 9 and y <= 19:
        y_multiplier = 2
        y_subtracion = 10
    elif y > 19 and y <= 29:
        y_multiplier = 3
        y_subtracion = 20
    elif y > 29 and y <= 39:
        y_multiplier = 4
        y_subtracion = 30
    elif y > 39 and y <= 49:
        y_multiplier = 5
        y_subtracion = 40

    if x >= 0 and x <= 9:
        x_multiplier = 1
        x_subtracion = 0
    elif x > 9 and x <= 19:
        x_multiplier = 2
        x_subtracion = 10
    elif x > 19 and x <= 29:
        x_multiplier = 3
        x_subtracion = 20
    elif x > 29 and x <= 39:
        x_multiplier = 4
        x_subtracion = 30
    elif x > 39 and x <= 49:
        x_multiplier = 5
        x_subtracion = 40

    multiplier = x_multiplier + y_multiplier
    position = (x-y_subtracion, y-x_subtracion)
    x,y = position
    neighbors = [(x+1, y), (x, y+1), (x-1, y), (x-1, y)]

    for n in neighbors:
        multiplier = x_multiplier + y_multiplier
        position = (x-y_subtracion, y-x_subtracion)

        if grid.get(n, None) != None:
            # our tentative distance + their size
            distance = (grid[position][1]*multiplier) + (grid[n][0]*multiplier)
            # Get smaller of us+them, or them
            smaller = min([distance, grid[n][1]*multiplier])
            grid[n] = (grid[n][0], smaller)

    unvisited_set.discard(position)

while grid.get((max_x, max_y)[1], sys.maxsize) == sys.maxsize:
    smallest = sys.maxsize
    smallest_tuple = None
    for node in tuple(unvisited_set):
        if grid[node][1] < smallest:
            smallest = grid[node][1]
            smallest_tuple = node

    current_node(smallest_tuple)


print(grid)
print(grid[max_x, max_y])

