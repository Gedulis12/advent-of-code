from game import Game

INPUT = './input'


def part_1(input):
    part_1_sum = 0
    with open(input, 'r') as f:
        games = f.readlines()
        for g in games:
            game = Game(g)
            if game.game_is_possible():
                part_1_sum += game.game_id()
        print(f'part 1 sum: {part_1_sum}')


def part_2(input):
    part_2_sum = 0
    with open(input, 'r') as f:
        games = f.readlines()
        for g in games:
            game = Game(g)
            red = game.game_max_color('red')
            blue = game.game_max_color('blue')
            green = game.game_max_color('green')
            pow = red*blue*green
            part_2_sum += pow
        print(f'part 1 sum: {part_2_sum}')


def main():
    part_1(INPUT)
    part_2(INPUT)


if __name__ == '__main__':
    main()
