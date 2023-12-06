from collections import defaultdict

with open("input.txt") as file:
#ith open("sample.txt") as file:
    input = file.readlines()

filesystem = defaultdict(int)
directory_structure = defaultdict(int)
current_location = []

total_disk_space = 70000000
needs            = 30000000

for i in input:
    i = i.strip()
    if i.startswith('$'):
        _, command = i.split(' ', 1)
        if command.startswith('cd'):
            _, path = command.split()
            if path == '..':
                current_location.pop()
            else:
                path_string = '/'.join(current_location) + '/' + path
                path_string = path_string[1:]
                directory_structure[path_string] = 1
                current_location.append(path)
    else:
        size, text = i.split()
        if size == 'dir':
            pass
        else:
            path_string = '/'.join(current_location) + '/' + text
            path_string = path_string[1:]
            filesystem[path_string] += int(size)


#print(filesystem)
#print(directory_structure)

used_space = sum(filesystem.values())
free_space = total_disk_space - used_space
# need to get this one
delta = 30000000 - free_space

#print(delta)

deltas = []
for d in sorted(directory_structure.keys()):
    total = 0
    for dd in sorted(filesystem.keys()):
        if dd.startswith(d):
            print(dd, d)
            total += filesystem[dd]

    if total > delta:
        deltas.append(total)

print(min(deltas))
