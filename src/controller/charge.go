package controller

import (
	"fmt"
	"net/http"

	Service "stripe.com/s/src/service"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"stripe.com/s/src/models"
)

func Payment(c *gin.Context) {
	var payment models.Charge
	c.BindJSON(&payment)

	apiKey := "sk_test_51HHWVYBB3FZLOZ1moSSahrY2wOpWXWmAsDNWHBJzxUJSocWEACxNs3e2SwIXVFJlV3Sp1HsWcinVpuA4xs0X5kMg00VQ29dSwD"
	fmt.Println(apiKey + "asdasd")
	stripe.Key = apiKey
	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(payment.Amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(payment.ProductName),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(payment.ReceiptEmail)})

	if err != nil {
		c.String(http.StatusBadRequest, "Payment Unsuccessfull")
		return
	}
	err1 := Service.SavePayment(&payment)
	if err1 != nil {
		c.String(http.StatusNotFound, "error occured")

	}

}
