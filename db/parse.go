package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct which contains
// an array of users
type Objects struct {
	Objects []User `json:"user"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	LastName   string `json:"last_name"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Gender     string `json:"gender"`
	Birth_date string `json:"birth_date"`
}



func Parse() {
	// Open our jsonFile
	jsonFile, err := os.Open("users_go.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Objects

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Objects); i++ {
		//fmt.Println("User Type: " + users.Users[i].Type)
		//fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		//fmt.Println("User Name: " + users.Users[i].Name)
		//fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)

	}

}