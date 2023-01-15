package packages

import (
	"bufio"
	"fmt"
	"net/http"
)

func CheckAns(arg1 string, arg2 string){
	fmt.Printf("\n ans \n \n")
	ansURL := fmt.Sprintf("https://tsserv.tinkermode.dev/hourly?begin=%s&end=%s", arg1, arg2 )
	responseAns , err := http.Get(ansURL)
	CheckError(err)
	ansBody := bufio.NewScanner(responseAns.Body)
	for ansBody.Scan(){
		fmt.Println(ansBody.Text())
	}
	fmt.Println("-----Ans Done-----")
}

func CheckError(e error){
	if e != nil{
		panic(e)
	}
}