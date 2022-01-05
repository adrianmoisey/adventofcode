#with open("input.txt") as file:
with open("input-real.txt") as file:
    input_s = file.read()

def find_all(a_string, sub):
    result = []
    k = 0
    while k < len(a_string):
        k = a_string.find(sub, k)
        if k == -1:
            return result
        else:
            result.append(k)
            k += 1 #change to k += len(sub) to not search overlapping results
    return result


width = len(input_s.split('\n')[0])

input = input_s.replace('\n', '')
length = len(input)

mins = []

for i in range(0,9):
    i_str = str(i)
    i_int = i
    positions = find_all(input, i_str)
    for pos in positions:
        too_big = False
        # check left:
        if pos % width > 0:
            if int(input[pos-1]) <= i_int:
                too_big = True

        #check right:
        if pos % width < width-1:
            if int(input[pos+1]) <= i_int:
                too_big = True
        #check up

        if pos-width >= 0:
            if int(input[pos-width]) <= i_int:
                too_big = True
        # check down:
        if pos+width <= length-1:
            if int(input[pos+width]) <= i_int:
                too_big = True

        if too_big is False:
            mins.append((pos, input[pos]))

print(mins)

def fetch_neighbours(ns):
    neighbours = []
    for position, value in ns:
        #print(position, value)
        i_str = str(position)
        i_int = int(value)
        # check left:
        if position % width > 0:
            if int(input[position-1]) >= i_int+1 and int(input[position-1]) != 9:
                neighbours.append((position-1, input[position-1]))

        #check right:
        if position % width < width-1:
            if int(input[position+1]) >= i_int+1 and int(input[position+1]) != 9:
                neighbours.append((position+1, input[position+1]))

        if position-width >= 0:
            if int(input[position-width]) >= i_int+1 and int(input[position-width]) != 9:
                neighbours.append((position-width, input[position-width]))
        # check down:
        if position+width <= length-1:
            if int(input[position+width]) >= i_int+1 and int(input[position+width]) != 9:
                neighbours.append((position+width, input[position+width]))

    if len(neighbours) == 0:
        return []
    else:
        return neighbours + fetch_neighbours(neighbours)

smallest = []
for min in mins:
    a = fetch_neighbours([min]) + [min]
    smallest.append(len(set(a)))

smallest.sort()
s = smallest[-3::]

result = s[0] * s[1] * s[2]
print(result)
