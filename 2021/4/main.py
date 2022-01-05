from functools import reduce

with open('input-real.txt') as input_file:
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

max_bingos = len(bingos)

def check_winner(bingo):
    for line in bingo:
        if line.count('x') > 0 and line.count('x') == len(line):
            return(True)
    flip = list(zip(*reversed(bingo)))
    for line in flip:
        if line.count('x') > 0 and line.count('x') == len(line):
            return(True)

winner = []
drawers = []
number = False

for draw in draw_numbers:
    if len(winner) == max_bingos:
        break
    for idx, bingo in enumerate(bingos):
        for line in bingo:
            if draw in line:
                line[line.index(draw)] = 'x'
        if check_winner(bingo):
            if idx not in winner:
                winner.append(idx)
                drawers.append(draw)

print(winner[-1])
print(drawers[-1])

winner_board = bingos[winner[-1]]

count = 0
for c in reduce(lambda x, y: x+y, winner_board):
    if c != 'x':
        count += int(c)

print(count * int(drawers[-1]))
