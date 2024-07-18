package data

import (
	"errors"
)

var (
	lightingDmg   [12]float32
	earthquakeDmg [6]float32
	giantarrowDmg [19]float32
	fireballDmg   [28]float32

	cannonHealth           [22]float32
	archerTowerHealth      [22]float32
	mortarHealth           [17]float32
	airDefenseHealth       [15]float32
	wizardTowerHealth      [17]float32
	airSweeperHealth       [8]float32
	hiddenTeslaHealth      [16]float32
	bombTowerHealth        [12]float32
	xBowHealth             [12]float32
	infernoTowerHealth     [11]float32
	eagleArtilleryHealth   [8]float32
	scattershotHealth      [6]float32
	builderHutHealth       [7]float32
	spellTowerHealth       [4]float32
	monolithHealth         [4]float32
	multiArcherTowerHealth [3]float32
	ricochetCannonHealth   [3]float32
	townHallHealth         [17]float32
	clanCastleHealth       [13]float32

	buildingNameMap map[string][]float32
)

func init() {
	// Spells & Equipment
	lightingDmg = [12]float32{0, 150, 180, 210, 240, 270, 320, 400, 480, 560, 600, 640}
	earthquakeDmg = [6]float32{0, 14.5, 17, 21, 25, 29}
	giantarrowDmg = [19]float32{0, 750, 750, 850, 850, 850, 1000, 1000, 1000, 1200, 1200, 1200, 1500, 1500, 1500, 1750, 1750, 1750, 1950}
	fireballDmg = [28]float32{0, 1500, 1500, 1700, 1700, 1800, 1950, 1950, 2050, 2200, 2200, 2350, 2650, 2650, 2750, 3100, 3100, 3250, 3400, 3400, 3500, 3650, 3650, 3750, 3900, 3900, 3950, 4100}

	// Buildings
	cannonHealth = [22]float32{0, 420, 470, 520, 570, 620, 670, 730, 800, 880, 960, 1060, 1160, 1260, 1380, 1500, 1620, 1740, 1870, 2000, 2150, 2250}
	archerTowerHealth = [22]float32{0, 380, 420, 460, 500, 540, 580, 630, 690, 750, 810, 890, 970, 1050, 1130, 1230, 1310, 1390, 1510, 1600, 1700, 1800}
	mortarHealth = [17]float32{0, 400, 450, 500, 550, 600, 650, 700, 800, 950, 1100, 1300, 1500, 1700, 1950, 2150, 2300}
	airDefenseHealth = [15]float32{0, 800, 850, 900, 950, 1000, 1050, 1100, 1210, 1300, 1400, 1500, 1650, 1750, 1850}
	wizardTowerHealth = [17]float32{0, 620, 650, 680, 730, 840, 960, 1200, 1440, 1600, 1900, 2120, 2240, 2500, 2800, 3000, 3150}
	airSweeperHealth = [8]float32{0, 750, 800, 850, 900, 950, 1000, 1050}
	hiddenTeslaHealth = [16]float32{0, 600, 630, 660, 690, 730, 770, 810, 850, 900, 980, 1100, 1200, 1350, 1450, 1550}
	bombTowerHealth = [12]float32{0, 650, 700, 750, 850, 1050, 1300, 1600, 1900, 2300, 2500, 2700}
	xBowHealth = [12]float32{0, 1500, 1900, 2300, 2700, 3100, 3400, 3700, 4000, 4200, 4400, 4600}
	infernoTowerHealth = [11]float32{0, 1500, 1800, 2100, 2400, 2700, 3000, 3300, 3700, 4000, 4400}
	eagleArtilleryHealth = [8]float32{0, 4000, 4400, 4800, 5200, 5600, 5900, 6200}
	scattershotHealth = [6]float32{0, 3600, 4200, 4800, 5100, 5410}
	builderHutHealth = [7]float32{0, 250, 1000, 1300, 1600, 1800, 1900}
	spellTowerHealth = [4]float32{0, 2500, 2800, 3100}
	monolithHealth = [4]float32{0, 4747, 5050, 5353}
	multiArcherTowerHealth = [3]float32{0, 5000, 5200}
	ricochetCannonHealth = [3]float32{0, 5400, 5700}
	townHallHealth = [17]float32{0, 450, 1600, 1850, 2100, 2400, 2800, 3300, 3900, 4600, 5500, 6800, 7500, 8200, 8900, 9600, 10000}
	clanCastleHealth = [13]float32{0, 1000, 1400, 2000, 2600, 3000, 3400, 4000, 4400, 4800, 5200, 5400, 5600}

	// Get Func Map
	buildingNameMap = map[string][]float32{
		"cannon":           cannonHealth[:],
		"archerTower":      archerTowerHealth[:],
		"mortar":           mortarHealth[:],
		"airDefense":       airDefenseHealth[:],
		"wizardTower":      wizardTowerHealth[:],
		"airSweeper":       airSweeperHealth[:],
		"hiddenTesla":      hiddenTeslaHealth[:],
		"bombTower":        bombTowerHealth[:],
		"xBow":             xBowHealth[:],
		"infernoTower":     infernoTowerHealth[:],
		"eagleArtillery":   eagleArtilleryHealth[:],
		"scattershot":      scattershotHealth[:],
		"builderHut":       builderHutHealth[:],
		"spellTower":       spellTowerHealth[:],
		"monolith":         monolithHealth[:],
		"multiArcherTower": multiArcherTowerHealth[:],
		"ricochetCannon":   ricochetCannonHealth[:],
		"townHall":         townHallHealth[:],
		"clanCastle":       clanCastleHealth[:],
	}
}

// Spell & Equipment Damage Getters
func GetLightingDmg(level int8) float32 {
	return lightingDmg[level]
}
func GetEarthquakeDmg(level int8) float32 {
	return earthquakeDmg[level]
}
func GetGiantarrowDmg(level int8) float32 {
	return giantarrowDmg[level]
}
func GetFireballDmg(level int8) float32 {
	return fireballDmg[level]
}

// Validate level value against buidling array length
func ValidateLevelValue(name string, level int8) error {
	buildingArray := buildingNameMap[name]
	if int(level) >= len(buildingArray) || level < 0 {
		return errors.New("level exceeded array length")
	}
	return nil
}

func GetBuildingArrayLength(name string) int {
	buildingArray := buildingNameMap[name]
	return len(buildingArray)
}

// Get building health based on building name
func GetBuildingHealth(name string, level int8) (float32, error) {
	err := ValidateLevelValue(name, level)
	if err != nil {
		return 0, errors.New("error: level value failed validation")
	}
	buildingArray := buildingNameMap[name]
	return buildingArray[level], nil
}
