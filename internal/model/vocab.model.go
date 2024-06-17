package model

type Vocab struct {
	Reading          string   `mapstructure:"reading"`
	Meaning          string   `mapstructure:"meaning"`
	Context          []string `mapstructure:"context"`
	KanjiComposition []string `mapstructure:"kanji_composition"`
}
