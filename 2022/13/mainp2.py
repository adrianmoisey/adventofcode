from itertools import zip_longest

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.read().strip().split('\n\n')

NoneType = type(None)

# True == Inputs are in the right order
# False == Inputs are not in the right order

def lister_tester(left, right):
    if isinstance(left, list) and isinstance(right, list):
        for l, r in zip_longest(left, right):
            result = lister_tester(l, r)
            if result == None:
                continue
            else:
                return result

    if isinstance(left, int) and isinstance(right, int):
        if left > right:
            return False
        elif left < right:
            return True
        else:
            return None

    if isinstance(left, list) and isinstance(right, int):
        return lister_tester(left, [right])

    if isinstance(left, int) and isinstance(right, list):
        return lister_tester([left], right)

    if isinstance(left, NoneType):
        return True

    if isinstance(right, NoneType):
        return False

packets = []
packets.append([[2]])
packets.append([[6]])

for index, i in enumerate(input):
    left, right = i.split('\n')
    left = eval(left)
    right = eval(right)

    packets.append(left)
    packets.append(right)

#    if lister_tester(left, right) is True:
#        sum += index+1

swapped = True
while swapped == True:
    swapped = False
    for i in range(len(packets)):
        if i == len(packets) - 1:
            break

        left = packets[i]
        right = packets[i+1]

        if lister_tester(left, right) is False:
            packets[i] = right
            packets[i+1] = left
            swapped = True


pos = 1

for index, i in enumerate(packets):
    if i == [[2]]:
        print(pos)
        pos = pos * (index+1)
        print(index+1)
    if i == [[6]]:
        print(pos)
        pos = pos * (index+1)
        print(index+1)

print(pos)
