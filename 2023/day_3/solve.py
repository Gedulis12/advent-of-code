INPUT = './input'

numbers = []
symbols = []
part_1_sum = 0
part_2_sum = 0

with open(INPUT, 'r') as f:
    line_len = len(f.readline())
    f.close()

with open(INPUT, 'r') as f:
    current_line = 0

    for line in f.readlines():
        current_number = ""

        for c in range(line_len):
            cur_char = line[c]

            if (cur_char.isascii() and not
                    cur_char.isdigit() and
                    cur_char != '.' and
                    cur_char != '\n'):
                symbols.append({
                    'symbol': cur_char,
                    'line': current_line,
                    'position': c
                    })

            if cur_char.isdigit():
                current_number = current_number + cur_char

            if cur_char.isdigit() and (not line[c-1].isdigit() or c == 0):
                start_pos = c

            if not cur_char.isdigit() and (line[c-1].isdigit() or c == line_len):
                end_pos = c-1
                numbers.append({
                    'number': current_number,
                    'start_pos': start_pos,
                    'end_pos': end_pos,
                    'line': current_line
                    })
                current_number = ""
        current_line += 1
    f.close()


def has_adj_symbols(number, symbols):
    adj_symbols = [
            symbol for
            symbol in
            symbols if
            symbol['line'] - number['line'] >= -1 and
            symbol['line'] - number['line'] <= 1 and
            symbol['position'] >= number['start_pos'] - 1 and
            symbol['position'] <= number['end_pos'] + 1
            ]
    if len(adj_symbols) > 0:
        return True
    return False


def get_gear_ratio_for_numbers(numbers, symbol):
    part_nums = [
            number for
            number in
            numbers if
            symbol['line'] - number['line'] >= -1 and
            symbol['line'] - number['line'] <= 1 and
            symbol['position'] >= number['start_pos'] - 1 and
            symbol['position'] <= number['end_pos'] + 1
            ]

    if len(part_nums) == 2:
        print(f'part nums: {part_nums} for gear: {symbol}')
        return int(part_nums[0]['number']) * int(part_nums[1]['number'])
    return 0


for n in numbers:
    if has_adj_symbols(n, symbols):
        part_1_sum += int(n['number'])

for s in symbols:
    if s['symbol'] == '*':
        ratio = get_gear_ratio_for_numbers(numbers, s)
        part_2_sum += ratio


print(part_1_sum)
print(part_2_sum)
