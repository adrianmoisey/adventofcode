from collections import defaultdict

with open("input.txt") as file:
#with open("input-real.txt") as file:
    source = file.read()

template, mapping = source.split('\n\n')

mapping_s = mapping.strip().split('\n')

mapping = {}

for m in mapping_s:
    l, _, r = m.split()
    mapping[l] = r

for _ in range(0, 40):
    new_string = ''
    max_length = len(template)
    for i, c in enumerate(template):
        if i == max_length - 1:
            new_string += c
            break
        new_string += c
        pair = c + template[i+1]
        new_string += mapping[pair]
    template = new_string

print(template)

counter = defaultdict(int)

for c in template:
    counter[c] += 1

print(counter)

smallest = min(counter.values())
biggest = max(counter.values())

print(smallest, biggest)

print(biggest - smallest)
