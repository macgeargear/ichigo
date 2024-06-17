package model

type Radical struct {
	Reading      string   `mapstructure:"reading"`
	Name         string   `mapstructure:"name"`
	Meaning      string   `mapstructure:"meaning"`
	FoundInKanji []string `mapstructure:"found_in_kanji"`
}
