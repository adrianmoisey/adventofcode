with open("input.txt") as input:
    positions_s = input.read()

print("This is the grid")
print(positions_s)
print("This is the grid </end>")

grid = {}

max_x = len(positions_s.splitlines()) - 1
max_y = len(positions_s.splitlines()[0]) - 1

for x, line in enumerate(positions_s.splitlines()):
    for y, c in enumerate(line):
        grid[(x, y)] = int(c)

#print(grid)
print("Max X and Y")
print(max_x, max_y)

def fetch_next(position, tally):
    x, y = position
    print(f"{x}, {y}")
    if x >= max_x and y >= max_y:
        tally += grid[(x,y)]
        print(tally)
        pass

    elif x >= max_x and y < max_y:
        _, tally_next = fetch_next((x, y+1), tally)
        tally += tally_next
    elif x < max_x and y >= max_y:
        _, tally_next = fetch_next((x+1, y), tally)
        tally += tally_next

    elif x < max_x and y < max_y:
        _, tally_next = fetch_next((x, y+1), tally)
        tally += tally_next
        _, tally = fetch_next((x+1, y), tally)
        tally += tally_next

    else:
        print("H")
    return ((x, y), tally)

print(fetch_next((0, 0), 1))
