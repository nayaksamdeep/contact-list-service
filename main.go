/*
 * This program creates a webapp to store and retrieve visiting cards
 */

package main

import  (
     "fmt"
     "log"
     "github.com/nayaksamdeep/contact-list-service/Config"
     "github.com/nayaksamdeep/contact-list-service/Models"
     "github.com/nayaksamdeep/contact-list-service/Routes"
     "github.com/jinzhu/gorm"
     _"github.com/mattn/go-sqlite3"
)

var err error

func main() {

	// Creating a connection to the database
	Config.DB, err = gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
                log.Fatal(err)
	}

	defer Config.DB.Close()

	// run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Contact{})

        fmt.Println("hi there! Welcome to Contact List Service\n");

        r := Routes.SetupRouter()
        // running
        r.Run() 
}
