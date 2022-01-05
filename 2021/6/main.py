from collections import Counter
with open('input-real.txt') as input_file:
#with open('input.txt') as input_file:
    raw_input = input_file.read()

lantern_fish_s = raw_input.strip().split(',')

lantern_fish = [int(fish) for fish in lantern_fish_s]
c = Counter(lantern_fish)

print(lantern_fish)

days = 256
ages = list(range(8,-1,-1))
reborn = 0

for day in range(days):
    n_c = Counter()
    for age in ages:
        number_of_fish = c[age]
        if age == 0:
            n_c[6] += number_of_fish
            n_c[8] += number_of_fish
        #elif age == 6:
        #    n_c[age-1] = c[age] + reborn
        #    reborn = 0
        else:
            n_c[age-1] = c[age]
    c = n_c

#print(lantern_fish)
print(c)
print(c.total())
