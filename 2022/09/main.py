from collections import OrderedDict

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.readlines()

T = [0, 0]
H = [0, 0]

moves = []
pos = []

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

        print("H had a turn")
        print(f"H: {H}, T: {T}")

        step_size_up_down = max(H[0], T[0]) - min(H[0], T[0])
        step_size_left_right = max(H[1], T[1]) - min(H[1], T[1])
        print("Step: ", step_size_up_down, step_size_left_right)

        if step_size_up_down + step_size_left_right == 3:
            if step_size_left_right == 1:
                if H[1] > T[1]:
                    T[1] += 1
                else:
                    T[1] -= 1
            if step_size_up_down == 1:
                if H[0] > T[0]:
                    T[0] += 1
                else:
                    T[0] -= 1

        if step_size_left_right == 2:
            # We're moving right
            if H[1] > T[1]:
                T[1] += 1
            else:
                T[1] -= 1

        if step_size_up_down == 2:
            # We're moving up
            if H[0] > T[0]:
                T[0] += 1
            else:
                T[0] -= 1

        print("T had a turn")

        print(f"H: {H}, T: {T}")
        pos.append((T[0], T[1]),)
        print()

print(len(set(pos)))
# sample == 13
