class Game:
    def __init__(self, raw_data):
        self.raw_data = raw_data
        self.red_constrain = 12
        self.green_constrain = 13
        self.blue_constrain = 14

    def game_id(self):
        return int(self.raw_data.split(':')[0].split(' ')[1].strip())

    def _game_data(self):
        return self.raw_data.split(':')[1].strip()

    def _game_sets(self):
        return self._game_data().split(';')

    def _game_sets_count(self):
        return len(self._game_data().split(';'))

    def game_data(self):
        sets = self._game_sets()
        sets_count = self._game_sets_count()
        game_data = {}

        for i in range(sets_count):
            set_key = f'set_{i}'
            game_data[set_key] = {}

            set = sets[i]
            cubes = set.split(',')
            for cube in cubes:
                if 'blue' in cube:
                    blue_count = int(cube.strip().split(' ')[0])
                    game_data[set_key]['blue'] = blue_count
                if 'red' in cube:
                    red_count = int(cube.strip().split(' ')[0])
                    game_data[set_key]['red'] = red_count
                if 'green' in cube:
                    green_count = int(cube.strip().split(' ')[0])
                    game_data[set_key]['green'] = green_count

        return game_data

    def _set_is_possible(self, set):
        possible = True

        if 'blue' in set.keys():
            if set['blue'] > self.blue_constrain:
                possible = False
        if 'red' in set.keys():
            if set['red'] > self.red_constrain:
                possible = False
        if 'green' in set.keys():
            if set['green'] > self.green_constrain:
                possible = False

        return possible

    def game_is_possible(self):
        possible = True
        data = self.game_data()
        sets = data.keys()
        for set in sets:
            possible = possible and self._set_is_possible(data[set])

        return possible

    def game_max_color(self, color):
        data = self.game_data()
        sets = data.keys()
        max = 0

        for set in sets:
            if color not in data[set].keys():
                continue

            if data[set][color] > max:
                max = data[set][color]

        return max


