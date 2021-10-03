package ntw

// tokens variable contains maps of all English words
// used in translating number to words
var (
	tokens = struct {
		tenRight         map[string]string
		tenLeft          map[string]string
		tenDouble        map[string]string
		tenRightOrdinal  map[string]string
		tenLeftOrdinal   map[string]string
		tenDoubleOrdinal map[string]string
		other            map[string]string
		thousand         []string
		thousandOrdinal  []string
	}{
		tenRight: map[string]string{
			"1": "one",
			"2": "two",
			"3": "three",
			"4": "four",
			"5": "five",
			"6": "six",
			"7": "seven",
			"8": "eight",
			"9": "nine",
		},
		tenLeft: map[string]string{
			"2": "twenty",
			"3": "thirty",
			"4": "forty",
			"5": "fifty",
			"6": "sixty",
			"7": "seventy",
			"8": "eighty",
			"9": "ninety",
		},
		tenDouble: map[string]string{
			"10": "ten",
			"11": "eleven",
			"12": "twelve",
			"13": "thirteen",
			"14": "fourteen",
			"15": "fifteen",
			"16": "sixteen",
			"17": "seventeen",
			"18": "eighteen",
			"19": "nineteen",
		},
		tenRightOrdinal: map[string]string{
			"1": "first",
			"2": "second",
			"3": "third",
			"4": "fourth",
			"5": "fifth",
			"6": "sixth",
			"7": "seventh",
			"8": "eighth",
			"9": "ninth",
		},
		tenLeftOrdinal: map[string]string{
			"2": "twentieth",
			"3": "thirtieth",
			"4": "fortieth",
			"5": "fiftieth",
			"6": "sixtieth",
			"7": "seventieth",
			"8": "eightieth",
			"9": "ninetieth",
		},
		tenDoubleOrdinal: map[string]string{
			"10": "tenth",
			"11": "eleventh",
			"12": "twelfth",
			"13": "thirteenth",
			"14": "fourteenth",
			"15": "fifteenth",
			"16": "sixteenth",
			"17": "seventeenth",
			"18": "eighteenth",
			"19": "nineteenth",
		},
		other: map[string]string{
			"100":   "hundred",
			"100th": "hundredth",
			"and":   "and",
		},
		thousand: []string{
			"",
			"thousand",
			"million",
			"billion",
		},
		thousandOrdinal: []string{
			"",
			"thousandth",
			"millionth",
			"billionth",
		},
	}
)
