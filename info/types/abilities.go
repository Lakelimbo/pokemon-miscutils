package types

import (
	"github.com/Lakelimbo/pokemon-miscutils/info/ability"
)

type AbilityModifier func(map[Type]float64)

// Takes possible abilities that affect damage output and overrides
// its values. Should be used AFTER base calculation is done, otherwise
// these values won't be properly calculated.
//
// Best to keep the values as lambdas instead of hardcoded values.
var AbilityEffects = map[ability.Abilities]AbilityModifier{
	// Dry Skin makes it immune to Water, and increases Fire
	// super-effectiveness by 25%
	ability.DrySkin: func(m map[Type]float64) {
		m[Water] = 0
		if m[Fire] != 0 {
			m[Fire] *= 1.25
		}
	},

	// Earth Eater makes it immune to Ground types
	ability.EarthEater: func(m map[Type]float64) {
		m[Ground] = 0
	},

	// Filter reduces any super-effective move by 25%
	ability.Filter: filterVariants,

	// Flash Fire makes it immune to Fire types
	ability.FlashFire: fireImmunityVariants,

	// Heatproof reduces Fire type attacks by 50% (essentially
	// cutting Fire effectiveness in half)
	ability.Heatproof: func(m map[Type]float64) {
		m[Fire] *= 0.5
	},

	// Levitate makes it immune to Ground type moves (except Thousand
	// Arrows), as well as making it ungrounded (e.g wouldn't be affected by
	// Toxic Spikes)
	ability.Levitate: func(m map[Type]float64) {
		m[Ground] = 0
	},

	// Lightning Rod makes it immune to Electric types
	ability.LightningRod: electricImmunityVariants,

	// Prism Armor reduces any super-effective moves by 25%
	ability.PrismArmor: filterVariants,

	// Solid Rock reduces any super-effective moves by 25%
	ability.SolidRock: filterVariants,

	// Thick Fat reduces Fire and Ice type moves by half
	ability.ThickFat: func(m map[Type]float64) {
		m[Fire] *= 0.5
		m[Ice] *= 0.5
	},

	// Volt Absorb makes it immune to Electric types
	ability.VoltAbsorb: electricImmunityVariants,

	// Water Absorb makes it immune to Water types
	ability.WaterAbsorb: func(m map[Type]float64) {
		m[Water] *= 0
	},

	// Well-Baked Body makes it immune to Fire types
	ability.WellBakedBody: fireImmunityVariants,

	// Wonder Guard makes it immune to any non super-effective
	// moves
	ability.WonderGuard: func(m map[Type]float64) {
		for t, v := range m {
			if v <= 1 {
				m[t] = 0
			}
		}
	},
}

func filterVariants(m map[Type]float64) {
	for t, v := range m {
		if v > 1 {
			m[t] *= 0.75
		}
	}
}

func electricImmunityVariants(m map[Type]float64) {
	m[Electric] = 0
}

func fireImmunityVariants(m map[Type]float64) {
	m[Fire] = 0
}
