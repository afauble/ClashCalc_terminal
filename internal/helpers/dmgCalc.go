package helpers

// Delete and replace with appropriate get method
func getDmgNmbr(level int8) float32 {
	return float32(level)
}

type UserInput struct {
	giantarrowLvl int8
	fireballLvl   int8
	lightningLvl  int8
	earthquakeLvl int8
}

func DealLightingDmg(level int8, currentHealth float32) float32 {
	currentHealth = currentHealth - getDmgNmbr(level)
	if currentHealth > 0 {
		return currentHealth
	}
	return 0
}

func DealEarthquakeDmg(level int8, currentHealth float32, maxHealth float32, index int) float32 {
	percentDmg := getDmgNmbr(level) / 100
	dmg := maxHealth * percentDmg * float32(1/(1+(index*2))) // index is 0 based
	currentHealth = currentHealth - dmg
	if currentHealth > 0 {
		return currentHealth
	}
	return 0
}

// Deals the damage of count # of earthquake spells and returns the remaining health
func ApplyMultiEq(level int8, currentHealth float32, maxHealth float32, count int) float32 {
	for i := 0; i < count; i++ {
		currentHealth = DealEarthquakeDmg(level, currentHealth, maxHealth, i)
	}
	return currentHealth
}

func ApplyMultiLtn(level int8, currentHealth float32) int8 {
	var count int8 = 0
	for {
		if currentHealth < 0 {
			break
		}
		currentHealth = DealLightingDmg(level, currentHealth)
	}
	return count
}

func FindOptimalSpells(maxHealth float32, userInput UserInput) [5]int8 {
	var resultSlice [5]int8
	var currentHealth float32 = maxHealth

	currentHealth = currentHealth - getDmgNmbr(userInput.giantarrowLvl) // giantarrow
	currentHealth = currentHealth - getDmgNmbr(userInput.fireballLvl)   // fireball

	for eqUsed := 0; eqUsed < 5; eqUsed++ {
		currentHealth = ApplyMultiEq(userInput.earthquakeLvl, currentHealth, maxHealth, eqUsed)
		ltnUsed := ApplyMultiLtn(userInput.lightningLvl, currentHealth)
		resultSlice[eqUsed] = ltnUsed
	}

	return resultSlice
}
