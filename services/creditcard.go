package services

import (
	"fmt"

	"github.com/welittonjr/cielo/constants"
	"github.com/welittonjr/cielo/requests"
	"github.com/welittonjr/cielo/types/capture"
)

type CreditCard struct {
	request *requests.Request
}

func Creditcard(r requests.Request) CreditCard {
	creditCard := CreditCard{
		request: &r,
	}
	return creditCard
}

func (c *CreditCard) Transaction(payload interface{}) (map[string]interface{}, error) {

	headers := map[string]string{
		"MerchantId":  c.request.MerchantId,
		"MerchantKey": c.request.MerchantKey,
		"RequestId":   c.request.RequestId,
	}

	host := c.request.HostnameTransacao

	baseURL := fmt.Sprintf("%s%s", host, constants.TRANSACTION)

	// Realizar Request
	return c.request.Post(baseURL, payload, headers)

}

func (c *CreditCard) CaptureSaleTransaction(payload capture.Sale) (map[string]interface{}, error) {

	headers := map[string]string{
		"MerchantId":  c.request.MerchantId,
		"MerchantKey": c.request.MerchantKey,
		"RequestId":   c.request.RequestId,
	}

	host := c.request.HostnameTransacao

	baseURL := fmt.Sprintf("%s/1/sales/%s/capture", host, payload.PaymentId)

	return c.request.Put(baseURL, payload, headers)

}
