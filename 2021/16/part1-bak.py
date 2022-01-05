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


def parse_packet(input):
    global total
    print()
    value, input, binary, count = input
    binary_string = ''
    if binary:
        binary_string = input
    else:
        for c in input:
            binary_string += mapping[c]
    print("Input", input)
    print("Input binary string:", binary_string)
    version = int(binary_string[:3], 2)
    total += version
    type_id = int(binary_string[3:6], 2)
    remaining = binary_string[6:]
    print("Version:", version)
    print("Type id:", type_id)
    print("Remaining:", remaining)

    if type_id == 4:
        position = 5
        new_string = ''
        while True:
            chunk = remaining[position-5:position]
            new_string += chunk[1:]
            if chunk[0] == '0':
                remaining = remaining[position:]
                break
            position += 5

        literal_value = int(new_string, 2)
        value += literal_value
        print("New string", new_string)
        print("New string value", int(new_string, 2))
        if remaining and count != True:
            literal_value += parse_packet((value, remaining, True, count))[0]
        value += literal_value
        return(literal_value, remaining, True, False)
    # operator packet
    else:
        length_id = remaining[0]
        print("Length ID:", length_id)
        if int(length_id) == 0:
            print('length id of zero')
            subpacket_length = remaining[1:16]
            subpacket_length = int(subpacket_length, 2)
            subpackets = remaining[16:16+subpacket_length]
            print("Subpacket length:", subpacket_length, subpackets)
            input_value = 0
            print("Subpackets", subpackets)
            if subpackets:
                input_value = parse_packet([value, subpackets, True, False])[0]
            value += input_value
            return(value, ())
        else: # length is 1
            print('length id of one')
            subpacket_length = remaining[1:12]
            subpacket_length = int(subpacket_length, 2)
            subpackets = remaining[12:]
            print("Subpacket length:", subpacket_length, subpackets, len(subpackets))
            input_value = 0
            print("Subpackets", subpackets)
            for _ in range(subpacket_length):
                print('loop')
                a = parse_packet([value, subpackets, True, True])
            value += input_value
            return(value, ())

total = 0
# [Value, input, boolean, count]
#input = 'D2FE28'
# operator example
#input = (0, ('38006F45291200', False))
#input = ('110100010100101001000100100', True)
#input = (0, ('110100010100101001000100100', True))
#input = 'D14A448'
#input = "A448"
input = (0, '620080001611562C8802118E34', False, False)
#input = (0, '01010000001100100000100011000001100000', True, False)
print(parse_packet(input))


print(total)
