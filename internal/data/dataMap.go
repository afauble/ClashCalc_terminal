package data

var lightingDmg [12]float32
var earthquakeDmg [6]float32
var giantarrowDmg [19]float32
var fireballDmg [28]float32

var cannonHealth [21]float32
var archerTowerHealth [21]float32
var mortarHealth [16]float32
var airDefenseHealth []float32
var wizardTowerHealth []float32
var airSweeperHealth []float32
var hiddenTeslaHealth []float32
var bombTowerHealth []float32
var xBowHealth []float32
var infernoTowerHealth []float32
var eagleArtilleryHealth []float32
var scattershotHealth []float32
var builderHuntHealth []float32
var spellTowerHealth []float32
var monolithHealth []float32
var multiArcherTowerHealth []float32
var ricochetCannonHealth []float32
var townHallHealth []float32
var clanCastleHealth []float32

func init() {

	// Spells & Abilities
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
	builderHuntHealth = []float32{}
	spellTowerHealth = []float32{}
	monolithHealth = []float32{}
	multiArcherTowerHealth = []float32{}
	ricochetCannonHealth = []float32{}
	townHallHealth = []float32{}
	clanCastleHealth = []float32{}
}

func GetLightingDmg(level int8) float32 {
	if int(level) >= len(lightingDmg) {
		// TODO create error
		return -1
	}
	return lightingDmg[level]
}
func GetEarthquakeDmg(level int8) float32 {
	if int(level) >= len(earthquakeDmg) {
		// TODO create error
		return -1
	}
	return earthquakeDmg[level]
}
func GetGiantarrowDmg(level int8) float32 {
	if int(level) >= len(giantarrowDmg) {
		// TODO create error
		return -1
	}
	return giantarrowDmg[level]
}
func GetFireballDmg(level int8) float32 {
	if int(level) >= len(fireballDmg) {
		// TODO create error
		return -1
	}
	return fireballDmg[level]
}
