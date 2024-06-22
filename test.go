package main

import (
	"log"
	"math/rand"
	"time"
)

// lethal hits - crit hits auto wound
// sustained hits - crit hits score additional hits
// crit hits - +/-1 modifier will not affect crit rolls

// twin linked - reroll failed wounds
// devastating wounds - crit wounds inflict mortal wounds
// anti x - crit wounds on x+
// mortal wounds

type HitRollParams struct {
	Attacks       int
	Skill         int
	LethalHits    bool
	SustainedHits bool
}

type HitRollResult struct {
	Roll         []int
	RegularHits  int
	CriticalHits int
	Wounds       int // eg from lethal hits
}

func rollDice(n int) []int {
	log.Printf("Rolling %d dice\n", n)
	rand.Seed(time.Now().UnixNano())
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = rand.Intn(6) + 1
	}
	return r
}

func hitRoll(params HitRollParams) HitRollResult {

}

func countHits(roll []int, skill) int {
	hits := 0
	for i := 0; i < len(roll); i++ {
		if roll[i] >= attack.Skill {
			hits++
		}
	}
	return hits
}
