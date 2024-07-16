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

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&giantarrowLvlStr).
				Title("Giant Arrow Level").
				Placeholder("0-18").
				Validate(func(s string) error {
					i, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("enter a valid level (0-18)")
					}
					if i > 18 || i < 0 {
						return errors.New("enter a valid level (0-18)")
					}
					return nil
				}).
				Description("Giant Arrow description"),

			huh.NewInput().
				Value(&fireballLvlStr).
				Title("Fireball Level").
				Placeholder("0-27").
				Validate(func(s string) error {
					i, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("enter a valid level (0-27)")
					}
					if i > 27 || i < 0 {
						return errors.New("enter a valid level (0-27)")
					}
					return nil
				}).
				Description("Fireball description"),
		),
	)

	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}
