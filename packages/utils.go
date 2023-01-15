package packages

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func GetFirstHour (rfc string) int{
	splitLine := strings.Split(rfc, " ")
	rfcGet, err := time.Parse(time.RFC3339, splitLine[0])
	CheckError(err)
	hour := rfcGet.Hour()
	return hour
}

func GetHour (rfc string) int{
	splitLine := strings.Split(rfc, " ")
	rfcGet, err := time.Parse(time.RFC3339, splitLine[0])
	CheckError(err)
	hour := rfcGet.Hour()
	return hour
}

func RFCRoundUp(rfc string) string {
    t, err := time.Parse(time.RFC3339, rfc)
    if err != nil {
        panic(err)
    }
    t = t.Truncate(time.Hour).Add(time.Hour - 1)
	// fmt.Println("ROUNDUP",t.Format(time.RFC3339))
    return t.Format(time.RFC3339)
}

func RFCRoundDown(rfc string) string {
    t, err := time.Parse(time.RFC3339, rfc)
    if err != nil {
        panic(err)
    }
    t = t.Truncate(time.Hour)
	// fmt.Println("ROUNDDown",t.Format(time.RFC3339))
    return t.Format(time.RFC3339)
}

func rfcCustom(year int, month time.Month, day int, hour int) string{

	rfcReturn := time.Date(year, month, day, hour, 0, 0, 0, time.UTC)
	return rfcReturn.Format(time.RFC3339)

}


func CreateBucket(year int, month time.Month, day int, valueBucket []float64, hour int){
	if len(valueBucket) <= 0{
		return
	}
	bucLen, curr := len(valueBucket), 0.0
	for _, value := range valueBucket{
		curr += value
	}
	
	
	curr /= float64(bucLen)
	// ratio := math.Pow(10,float64(4))
	// fmt.Println(rfcCustom(year,month, day,hour), math.Round(curr*ratio)/ratio)

	roundedValue := math.Round(curr*10000)/10000
	formatted := fmt.Sprintf("%.4f", roundedValue)
	fmt.Println(rfcCustom(year,month, day,hour),formatted)
}