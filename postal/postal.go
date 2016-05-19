package postal

func GetLibpostalDefaultOptions() NormalizeOptions {
	return NormalizeOptions{}
}

func ExpandAddress(input []byte, options NormalizeOptions, n []uint) **byte {
	return nil
}

func ExpansionArrayDestroy(expansions [][]byte, n uint) {
}

func AddressParserResponseDestroy(self *AddressParserResponse) {
}

func GetLibpostalAddressParserDefaultOptions() AddressParserOptions {
	return AddressParserOptions{}
}

func ParseAddress(address []byte, options AddressParserOptions) *AddressParserResponse {
	return &AddressParserResponse{}
}

func Setup() bool {
	return false
}

func Teardown() {
}

func SetupParser() bool {
	return false
}

func TeardownParser() {
}

func SetupLanguageClassifier() bool {
	return false
}

func TeardownLanguageClassifier() {
}
