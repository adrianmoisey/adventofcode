def read_input():
    with open('input.txt') as input_file:
        input = input_file.read()

    return input.split('\n\n')

input = read_input()

fields = [
    'byr',
    'iyr',
    'eyr',
    'hgt',
    'hcl',
    'ecl',
    'pid',
    'cid',
]

valid = 0

for i in input:
    i = i.replace('\n', ' ')
    #i.split()
    count = i.count(':')
    if count == 7:
        if 'cid' not in i:
            valid += 1
    if count == 8:
        valid += 1

print(valid)



valid = 0

for i in input:
    i = i.replace('\n', ' ')
    passport = {}
    for field in i.split():
        f, v = field.split(':')
        passport[f] = v
    difference = set(fields) - set(passport.keys())
    if list(difference) == ['cid'] or len(difference) == 0:
        if passport.get()



        valid += 1

print(valid)

