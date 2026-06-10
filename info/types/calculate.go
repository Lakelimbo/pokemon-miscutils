package types

import (
	"slices"

	"github.com/Lakelimbo/pokemon-miscutils/info/ability"
)

// CalculateEffectiveness calculates the effectiveness in a defensive way.
// It takes a mandatory first type, and a slice for secondary type (though it should
// work with any amount of variants, if they ever come to exist).
//
// It also takes a pointer to ability.Abilities for any ability that can modify the
// effectiveness output. If no ability matching the criteria is found, ability will be
// nil.
func CalculateEffectiveness(ability *ability.Abilities, type1 Type, type2 ...Type) map[Type]float64 {
	res := make(map[Type]float64)
	for attacking := range Effect {
		res[attacking] = 1
	}

	defending := make([]Type, 1, len(type2))
	defending = append(defending, type2...)

	for _, def := range defending {
		e := Effect[def]

		for atk := range res {
			if slices.Contains(e.Immunities, atk) {
				res[atk] = 0
				continue
			}

			if res[atk] == 0 {
				continue
			}

			if slices.Contains(e.Weaknesses, atk) {
				res[atk] *= 2
			}
			if slices.Contains(e.Resistances, atk) {
				res[atk] *= 0.5
			}
		}
	}

	if modifier, ok := AbilityEffects[*ability]; ok {
		modifier(res)
	}

	return res
}
