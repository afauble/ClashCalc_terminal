package frontend

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

func CreateTerminalUI() {

	var giantarrowLvlStr string
	var fireballLvlStr string
	var lightningLvlStr string
	var earthquakeLvlStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("ClashCalc™").
				Description("Enter your spell/equipment levels below."),

			huh.NewInput().
				Value(&lightningLvlStr).
				Title("Lightning Spell Level").
				Placeholder("1-11").
				Validate(func(s string) error {
					i, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("enter a valid level (1-11)")
					}
					if i > 11 || i < 0 {
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
					if err != nil {
						return errors.New("enter a valid level (1-5)")
					}
					if i > 5 || i < 1 {
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
						s = "0"
						return nil
					}
					i, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("enter a valid level (0-18)")
					}
					if i > 18 || i < 0 {
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
						s = "0"
						return nil
					}
					i, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("enter a valid level (0-27)")
					}
					if i > 27 || i < 0 {
						return errors.New("enter a valid level (0-27)")
					}
					return nil
				}),
		),
	)
	form.WithTheme(huh.ThemeBase16())
	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}
