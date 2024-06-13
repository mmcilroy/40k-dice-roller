package main

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type Attack struct {
	Attacks         int
	Skill           int
	Strength        int
	Toughness       int
	Penetration     int
	Damage          int
	WoundOn         int
	CriticalHitOn   int
	CriticalWoundOn int
}

type RollResult struct {
	Roll      []int
	SuccessOn int
	Successes int
	Criticals int
}

type countDieFunc func(Attack, []int) int

func logAttackRoll(attack Attack) {
	log.Printf("Hit Roll - %d attacks hitting on %d+\n", attack.Attacks, attack.Skill)
}

func logWoundRoll(numAttacks int, attack Attack) {
	log.Printf("Wound Roll - %d attacks hitting on %d+\n", numAttacks, attack.WoundOn)
}

func logRollResult(rollResult RollResult) {
	var sb strings.Builder
	for i := 0; i < len(rollResult.Roll); i++ {
		sb.WriteString(strconv.Itoa(rollResult.Roll[i]))
		sb.WriteString(" ")
	}
	log.Printf("RollResult - Roll: %s, SuccessOn: %d, Successes: %d, Criticals: %d\n", sb.String(),
		rollResult.SuccessOn,
		rollResult.Successes,
		rollResult.Criticals)
}

func countSuccesfulHits(attack Attack, roll []int) int {
	hits := 0
	for i := 0; i < len(roll); i++ {
		if roll[i] >= attack.Skill {
			hits++
		}
	}
	return hits
}

func woundRollNeeded(strength int, toughness int) int {
	if strength >= 2*toughness {
		// Strength is TWICE (or more than twice) the Toughness
		return 2
	} else if toughness >= 2*strength {
		// Strength is HALF (or less than half) the Toughness
		return 6
	} else if strength > toughness {
		// Strength is GREATER than the Toughness
		return 3
	} else if strength == toughness {
		// Strength is EQUAL to the Toughness
		return 4
	} else if strength < toughness {
		// Strength is LESS than the Toughness
		return 5
	}
	return 7
}

func countSuccesfulWounds(attack Attack, roll []int) int {
	hits := 0
	for i := 0; i < len(roll); i++ {
		if roll[i] >= attack.WoundOn {
			hits++
		}
	}
	return hits
}

func countCriticalHits(attack Attack, roll []int) int {
	hits := 0
	for i := 0; i < len(roll); i++ {
		if roll[i] >= attack.CriticalHitOn {
			hits++
		}
	}
	return hits
}

func countCriticalWounds(attack Attack, roll []int) int {
	hits := 0
	for i := 0; i < len(roll); i++ {
		if roll[i] >= attack.CriticalWoundOn {
			hits++
		}
	}
	return hits
}

func makeRoll(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = rand.Intn(6) + 1
	}
	return r
}

func makeRollResult(attack Attack, roll []int, successOn int, successes countDieFunc, criticals countDieFunc) RollResult {
	rollResult := RollResult{
		Roll:      roll,
		SuccessOn: successOn,
		Successes: successes(attack, roll),
		Criticals: criticals(attack, roll),
	}
	return rollResult
}

func rerollResultDie(attack Attack, roll []int, die int, successOn int, successes countDieFunc, criticals countDieFunc) RollResult {
	roll[die] = makeRoll(1)[0]
	rollResult := RollResult{
		Roll:      roll,
		SuccessOn: successOn,
		Successes: successes(attack, roll),
		Criticals: criticals(attack, roll),
	}
	return rollResult
}

func modifyResultDie(attack Attack, roll []int, change int, successOn int, successes countDieFunc, criticals countDieFunc) RollResult {
	for i := 0; i < len(roll); i++ {
		die := roll[i]
		die += change
		if die < 1 {
			die = 1
		} else if die > 6 {
			die = 6
		}
		roll[i] = die
	}
	rollResult := RollResult{
		Roll:      roll,
		SuccessOn: successOn,
		Successes: successes(attack, roll),
		Criticals: criticals(attack, roll),
	}
	return rollResult
}

func rerollFailedResultDie(attack Attack, roll []int, successOn int, successes countDieFunc, criticals countDieFunc) RollResult {
	for i := 0; i < len(roll); i++ {
		if roll[i] < successOn {
			roll[i] = makeRoll(1)[0]
		}
	}
	rollResult := RollResult{
		Roll:      roll,
		SuccessOn: successOn,
		Successes: successes(attack, roll),
		Criticals: criticals(attack, roll),
	}
	return rollResult
}

/*
func main() {
	attack := Attack{
		Attacks:         10,
		Skill:           3,
		Strength:        3,
		Toughness:       3,
		Penetration:     1,
		Damage:          1,
		CriticalHitOn:   6,
		CriticalWoundOn: 6,
		WoundOn:         woundRollNeeded(4, 2),
	}

	logAttackRoll(attack)
	roll := makeRoll(attack.Attacks)

	result := makeRollResult(attack, roll, countSuccesfulHits, countCriticalHits)
	logRollResult(result)

	roll[0]++
	result = makeRollResult(attack, roll, countSuccesfulHits, countCriticalHits)
	logRollResult(result)

	logWoundRoll(result.Successes, attack)
	roll = makeRoll(result.Successes)
	result = makeRollResult(attack, roll, countSuccesfulWounds, countCriticalWounds)
	logRollResult(result)
}
*/
