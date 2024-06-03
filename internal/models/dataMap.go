package models

var lightingDmg [12]float64
var earthquakeDmg [6]float64
var giantarrowDmg [19]float64
var fireballDmg [28]float64

var cannonHealth [21]float64
var archerTowerHealth [21]float64
var mortarHealth [16]float64
var airDefenseHealth []float64
var wizardTowerHealth []float64
var airSweeperHealth []float64
var hiddenTeslaHealth []float64
var bombTowerHealth []float64
var xBowHealth []float64
var infernoTowerHealth []float64
var eagleArtilleryHealth []float64
var scattershotHealth []float64
var builderHuntHealth []float64
var spellTowerHealth []float64
var monolithHealth []float64
var multiArcherTowerHealth []float64
var ricochetCannonHealth []float64
var townHallHealth []float64

func init() {

	// Spells & Abilities
	lightingDmg = [12]float64{0, 150, 180, 210, 240, 270, 320, 400, 480, 560, 600, 640}
	earthquakeDmg = [6]float64{0, 14.5, 17, 21, 25, 29}
	giantarrowDmg = [19]float64{0, 750, 750, 850, 850, 850, 1000, 1000, 1000, 1200, 1200, 1200, 1500, 1500, 1500, 1750, 1750, 1750, 1950}
	fireballDmg = [28]float64{0, 1500, 1500, 1700, 1700, 1800, 1950, 1950, 2050, 2200, 2200, 2350, 2650, 2650, 2750, 3100, 3100, 3250, 3400, 3400, 3500, 3650, 3650, 3750, 3900, 3900, 3950, 4100}

	// Buildings
	cannonHealth = [21]float64{420, 470, 520, 570, 620, 670, 730, 800, 880, 960, 1060, 1160, 1260, 1380, 1500, 1620, 1740, 1870, 2000, 2150, 2250}
	archerTowerHealth = [21]float64{380, 420, 460, 500, 540, 580, 630, 690, 750, 810, 890, 970, 1050, 1130, 1230, 1310, 1390, 1510, 1600, 1700, 1800}
	mortarHealth = [16]float64{400, 450, 500, 550, 600, 650, 700, 800, 950, 1100, 1300, 1500, 1700, 1950, 2150, 2300}
	airDefenseHealth = []float64{}
	wizardTowerHealth = []float64{}
	airSweeperHealth = []float64{}
	hiddenTeslaHealth = []float64{}
	bombTowerHealth = []float64{}
	xBowHealth = []float64{}
	infernoTowerHealth = []float64{}
	eagleArtilleryHealth = []float64{}
	scattershotHealth = []float64{}
	builderHuntHealth = []float64{}
	spellTowerHealth = []float64{}
	monolithHealth = []float64{}
	multiArcherTowerHealth = []float64{}
	ricochetCannonHealth = []float64{}
	townHallHealth = []float64{}
}

func GetLightingDmg(level int8) float64 {
	return lightingDmg[level]
}
func GetEarthquakeDmg(level int8) float64 {
	return earthquakeDmg[level]
}
func GetGiantarrowDmg(level int8) float64 {
	return giantarrowDmg[level]
}
func GetFireballDmg(level int8) float64 {
	return fireballDmg[level]
}
