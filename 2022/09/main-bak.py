from collections import OrderedDict

#with open("input.txt") as file:
with open("sample.txt") as file:
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
        print(H, T)

        if (H[0] + H[1]) - (T[0] + T[1]) == 3:
            print("Big move time")

            if H[0] - T[0] == 1:
                print("T moved up - big move")
                T[0] += 1

            elif H[0] - T[0] == -1:
                print("T moved down - big move")
                T[0] -= 1


            elif H[1] - T[1] == 1:
                print("T moved right - big move")
                T[1] += 1

            elif H[1] - T[1] == -1:
                print("T moved left - big move")
                T[1] -= 1

        if  H[0] - T[0] == 0:
            pass

        elif H[0] - T[0] > 1:
            print("T moved up")
            T[0] += 1

        elif H[0] - T[0] < 1:
            print("T moved down")
            T[0] -= 1


        if  H[1] - T[1] == 0:
            pass

        elif H[1] - T[1] > 1:
            print("T moved right")
            T[1] += 1

        elif H[1] - T[1] < 1:
            print("T moved left")
            T[1] -= 1

        print("T had a turn")
        print(H, T)
        pos.append((T[0], T[1]),)
        print()

print(len(set(pos)))
for i  in pos:
    print(i)
# sample == 13
