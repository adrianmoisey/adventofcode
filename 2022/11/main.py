from collections import deque

with open("input.txt") as file:
#with open("sample.txt") as file:
    input = file.read().strip()

class Monkey():
    def __init__(self, name, items, operation, test, true, false):
        self.name = name
        self.items = items
        self.operation = operation
        self.test = test
        self.true = true
        self.false = false
        self.inspections = 0

monkeys = []

for monkey in input.split('\n\n'):
    split = monkey.split('\n')

    name = split[0].split()[-1].strip(':')
    items = split[1].split(':')[-1].strip().split(',')
    items = deque([int(x) for x in items])
    operation = split[2].split(':')[-1].split('=')[-1].strip()
    test = int(split[3].split(':')[-1].split()[-1])
    true = int(split[4].split(':')[-1][-1])
    false = int(split[5].split(':')[-1][-1])

    monkey = Monkey(name, items, operation, test, true, false)
    monkeys.append(monkey)


for _ in range(20):
    for monkey in monkeys:
        while len(monkey.items) != 0:
            item = monkey.items.popleft()
            monkey.inspections += 1
            # needed for eval. Gross.
            old = item
            worry_level = eval(monkey.operation)
            worry_level = int(worry_level / 3)
            if worry_level % monkey.test == 0:
                monkeys[monkey.true].items.append(worry_level)
            else:
                monkeys[monkey.false].items.append(worry_level)

for i in monkeys:
    print(i.items)
    print(i.inspections)

inspections = [i.inspections for i in monkeys]
top_2 = sorted(inspections)[-2:]
print(top_2[0] * top_2[1])


