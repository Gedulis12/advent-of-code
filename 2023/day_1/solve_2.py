INPUT = "./input"

lines = []
lines_digitized = []
digits = []
first_last = []
sum = 0

alphadigits = {
    "one": "one1one",
    "two": "two2two",
    "three": "three3three",
    "four": "four4four",
    "five": "five5five",
    "six": "six6six",
    "seven": "seven7seven",
    "eight": "eight8eight",
    "nine": "nine9nine",
        }

with open(INPUT, 'r') as f:
    for line in f.readlines():
        lines.append(line)

for line in lines:
    newline = line
    for k in alphadigits.keys():
        if k in line:
            newline = newline.replace(k, alphadigits[k])
    lines_digitized.append(newline)


for line in lines_digitized:
    cur = ""
    for c in line.lower():
        if c.isdigit():
            cur = cur + c
    digits.append(cur)


for digit in digits:
    first = digit[0]
    last = digit[-1]
    first_and_last = str(first) + str(last)
    first_last.append(first_and_last)


for i in first_last:
    sum += int(i)

print(sum)
