package ntw

import (
	"errors"
	"strconv"
	"strings"
)

// const represent parsing segment of target number
const (
	// last 2 digits of each three digit group
	tens = iota + 1
	// first digit of three digit group
	hundreds
	// position of three digit group, last is 0, one before last 1...
	thousands
)

// CardinalToWords converts integer number into a word expression
// result is cardinal expression of number in English language
func CardinalToWords(i int) (string, error) {
	if i < 0 || i > 9999999999 {
		return "", errors.New("expected integer between 0 and 9,999,999,999")
	}
	if i == 0 {
		return "zero", nil
	}

	return parse(strconv.Itoa(i), false), nil
}

// OrdinalToWords converts integer number into a word expression
// result is ordinal expression of number in English language
func OrdinalToWords(i int) (string, error) {
	if i < 1 || i > 9999999999 {
		return "", errors.New("expected integer between 1 and 9,999,999,999")
	}

	return parse(strconv.Itoa(i), true), nil
}

// parse is here to give overall readability to
// what and how code is working
func parse(input string, ordinal bool) string {
	// 1. call build and get tokens
	ts := build(input, 0)

	// transform last word to ordinal form
	if ordinal {
		setOrdinal(ts)
	}

	// insert word "and" in front of "tens" words
	setAnd(ts)

	s := ""
	// concat all non empty tokens into a sentence
	for _, t := range ts {
		if t.isEmpty() {
			continue
		}

		s = t.value() + " " + s
	}

	// trim unwanted prefixes and suffixes
	s = strings.TrimPrefix(s, tokens.other["and"]+" ")
	s = strings.TrimSuffix(s, " ")

	return s
}

// this is core function used to build words
// it parses maximum of 3 digits
// passing remaining digits to recursive call of build
// it's done this way to theoretically
// enable parsing numbers to sextillions
// (you only have to add words to metadata.go, line 90)
func build(input string, pos int) []iToken {
	l := len(input)
	i := input
	if l > 3 {
		i = i[l-3:]
	}

	// code builds all three words,
	// even if digits don't exist,
	// and empty thousands token for last three digits
	// tokens are created in reverse order
	// IMHO this approach has biggest code readability
	ts := []iToken{
		newThousandsToken(i, pos),
		newTensToken(i),
		newHundredsToken(i),
	}

	// if there are more than 3 digits left to process
	// pass them to next recursion
	if l > 3 {
		return append(ts, build(input[:l-3], pos+1)...)
	}

	return ts
}

// setOrdinal represents the formula
// to convert token to ordinal form
func setOrdinal(ts []iToken) {
	for k, t := range ts {
		if k == 0 {
			continue
		}
		if !t.isEmpty() {
			t.setOrdinal()
			t.parse()
			break
		}
	}
}

// setAnd will append word "and"
// to all "tens" tokens as specified in assignment.
// Given 110105 output one hundred and ten thousand one hundred and five
func setAnd(ts []iToken) {
	for _, t := range ts {
		if !t.isEmpty() && t.kind() == tens {
			t.setAnd()
		}
	}
}

// this is initial Merriam-Webster version where "and" is placed only in front of last token,
// but including hundreds. For example:
// 111105 - one hundred eleven thousand one hundred "and" five
// 111100 - one hundred eleven thousand "and" one hundred
// ******IT'S NOT HERE TO JUDGE ASSIGNMENT, BUT TO SHOW HOW EASY IS TO FINE TUNE :) ******
//func setAnd(ts []iToken) {
//	for _, t := range ts {
//		if !t.isEmpty() && t.kind() != thousands {
//			t.setAnd()
//			break
//		}
//	}
//}

// helper function used to select Cardinal or Ordinal variant of spelling number,
// it's only purpose is to make code more readable
func carord(mc map[string]string, mo map[string]string, ordinal bool) map[string]string {
	if ordinal {
		return mo
	}
	return mc
}
