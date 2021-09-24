package simple

type Customer struct {
	Name string `json:"Name"`
}

type Payment struct {
	Type                        string `json:"Type"`
	Amount                      int    `json:"Amount"`
	Installments                int    `json:"Installments"`
	SoftDescriptor              string `json:"SoftDescriptor"`
	CreditCard                  *CreditCard
	IsCryptoCurrencyNegotiation bool `json:"IsCryptoCurrencyNegotiation"`
}

type CreditCard struct {
	CardNumber     string `json:"CardNumber"`
	Holder         string `json:"Holder"`
	ExpirationDate string `json:"ExpirationDate"`
	SecurityCode   string `json:"SecurityCode"`
	Brand          string `json:"Brand"`
	CardOnFile     *CardOnFile
}

type CardOnFile struct {
	Usage  string `json:"Usage"`
	Reason string `json:"Reason"`
}

type Transaction struct {
	MerchantOrderID string `json:"MerchantOrderId"`
	Customer        *Customer
	Payment         *Payment
}
