import sys

with open("input-real.txt") as input:
    positions_s = input.read()

print("This is the grid")
print(positions_s)
print("This is the grid </end>")

grid = {}

max_x = len(positions_s.splitlines()) - 1
max_y = len(positions_s.splitlines()[0]) - 1


for x, line in enumerate(positions_s.splitlines()):
    for y, c in enumerate(line):
                      # (value, tentative distance value)
        grid[(x, y)] = (int(c), sys.maxsize)

#print(grid)
print("Max X and Y")
print(max_x, max_y)
unvisited_set = set(grid.keys())
grid[(0,0)] = (1,0)

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

    unvisited_set.discard(position)

while grid[(max_x, max_y)][1] == sys.maxsize:
    smallest = sys.maxsize
    smallest_tuple = None
    for node in tuple(unvisited_set):
        if grid[node][1] < smallest:
            smallest = grid[node][1]
            smallest_tuple = node

    current_node(smallest_tuple)


print(grid)
print(grid[max_x, max_y])

