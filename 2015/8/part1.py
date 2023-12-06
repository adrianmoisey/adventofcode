with open("input.txt") as file:
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

    memory_length = 0

    line = line[1:-1]

    for count, character in enumerate(line):
        if skip > 0:
            skip = skip - 1
            continue
        print(character, memory_length)

        if character == "\\":
            if line[count+1] == "x":
                skip = 3
                print("Skipping three")
                memory_length += 1
                continue

            elif line[count+1] == "\\":
                skip = 1
                print("Skipping one")
                memory_length += 1
                continue
            elif line[count+1] == '"':
                skip = 1
                print("Skipping one")
                memory_length += 1
                continue
        memory_length += 1

    print(f"Memory: {memory_length}")

    total_memory_length += memory_length
    total_literal_length += literal_length

print("==========")
print(total_literal_length - total_memory_length)
