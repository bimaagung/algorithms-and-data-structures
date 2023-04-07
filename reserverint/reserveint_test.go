package reserverint_test

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func reverse(x int) int {

	if (x % 10) == 0 {
        x = x / 10
    }

    arrNumb := []string{}

    itoaNumb := strconv.Itoa(x)
    splitNumb := strings.Split(itoaNumb,"")

    var negative string

    for _, v := range splitNumb {
        if v == "-" {
            negative = "-"
        }else{
            arrNumb = append(arrNumb, v)
        }
    }

    var numberRemind string 
    pointerY := len(arrNumb)-1
    var lenNumb float64 = float64(len(arrNumb) / 2) 
    floorNumb := math.Floor(lenNumb)

    for i := 0; i < int(floorNumb); i++ {
        numberRemind = arrNumb[i] 
        arrNumb[i] = arrNumb[pointerY] 
        arrNumb[pointerY] = numberRemind
        pointerY-- 
    }


    strJoin := strings.Join(arrNumb, "")
    result, err := strconv.Atoi(strJoin)
    
    if err != nil {
        panic(err)
    }

    if negative == "-" {
        result = result * -1
    }

    if result > int(math.Pow(2, 31)) || result < int(math.Pow(-2, 31)) {
        return 0
    }


    return result
}


func TestReserveInt(t *testing.T){

	var testCase = []struct {
		x, output int
	}{	
		{
			432,234,
		},
		{
			-761,-167,
		},
		{
			150,51,
		},
		{
			230,32,
		},
		{
			94859483726374, 0,
		},
	}

	for _, v := range testCase {
		testName := fmt.Sprintf("function reserver int with x = %d, output = %d", v.x, v.output)
		
		t.Run(testName, func(t *testing.T) {
			result := reverse(v.x)

			assert.Equal(t, v.output, result)
		})
	}
}