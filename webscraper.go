package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"errors"
)

type User struct {
	Login, Html_url string
}

func main() {

	var userid string
	fmt.Println("Enter a valid github username")
	fmt.Scanln(&userid)

	url := "https://api.github.com/users/" + userid + "/followers"

	content := serverContent(url)
	users := getUserInfo(content)

	for _,user := range users {
			//Parsing prices / numbers "£24.00" with ($%.2f)
			//price, _, _ := big.ParseFloat(user.Price, 10 ,2, big.ToZero)
			fmt.Printf("\n%v %v\n", user.Login, user.Html_url)
	}

}

func checkError (err error){
	if err !=nil{
		invUser := errors.New("Username is Invalid.")
		panic(invUser)
	}
}

func serverContent(url string) string {
	resp, err := http.Get(url)
	checkError(err)


	defer resp.Body.Close()

	bytes, err :=  ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func getUserInfo(content string) []User {
	users := make([]User,0,10)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var user User
	for decoder.More() {
		err := decoder.Decode(&user)
		checkError(err)
		users = append(users, user)
	}

	return users
}