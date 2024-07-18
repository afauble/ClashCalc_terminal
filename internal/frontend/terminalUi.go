package frontend

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/afauble/ClashCalc/internal/data"
	"github.com/afauble/ClashCalc/internal/helpers"
	"github.com/afauble/ClashCalc/internal/models"
	"github.com/charmbracelet/huh"
)

const headerStr string = `
  /$$$$$$  /$$                     /$$        /$$$$$$           /$$          
 /$$__  $$| $$                    | $$       /$$__  $$         | $$          
| $$  \__/| $$  /$$$$$$   /$$$$$$$| $$$$$$$ | $$  \__/ /$$$$$$ | $$  /$$$$$$$
| $$      | $$ |____  $$ /$$_____/| $$__  $$| $$      |____  $$| $$ /$$_____/
| $$      | $$  /$$$$$$$|  $$$$$$ | $$  \ $$| $$       /$$$$$$$| $$| $$      
| $$    $$| $$ /$$__  $$ \____  $$| $$  | $$| $$    $$/$$__  $$| $$| $$      
|  $$$$$$/| $$|  $$$$$$$ /$$$$$$$/| $$  | $$|  $$$$$$/  $$$$$$$| $$|  $$$$$$$
\______/ |__/ \_______/|_______/ |__/  |__/ \______/ \_______/|__/ \_______/
`

var (
	giantarrowLvlStr string
	fireballLvlStr   string
	lightningLvlStr  string
	earthquakeLvlStr string

	selectedBuilding    string
	selectedBuildingLvl string
	resultsStr          string
)

func CreateTerminalUI() {

	form := huh.NewForm(
		// Gather User Input
		huh.NewGroup(
			huh.NewNote().
				Title(headerStr).
				Description("Enter your spell/equipment levels"),

			huh.NewInput().
				Value(&lightningLvlStr).
				Title("Lightning Spell Level").
				Placeholder("1-11").
				Validate(func(s string) error {
					i, err := strconv.Atoi(s)
					if err != nil || i > 11 || i < 0 {
						return errors.New("enter a valid level (1-11)")
					}
					return nil
				}),
			huh.NewInput().
				Value(&earthquakeLvlStr).
				Title("Earthquake Spell Level").
				Placeholder("1-5").
				Validate(func(s string) error {
					i, err := strconv.Atoi(s)
					if err != nil || i > 5 || i < 1 {
						return errors.New("enter a valid level (1-5)")
					}
					return nil
				}),
			huh.NewInput().
				Value(&giantarrowLvlStr).
				Title("Giant Arrow Level").
				Placeholder("0-18").
				Validate(func(s string) error {
					if s == "" {
						return nil
					}
					i, err := strconv.Atoi(s)
					if err != nil || i > 18 || i < 0 {
						return errors.New("enter a valid level (0-18)")
					}
					return nil
				}),
			huh.NewInput().
				Value(&fireballLvlStr).
				Title("Fireball Level").
				Placeholder("0-27").
				Validate(func(s string) error {
					if s == "" {
						return nil
					}
					i, err := strconv.Atoi(s)
					if err != nil || i > 27 || i < 0 {
						return errors.New("enter a valid level (0-27)")
					}
					return nil
				}),
		),
		// Selecting buildings to calculate
		huh.NewGroup(
			huh.NewNote().
				Title(headerStr).
				Description("Select desired building"),

			huh.NewSelect[string]().
				Title("Buildings").
				Options(
					huh.NewOption("Cannon", "cannon"),
					huh.NewOption("Archer Tower", "archerTower"),
					huh.NewOption("Mortar", "mortar"),
					huh.NewOption("Air Defense", "airDefense"),
					huh.NewOption("Wizard Tower", "wizardTower"),
					huh.NewOption("Air Sweeper", "airSweeper"),
					huh.NewOption("Hidden Tesla", "hiddenTesla"),
					huh.NewOption("Bomb Tower", "bombTower"),
					huh.NewOption("XBow", "xBow"),
					huh.NewOption("InfernoTower", "infernoTower"),
					huh.NewOption("Eagle Artillery", "eagleArtillery"),
					huh.NewOption("Scattershot", "scattershot"),
					huh.NewOption("Builder Hut", "builderHut"),
					huh.NewOption("Spell Tower", "spellTower"),
					huh.NewOption("Monolith", "monolith"),
					huh.NewOption("Multi-Archer Tower", "multiArcherTower"),
					huh.NewOption("Ricochet Cannon", "ricochetCannon"),
					huh.NewOption("Town Hall", "townHall"),
					huh.NewOption("Clan Castle", "clanCastle"),
				).
				Value(&selectedBuilding),
			huh.NewInput().
				Description("Input level for "+selectedBuilding).
				Value(&selectedBuildingLvl).
				Validate(func(s string) error {
					level, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("enter a valid building level")
					}
					err = data.ValidateLevelValue(selectedBuilding, int8(level))
					if err != nil {
						return errors.New("enter a valid building level")
					}
					return nil
				}).
				PlaceholderFunc(func() string {
					length := data.GetBuildingArrayLength(selectedBuilding) - 1
					lengthStr := strconv.Itoa(length)
					return "1-" + lengthStr
				}, &selectedBuilding),
			huh.NewNote().
				Title("Results").
				DescriptionFunc(func() string {
					if selectedBuildingLvl == "" {
						return ""
					}
					calculateBuildings()
					return resultsStr
				}, []*string{&selectedBuilding, &selectedBuildingLvl, &lightningLvlStr, &earthquakeLvlStr, &fireballLvlStr, &giantarrowLvlStr}).
				Height(3),
		),
		huh.NewGroup(
			huh.NewNote().
				Title(headerStr).
				Description("\nPress *enter* to exit.\n*Shift + Tab* to return.\n"),
		),
	)

	form.WithTheme(huh.ThemeBase16())
	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}

func calculateBuildings() {
	var gaLevelInt int
	var fbLevelInt int
	var ltLevelInt int
	var eqLevelInt int
	var bLevelInt int

	// These strings were allowed to be empty
	// Convert the empty strings to 0
	if giantarrowLvlStr == "" {
		gaLevelInt = 0
	} else {
		gaLevelInt, _ = strconv.Atoi(giantarrowLvlStr)
	}
	if fireballLvlStr == "" {
		fbLevelInt = 0
	} else {
		fbLevelInt, _ = strconv.Atoi(fireballLvlStr)
	}

	// Convert rest of user input strings
	ltLevelInt, _ = strconv.Atoi(lightningLvlStr)
	eqLevelInt, _ = strconv.Atoi(earthquakeLvlStr)
	bLevelInt, _ = strconv.Atoi(selectedBuildingLvl)

	// Create user input model
	userInput := models.UserInput{
		GiantarrowLvl: int8(gaLevelInt),
		FireballLvl:   int8(fbLevelInt),
		LightningLvl:  int8(ltLevelInt),
		EarthquakeLvl: int8(eqLevelInt),
	}

	// Get building maz health
	bHealth, err := data.GetBuildingHealth(selectedBuilding, int8(bLevelInt))
	if err != nil {
		resultsStr = "*enter a valid building level"
		return
	}

	// Call helper method to calculate results
	selectedBuildingResults := helpers.FindOptimalSpells(bHealth, userInput)

	// Set global variable for huh.form
	resultsStr = selectedBuildingResults.ToString()
}
