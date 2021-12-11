package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	me := &player{}
	spells := []*spell{
		{me.magicMissile, 53},
		{me.drain, 73},
		{me.shield, 113},
		{me.poison, 173},
		{me.recharge, 229},
	}

	random := func() *spell {
		invalid := make(map[int]struct{})
		for {
			if len(invalid) == len(spells) {
				return nil
			}
			i := rand.Intn(len(spells))
			sp := spells[i]
			if me.mana <= sp.mana ||
				(i == 2 && me.shieldTimer > 0) ||
				(i == 3 && me.poisonTimer > 0) ||
				(i == 4 && me.rechargeTimer > 0) {
				invalid[i] = struct{}{}
				continue
			}
			return sp
		}
	}

	fmt.Println(applyStrategy(me, random, time.Second, false))
	fmt.Println(applyStrategy(me, random, time.Second, true))
}

func applyStrategy(me *player, strategy func() *spell, timeout time.Duration, hard bool) (minMana int) {
	minMana = math.MaxInt
	t := time.After(timeout)
	for {
		select {
		case <-t:
			return
		default:
			me.reset()
			b := boss()
			winner := play(me, b, strategy, minMana, hard)
			if winner == me && me.manaSpent < minMana {
				minMana = me.manaSpent
			}
		}
	}
}

func play(me, boss *player, strategy func() *spell, minMana int, hard bool) *player {
	for {
		// my turn
		if hard {
			me.hitPoints--
			if me.hitPoints <= 0 {
				return boss
			}
		}
		s := strategy()
		// No spell available
		if s == nil {
			return boss
		}
		// Cast spell
		s.cast(me, boss)
		// Return early if we already spent too much mana
		if me.manaSpent > minMana {
			return boss
		}
		// Check if I won
		if boss.hitPoints <= 0 {
			return me
		}
		effects(me, boss)
		// Check if boss won
		if boss.hitPoints <= 0 {
			return me
		}

		// Boss' turn
		// Boss attacks
		damage := boss.damage - me.armor()
		if damage < 1 {
			damage = 1
		}
		me.hitPoints -= damage
		// Check if boss won
		if me.hitPoints <= 0 {
			return boss
		}
		effects(me, boss)
		// Check if I won
		if boss.hitPoints <= 0 {
			return me
		}
	}
}

func effects(me, boss *player) {
	if me.poisonTimer > 0 {
		boss.hitPoints -= 3
		me.poisonTimer--
	}
	if me.rechargeTimer > 0 {
		me.mana += 101
		me.rechargeTimer--
	}
	if me.shieldTimer > 0 {
		me.shieldTimer--
	}
}

var bossHitPoints, bossDamage int

func init() {
	inputBytes, err := os.ReadFile("2015/22/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputBytes), "\n")
	bossHitPointsString := strings.Fields(lines[0])[2]
	bossDamageString := strings.Fields(lines[1])[1]

	bossHitPoints, _ = strconv.Atoi(bossHitPointsString)
	bossDamage, _ = strconv.Atoi(bossDamageString)
}

func boss() *player {
	return &player{
		hitPoints: bossHitPoints,
		damage:    bossDamage,
	}
}

type player struct {
	hitPoints, damage, mana                 int
	shieldTimer, poisonTimer, rechargeTimer int
	manaSpent                               int
}

type spell struct {
	f    func(*player)
	mana int
}

func (s *spell) cast(caster, victim *player) {
	caster.mana -= s.mana
	caster.manaSpent += s.mana
	s.f(victim)
}

func (p *player) reset() {
	p.hitPoints = 50
	p.damage = 0
	p.mana = 500
	p.shieldTimer = 0
	p.poisonTimer = 0
	p.rechargeTimer = 0
	p.manaSpent = 0
}

func (p *player) armor() int {
	if p.shieldTimer > 0 {
		return 7
	}
	return 0
}

func (p *player) magicMissile(other *player) {
	other.hitPoints -= 4
}

func (p *player) drain(other *player) {
	other.hitPoints -= 2
	p.hitPoints += 2
}

func (p *player) shield(*player) {
	p.shieldTimer = 6
}

func (p *player) poison(*player) {
	p.poisonTimer = 6
}

func (p *player) recharge(*player) {
	p.rechargeTimer = 5
}
