package main

import (
	"fmt"
	"math/rand"
)

type Player string

type Team struct {
	Name    string
	Players []Player
}

type TeamScore struct {
	Name  string
	Score int
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(teamA string, scoreA int, teamB string, scoreB int) {
	updateWin := func(team string) {
		l.Wins[team]++
	}

	if scoreA > scoreB {
		updateWin(teamA)
		return
	}

	if scoreA < scoreB {
		updateWin(teamB)
		return
	}
}

func (l *League) Ranking() []string {
	var teams []TeamScore

	for key, value := range l.Wins {
		teams = append(teams, TeamScore{Name: key, Score: value})
	}

	var teamNameSortedByWin []string
	for _, t := range QuickSort(teams) {
		teamNameSortedByWin = append(teamNameSortedByWin, t.Name)
	}

	return teamNameSortedByWin
}

func QuickSort(teams []TeamScore) []TeamScore {
	i := 1
	for i < len(teams) {
		x := teams[i]
		j := i
		for j > 0 && teams[j-1].Score < x.Score {
			teams[j] = teams[j-1]
			j--
		}

		teams[j] = x
		i++
	}
	return teams
}

func main() {
	laker := Team{Name: "Laker", Players: []Player{"Lebron", "Davis", "Reaves"}}
	bulls := Team{Name: "Bulls", Players: []Player{"Ball", "DeRozan", "Terry"}}
	warrior := Team{Name: "Warrior", Players: []Player{"Green", "Thomson", "Reaves"}}

	nba := League{
		Teams: []Team{laker, bulls, warrior},
		Wins:  make(map[string]int),
	}

	totalMatch := 20

	makeMatch := func(teamA, teamB *Team) {
		scoreA, scoreB := rand.Intn(100), rand.Intn(100)
		nba.MatchResult(teamA.Name, scoreA, teamB.Name, scoreB)
	}

	for i := 0; i < totalMatch; i++ {
		makeMatch(&laker, &bulls)
		makeMatch(&laker, &warrior)

		makeMatch(&bulls, &laker)
		makeMatch(&bulls, &warrior)

		makeMatch(&warrior, &laker)
		makeMatch(&warrior, &warrior)

	}

	fmt.Println("Result of the league:")
	for i, t := range nba.Ranking() {
		wins, ok := nba.Wins[t]
		if ok == false {
			fmt.Printf("Team %s do not exist.\n", t)
			continue
		}
		fmt.Printf("Rank %d  %s with %d wins\n", i+1, t, wins)
	}
}
