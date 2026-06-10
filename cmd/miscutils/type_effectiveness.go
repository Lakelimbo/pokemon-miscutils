package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/Lakelimbo/pokemon-miscutils/info/ability"
	"github.com/Lakelimbo/pokemon-miscutils/info/types"
	"github.com/Lakelimbo/pokemon-miscutils/internal/colors"
	"github.com/Lakelimbo/pokemon-miscutils/internal/utils"
	"github.com/urfave/cli/v3"
)

var effectivenessCmd = cli.Command{
	Name:    "effectiveness",
	Aliases: []string{"types", "effect"},
	Usage:   "calculates type effectiveness",
	Arguments: []cli.Argument{
		&cli.StringArgs{
			Name: "types",
			Min:  1,
			Max:  2,
		},
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "ability",
			Aliases: []string{"a"},
			Usage:   "possible ability that affects the type effectiveness",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		args := c.StringArgs("types")
		abilityName := c.String("ability")

		t1, err := parseType(args[0])
		if err != nil {
			return err
		}

		t := []types.Type{t1}
		if len(args) > 1 {
			t2, err := parseType(args[1])
			if err != nil {
				return err
			}

			if t2 != t1 {
				t = append(t, t2)
			}
		}

		var a ability.Abilities
		if abilityName != "" {
			a = ability.Abilities(utils.Trimmer(abilityName))
		}

		effect := types.CalculateEffectiveness(&a, t[0], t[1:]...)

		var rows = make([][]string, 0, len(effect))
		for t, v := range effect {
			sType := lipgloss.NewStyle().
				Foreground(colors.TypeColors[t].Fg).
				Background(colors.TypeColors[t].Bg).
				Width(12).
				Padding(0, 1)

			sVal := lipgloss.NewStyle().Align(lipgloss.Center)

			switch {
			case v >= 4:
				sVal = sVal.Foreground(lipgloss.Red).Bold(true)
			case v >= 2:
				sVal = sVal.Foreground(lipgloss.Red)
			case v > 1:
				sVal = sVal.Foreground(lipgloss.Yellow)
			case v == 1:
				sVal = sVal.Foreground(lipgloss.NoColor{})
			case v >= 0.25:
				sVal = sVal.Foreground(lipgloss.Green)
			case v >= 0:
				sVal = sVal.Foreground(lipgloss.Blue).Bold(true)
			case v == 0:
				sVal = sVal.Bold(true)
			}

			rows = append(rows, []string{
				sType.Render(string(t)),
				sVal.Render(strconv.FormatFloat(v, 'f', 2, 64)),
			})
		}

		tui := table.New().
			Border(lipgloss.RoundedBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("241"))).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch row {
				case table.HeaderRow:
					return lipgloss.NewStyle().Align(lipgloss.Center)
				default:
					return lipgloss.NewStyle().
						Padding(0, 1).
						Width(14).
						Align(lipgloss.Center)
				}
			}).
			Headers("TYPE", "VALUE").
			Rows(rows...)

		sT := lipgloss.NewStyle().Bold(true)
		sT1 := sT.Foreground(colors.TypeColors[t1].Bg).SetString(string(t1))
		if len(t) > 1 {
			sT2 := sT.Foreground(colors.TypeColors[t[1]].Bg).SetString(string(t[1]))
			fmt.Printf("The effectiveness for the %s and %s types are:\n", sT1, sT2)
		} else {
			fmt.Printf("The effectiveness for the %s type is:\n", sT1)
		}

		lipgloss.Println(tui)

		return nil
	},
}

// parseType takes aliases for multiple types for easier CLI usage (slightly less
// typo-prone).
func parseType(str string) (types.Type, error) {
	styledStr := lipgloss.NewStyle().Foreground(lipgloss.Red)

	switch strings.ToLower(str) {
	case "bug":
		return types.Bug, nil
	case "dark":
		return types.Dark, nil
	case "dragon":
		return types.Dragon, nil
	case "electric", "electr", "elec":
		return types.Electric, nil
	case "fairy":
		return types.Fairy, nil
	case "fighting", "fight":
		return types.Fighting, nil
	case "fire":
		return types.Fire, nil
	case "flying", "fly":
		return types.Flying, nil
	case "ghost":
		return types.Ghost, nil
	case "grass", "plant":
		return types.Grass, nil
	case "ground", "earth":
		return types.Ground, nil
	case "ice":
		return types.Ice, nil
	case "normal":
		return types.Normal, nil
	case "poison", "toxic":
		return types.Poison, nil
	case "psychic", "psych", "psy":
		return types.Psychic, nil
	case "rock", "stone":
		return types.Rock, nil
	case "steel", "metal", "iron":
		return types.Steel, nil
	case "water":
		return types.Water, nil
	default:
		return "", fmt.Errorf("unknown type '%s'. See if you typed correctly", styledStr.Render(str))
	}
}
