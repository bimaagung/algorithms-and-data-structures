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


END
**/


func theBirthdayBar(s []int32, d, m int32) int32 {
	return 0
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
	}

	for _, v := range testCase {
		testName := fmt.Sprintf("function TheBirthday with s = %v, d = %d, m = %d, output = %d", v.s, v.d, v.m, v.output)
		
		t.Run(testName, func(t *testing.T) {
			result := theBirthdayBar(v.s, v.d, v.m)
			assert.Equal(t, result, v.output)
		})
	}
}