package complete

type Customer struct {
	Name            string `json:"Name"`
	Email           string `json:"Email"`
	Birthdate       string `json:"Birthdate"`
	Address         *Address
	DeliveryAddress *DeliveryAddress
}

type Address struct {
	Street     string `json:"Street"`
	Number     string `json:"Number"`
	Complement string `json:"Complement"`
	ZipCode    string `json:"ZipCode"`
	City       string `json:"City"`
	State      string `json:"State"`
	Country    string `json:"Country"`
}

type DeliveryAddress struct {
	Street     string `json:"Street"`
	Number     string `json:"Number"`
	Complement string `json:"Complement"`
	ZipCode    string `json:"ZipCode"`
	City       string `json:"City"`
	State      string `json:"State"`
	Country    string `json:"Country"`
}

type Payment struct {
	Currency                    string `json:"Currency"`
	Country                     string `json:"Country"`
	ServiceTaxAmount            int    `json:"ServiceTaxAmount"`
	Installments                int    `json:"Installments"`
	Interest                    string `json:"Interest"`
	Capture                     bool   `json:"Capture"`
	Authenticate                bool   `json:"Authenticate"`
	SoftDescriptor              string `json:"SoftDescriptor"`
	CreditCard                  *CreditCard
	IsCryptoCurrencyNegotiation bool   `json:"IsCryptoCurrencyNegotiation"`
	Type                        string `json:"Type"`
	Amount                      int    `json:"Amount"`
	AirlineData                 *AirlineData
}

type AirlineData struct {
	TicketNumber string `json:"TicketNumber"`
}

type CardOnFile struct {
	Usage  string `json:"Usage"`
	Reason string `json:"Reason"`
}

type CreditCard struct {
	CardNumber     string `json:"CardNumber"`
	Holder         string `json:"Holder"`
	ExpirationDate string `json:"ExpirationDate"`
	SecurityCode   string `json:"SecurityCode"`
	SaveCard       string `json:"SaveCard"`
	Brand          string `json:"Brand"`
	CardOnFile     *CardOnFile
}

type Transaction struct {
	MerchantOrderID string `json:"MerchantOrderId"`
	Customer        *Customer
	Payment         *Payment
}
