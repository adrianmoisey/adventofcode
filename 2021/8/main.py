with open('input.txt') as input_file:
    INPUT_S = input_file.read()
input_lines = INPUT_S.strip().split('\n')

total = 0

def deduce_mapping(patterns):
    # pattern to digit mapping
    d2p = {}

    for p, plen in patterns:
        if plen == 2: # 1
            d2p[1] = p
        elif plen == 3: # 7
            d2p[7] = p
        elif plen == 4: # 4
            d2p[4] = p
        elif plen == 7: # 8
            d2p[8] = p

    p2d = {v: k for k, v in d2p.items()}

    for p, plen in patterns:
        if p in p2d:
            continue

        if plen == 5:
            # 2 or 3 or 5
            if len(p & d2p[1]) == 2:
                # 3 has 2 ON segments in common with 1
                p2d[p] = 3
            elif len(p & d2p[4]) == 3:
                # 5 has 3 ON segments in common with 4
                p2d[p] = 5
            else:
                p2d[p] = 2
        else:
            # 0 or 6 or 9
            if len(p & d2p[4]) == 4:
                # 9 has 4 ON segments in common with 4
                p2d[p] = 9
            elif len(p & d2p[7]) == 2:
                # 6 has 2 ON segments in common with 7
                p2d[p] = 6
            else:
                p2d[p] = 0


    return p2d



for line in input_lines:
    patterns, digits = [], []
    left, right = line.split('|')

    right_l = right.split()
    left_l = left.split()

    for p in left_l:
        patterns.append((frozenset(p), len(p)))

    for d in right_l:
        digits.append((frozenset(d), len(d)))

    p2d = deduce_mapping(patterns)

    total += p2d[digits[0][0]] * 1000
    total += p2d[digits[1][0]] * 100
    total += p2d[digits[2][0]] * 10
    total += p2d[digits[3][0]]


print(total)
