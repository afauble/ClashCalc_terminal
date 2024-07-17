package data

import (
	"fmt"
	"os"
)

var (
	lightingDmg   [12]float32
	earthquakeDmg [6]float32
	giantarrowDmg [19]float32
	fireballDmg   [28]float32

	cannonHealth           [21]float32
	archerTowerHealth      [21]float32
	mortarHealth           [16]float32
	airDefenseHealth       []float32
	wizardTowerHealth      []float32
	airSweeperHealth       []float32
	hiddenTeslaHealth      []float32
	bombTowerHealth        []float32
	xBowHealth             []float32
	infernoTowerHealth     []float32
	eagleArtilleryHealth   []float32
	scattershotHealth      []float32
	builderHutHealth       []float32
	spellTowerHealth       []float32
	monolithHealth         []float32
	multiArcherTowerHealth []float32
	ricochetCannonHealth   []float32
	townHallHealth         []float32
	clanCastleHealth       []float32

	getFuncMap map[string]func(int8) float32
)

func init() {
	// Spells & Equipment
	lightingDmg = [12]float32{0, 150, 180, 210, 240, 270, 320, 400, 480, 560, 600, 640}
	earthquakeDmg = [6]float32{0, 14.5, 17, 21, 25, 29}
	giantarrowDmg = [19]float32{0, 750, 750, 850, 850, 850, 1000, 1000, 1000, 1200, 1200, 1200, 1500, 1500, 1500, 1750, 1750, 1750, 1950}
	fireballDmg = [28]float32{0, 1500, 1500, 1700, 1700, 1800, 1950, 1950, 2050, 2200, 2200, 2350, 2650, 2650, 2750, 3100, 3100, 3250, 3400, 3400, 3500, 3650, 3650, 3750, 3900, 3900, 3950, 4100}

	// Buildings
	cannonHealth = [21]float32{420, 470, 520, 570, 620, 670, 730, 800, 880, 960, 1060, 1160, 1260, 1380, 1500, 1620, 1740, 1870, 2000, 2150, 2250}
	archerTowerHealth = [21]float32{380, 420, 460, 500, 540, 580, 630, 690, 750, 810, 890, 970, 1050, 1130, 1230, 1310, 1390, 1510, 1600, 1700, 1800}
	mortarHealth = [16]float32{400, 450, 500, 550, 600, 650, 700, 800, 950, 1100, 1300, 1500, 1700, 1950, 2150, 2300}
	airDefenseHealth = []float32{}
	wizardTowerHealth = []float32{}
	airSweeperHealth = []float32{}
	hiddenTeslaHealth = []float32{}
	bombTowerHealth = []float32{}
	xBowHealth = []float32{}
	infernoTowerHealth = []float32{}
	eagleArtilleryHealth = []float32{}
	scattershotHealth = []float32{}
	builderHutHealth = []float32{}
	spellTowerHealth = []float32{}
	monolithHealth = []float32{}
	multiArcherTowerHealth = []float32{}
	ricochetCannonHealth = []float32{}
	townHallHealth = []float32{}
	clanCastleHealth = []float32{}

	// Get Func Map
	getFuncMap = map[string]func(int8) float32{
		"cannonHealth":           GetCannonHealth,
		"archerTowerHealth":      GetArcherTowerHealth,
		"mortarHealth":           GetMortarHealth,
		"airDefenseHealth":       GetAirDefenseHealth,
		"wizardTowerHealth":      GetWizardTowerHealth,
		"airSweeperHealth":       GetAirSweeperHealth,
		"hiddenTeslaHealth":      GetHiddenTeslaHealth,
		"bombTowerHealth":        GetBombTowerHealth,
		"xBowHealth":             GetXBowHealth,
		"infernoTowerHealth":     GetInfernoTowerHealth,
		"eagleArtilleryHealth":   GetEagleArtilleryHealth,
		"scattershotHealth":      GetScattershotHealth,
		"builderHutHealth":       GetBuilderHutHealth,
		"spellTowerHealth":       GetSpellTowerHealth,
		"monolithHealth":         GetMonolithHealth,
		"multiArcherTowerHealth": GetMultiArcherTowerHealth,
		"ricochetCannonHealth":   GetRicochetCannonHealth,
		"townHallHealth":         GetTownHallHealth,
		"clanCastleHealth":       GetClanCastleHealth,
	}
}

func validateLength(level int8, length []float32) {
	if int(level) >= len(length) || level < 0 {
		fmt.Println("error: validateLength failed")
		os.Exit(1)
	}
}

// Spell & Equipment Damage Getters
func GetLightingDmg(level int8) float32 {
	validateLength(level, lightingDmg[:])
	return lightingDmg[level]
}
func GetEarthquakeDmg(level int8) float32 {
	validateLength(level, earthquakeDmg[:])
	return earthquakeDmg[level]
}
func GetGiantarrowDmg(level int8) float32 {
	validateLength(level, giantarrowDmg[:])
	return giantarrowDmg[level]
}
func GetFireballDmg(level int8) float32 {
	validateLength(level, fireballDmg[:])
	return fireballDmg[level]
}

// Building Health Getters
func GetCannonHealth(level int8) float32 {
	validateLength(level, cannonHealth[:])
	return cannonHealth[level]
}
func GetArcherTowerHealth(level int8) float32 {
	validateLength(level, archerTowerHealth[:])
	return archerTowerHealth[level]
}
func GetMortarHealth(level int8) float32 {
	validateLength(level, mortarHealth[:])
	return mortarHealth[level]
}
func GetAirDefenseHealth(level int8) float32 {
	validateLength(level, airDefenseHealth[:])
	return airDefenseHealth[level]
}
func GetWizardTowerHealth(level int8) float32 {
	validateLength(level, wizardTowerHealth[:])
	return wizardTowerHealth[level]
}
func GetAirSweeperHealth(level int8) float32 {
	validateLength(level, airSweeperHealth[:])
	return airSweeperHealth[level]
}
func GetHiddenTeslaHealth(level int8) float32 {
	validateLength(level, hiddenTeslaHealth[:])
	return hiddenTeslaHealth[level]
}
func GetBombTowerHealth(level int8) float32 {
	validateLength(level, bombTowerHealth[:])
	return bombTowerHealth[level]
}
func GetXBowHealth(level int8) float32 {
	validateLength(level, xBowHealth[:])
	return xBowHealth[level]
}
func GetInfernoTowerHealth(level int8) float32 {
	validateLength(level, infernoTowerHealth[:])
	return infernoTowerHealth[level]
}
func GetEagleArtilleryHealth(level int8) float32 {
	validateLength(level, eagleArtilleryHealth[:])
	return eagleArtilleryHealth[level]
}
func GetScattershotHealth(level int8) float32 {
	validateLength(level, scattershotHealth[:])
	return scattershotHealth[level]
}
func GetBuilderHutHealth(level int8) float32 {
	validateLength(level, builderHutHealth[:])
	return builderHutHealth[level]
}
func GetSpellTowerHealth(level int8) float32 {
	validateLength(level, spellTowerHealth[:])
	return spellTowerHealth[level]
}
func GetMonolithHealth(level int8) float32 {
	validateLength(level, monolithHealth[:])
	return monolithHealth[level]
}
func GetMultiArcherTowerHealth(level int8) float32 {
	validateLength(level, multiArcherTowerHealth[:])
	return multiArcherTowerHealth[level]
}
func GetRicochetCannonHealth(level int8) float32 {
	validateLength(level, ricochetCannonHealth[:])
	return ricochetCannonHealth[level]
}
func GetTownHallHealth(level int8) float32 {
	validateLength(level, townHallHealth[:])
	return townHallHealth[level]
}
func GetClanCastleHealth(level int8) float32 {
	validateLength(level, clanCastleHealth[:])
	return clanCastleHealth[level]
}

// Converting strings to the correct get func
func GetCorrespondingHealthFunc(name string) func(int8) float32 {
	return getFuncMap[name]
}
