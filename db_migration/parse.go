package main

import (
	"context"
	"date_users_app/config"
	"date_users_app/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Users struct which contains
// an array of users
type Objects struct {
	Objects []User `json:"objects"`
}

type User struct {
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
}

func parse() error {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	conn := services.InitDB()
	collection := conn.Collection("users")

	// Open our jsonFile
	jsonFile, err := os.Open("db_migration/users_go.json")
	// if we os.Open returns an error then handle it
	if err != nil {

		return err
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
	err = json.Unmarshal(byteValue, &users)
	if err != nil {

		return err

	}

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for _, user := range users.Objects {
		_, err := collection.InsertOne(context.Background(), user, nil)
		if err != nil {

			return err
		}

	}
	return nil
}

func main() {
	parse()

}
