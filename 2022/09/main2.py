with open("input.txt") as file:
#with open("sample2.txt") as file:
    input = file.readlines()

H = [0, 0]
knots = []

for i in range(9):
    knots.append([0,0], )

moves = []
part1_position = []
part2_position = []

for i in input:
    direction, distance = i.strip().split()
    distance = int(distance)
    moves.append((direction, distance))

for direction, distance in moves:
    print("Start loop")
    print(direction, distance)
    for _ in range(distance):
        match direction:
            case "U":
                print("H moved up")
                H[0] += 1
            case "D":
                print("H moved down")
                H[0] -= 1
            case "L":
                print("H moved left")
                H[1] -= 1
            case "R":
                print("H moved right")
                H[1] += 1

        for i in range(len(knots)):
            if i == 0:
                previous = H
            else:
                previous = knots[i-1]
            current = knots[i]

            step_size_up_down = max(previous[0], current[0]) - min(previous[0], current[0])
            step_size_left_right = max(previous[1], current[1]) - min(previous[1], current[1])

            if step_size_up_down + step_size_left_right == 3:
                if step_size_left_right == 1:
                    if previous[1] > current[1]:
                        current[1] += 1
                    else:
                        current[1] -= 1
                if step_size_up_down == 1:
                    if previous[0] > current[0]:
                        current[0] += 1
                    else:
                        current[0] -= 1

            if step_size_left_right == 2:
                # We're moving right
                if previous[1] > current[1]:
                    current[1] += 1
                else:
                    current[1] -= 1

            if step_size_up_down == 2:
                # We're moving up
                if previous[0] > current[0]:
                    current[0] += 1
                else:
                    current[0] -= 1

            # First segment after the head
            if i == 0:
                part1_position.append((current[0], current[1]),)
            # Tail
            if len(knots)-1 == i:
                part2_position.append((current[0], current[1]),)

print(len(set(part1_position)))
print(len(set(part2_position)))
