package types

// Pokémon types enum
type Type string

var (
	Bug      Type = "Bug"
	Dark     Type = "Dark"
	Dragon   Type = "Dragon"
	Electric Type = "Electric"
	Fairy    Type = "Fairy"
	Fighting Type = "Fighting"
	Fire     Type = "Fire"
	Flying   Type = "Flying"
	Ghost    Type = "Ghost"
	Grass    Type = "Grass"
	Ground   Type = "Ground"
	Ice      Type = "Ice"
	Normal   Type = "Normal"
	Poison   Type = "Poison"
	Psychic  Type = "Psychic"
	Rock     Type = "Rock"
	Steel    Type = "Steel"
	Water    Type = "Water"
)

type Effectiveness struct {
	Weaknesses  []Type
	Resistances []Type
	Immunities  []Type
}

var Effect = map[Type]Effectiveness{
	Bug: {
		Weaknesses:  []Type{Fire, Flying, Rock},
		Resistances: []Type{Fighting, Grass, Ground},
	},
	Dark: {
		Weaknesses:  []Type{Bug, Fairy, Fighting},
		Resistances: []Type{Ghost, Dark},
		Immunities:  []Type{Psychic},
	},
	Dragon: {
		Weaknesses:  []Type{Dragon, Fairy, Ice},
		Resistances: []Type{Electric, Fire, Grass, Water},
	},
	Electric: {
		Weaknesses:  []Type{Ground},
		Resistances: []Type{Electric, Flying, Steel},
	},
	Fairy: {
		Weaknesses:  []Type{Poison, Steel},
		Resistances: []Type{Bug, Dark, Fighting},
		Immunities:  []Type{Dragon},
	},
	Fighting: {
		Weaknesses:  []Type{Fairy, Flying, Psychic},
		Resistances: []Type{Bug, Dark, Rock},
	},
	Fire: {
		Weaknesses:  []Type{Ground, Rock, Water},
		Resistances: []Type{Bug, Fairy, Fire, Grass, Ice, Steel},
	},
	Flying: {
		Weaknesses:  []Type{Electric, Ice, Rock},
		Resistances: []Type{Bug, Fighting, Grass},
	},
	Ghost: {
		Weaknesses:  []Type{Dark, Ghost},
		Resistances: []Type{Bug, Poison},
		Immunities:  []Type{Fighting, Normal},
	},
	Grass: {
		Weaknesses:  []Type{Bug, Fire, Flying, Ice, Poison},
		Resistances: []Type{Electric, Grass, Ground, Water},
	},
	Ground: {
		Weaknesses:  []Type{Grass, Ice, Water},
		Resistances: []Type{Fire, Poison, Rock},
		Immunities:  []Type{Electric},
	},
	Ice: {
		Weaknesses:  []Type{Fighting, Fire, Rock, Steel},
		Resistances: []Type{Ice},
	},
	Normal: {
		Weaknesses: []Type{Fighting},
		Immunities: []Type{Ghost},
	},
	Poison: {
		Weaknesses:  []Type{Ground, Psychic},
		Resistances: []Type{Bug, Fairy, Fighting, Grass, Poison},
	},
	Psychic: {
		Weaknesses:  []Type{Bug, Dark, Ghost},
		Resistances: []Type{Fighting, Psychic},
	},
	Rock: {
		Weaknesses:  []Type{Fighting, Grass, Ground, Rock, Steel},
		Resistances: []Type{Fire, Flying, Normal, Poison},
	},
	Steel: {
		Weaknesses:  []Type{Fighting, Fire, Ground},
		Resistances: []Type{Bug, Dragon, Fairy, Flying, Grass, Ice, Normal, Psychic, Rock},
		Immunities:  []Type{Poison},
	},
	Water: {
		Weaknesses:  []Type{Electric, Grass},
		Resistances: []Type{Fire, Ice, Steel, Water},
	},
}
