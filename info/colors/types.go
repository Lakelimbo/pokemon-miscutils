package colors

import (
	"github.com/Lakelimbo/pokemon-miscutils/info/types"
	"github.com/charmbracelet/x/ansi"
)

type TypeColor struct {
	Bg ansi.HexColor
	Fg ansi.HexColor
}

var TypeColors = map[types.Type]TypeColor{
	types.Bug:      {ansi.HexColor("#A3C331"), ansi.HexColor("#FFFFFF")},
	types.Dark:     {ansi.HexColor("#73798A"), ansi.HexColor("#FFFFFF")},
	types.Dragon:   {ansi.HexColor("#037BC7"), ansi.HexColor("#FFFFFF")},
	types.Electric: {ansi.HexColor("#F3DA56"), ansi.HexColor("#000000")},
	types.Fairy:    {ansi.HexColor("#EC8FE4"), ansi.HexColor("#000000")},
	types.Fighting: {ansi.HexColor("#E2414B"), ansi.HexColor("#FFFFFF")},
	types.Fire:     {ansi.HexColor("#FBA749"), ansi.HexColor("#FFFFFF")},
	types.Flying:   {ansi.HexColor("#97B0E2"), ansi.HexColor("#FFFFFF")},
	types.Ghost:    {ansi.HexColor("#556AB0"), ansi.HexColor("#FFFFFF")},
	types.Grass:    {ansi.HexColor("#5EBD57"), ansi.HexColor("#FFFFFF")},
	types.Ground:   {ansi.HexColor("#D68555"), ansi.HexColor("#FFFFFF")},
	types.Ice:      {ansi.HexColor("#76CFC2"), ansi.HexColor("#FFFFFF")},
	types.Normal:   {ansi.HexColor("#9498A3"), ansi.HexColor("#FFFFFF")},
	types.Poison:   {ansi.HexColor("#B462CD"), ansi.HexColor("#FFFFFF")},
	types.Psychic:  {ansi.HexColor("#F77676"), ansi.HexColor("#FFFFFF")},
	types.Rock:     {ansi.HexColor("#C6B689"), ansi.HexColor("#000000")},
	types.Steel:    {ansi.HexColor("#528B9F"), ansi.HexColor("#FFFFFF")},
	types.Water:    {ansi.HexColor("#4D94DD"), ansi.HexColor("#FFFFFF")},
}
