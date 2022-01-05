with open("input.txt") as file:
    input = file.read().strip()

mapping = {
    '0': '0000',
    '1': '0001',
    '2': '0010',
    '3': '0011',
    '4': '0100',
    '5': '0101',
    '6': '0110',
    '7': '0111',
    '8': '1000',
    '9': '1001',
    'A': '1010',
    'B': '1011',
    'C': '1100',
    'D': '1101',
    'E': '1110',
    'F': '1111',
}

def type_four(binary_s, value):
    position = 5
    new_string = ''
    while True:
        chunk = binary_s[position-5:position]
        new_string += chunk[1:]
        if chunk[0] == '0':
            remaining = binary_s[position:]
            binary_s = binary_s[position:]
            break
        position += 5
    value = int(new_string, 2)

    print(f"{remaining=}")
    print(f"{value=}")
    return value, remaining


def type_operator(binary_s, value):
    # 0 represents that length is total length of sub packets
    # 1 means that length is the number of subpackets
    length_type = binary_s[0]

    if length_type == '0':
        length = binary_s[1:16]
        length = int(length, 2)
        print(f"{length_type=}, {length=}")

        subpackets = binary_s[16:16+int(length)]
        print("Processing length type of 0")
        print(f"{subpackets=}")
        processed = 0
        while processed != length:
            print("in while loop")
            value, remaining = parse_string(subpackets, value)
            print(f"{remaining}")
            processed += len(remaining)

    elif length_type == '1':
        length = binary_s[1:12]
        length = int(length, 2)
        print(f"{length_type=}, {length=}")

        subpackets = binary_s[12:]
        print("Processing length type of 1")
        print(f"{subpackets=}")
        for _ in range(length):
            print('in loop')
            print(f"{subpackets=}")
            value, subpackets = parse_string(subpackets, value)

    else:
        raise "Whoops"
    return value, ''


def parse_string(binary_s, value):
    global total
    print()
    if len(binary_s) == 0:
        return value, ''
    print(f"Parsing string: {binary_s}")
    version = int(binary_s[:3], 2)
    # Part 1
    total += version

    type_id = int(binary_s[3:6], 2)
    remaining = binary_s[6:]
    print(f"{version=}, {type_id=}, {remaining=}")

    if type_id == 4:
        print("Processing value...")
        value, remaining = type_four(remaining, value)

    # type if of other than 4 is an operator
    else:
        print("Processing operator...")
        value, remaining =  type_operator(remaining, value)

    return value, remaining

total = 0
input = '620080001611562C8802118E34'
binary_string = ''
for c in input:
    binary_string += mapping[c]

print(parse_string(binary_string, 0))
print("Total")
print(total)
