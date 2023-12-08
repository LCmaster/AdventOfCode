package main

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

type Hand struct {
    cards []rune
    strength int
    bid int
}

func getCardValue(card rune, joker rune) int {
    cardValue := 0
    switch card {
    case 'A':
        cardValue = 14
    case 'K':
        cardValue = 13
    case 'Q':
        cardValue = 12
    case 'J':
        cardValue = 11
    case 'T':
        cardValue = 10
    default:
        value, _ := strconv.Atoi(string(card))
        cardValue = value
    }

    if joker == card {
        cardValue = 1
    }

    return cardValue
}

func getHandStrengh(cards []rune) int {
    cardCount := map[rune]int{
        '2':0, '3':0, '4':0, 
        '5':0, '6':0, '7':0, 
        '8':0, '9':0, 'T':0, 
        'J':0, 'Q':0, 'K':0, 
        'A':0,
    }
    for _, ch := range cards {
        cardCount[ch] += 1
    }
        strength := 0
    for _, count := range cardCount {
        if count >= 2 {
            strength += int(math.Pow(10, float64(count)))
        } 
    }
    return strength
}

func part1(input []string) int {
    output := 0
    hands := []Hand{} 
    for _, line := range input {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }

        token := strings.Split(line, " ")

        cards := make([]rune, 5)
        for _, ch := range strings.TrimSpace(token[0]) {
            cards = append(cards, ch)
        }
        bid, _ := strconv.Atoi(token[1])
        strength := getHandStrengh(cards)
                hands = append(hands, Hand{
            cards: cards,
            strength: strength,
            bid: bid,
        })
    }
    slices.SortFunc(
        hands, 
        func(a, b Hand) int { 
            result := cmp.Compare(a.strength, b.strength)
            if result == 0 {
                for i := 0; i < len(a.cards); i++ {
                    result = cmp.Compare(
                        getCardValue(a.cards[i], ' '),
                        getCardValue(b.cards[i], ' '),
                    )
                    if result != 0 {
                        break
                    }
                }
            }
            return result
        },
    )
    for i, hand := range hands {
        output += (i+1) * hand.bid
    }

    return output
}

func part2(input []string) int {
    deck := []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
    joker := 'J'

    output := 0
    hands := []Hand{} 
    for _, line := range input {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }

        token := strings.Split(line, " ")

        cards := make([]rune, 5)
        for _, ch := range strings.TrimSpace(token[0]) {
            cards = append(cards, ch)
        }
        bid, _ := strconv.Atoi(token[1])
        strength := getHandStrengh(cards)

        for _, replacement := range deck {
            newCards := append([]rune{}, cards...)
            for i, ch := range newCards {
                if ch == joker {
                    newCards[i] = replacement
                }
            }
            newCardsStrength := getHandStrengh(newCards)
            if newCardsStrength > strength {
                strength = newCardsStrength
            }
        }

        hands = append(hands, Hand{
            cards: cards,
            strength: strength,
            bid: bid,
        })
    }
    slices.SortFunc(
        hands, 
        func(a, b Hand) int { 
            result := cmp.Compare(a.strength, b.strength)
            if result == 0 {
                for i := 0; i < len(a.cards); i++ {
                    result = cmp.Compare(
                        getCardValue(a.cards[i], joker),
                        getCardValue(b.cards[i], joker),
                    )
                    if result != 0 {
                        break
                    }
                }
            }
            return result
        },
    )
    for i, hand := range hands {
        output += (i+1) * hand.bid
    }

    return output
}

func main() {
    input, err := utils.ReadInputFile(2023, 7)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1:", part1(input))
    fmt.Println("Answer to part 2:", part2(input))
}
