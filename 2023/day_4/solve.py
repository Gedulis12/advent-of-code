from card import Card
from collections import defaultdict

INPUT = './input'
part_1_sum = 0


def get_cards(input):
    with open(INPUT, 'r') as f:
        return [card for card in f.readlines()]
    f.close()


for i in get_cards(INPUT):
    card = Card(i)
    part_1_sum += card.card_points()
print(part_1_sum)

N = defaultdict(int)
for i, line in enumerate(get_cards(INPUT)):
    N[i] += 1
    winning_num = int(Card(line)._card_my_winning_numbers_count())
    for j in range(winning_num):
        N[i+1+j] += N[i]
print(sum(N.values()))


