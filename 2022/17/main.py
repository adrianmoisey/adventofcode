from dataclasses import dataclass


file = 'input.txt'
file = 'sample.txt'

with open(file) as file:
    input = file.read().strip()


#rock_list = [dash, cross, l, i, square]
chamber = 7

#def add_row(y):
#    row = []
#    for x in range(chamber):
#        row.append((x, y))
#    return row

grid = {}

# starting
for i in range(3):
    grid[0,0] = 'e'
    grid[0,1] = 'e'
    grid[0,2] = 'e'

print(grid)
#while true:

