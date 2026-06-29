package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

func RollDice() uint8 {
	return uint8(rand.Intn(6)+1)
}

func Modifier(score int) int {
    calculated := float64(score - 10) / 2.0
	return int(math.Floor(calculated))
}

func Ability() int {
	nums := make([]int, 4)
	for i:= range nums{
		nums[i] = int(RollDice())
	}
	min := slices.Min(nums)

	excluded := false

	sum := 0

	for _ , v := range nums {
		if ! excluded {
			if v == min {
				excluded = true
				continue
			}
		}
		sum += v
	}
	return sum
}

func GenerateCharacter() Character {
	constit := Ability()
	return Character{
		Strength: Ability(),
		Dexterity: Ability(),
		Constitution: constit,
		Intelligence: Ability(),
		Wisdom: Ability(),
		Charisma: Ability(),
		Hitpoints: Modifier(constit) + 10,
	}
}
