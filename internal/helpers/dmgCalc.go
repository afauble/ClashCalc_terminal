package helpers

import (
	"github.com/afauble/ClashCalc/internal/data"
	"github.com/afauble/ClashCalc/internal/models"
)

func dealLightingDmg(level int8, currentHealth float32) float32 {
	currentHealth = currentHealth - data.GetLightingDmg(level)
	if currentHealth > 0 {
		return currentHealth
	}
	return 0
}

func dealFireballDmg(level int8, currentHealth float32) float32 {
	currentHealth = currentHealth - data.GetFireballDmg(level)
	if currentHealth > 0 {
		return currentHealth
	}
	return 0
}

func dealGiantarrowDmg(level int8, currentHealth float32) float32 {
	currentHealth = currentHealth - data.GetGiantarrowDmg(level)
	if currentHealth > 0 {
		return currentHealth
	}
	return 0
}

func dealEarthquakeDmg(level int8, currentHealth float32, maxHealth float32, index int) float32 {
	percentDmg := data.GetEarthquakeDmg(level) / 100
	dmg := maxHealth * percentDmg * float32(1/(1+(index*2))) // index is 0 based
	currentHealth = currentHealth - dmg
	if currentHealth > 0 {
		return currentHealth
	}
	return 0
}

// Deals the damage of count # of earthquake spells and returns the remaining health
func applyMultiEq(level int8, currentHealth float32, maxHealth float32, count int) float32 {
	for i := 0; i < count; i++ {
		currentHealth = dealEarthquakeDmg(level, currentHealth, maxHealth, i)
	}
	return currentHealth
}

// Determines the # of lightning spells needed to kill at the given health
func applyMultiLtn(level int8, currentHealth float32) int8 {
	var count int8 = 0
	for currentHealth > 0 {
		count++
		currentHealth = dealLightingDmg(level, currentHealth)
	}
	return count
}

func FindOptimalSpells(maxHealth float32, userInput models.UserInput) [5]int8 {
	var resultSlice [5]int8
	var currentHealth float32 = maxHealth

	// giantarrow damage
	currentHealth = dealGiantarrowDmg(userInput.GiantarrowLvl, currentHealth)

	// fireball damage
	currentHealth = dealFireballDmg(userInput.FireballLvl, currentHealth)

	// calculate zapquake at each quake amount and put in result array
	for eqUsed := 0; eqUsed < 5; eqUsed++ {
		currentHealth = applyMultiEq(userInput.EarthquakeLvl, currentHealth, maxHealth, eqUsed)
		ltnUsed := applyMultiLtn(userInput.LightningLvl, currentHealth)
		resultSlice[eqUsed] = ltnUsed
	}

	return resultSlice
}
