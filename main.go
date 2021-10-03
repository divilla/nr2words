package main

import (
	"fmt"
	"github.com/divilla/nr2words/ntw"
)

func main() {
	nrs := []int{
		-1,
		0,
		1,
		10,
		11,
		21,
		30,
		100,
		105,
		111,
		130,
		131,
		1131,
		13131,
		30131,
		31131,
		110105,
		1000000,
		1000099,
		1000100,
		1100000,
		56945781,
		999999999,
		1001000001,
		9999999999,
	}

	for k, v := range nrs {
		val, err := ntw.CardinalToWords(v)
		if err != nil {
			fmt.Println(k+1, "Error: ", err.Error())
		} else {
			fmt.Println(k+1, v, "\""+val+"\"")
		}
	}

	for k, v := range nrs {
		val, err := ntw.OrdinalToWords(v)
		if err != nil {
			fmt.Println(k+1, "Error: ", err.Error())
		} else {
			fmt.Println(k+1, v, "\""+val+"\"")
		}
	}
}
