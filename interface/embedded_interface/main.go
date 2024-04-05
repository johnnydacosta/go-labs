package main

import "fmt"

type Attacker interface {
	Attack(target *Character, damage int)
}

type Defenser interface {
	Defense() int
}

type Player interface {
	Attacker
	Defenser
}

type CastSpell interface {
	Spell()
}

type Swimmer interface {
	Swim()
}

type Fly interface {
	FlyAway()
}

type Puncher interface {
	Punch()
}

type Kicker interface {
	Kick()
}

type Character struct {
	Name string
	Life int
}

type Wizard struct {
	c Character
}

func (w Wizard) CastSpell() {
	fmt.Println("Cast fire !")
}

type Warrior struct {
	c Character
}

func (w Warrior) String() {
	fmt.Printf("I'm %s, I have %d of life\n", w.c.Name, w.c.Life)
}

func (w Warrior) Kick() {
	fmt.Println("Throw a kick !")
}

func (w Warrior) Punch() {
	fmt.Println("Throw a punch !")
}

func (w Warrior) Swim() {
	fmt.Printf("%s jump in the water !\n", w.c.Name)
}

type Lezard struct {
	c Character
}

func (l Lezard) Swim() {
	fmt.Printf("%s jump in the water !\n", l.c.Name)
}

type Angel struct {
	c Character
}

func (a Angel) FlyAway() {
	fmt.Println("Flay away !!")
}

func main() {
	wizzard := Wizard{
		c: Character{Name: "The Wizzard", Life: 50},
	}
	warrior := Warrior{
		c: Character{Name: "The Warrior", Life: 50},
	}

	lezard := Lezard{
		c: Character{Name: "The Lezard", Life: 50},
	}
	angel := Angel{
		c: Character{Name: "The Angel", Life: 50},
	}

	wizzard.CastSpell()
	warrior.Kick()
	warrior.Punch()
	lezard.Swim()
	angel.FlyAway()

	var swimmers []Swimmer
	swimmers = append(swimmers, lezard)
	swimmers = append(swimmers, warrior)
}
