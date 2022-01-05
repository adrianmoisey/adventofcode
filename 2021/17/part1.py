import heapq


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
    big_y = 0

    while True:
        currentx, currenty = moves[-1]
        x = currentx + step[0]
        y = currenty + step[1]
        moves.append((x, y))
        if y > big_y:
            big_y = y
        # we hit the target, break out
        if (x, y) in target_area:
            break
        if x > target_area_max[0]:
            #print("we missed due to x")
            moves = []
            big_y = 0
            break
        if y < target_area_max[1]:
            #print("we missed due to y")
            moves = []
            big_y = 0
            break

        if step[0] == 0:
            nextx = 0
        elif step[0] > 0:
            nextx = step[0] -1
        else:
            nextx = step[0] +1

        nexty = step[1] -1
        step = (nextx, nexty)

    return(big_y)

possible_probes = []

for y in range(1000):
    for x in range(x_target[0]):
        possible_probes.append((x, y))

y = set()
for a, b in possible_probes:
    y.add(shoot(a, b))
    print(max(y))

print(max(y))

# 3916 is too low
# 15931 is it
