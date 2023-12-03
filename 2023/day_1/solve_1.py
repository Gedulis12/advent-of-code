INPUT = "./input"

lines = []
digits = []
first_last_digits = []
sum = 0

with open(INPUT, 'r') as f:
    for line in f.readlines():
        lines.append(line)

for line in lines:
    cur_line = ""
    for c in line:
        if c.isdigit():
            cur_line = cur_line + str(c)
    digits.append(cur_line)

for d in digits:
    d_modified = ""
    d_modified = d_modified + d[0]
    d_modified = d_modified + d[-1]
    first_last_digits.append(d_modified)

for i in first_last_digits:
    sum += int(i)

print(sum)
