package model

type Kanji struct {
	RadicalCombination []string `mapstructure:"radical_combination"`
	Name               string   `mapstructure:"name"`
	Reading            string   `mapstructure:"reading"`
	FoundInVocab       []string `mapstructure:"found_in_vocab"`
}
