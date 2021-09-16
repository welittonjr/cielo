## cielo

Client para a API 3.0 da Cielo em Golang **[ Em Desenvolvimento ]**

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![Open Source Love svg2](https://img.shields.io/static/v1?label=Open&message=Source&color=green)](https://img.shields.io/)


## Índice
#### [Cartão de Crédito](#creditCard)
+ [Criando uma transação simples](#creditSimpleTransaction)

## <a name="creditCard"></a> Cartão de Crédito

### <a name="creditSimpleTransaction"></a>  Criando uma transação


```golang
data := `{
		"MerchantOrderId":"2014111703",
		"Customer":{
			 "Name":"Comprador crédito simples"
		},
		"Payment":{
			"Type":"CreditCard",
			"Amount":15700,
			"Installments":1,
			"SoftDescriptor":"123456789ABCD",
			"CreditCard":{
					"CardNumber":"1234123412341231",
					"Holder":"Teste Holder",
					"ExpirationDate":"12/2030",
					"SecurityCode":"123",
					"Brand":"Visa",
					"CardOnFile":{
						 "Usage": "Used",
						 "Reason":"Unscheduled"
					}
			},
			"IsCryptoCurrencyNegotiation": true
		}
 }` 

data, err := cielo.Creditcard.Transaction(data)

if err != nil {
  panic(err)
}

```