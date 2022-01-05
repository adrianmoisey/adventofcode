
"""
Addition:

For example, [1,2] + [[3,4],5] becomes [[1,2],[[3,4],5]].
"""

"""
Reduction:

- If any pair is nested inside four pairs, the leftmost such pair explodes.
- If any regular number is 10 or greater, the leftmost such regular number splits.

During reduction, at most one action applies, after which the process returns to the top of the list of actions. For example, if split produces a pair that meets the explode criteria, that pair explodes before other splits occur.
"""

input = [[[[[9,8],1],2],3],4]
input = [7,[6,[5,[4,[3,2]]]]]
input = [[6,[5,[4,[3,2]]]],1]
input = [[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]


# explode
print(input)

for i1, a in enumerate(input):
    print(f"{a=}")
    if isinstance(a, list):
        for i2, b in enumerate(a):
            print(f"{b=}")
            if isinstance(b, list):
                for i3, c in enumerate(b):
                    print(f"{c=}")
                    if isinstance(c, list):
                        for i4, d in enumerate(c):
                            if isinstance(d, list):

                                print('Here is where we do the explode')
                                print(f"{d=}")

                                # Looking left
                                if i3 > 0:
                                    d[0] = d[0] + c[i3 -1]
                                else:
                                    d[0] = 0

                                # Looking right
                                if i3 + 1 < len(c):
                                    d[1] += c[i3 + 1]
                                else:
                                    print('here')
                                    if i3+1 <= len(b):
                                        if i2+1 <= len(a):
                                            if i1+1 < len(input):
                                                if isinstance(input[i1+1], list):
                                                    input[i1+1][0] += d[1]
                                                else:
                                                    input[i1+1] += d[1]
                                    d[1] = 0

                                b[i2] = d


print(input)
