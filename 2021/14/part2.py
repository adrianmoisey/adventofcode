from collections import defaultdict

#with open("input.txt") as file:
with open("input-real.txt") as file:
    source = file.read()

template, mapping = source.split('\n\n')
mapping_s = mapping.strip().split('\n')
mapping = {}

for m in mapping_s:
    l, _, r = m.split()
    mapping[l] = r

templates = []
for i, c in enumerate(template):
    max_length = len(template)
    if i == max_length - 1:
        break
    templates.append(c + template[i+1])

last = template[-1]

counter = defaultdict(int)
templates_d = defaultdict(int)

for i, c in enumerate(template):
    max_length = len(template)
    if i == max_length - 1:
        break
    templates_d[c + template[i+1]] += 1


for _ in range(0, 40):
    new_template = defaultdict(int)
    for i in tuple(templates_d):
        new_template[i[0] + mapping[i]] += templates_d[i]
        new_template[mapping[i] + i[1]] += templates_d[i]

    templates_d = new_template

print(templates_d)

counter = defaultdict(int)

for i in templates_d:
    value = templates_d[i]
    l, r = i[0], i[1]

    counter[l] += value

counter[last] += 1

print(counter)

smallest = min(counter.values())
biggest = max(counter.values())

print(smallest, biggest)

# too high 3692219987039
print(biggest - smallest)

