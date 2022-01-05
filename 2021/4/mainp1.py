from functools import reduce

with open('input.txt') as input_file:
    raw_input = input_file.read()


draw_numbers = raw_input.split('\n', 1)[0].split(',')
raw_bingos = raw_input.split('\n', 1)[1].split('\n\n')

bingos_string = [bingo.split('\n') for bingo in raw_bingos]

bingos = []
for bingo in bingos_string:
    new_bingo = []
    for line in bingo:
        new_bingo.append(line.split())
    bingos.append(new_bingo)

def check_winner(bingo):
    for line in bingo:
        if line.count('x') > 0 and line.count('x') == len(line):
            return(True)
    flip = list(zip(*reversed(bingo)))
    for line in flip:
        if line.count('x') > 0 and line.count('x') == len(line):
            return(True)

winner = []
number = False

for draw in draw_numbers:
    if len(winner) == 2:
        break
    for bingo in bingos:
        for line in bingo:
            if draw in line:
                line[line.index(draw)] = 'x'
        if check_winner(bingo):
            winner.append(bingo)
            winner.append(draw)

print(winner[0])

count = 0
for c in reduce(lambda x, y: x+y, winner[0]):
    if c != 'x':
        count += int(c)

print(count * int(winner[1]))
