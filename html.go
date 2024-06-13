package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type AttackRollResult struct {
	A Attack
	R RollResult
	J string
	D []string
}

func renderDice(result RollResult) []string {
	d := []string{}
	for i := 0; i < len(result.Roll); i++ {
		if result.Roll[i] >= result.SuccessOn {
			d = append(d, fmt.Sprintf("dice%d", result.Roll[i]))
		} else {
			d = append(d, fmt.Sprintf("dice%dred", result.Roll[i]))
		}
	}
	return d
}

func renderHitRollResults(w http.ResponseWriter, attack Attack, result RollResult) {
	t, _ := template.ParseFiles("static/html/hitRollResults.html")
	json, _ := json.Marshal(result.Roll)
	t.Execute(w, AttackRollResult{
		A: attack,
		R: result,
		J: string(json),
		D: renderDice(result),
	})
}

func renderWoundRollResults(w http.ResponseWriter, attack Attack, result RollResult) {
	t, _ := template.ParseFiles("static/html/woundRollResults.html")
	json, _ := json.Marshal(result.Roll)
	t.Execute(w, AttackRollResult{
		A: attack,
		R: result,
		J: string(json),
		D: renderDice(result),
	})
}
