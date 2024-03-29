package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type ThermicEngine struct{}

func (t ThermicEngine) Start() {
	fmt.Println("Starting thermic engine...")
	fmt.Println("Thermic engine started.")
}

func (t ThermicEngine) Stop() {
	fmt.Println("Stopping thermic engine...")
	fmt.Println("Thermic engine stopped.")
}

type HybridEngine struct{}

func (h HybridEngine) Start() {
	fmt.Println("Starting hybrid engine...")
	fmt.Println("Hybrid engine started.")
}

func (h HybridEngine) Stop() {
	fmt.Println("Stopping hybrid engine...")
	fmt.Println("Hybrid engine stopped.")
}

type ElectricEngine struct{}

func (e ElectricEngine) Start() {
	fmt.Println("Starting electric engine...")
	fmt.Println("Thermic engine started.")
}

func (e ElectricEngine) Stop() {
	fmt.Println("Stopping electric engine...")
	fmt.Println("Electric engine stopped.")
}

type Car struct {
	E Engine
}

func main() {
	car := Car{E: ThermicEngine{}}
	car.E.Start()
	car.E.Stop()

	car.E = HybridEngine{}
	car.E.Start()
	car.E.Stop()

	car.E = ElectricEngine{}
	car.E.Start()
	car.E.Stop()
}
