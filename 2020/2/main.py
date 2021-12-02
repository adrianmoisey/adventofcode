

def read_input():
    with open('input.txt') as input_file:
        input = input_file.read()

    return input.splitlines()

# Part 1
match = 0
for password in read_input():
    policy, letter, password = password.split()
    letter = letter[0]

    count = password.count(letter)

    policy_min, policy_max = policy.split('-')

    if count >= int(policy_min) and count <= int(policy_max):
        match += 1

print(match)

# Part 2
match = 0
for password in read_input():
    policy, letter, password = password.split()
    letter = letter[0]

    policy_left, policy_right = policy.split('-')

    policy_left = int(policy_left)
    policy_right = int(policy_right)

    begin = password[policy_left-1] == letter
    end = password[policy_right-1] == letter

    #if password[policy_left-1] == letter and password[policy_right-1] == letter:
    #    continue
    if begin ^ end:
    #if password[policy_left-1] == letter ^ password[policy_right-1] == letter:
        match += 1

print(match)
