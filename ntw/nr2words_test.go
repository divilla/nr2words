package ntw_test

import (
	"github.com/divilla/nr2words/ntw"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardinalToWords(t *testing.T) {
	ast := assert.New(t)
	tt := []struct {
		nr  int
		res string
	}{
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{14, "fourteen"},
		{15, "fifteen"},
		{16, "sixteen"},
		{17, "seventeen"},
		{18, "eighteen"},
		{19, "nineteen"},
		{20, "twenty"},
		{21, "twenty-one"},
		{30, "thirty"},
		{40, "forty"},
		{50, "fifty"},
		{60, "sixty"},
		{70, "seventy"},
		{80, "eighty"},
		{90, "ninety"},
		{33, "thirty-three"},
		{100, "one hundred"},
		{105, "one hundred and five"},
		{111, "one hundred and eleven"},
		{133, "one hundred and thirty-three"},
		{1000, "one thousand"},
		{1003, "one thousand and three"},
		{1011, "one thousand and eleven"},
		{1033, "one thousand and thirty-three"},
		{1100, "one thousand one hundred"},
		{1103, "one thousand one hundred and three"},
		{1111, "one thousand one hundred and eleven"},
		{1130, "one thousand one hundred and thirty"},
		{1133, "one thousand one hundred and thirty-three"},
		{10133, "ten thousand one hundred and thirty-three"},
		{11133, "eleven thousand one hundred and thirty-three"},
		{30133, "thirty thousand one hundred and thirty-three"},
		{33133, "thirty-three thousand one hundred and thirty-three"},
		{33100, "thirty-three thousand one hundred"},
		{110105, "one hundred and ten thousand one hundred and five"},
		{300000, "three hundred thousand"},
		{333000, "three hundred and thirty-three thousand"},
		{333001, "three hundred and thirty-three thousand and one"},
		{1000000, "one million"},
		{1000003, "one million and three"},
		{1000011, "one million and eleven"},
		{1000033, "one million and thirty-three"},
		{1000099, "one million and ninety-nine"},
		{1000100, "one million one hundred"},
		{1000103, "one million one hundred and three"},
		{1000111, "one million one hundred and eleven"},
		{1000130, "one million one hundred and thirty"},
		{1000133, "one million one hundred and thirty-three"},
		{1333133, "one million three hundred and thirty-three thousand one hundred and thirty-three"},
		{56945781, "fifty-six million nine hundred and forty-five thousand seven hundred and eighty-one"},
		{999999999, "nine hundred and ninety-nine million nine hundred and ninety-nine thousand nine hundred and ninety-nine"},
		{9999999999, "nine billion nine hundred and ninety-nine million nine hundred and ninety-nine thousand nine hundred and ninety-nine"},
	}

	for _, v := range tt {
		res, _ := ntw.CardinalToWords(v.nr)
		ast.Equal(v.res, res, "expected:", v.nr, res)
	}

	// test for errors
	te := []struct {
		nr  int
		err string
	}{
		{-1, "expected integer between 0 and 9,999,999,999"},
		{10000000000, "expected integer between 0 and 9,999,999,999"},
	}

	for _, v := range te {
		_, err := ntw.CardinalToWords(v.nr)
		ast.EqualError(err, v.err, "nr:", v.nr, err)
	}
}

func TestOrdinalToWords(t *testing.T) {
	ast := assert.New(t)
	tt := []struct {
		nr  int
		res string
	}{
		{1, "first"},
		{2, "second"},
		{3, "third"},
		{4, "fourth"},
		{5, "fifth"},
		{6, "sixth"},
		{7, "seventh"},
		{8, "eighth"},
		{9, "ninth"},
		{10, "tenth"},
		{11, "eleventh"},
		{12, "twelfth"},
		{13, "thirteenth"},
		{14, "fourteenth"},
		{15, "fifteenth"},
		{16, "sixteenth"},
		{17, "seventeenth"},
		{18, "eighteenth"},
		{19, "nineteenth"},
		{20, "twentieth"},
		{21, "twenty-first"},
		{30, "thirtieth"},
		{40, "fortieth"},
		{50, "fiftieth"},
		{60, "sixtieth"},
		{70, "seventieth"},
		{80, "eightieth"},
		{90, "ninetieth"},
		{33, "thirty-third"},
		{100, "one hundredth"},
		{105, "one hundred and fifth"},
		{111, "one hundred and eleventh"},
		{133, "one hundred and thirty-third"},
		{1000, "one thousandth"},
		{1003, "one thousand and third"},
		{1011, "one thousand and eleventh"},
		{1033, "one thousand and thirty-third"},
		{1100, "one thousand one hundredth"},
		{1103, "one thousand one hundred and third"},
		{1111, "one thousand one hundred and eleventh"},
		{1130, "one thousand one hundred and thirtieth"},
		{1133, "one thousand one hundred and thirty-third"},
		{10133, "ten thousand one hundred and thirty-third"},
		{11133, "eleven thousand one hundred and thirty-third"},
		{30133, "thirty thousand one hundred and thirty-third"},
		{33133, "thirty-three thousand one hundred and thirty-third"},
		{33100, "thirty-three thousand one hundredth"},
		{110105, "one hundred and ten thousand one hundred and fifth"},
		{300000, "three hundred thousandth"},
		{333000, "three hundred and thirty-three thousandth"},
		{333001, "three hundred and thirty-three thousand and first"},
		{1000000, "one millionth"},
		{1000003, "one million and third"},
		{1000011, "one million and eleventh"},
		{1000033, "one million and thirty-third"},
		{1000099, "one million and ninety-ninth"},
		{1000100, "one million one hundredth"},
		{1000103, "one million one hundred and third"},
		{1000111, "one million one hundred and eleventh"},
		{1000130, "one million one hundred and thirtieth"},
		{1000133, "one million one hundred and thirty-third"},
		{1333133, "one million three hundred and thirty-three thousand one hundred and thirty-third"},
		{56945781, "fifty-six million nine hundred and forty-five thousand seven hundred and eighty-first"},
		{999999999, "nine hundred and ninety-nine million nine hundred and ninety-nine thousand nine hundred and ninety-ninth"},
		{9999999999, "nine billion nine hundred and ninety-nine million nine hundred and ninety-nine thousand nine hundred and ninety-ninth"},
	}

	for _, v := range tt {
		res, _ := ntw.OrdinalToWords(v.nr)
		ast.Equal(v.res, res, "expected:", v.nr, res)
	}

	// test for errors
	te := []struct {
		nr  int
		err string
	}{
		{0, "expected integer between 1 and 9,999,999,999"},
		{10000000000, "expected integer between 1 and 9,999,999,999"},
	}

	for _, v := range te {
		_, err := ntw.OrdinalToWords(v.nr)
		ast.EqualError(err, v.err, "nr:", v.nr, err)
	}
}
