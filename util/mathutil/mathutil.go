package mathutil

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

var Wei18 =big.NewInt(int64(math.Pow10(18)))

func FloatToString(f float64) string {

	fStr := strconv.FormatFloat(f, 'f', -1, 64)

	return fStr

}

func StringToFloat(f string) float64 {

	n, err := strconv.ParseFloat(f, 64)

	if err != nil {
		return 0
	}
	return n

}

func WeiToFloat(w *big.Int) float64 {

	re :=new(big.Rat).Quo(new(big.Rat).SetInt(w),new(big.Rat).SetInt(Wei18))

	fe,_ := re.Float64()


	return fe

}

func FloatToWei(f float64) (w *big.Int) {

	weiCount := 18

	decimal := GetFloatDecimal(f)

	powCount := weiCount - decimal

	fBigInt := big.NewInt(int64(math.Pow(10, float64(decimal)) * f))

	hBigInt := big.NewInt(int64(math.Pow(10, float64(powCount))))

	return fBigInt.Mul(fBigInt, hBigInt)

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
