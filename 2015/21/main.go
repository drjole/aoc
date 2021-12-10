package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(first())
	fmt.Println(second())
}

func first() int {
	equipments := itemCombinations()

	minCost := math.MaxInt
	for _, equipment := range equipments {
		cost := 0
		me := &player{100, 0, 0}
		for _, i := range equipment {
			cost += i.cost
			me.damage += i.damage
			me.armor += i.armor
		}

		if cost >= minCost {
			continue
		}

		winner := play(me, boss())
		if winner == me {
			minCost = cost
		}
	}

	return minCost
}

func second() int {
	equipments := itemCombinations()

	maxCost := 0
	for _, equipment := range equipments {
		cost := 0
		me := &player{100, 0, 0}
		for _, i := range equipment {
			cost += i.cost
			me.damage += i.damage
			me.armor += i.armor
		}

		if cost <= maxCost {
			continue
		}

		winner := play(me, boss())
		if winner != me {
			maxCost = cost
		}
	}

	return maxCost
}

type player struct {
	hitPoints, damage, armor int
}

func play(me, boss *player) *player {
	attacker, defender := me, boss
	for {
		if attacker.hitPoints <= 0 {
			return defender
		}
		if defender.hitPoints <= 0 {
			return attacker
		}

		damage := attacker.damage - defender.armor
		if damage <= 0 {
			damage = 1
		}

		defender.hitPoints -= damage
		attacker, defender = defender, attacker
	}
}

func itemCombinations() (combinations [][]item) {
	for w := 0; w < len(weapons); w++ {
		for a := 0; a < 1+len(armor); a++ {
			for r1 := 0; r1 < 1+len(rings); r1++ {
				for r2 := 0; r2 < 1+len(rings); r2++ {
					equipment := make([]item, 0)
					equipment = append(equipment, weapons[w])
					if a != 0 {
						equipment = append(equipment, armor[a-1])
					}
					if r1 != 0 {
						equipment = append(equipment, rings[r1-1])
					}
					if r2 != 0 {
						equipment = append(equipment, rings[r2-1])
					}
					combinations = append(combinations, equipment)
				}
			}
		}
	}
	return
}

var bossHitPoints, bossDamage, bossArmor int

func init() {
	inputBytes, err := os.ReadFile("2015/21/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputBytes), "\n")

	bossHitPoints, _ = strconv.Atoi(strings.Fields(lines[0])[2])
	bossDamage, _ = strconv.Atoi(strings.Fields(lines[1])[1])
	bossArmor, _ = strconv.Atoi(strings.Fields(lines[2])[1])
}

func boss() *player {
	return &player{bossHitPoints, bossDamage, bossArmor}
}

type item struct {
	cost, damage, armor int
}

const (
	weaponsString = `Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0`
	armorString = `Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5`
	ringsString = `Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`
)

var (
	weapons, armor, rings []item
)

func init() {
	for _, line := range strings.Split(weaponsString, "\n") {
		weapons = append(weapons, parseItem(strings.Fields(line)))
	}
	for _, line := range strings.Split(armorString, "\n") {
		armor = append(armor, parseItem(strings.Fields(line)))
	}
	for _, line := range strings.Split(ringsString, "\n") {
		fields := strings.Fields(line)
		rings = append(rings, parseItem(fields[1:]))
	}
}

func parseItem(fields []string) item {
	c, _ := strconv.Atoi(fields[1])
	d, _ := strconv.Atoi(fields[2])
	a, _ := strconv.Atoi(fields[3])
	return item{c, d, a}
}
