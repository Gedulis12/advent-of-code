package main

import (
	"bufio"
	"fmt"
	"cmp"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
    HIGH_CARD int = iota
    ONE_PAIR
    TWO_PAIR
    THREE_OF_KIND
    FULL_HOUSE
    FOUR_OF_KIND
    FIVE_OF_KIND
)

type hand struct {
    card  string
    bid   int
    htype int
    rank  int
}

var strength = map[byte]int{
    'A': 13,
    'K': 12,
    'Q': 11,
    'J': 10,
    'T': 9,
    '9': 8,
    '8': 7,
    '7': 6,
    '6': 5,
    '5': 4,
    '4': 3,
    '3': 2,
    '2': 1,
}

func main () {
    res1 := SolvePart1("./input")
    res2 := SolvePart2("./input")

    fmt.Println("Part 1 result: ", res1)
    fmt.Println("Part 2 result: ", res2)
}

func SolvePart1 (inputPath string) int {
    file, err := os.Open(inputPath)
    if err != nil {
        panic(err)
    }
    hands := []hand{}
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        c := strings.Split(scanner.Text(), " ")[0]
        b, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])

        if err != nil {
            panic(err)
        }

        h := hand{card: c, bid: b,}
        hands = append(hands, h)
    }
    for i, h := range(hands) {
        hands[i].htype = getType(h)
    }
    hands = rankHands(hands)
    result := 0
    for _, h := range(hands) {
        result += h.rank * h.bid
    }
    return result
}

func SolvePart2 (inputPath string) int {
    return 0
}

func rankHands(hands []hand) []hand {
    slices.SortFunc(hands, func (a, b hand) int {
        if n := cmp.Compare(a.htype, b.htype); n != 0 {
            return n
        }

        for i := 0; i < 5; i++ {
            if n := cmp.Compare(strength[a.card[i]], strength[b.card[i]]); n != 0 {
                return n
            }
        }
        return 0
    })
    newHands := []hand{}
    for i, h := range(hands) {
        h.rank = i+1
        newHands = append(newHands, h)
    }
    return newHands
}

func getType(h hand) int {
    typeMap := map[rune]int{}
    for _, i := range(h.card) {
        _, ok := typeMap[i]
        if ok {
            typeMap[i]++
        } else {
            typeMap[i] = 1
        }
    }
    values := []int{}
    for _, v := range(typeMap) {
        values = append(values, v)
    }
    if len(values) == 1 && values[0] == 5 {
        return FIVE_OF_KIND
    }
    if len(values) == 2 && slices.Contains(values, 4) {
        return FOUR_OF_KIND
    }
    if len(values) == 2 && slices.Contains(values, 3) && slices.Contains(values, 2) {
        return FULL_HOUSE
    }
    if len(values) == 3 && slices.Contains(values, 3) && slices.Contains(values, 1) {
        return THREE_OF_KIND
    }
    if slices.Equal(values, []int{1, 2, 2}) || slices.Equal(values, []int{2, 1, 2}) || slices.Equal(values, []int{2, 2, 1}) {
        return TWO_PAIR
    }
    if slices.Equal(values, []int{2, 1, 1, 1}) || slices.Equal(values, []int{1, 2, 1, 1}) || slices.Equal(values, []int{1, 1, 2, 1}) || slices.Equal(values, []int{1, 1, 1, 2}){
        return ONE_PAIR
    }
    return HIGH_CARD
}
