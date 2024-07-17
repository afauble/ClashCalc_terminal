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
	"github.com/charmbracelet/huh/spinner"
)

var (
	giantarrowLvlStr  string
	fireballLvlStr    string
	lightningLvlStr   string
	earthquakeLvlStr  string
	selectedBuildings []string
	buidlingLevelMap  map[string]string
)

func CreateTerminalUI() {

	form := huh.NewForm(
		// Gather User Input
		huh.NewGroup(
			huh.NewNote().
				Title(`

  /$$$$$$  /$$                     /$$        /$$$$$$           /$$          
 /$$__  $$| $$                    | $$       /$$__  $$         | $$          
| $$  \__/| $$  /$$$$$$   /$$$$$$$| $$$$$$$ | $$  \__/ /$$$$$$ | $$  /$$$$$$$
| $$      | $$ |____  $$ /$$_____/| $$__  $$| $$      |____  $$| $$ /$$_____/
| $$      | $$  /$$$$$$$|  $$$$$$ | $$  \ $$| $$       /$$$$$$$| $$| $$      
| $$    $$| $$ /$$__  $$ \____  $$| $$  | $$| $$    $$/$$__  $$| $$| $$      
|  $$$$$$/| $$|  $$$$$$$ /$$$$$$$/| $$  | $$|  $$$$$$/  $$$$$$$| $$|  $$$$$$$
 \______/ |__/ \_______/|_______/ |__/  |__/ \______/ \_______/|__/ \_______/
	`).
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
				Title(`

  /$$$$$$  /$$                     /$$        /$$$$$$           /$$          
 /$$__  $$| $$                    | $$       /$$__  $$         | $$          
| $$  \__/| $$  /$$$$$$   /$$$$$$$| $$$$$$$ | $$  \__/ /$$$$$$ | $$  /$$$$$$$
| $$      | $$ |____  $$ /$$_____/| $$__  $$| $$      |____  $$| $$ /$$_____/
| $$      | $$  /$$$$$$$|  $$$$$$ | $$  \ $$| $$       /$$$$$$$| $$| $$      
| $$    $$| $$ /$$__  $$ \____  $$| $$  | $$| $$    $$/$$__  $$| $$| $$      
|  $$$$$$/| $$|  $$$$$$$ /$$$$$$$/| $$  | $$|  $$$$$$/  $$$$$$$| $$|  $$$$$$$
 \______/ |__/ \_______/|_______/ |__/  |__/ \______/ \_______/|__/ \_______/`).
				Description("Select desired buildings"),

			huh.NewMultiSelect[string]().
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
				Validate(func(t []string) error {
					if len(t) <= 0 {
						return fmt.Errorf("at least one building is required")
					}
					return nil
				}).
				Value(&selectedBuildings).
				Filterable(true),
		),
	)

	form.WithTheme(huh.ThemeBase16())
	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	_ = spinner.New().Title("Calculating").Action(calculateBuildings).Run()

}

func calculateBuildings() {

	gaLevelInt, _ := strconv.Atoi(giantarrowLvlStr)
	gaLevel := int8(gaLevelInt)
	fbLevelInt, _ := strconv.Atoi(fireballLvlStr)
	fbLevel := int8(fbLevelInt)
	ltLevelInt, _ := strconv.Atoi(lightningLvlStr)
	ltLevel := int8(ltLevelInt)
	eqLevelInt, _ := strconv.Atoi(earthquakeLvlStr)
	eqLevel := int8(eqLevelInt)

	userInput := models.UserInput{
		GiantarrowLvl: gaLevel,
		FireballLvl:   fbLevel,
		LightningLvl:  ltLevel,
		EarthquakeLvl: eqLevel,
	}

	for _, b := range selectedBuildings {
		getHealthFunction := data.GetCorrespondingHealthFunc(b)
		bLevelInt, _ := strconv.Atoi(buidlingLevelMap[b])
		bLevel := int8(bLevelInt)
		bHealth := getHealthFunction(bLevel)

		helpers.FindOptimalSpells(bHealth, userInput)

	}
}
