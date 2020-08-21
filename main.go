package main

import (
	"fmt"

	"stripe.com/s/src/models"
	"stripe.com/s/src/router"

	"github.com/jinzhu/gorm"
	Config "stripe.com/s/src/config"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("error", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&models.Charge{})
	r := router.ChargeRouter()
	r.Run()

}
