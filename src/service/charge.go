package Service

import (
	_ "github.com/go-sql-driver/mysql"
	Config "stripe.com/s/src/config"
	"stripe.com/s/src/models"
)

func SavePayment(charge *models.Charge) (err error) {
	if err = Config.DB.Create(charge).Error; err != nil {
		return err
	}
	return nil

}
