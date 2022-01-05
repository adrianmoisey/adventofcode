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
            mins.append(input[pos])

mins = [int(min)+1 for min in mins]
print(sum(mins))
