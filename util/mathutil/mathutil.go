package mathutil

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

func FloatToWei(f float64) (w *big.Int) {

	weiCount := 18

	decimal := GetFloatDecimal(f)

	powCount := weiCount - decimal

	fBigInt := big.NewInt(int64(math.Pow(10, float64(decimal)) * f) )

	hBigInt := big.NewInt(int64(math.Pow(10, float64(powCount))))


	return fBigInt.Mul(fBigInt,hBigInt)

}

func GetFloatDecimal(f float64) int {

	fStr := strconv.FormatFloat(f, 'f', -1, 64)

	_, endCount := GetFloatStrDecimal(fStr)

	return endCount

}

func GetFloatStrDecimal(f_str string) (headCount int, endCount int) {

	strLen := len(f_str)

	pIndex := strings.Index(f_str, ".")

	if pIndex >= 0 {

		return pIndex, strLen - pIndex

	}

	return strLen, 0
}
