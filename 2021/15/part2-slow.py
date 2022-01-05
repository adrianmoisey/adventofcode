import sys

with open("input-real.txt") as input:
    positions_s = input.read()

#print("This is the grid")
#print(positions_s)
#print("This is the grid </end>")

grid = {}

max_x = len(positions_s.splitlines() * 5) - 1
max_y = len(positions_s.splitlines()[0] * 5) - 1

short_x = len(positions_s.splitlines())
short_y = len(positions_s.splitlines()[0])

for x_m in range(0,5):
     for y_m in range(0,5):
         for x, line in enumerate(positions_s.splitlines()):
             for y, c in enumerate(line):
                               # (value, tentative distance value)
                 multiplier = x_m + y_m
                 value = int(c) + multiplier
                 value = value % 10
                 grid[(x + (x_m * short_x), y + (y_m  * short_y))] = (value, sys.maxsize)

#print(grid)
print("Max X and Y")
print(max_x, max_y)
grid[(0,0)] = (1,0)
#unvisited_set = set(grid.keys())
unvisited_set = grid.copy()

def current_node(position):
    x, y = position
    neighbors = [(x+1, y), (x, y+1), (x-1, y), (x-1, y)]

    for n in neighbors:
        if grid.get(n, None) != None:
            # our tentative distance + their size
            distance = grid[position][1] + grid[n][0]
            # Get smaller of us+them, or them
            smaller = min([distance, grid[n][1]])
            grid[n] = (grid[n][0], smaller)

    del unvisited_set[position]

while grid[(max_x, max_y)][1] == sys.maxsize:
    smallest = sys.maxsize
    smallest_tuple = None
    print(len(unvisited_set))
    node = sorted(unvisited_set, key=unvisited_set.get)[0]
    smallest = grid[node][1]
    smallest_tuple = node
    current_node(smallest_tuple)


print(grid)
print(grid[max_x, max_y])

