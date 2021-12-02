with open('input.txt') as input_file:
    input = input_file.read()
    input = input.splitlines()

count = 0
for i in input:
    l, w, h = i.split('x')
    l = int(l)
    w = int(w)
    h = int(h)

    sl = l*w
    sr = w*h
    sf = h*l

    area = 2*sl + 2*sr + 2*sf
    small_side = min(sl, sr, sf)
    count += area+small_side

print(count)
