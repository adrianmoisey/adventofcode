with open("input-sample.txt") as file:
    input = file.readlines()


skip = 0

total_memory_length = 0
total_literal_length = 0

for line in input:
    print("=================")
    line = line.strip()
    literal_length = len(line)
    print(line)
    print(f"Length: {literal_length}")
    new_string = []

    for i in line:
        new_string.append(i)
        if i == '"':
            new_string.append('\\"')
        elif i == "\\":
            new_string.append('\\\\')

    new_string = "".join(new_string)
    print(new_string, len(new_string))



