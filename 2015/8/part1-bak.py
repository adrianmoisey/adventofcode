with open("input-sample.txt") as file:
    input = file.readlines()

# \\ - single backslash
# \" - single quote
# \xXX - single ascii character

def findall(f, s):
    l = []
    i = -1
    while True:
        i = s.find(f, i+1)
        if i == -1:
            return l
        l.append(s.find(f, i))

total_length = 0
total_memory_length = 0

for line in input:
    print("=================")
    line = line.strip()
    length = len(line)
    print(line)
    print(f"Length: {length}")

    memory_length = length - 2  # remove trailing and leading "
    print(memory_length)

    memory_length = memory_length - len(findall('\\\\', line))  # remove for \\ characters
    line = line.replace('\\\\', '')
    print(memory_length)


    memory_length = memory_length - len(findall('\\"', line))  # remove for \# characters
    line = line.replace('\\#', '')
    print(memory_length)

    memory_length = memory_length - (len(findall('\\x', line)) ) # remove for \x<hex digit><hexdigit> characters

    total_length += length
    print(f"Memory Length: {memory_length}")
    total_memory_length += memory_length


print("=====")
print(total_length)
print(total_memory_length)
print(total_length-total_memory_length)

# 1252 too low
# 2117 not right
# 2978 too high
