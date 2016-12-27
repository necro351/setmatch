package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"sort"
)

type card struct {
	count byte
	color byte
	shade byte
	shape byte
}

const (
	wave = 'w'
	pill = 'p'
	diamond = 'd'

	one = '1'
	two = '2'
	three = '3'

	red = 'r'
	purple = 'p'
	green = 'g'

	pen = 'p'
	solid = 's'
	hollow = 'h'
)

func main() {
	cards := parseCards()
	if len(cards) != 12 && len(cards) != 15 {
		log.Printf("warning: number of cards (%d) is not 12 or 15\n", len(cards))
	}
	sets := findSets(cards)
	for set, _ := range sets {
		printSet(set, cards)
	}
}

func parseCards() (cards []card) {
	for {
		var cardStr string
		_, err := fmt.Fscanln(os.Stdin, &cardStr)
		if err == io.EOF {
			break
		} else if err != nil && err != io.EOF {
			log.Printf("warning: could not scan line: %v\n", err)
			continue
		}

		card, err := parseCard(cardStr)
		if err != nil {
			log.Print("error: ", err)
		} else {
			cards = append(cards, card)
		}
	}

	return
}

func findSets(cards []card) map[[3]int]bool {
	sets := make(map[[3]int]bool)
	for i := 0; i < len(cards); i += 1 {
		for j := 0; j < len(cards); j += 1 {
			if j == i {
				j += 1
				continue
			}
			cardNeeded := thirdCard(cards[i], cards[j])
			for k := 0; k < len(cards); k += 1 {
				if k == i || k == j {
					continue;
				}
				if cards[k] == cardNeeded {
					set := [3]int{i, j, k}
					sort.Ints(set[:])
					sets[set] = true
				}
			}
		}
	}
	return sets
}

func printSet(set [3]int, cards []card) {
	fmt.Printf("%s,%s,%s\n", cards[set[0]].String(), cards[set[1]].String(), cards[set[2]].String())
}

func (c *card) String() string {
	return string(c.shape) + string(c.count) + string(c.color) + string(c.shade)
}

func parseCard(cardStr string) (card, error) {
	if len(cardStr) != 4 {
		return card{}, fmt.Errorf("card '%s' has the wrong length", cardStr)
	}
	switch cardStr[0] {
	case one:
	case two:
	case three:
	default:
		return card{}, fmt.Errorf("could not parse count '%c', not 1, 2, or 3", cardStr[1])
	}

	switch cardStr[1] {
	case red:
	case purple:
	case green:
	default:
		return card{}, fmt.Errorf("could not parse color '%c', not r, p, or g", cardStr[2])
	}

	switch cardStr[2] {
	case pen:
	case solid:
	case hollow:
	default:
		return card{}, fmt.Errorf("could not parse shade '%c', not p, s, or n", cardStr[3])
	}

	switch cardStr[3] {
	case wave:
	case pill:
	case diamond:
	default:
		return card{}, fmt.Errorf("could not parse shape '%c', not w, p, or d", cardStr[0])
	}

	return card{
		count: cardStr[0],
		color: cardStr[1],
		shade: cardStr[2],
		shape: cardStr[3],
	}, nil
}

func thirdCard(first, second card) (third card) {
	if first.shape == second.shape {
		third.shape = first.shape
	} else {
		third.shape = remaining([]byte{wave, pill, diamond}, first.shape, second.shape)
	}

	if first.count == second.count {
		third.count = first.count
	} else {
		third.count = remaining([]byte{one, two, three}, first.count, second.count)
	}

	if first.color == second.color {
		third.color = first.color
	} else {
		third.color = remaining([]byte{red, purple, green}, first.color, second.color)
	}

	if first.shade == second.shade {
		third.shade = first.shade
	} else {
		third.shade = remaining([]byte{pen, solid, hollow}, first.shade, second.shade)
	}

	return
}

func remaining(set []byte, v1, v2 byte) byte {
	for _, el := range set {
		if el != v1 && el != v2 {
			return el
		}
	}
	panic(fmt.Sprintf("set '%v' contains no alternative to %v or %v", set, v1, v2))
	return byte(0)
}
