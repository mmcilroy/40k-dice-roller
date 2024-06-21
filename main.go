package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func parseAttack(r *http.Request) Attack {
	a, _ := strconv.Atoi(r.PostFormValue("A"))
	ws, _ := strconv.Atoi(r.PostFormValue("WS"))
	s, _ := strconv.Atoi(r.PostFormValue("S"))
	t, _ := strconv.Atoi(r.PostFormValue("T"))
	ap, _ := strconv.Atoi(r.PostFormValue("AP"))
	d, _ := strconv.Atoi(r.PostFormValue("D"))
	ch, _ := strconv.Atoi(r.PostFormValue("CH"))
	cw, _ := strconv.Atoi(r.PostFormValue("CW"))

	return Attack{
		Attacks:         a,
		Skill:           ws,
		Strength:        s,
		Toughness:       t,
		Penetration:     ap,
		Damage:          d,
		CriticalHitOn:   ch,
		CriticalWoundOn: cw,
		WoundOn:         woundRollNeeded(s, t),
	}
}

func parseRollResult(r *http.Request) RollResult {
	rr := parseRoll(r)
	rs, _ := strconv.Atoi(r.PostFormValue("RS"))
	rc, _ := strconv.Atoi(r.PostFormValue("RC"))

	return RollResult{
		Roll:      rr,
		Successes: rs,
		Criticals: rc,
	}
}

func parseRoll(r *http.Request) []int {
	var roll []int
	json.Unmarshal([]byte(r.PostFormValue("R")), &roll)
	return roll
}

func parseDie(r *http.Request) int {
	die, _ := strconv.Atoi(r.PostFormValue("die"))
	return die
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("/index")
	index, _ := os.ReadFile("static/html/index.html")
	fmt.Fprintln(w, string(index))
}

func handleHitRoll(w http.ResponseWriter, r *http.Request) {
	log.Println("/hitroll")
	attack := parseAttack(r)
	roll := makeRoll(attack.Attacks)
	result := makeRollResult(attack, roll, attack.Skill, countSuccesfulHits, countCriticalHits)
	renderHitRollResults(w, attack, result)
}

func handleHitReroll(w http.ResponseWriter, r *http.Request) {
	log.Println("/hitreroll")
	attack := parseAttack(r)
	roll := parseRoll(r)
	die := parseDie(r)
	result := rerollResultDie(attack, roll, die, attack.Skill, countSuccesfulHits, countCriticalHits)
	renderHitRollResults(w, attack, result)
}

func handleHitMinus1(w http.ResponseWriter, r *http.Request) {
	log.Println("/hitminus1")
	attack := parseAttack(r)
	roll := parseRoll(r)
	result := modifyResultDie(attack, roll, -1, attack.Skill, countSuccesfulHits, countCriticalHits)
	renderHitRollResults(w, attack, result)
}

func handleHitPlus1(w http.ResponseWriter, r *http.Request) {
	log.Println("/hitplus1")
	attack := parseAttack(r)
	roll := parseRoll(r)
	result := modifyResultDie(attack, roll, +1, attack.Skill, countSuccesfulHits, countCriticalHits)
	renderHitRollResults(w, attack, result)
}

func handleHitRerollFailed(w http.ResponseWriter, r *http.Request) {
	log.Println("/hitrerollfailed")
	attack := parseAttack(r)
	roll := parseRoll(r)
	result := rerollFailedResultDie(attack, roll, attack.Skill, countSuccesfulHits, countCriticalHits)
	renderHitRollResults(w, attack, result)
}

func handleWoundRoll(w http.ResponseWriter, r *http.Request) {
	log.Println("/woundroll")
	attack := parseAttack(r)
	hitResult := parseRollResult(r)
	roll := makeRoll(hitResult.Successes)
	woundResult := makeRollResult(attack, roll, attack.WoundOn, countSuccesfulWounds, countCriticalWounds)
	renderWoundRollResults(w, attack, woundResult)
}

func handleWoundReroll(w http.ResponseWriter, r *http.Request) {
	log.Println("/woundreroll")
	attack := parseAttack(r)
	roll := parseRoll(r)
	die := parseDie(r)
	result := rerollResultDie(attack, roll, die, attack.WoundOn, countSuccesfulWounds, countCriticalWounds)
	renderWoundRollResults(w, attack, result)
}

func handleWoundMinus1(w http.ResponseWriter, r *http.Request) {
	log.Println("/woundminus1")
	attack := parseAttack(r)
	roll := parseRoll(r)
	result := modifyResultDie(attack, roll, -1, attack.Skill, countSuccesfulWounds, countCriticalWounds)
	renderWoundRollResults(w, attack, result)
}

func handleWoundPlus1(w http.ResponseWriter, r *http.Request) {
	log.Println("/woundplus1")
	attack := parseAttack(r)
	roll := parseRoll(r)
	result := modifyResultDie(attack, roll, +1, attack.Skill, countSuccesfulWounds, countCriticalWounds)
	renderWoundRollResults(w, attack, result)
}

func handleWoundRerollFailed(w http.ResponseWriter, r *http.Request) {
	log.Println("/woundrerollfailed")
	attack := parseAttack(r)
	roll := parseRoll(r)
	result := rerollFailedResultDie(attack, roll, attack.Skill, countSuccesfulWounds, countCriticalWounds)
	renderWoundRollResults(w, attack, result)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/hitroll/", handleHitRoll)
	http.HandleFunc("/hitreroll/", handleHitReroll)
	http.HandleFunc("/hitminus1/", handleHitMinus1)
	http.HandleFunc("/hitplus1/", handleHitPlus1)
	http.HandleFunc("/hitrerollfailed/", handleHitRerollFailed)

	http.HandleFunc("/woundroll/", handleWoundRoll)
	http.HandleFunc("/woundreroll/", handleWoundReroll)
	http.HandleFunc("/woundminus1/", handleWoundMinus1)
	http.HandleFunc("/woundplus1/", handleWoundPlus1)
	http.HandleFunc("/woundrerollfailed/", handleWoundRerollFailed)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
