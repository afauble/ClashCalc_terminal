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
	var percentDmg float32 = data.GetEarthquakeDmg(level) / 100
	var modifier float32 = 1.0 / (1.0 + (float32(index) * 2.0))
	var dmg float32 = maxHealth * percentDmg * modifier // index is 0 based
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
func applyMultiLtn(level int8, currentHealth float32) int {
	var count int = 0
	for currentHealth > 0 {
		currentHealth = dealLightingDmg(level, currentHealth)
		count++
	}
	return count
}

func FindOptimalSpells(maxHealth float32, userInput models.UserInput) models.CalcResult {
	var bestResults models.CalcResult
	var resultSlice [5]int
	var currentHealth float32 = maxHealth
	var lowestSpellCount int = 99

	// giantarrow damage
	if userInput.GiantarrowLvl != 0 {
		currentHealth = dealGiantarrowDmg(userInput.GiantarrowLvl, currentHealth)
		bestResults.GiantarrowUsed = true
	}

	// fireball damage
	if userInput.FireballLvl != 0 {
		currentHealth = dealFireballDmg(userInput.FireballLvl, currentHealth)
		bestResults.FireballUsed = true
	}

	// calculate zapquake at each quake amount and put in result array
	for eqUsed := 0; eqUsed < 5; eqUsed++ {
		tempHealth := currentHealth
		tempHealth = applyMultiEq(userInput.EarthquakeLvl, tempHealth, maxHealth, eqUsed)
		ltnUsed := applyMultiLtn(userInput.LightningLvl, tempHealth)
		if eqUsed+ltnUsed < lowestSpellCount {
			lowestSpellCount = eqUsed + ltnUsed
		}
		resultSlice[eqUsed] = ltnUsed
	}

	bestResults.ZapQuakeCombos = selectBestCombos(resultSlice, lowestSpellCount)
	return bestResults
}

func selectBestCombos(resultSlice [5]int, lowestSpellCount int) map[int]int {
	var results = make(map[int]int)
	for eq, lt := range resultSlice {
		if eq+lt == lowestSpellCount {
			results[eq] = lt
		}
	}
	return results
}
