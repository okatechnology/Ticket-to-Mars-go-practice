package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	const SecondsInDay = 86400
	const Distance = 62100000
	const MaxSpeed = 30
	const MinSpeed = 16
	spaceline := ""
	tripType := ""

	fmt.Printf("%v\n%v\n", "Spaceline        Days Trip type  Price", "======================================")

	for count := 0; count < 10; count++ {
		speed := rand.Intn(MaxSpeed-(MinSpeed-1)) + MinSpeed
		price := speed + 20
		days := Distance / speed / SecondsInDay

		switch spacelineNum := rand.Intn(3); spacelineNum {
		case 0:
			spaceline = "Virgin Galactic"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Space Adventures"
		}

		switch tripTypeNum := rand.Intn(2); tripTypeNum {
		case 0:
			tripType = "One-way"
		case 1:
			tripType = "Round-trip"
			price *= 2
			days *= 2
		}
		fmt.Printf("%-16v %4v %-10v $ %3v\n", spaceline, strconv.Itoa(days), tripType, strconv.Itoa(price))
	}
}
