from dataclasses import dataclass
#https://stackoverflow.com/questions/2482602/a-general-tree-implementation
# https://pypi.org/project/anytree/

file = 'input.txt'
#file = 'sample.txt'

class Node(object):
    def __init__(self, name, rate, minutes_left, remaining_nodes, total_pressure):
        self.name = name
        self.rate = rate
        self.minutes_left = minutes_left
        self.remaining_nodes = remaining_nodes
        self.total_pressure = total_pressure
        self.children = []
        self.parent = None

    def __repr__(self):
        return self.name

    def add_child(self, obj):
        self.children.append(obj)

@dataclass
class Valve():
    valve: str
    rate: int
    position: str
    feeds: list


with open(file) as file:
    input = file.read().strip().split('\n')

valves = {}

for i in input:
    valve, tunnel = i.split(';')

    _, valve_name, _, _, rate = valve.split()
    rate = int(rate.split('=')[1])
    feeds = tunnel.split()[4:]
    feeds = [feed.strip(',') for feed in feeds]

    new_valve = Valve(valve_name, rate, 'closed', feeds)
    valves[valve_name] = new_valve


def get_distance(valve, nodes_to_visit): #start_position, nodes_to_visit, remaing_nodes, minutes_remaining):
    if valve.name == 'BB':
        print('----')
        print(valve.remaining_nodes)
    start_position = valve.name
    remaing_nodes = valve.remaining_nodes
    if valve.name in remaing_nodes:
        remaing_nodes.remove(valve.name)
    minutes_remaining = valve.minutes_left

    #start_position = start_position
    unvisited_set = nodes_to_visit
    tenative_distance = {key:9999 for key, value in valves.items()}
    tenative_distance[start_position] = 0

    while len(unvisited_set) != 0:

        smallest = 9999999
        for node in unvisited_set:
            if tenative_distance[node] < smallest:
                smallest = tenative_distance[node]
                smallest_node = node

        unvisited_neighbors = valves[smallest_node].feeds

        for neighbour in unvisited_neighbors:
            if tenative_distance[neighbour] > tenative_distance[smallest_node] + 1:
                tenative_distance[neighbour] = tenative_distance[smallest_node] + 1

        unvisited_set.remove(smallest_node)

    nodes_to_return = []

    for node in tenative_distance:
        if node in remaing_nodes:
            rate = valves[node].rate
            moves = tenative_distance[node]
            minutes_left = minutes_remaining - moves - 1
            total_pressure = rate * minutes_left
            total_pressure = valve.total_pressure + rate * minutes_left

            # name, rate, minutes_left, remaining_nodes, total_pressure
            node = Node(node, rate, minutes_left, remaing_nodes.copy(), total_pressure)
            #print(total_pressure)
            if minutes_left >= 0:
                nodes_to_return.append(node)
            #print(f"{node=} {rate=} {moves=} {minutes_left=} {total_pressure}")
    return(nodes_to_return)

remaing_nodes = [valve for valve in valves if valves[valve].rate != 0]
print(remaing_nodes)

seen_nodes = set()

nodes_to_work_on = [Node('AA', 0, 30, remaing_nodes, 0)]
print(nodes_to_work_on)

while len(nodes_to_work_on) != 0:
    valve = nodes_to_work_on.pop()
    print(f"{valve.remaining_nodes=}")
    returned = get_distance(valve, set(valves.keys()))
    print(f"{valve=}")
    print(f"{returned=}")
    for i in returned:
        valve.add_child(i)
        nodes_to_work_on.append(i)
    seen_nodes.add(valve)


## Solution!
#get_distance('AA', set(valves.keys()), remaing_nodes, 30)
#print('-')
#get_distance('DD', set(valves.keys()), remaing_nodes, 28)
#print('-')
#get_distance('BB', set(valves.keys()), remaing_nodes, 25)
#print('-')
#get_distance('JJ', set(valves.keys()), remaing_nodes, 21)
#print('-')
#get_distance('HH', set(valves.keys()), remaing_nodes, 13)
#print('-')
#get_distance('EE', set(valves.keys()), remaing_nodes, 9)
#print('-')
#get_distance('CC', set(valves.keys()), remaing_nodes, 6)
#print('-')

print(max([a.total_pressure for a in seen_nodes]))

#from IPython import embed; embed()



