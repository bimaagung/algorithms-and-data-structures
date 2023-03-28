package thebirthdaybar_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**

Two children, Lily and Ron, want to share a chocolate bar. Each of the squares has an integer on it.

Lily decides to share a contiguous segment of the bar selected such that:

1. The length of the segment matches Ron's birth month, and,
2. The sum of the integers on the squares is equal to his birth day.

Determine how many ways she can divide the chocolate.

* PSEUDOCODE

PROGRAM theBirthdayBar

INPUT
var squareChoco: array
var day, month

DEKLARASI
result: interger
sumResult: integer

ALGORITMA
IF S[0] equal day
	OUTPUT 1
ENDIF

FOR
	sumResult = 0

	IF long from squareChoco month
		BREAK
	ENDIF

	FOR i equal 0 to month itterate ++
		sumResult =+ squareChoco[i]
	ENDFOR

	IF sumResult equals day
		result++
	ENDIF

	squareChoco = array of squareChoco always slice 1 item from start

	CONTINUE

ENDFOR


OUTPUT result


END
**/


func theBirthdayBar(s []int32, d, m int32) int32 {

	var result int32 = 0

	if s[0] == d {
		return 1
	}

	for {
		var sumResult int32 = 0

		if int32(len(s)) < m {
			break
		}

		for i := int32(0); i < m; i++ {
			sumResult += s[i]
		}

		if sumResult == d {
			result++
		}

		s = s[1:]

		continue
	}


	return result
}


func TestTheBirthdayBar(t *testing.T){

	var testCase = []struct {
		s []int32
		d, m , output int32
	}{	
		{
			[]int32{2,2,1,3,2},4,2,2,
		},
		{
			[]int32{1,2,1,3,2},3,2,2,
		},
		{
			[]int32{1,1,1,1,1,1},3,2,0,
		},
		{
			[]int32{4},4,1,1,
		},
		{
			[]int32{2,5,1,3,4,4,3,5,1,1,2,1,4,1,3,3,4,2,1}, 18, 7, 3,
		},
	}

	for _, v := range testCase {
		testName := fmt.Sprintf("function TheBirthday with s = %v, d = %d, m = %d, output = %d", v.s, v.d, v.m, v.output)
		
		t.Run(testName, func(t *testing.T) {
			result := theBirthdayBar(v.s, v.d, v.m)

			assert.Equal(t, v.output, result)
		})
	}
}