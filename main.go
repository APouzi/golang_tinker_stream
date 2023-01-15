package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/APouzi/golang_project_folder/packages"
)

//2021-03-04T03:45:00Z 2021-03-04T04:17:00Z were the examples
func main() {

	argLen := len(os.Args)
	
	if argLen != 3 {
		fmt.Fprintln(os.Stderr,"please provide two arguments for this program")
		return
	}
	arg1, arg2 := os.Args[1], os.Args[2]
	// // This will check the answer against the stream
	packages.CheckAns(arg1, arg2)
	rfc, err := time.Parse(time.RFC3339, arg1)
	CheckError(err)
	rfc2 , err := time.Parse(time.RFC3339, arg2)
	CheckError(err)
	if rfc2.Before(rfc){
		fmt.Fprintln(os.Stderr, "the first argument must come before the 2nd argument")
		return
	}
	if rfc2.Minute() != 0 || rfc2.Second() != 0 || rfc.Minute() != 0 || rfc.Second() != 0{
		fmt.Fprintln(os.Stderr, "Warning! because Second and Minute are not '00' this means that the argument will be rounded up and rounded down respectively.")
		//Instructions says to roundup and rounddown arg1 and arg2 respectively to get the entire hours, as I understood it. 
	}

	arg1 = packages.RFCRoundDown(arg1)
	arg2 = packages.RFCRoundUp(arg2)
	reqURL := fmt.Sprintf("https://tsserv.tinkermode.dev/data?begin=%s&end=%s", arg1, arg2)

	TimeStampProccesser(reqURL, arg1, arg2)


}


func TimeStampProccesser(reqURL string, arg1 string, arg2 string){
	
	//request to endpoint
	response, err := http.Get(reqURL)
	CheckError(err)

	defer response.Body.Close()

	// turn the response into a scanner iterable, use scanner because of possible memory limitations.
	scanBody := bufio.NewScanner(response.Body)


	bucketValue := []float64{}
	firstHour := packages.GetFirstHour(arg1)
	hourCompare := firstHour
	var value float64
	var year, day int
	var month time.Month
	for scanBody.Scan(){
		line := scanBody.Text()
		splitLine := strings.Split(line, " ")

		rfc, err := time.Parse(time.RFC3339, splitLine[0])
		CheckError(err)
		year, month, day = rfc.Date()
		hour := rfc.Hour()
		// This takes care of the single space or double space formatting that happens randomly
		if splitLine[1] != ""{
			value, err = strconv.ParseFloat(splitLine[1], 64)
		}else{
			value, err = strconv.ParseFloat(splitLine[2], 64)
		}
		
		CheckError(err)

		//This is how I create buckets depending on different buckets. Only issue is the last bucket so far. 
		if hourCompare != hour{
			packages.CreateBucket(year,month, day, bucketValue,hourCompare)
			hourCompare = hour
			bucketValue = bucketValue[:0]
		}
		bucketValue = append(bucketValue, value)
		

	}
	// bucketValue = append(bucketValue, value)
	packages.CreateBucket(year,month, day, bucketValue,hourCompare)
	

	

}





func CheckError(e error){
	if e != nil{
		panic(e)
	}
}


