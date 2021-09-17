package services

import (
	"fmt"

	"github.com/welittonjr/cielo/requests"
)

type CreditCard struct {
	Request *requests.Request
}

func Creditcard(r requests.Request) CreditCard {
	creditCard := CreditCard{
		Request: &r,
	}
	return creditCard
}

func (c *CreditCard) Transaction(payload string) (map[string]interface{}, error) {

	headers := map[string]string{
		"MerchantId":  c.Request.MerchantId,
		"MerchantKey": c.Request.MerchantKey,
		"RequestId":   c.Request.RequestId,
	}

	host := c.Request.HostnameTransacao

	baseURL := fmt.Sprintf("%s%s", host, "/1/sales/")

	// Realizar Request
	return c.Request.Post(baseURL, payload, headers)

}
