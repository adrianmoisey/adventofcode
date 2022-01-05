# Sample input
#x_target = (20, 30)
#y_target = (-10, -5)


x_target = (70, 96)
y_target = (-179, -124)

target_area = []
for x in range(x_target[0], x_target[1]+1):
    for y in range(y_target[0], y_target[1]+1):
        target_area.append((x, y))

target_area_max = (x_target[1], y_target[0])


start = (0,0)

def shoot(x, y):
    step = (x,y)

    moves = []
    moves.append(start)
    hit = False

    while True:
        currentx, currenty = moves[-1]
        x = currentx + step[0]
        y = currenty + step[1]
        moves.append((x, y))
        # we hit the target, break out
        if (x, y) in target_area:
            hit = True
            break
        if x > target_area_max[0]:
            #print("we missed due to x")
            moves = []
            break
        if y < target_area_max[1]:
            #print("we missed due to y")
            moves = []
            break

        if step[0] == 0:
            nextx = 0
        elif step[0] > 0:
            nextx = step[0] -1
        else:
            nextx = step[0] +1

        nexty = step[1] -1
        step = (nextx, nexty)

    return(hit)

possible_probes = []


x_target = (70, 96)
y_target = (-179, -124)


for y in range(y_target[0], 1000+1):
    for x in range(x_target[1]+1):
        possible_probes.append((x, y))


y = set()
for a, b in possible_probes:
    if shoot(a, b):
        y.add((a, b))
        print(len(y))

#print(max(y))

# 2293 is too low
