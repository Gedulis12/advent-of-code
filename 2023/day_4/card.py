class Card:

    def __init__(self, raw_data):
        self.raw_data = raw_data

    def card_id(self):
        return self.raw_data.split(':')[0].split(' ')[-1]

    def _card_data(self):
        return self.raw_data.split(':')[1].strip()

    def _card_winning_numbers(self):
        numbers = self._card_data().split('|')[0].strip().split(' ')
        return [number for number in numbers if number != '']

    def _card_my_numbers(self):
        numbers = self._card_data().split('|')[1].strip().split(' ')
        return [number for number in numbers if number != '']

    def _card_my_winning_numbers(self):
        return [
                number for
                number in
                self._card_my_numbers() if
                number in
                self._card_winning_numbers()
                ]

    def _card_my_winning_numbers_count(self):
        return len(self._card_my_winning_numbers())

    def card_points(self):

        count = self._card_my_winning_numbers_count()
        if count > 0:
            score = 1
            for i in range(count-1):
                score = score * 2
            return score
        return 0
