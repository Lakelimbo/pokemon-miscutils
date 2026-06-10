package ability

// Ability enum. Lots of abilities have more than one word, so it's
// best to keep their string value without spaces and in lower case.
// (the CLI will trim and lowercase them if the --ability flag is provided)
type Abilities string

var (
	DrySkin       Abilities = "dryskin"
	EarthEater    Abilities = "eartheater"
	Filter        Abilities = "filter"
	FlashFire     Abilities = "flashfire"
	Heatproof     Abilities = "heatproof"
	Levitate      Abilities = "levitate"
	LightningRod  Abilities = "lightningrod"
	PrismArmor    Abilities = "prismarmor"
	SapSipper     Abilities = "sapsipper"
	SolidRock     Abilities = "solidrock"
	StormDrain    Abilities = "stormdrain"
	ThickFat      Abilities = "thickfat"
	VoltAbsorb    Abilities = "voltabsorb"
	WaterAbsorb   Abilities = "waterabsorb"
	WellBakedBody Abilities = "wellbakedbody"
	WonderGuard   Abilities = "wonderguard"
)
