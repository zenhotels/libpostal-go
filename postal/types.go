package postal

type NormalizeOptions struct {
	Languages              [][]byte
	NumLanguages           int32
	AddressComponents      uint16
	LatinAscii             bool
	Transliterate          bool
	StripAccents           bool
	Decompose              bool
	Lowercase              bool
	TrimString             bool
	DropParentheticals     bool
	ReplaceNumericHyphens  bool
	DeleteNumericHyphens   bool
	SplitAlphaFromNumeric  bool
	ReplaceWordHyphens     bool
	DeleteWordHyphens      bool
	DeleteFinalPeriods     bool
	DeleteAcronymPeriods   bool
	DropEnglishPossessives bool
	DeleteApostrophes      bool
	ExpandNumex            bool
	RomanNumerals          bool
}

type AddressParserResponse struct {
	NumComponents uint
	Components    [][]byte
	Labels        [][]byte
}

type AddressParserOptions struct {
	Language []byte
	Country  []byte
}
