with open("input-sample.txt") as file:
    input = file.readlines()


instructions = {}

for line in input:
    a, b = line.strip().split('->')
    instructions[b.strip()] = a.strip()

instructions['b'] = '16076'

for _ in range(1,1000):
    print("WINNING!", instructions['lx'])
    for i in instructions.keys():
        if i.isdigit():
            pass
        else:
            if "AND" in instructions[i]:
                left, _, right = instructions[i].split()
                if left.isdigit() and instructions.get(right, '').isdigit() or  instructions.get(left, '').isdigit() and instructions.get(right, '').isdigit():
                    if not left.isdigit():
                        a = int(instructions[left]) & int(instructions[right])
                    else:
                        a = int(left) & int(instructions[right])
                    instructions[i] = str(a)
                    print("processing...")

            elif "OR" in instructions[i]:
                left, _, right = instructions[i].split()
                if instructions[left].isdigit() and instructions[right].isdigit():
                    a = int(instructions[left]) | int(instructions[right])
                    instructions[i] = str(a)
                    print("processing...")

            elif "LSHIFT" in instructions[i]:
                left, _, right = instructions[i].split()
                if instructions[left].isdigit():
                    a = int(instructions[left])
                    shift = int(right)
                    a = a << shift
                    instructions[i] = str(a)
                    print("processing...")

            elif "RSHIFT" in instructions[i]:
                left, _, right = instructions[i].split()
                if instructions[left].isdigit():
                    a = int(instructions[left])
                    shift = int(right)
                    a = a >> shift
                    instructions[i] = str(a)
                    print("processing...")

            elif "NOT" in instructions[i]:
                left, right = instructions[i].split()
                if instructions[right].isdigit():
                    a = int(instructions[right])
                    a = ~a
                    a = 65536 + a
                    instructions[i] = str(a)
                    print("processing...")
            else:
                pass
                #print("Stuck!")
                #print(instructions[i])

