import sys
import heapq

with open("input2.txt") as input:
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
                 if value > 9:
                    value = value - 9
#                 print((x + (x_m * short_x), y + (y_m  * short_y)))
                 grid[(x + (x_m * short_x), y + (y_m  * short_y))] = (value, sys.maxsize)

#print(grid)
print("Max X and Y")
print(max_x, max_y)
grid[(0,0)] = (0,0)

#for x in range(0, max_x+1):
#    for y in range(0, max_y+1):
#        print(grid[(x, y)][0], end='')
#    print()

seen = set()
queue = [(0, (0,0))]

def current_node(position):
    seen.add(position)
    x, y = position
    neighbors = [(x+1, y), (x, y+1), (x-1, y), (x-1, y)]
    for n in neighbors:
        if grid.get(n, None) != None:
            # our tentative distance + their size
            #print(grid[position][1])
            #print(grid[n][0])
            distance = grid[position][1] + grid[n][0]
            #print("Distance:", distance)

            # Get smaller of us+them, or them
            smaller = min([distance, grid[n][1]])
            #print(n)
            #print(smaller)
            grid[n] = (grid[n][0], smaller)

            heapq.heappush(queue, (smaller, n))

#while grid[(max_x, max_y)][1] == sys.maxsize:
while True:
    if (max_x, max_y) in seen:
        break
    #print(queue)
    v, node = heapq.heappop(queue)
    if node in seen:
        continue
    current_node(node)

#print(grid)
print(grid[max_x, max_y])

# 2881 is too high
