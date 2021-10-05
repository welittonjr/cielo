## cielo

Client para a API 3.0 da Cielo em Golang **[ Em Desenvolvimento ]**

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![Open Source Love svg2](https://img.shields.io/static/v1?label=Open&message=Source&color=green)](https://img.shields.io/)


## Índice

#### [Início]()
+ [Como Utilizar](#howuse)
+ [Paramêtros de criação](#params)

#### [Cartão de Crédito](#creditCard)
+ [Criando uma transação simples](#creditSimpleTransaction)
+ [Criando uma transação completa](#creditCompleteTransaction)
+ [Capturando uma venda](#creditSaleCapture)

## <a name="howuse"></a> Como utilizar?

### Iniciando
```go
import "github.com/welittonjr/cielo"

credentials := cielo.Credentials{
		MerchantId:  "xxxxxxxxxxxxxxxxxxxxxxx",
		MerchantKey: "xxxxxxxxxxxxxxxxxxxxxxxxxx",
		RequestId:   "xxxxxxx", // Opcional - Identificação do Servidor na Cielo
		Sandbox:     true, // Opcional - Ambiente de Testes
		Debug:       true, // Opcional - Exibe os dados enviados na requisição para a Cielo
}

cielo := cielo.New(credentials)
```

### <a name="params"></a> Paramêtros de criação

| Campo | Descrição | Obrigatório? | Default |
| :-------------: |:-------------:| :-----:| :-----:|
| merchantId | Identificador da loja na Cielo. | Sim | null |
| merchantKey | Chave publica para autenticação dupla na Cielo. | Sim | null |
| requestId | Identificador do Request, utilizado quando o lojista usa diferentes servidores para cada GET/POST/PUT. | Não | null |
| sandbox | Ambiente de testes da Cielo | Não | false |
| debug | Exibe requisição da transação no console | Não | false |

## <a name="creditCard"></a> Cartão de Crédito

### <a name="creditSimpleTransaction"></a>  Criando uma transação simples

```golang
	import (
		"github.com/welittonjr/cielo"
		"github.com/welittonjr/cielo/types/simple"
	)

	ts := simple.Transaction{
		MerchantOrderID: "2014111703",
		Customer: &simple.Customer{
			Name: "Comprador crédito simples",
		},
		Payment: &simple.Payment{
			Type:           "CreditCard",
			Amount:         15700,
			Installments:   1,
			SoftDescriptor: "123456789ABCD",
			CreditCard: &simple.CreditCard{
				CardNumber:     "1234123412341231",
				Holder:         "Teste Holder",
				ExpirationDate: "12/2030",
				SecurityCode:   "123",
				Brand:          "Visa",
				CardOnFile: &simple.CardOnFile{
					Usage:  "Used",
					Reason: "Unscheduled",
				},
			},
			IsCryptoCurrencyNegotiation: true,
		},
	}

response, err := cielo.Creditcard.Transaction(ts)

if err != nil {
	log.Fatal(err)
}

fmt.Println(response)

```

### <a name="creditCompleteTransaction"></a>  Criando uma transação completa

```golang
	import (
		"github.com/welittonjr/cielo"
		"github.com/welittonjr/cielo/types/complete"
	)

	tc := complete.Transaction{
		MerchantOrderID: "2014111701",
		Customer: &complete.Customer{
			Name:      "Comprador crédito completo",
			Email:     "compradorteste@teste.com",
			Birthdate: "1991-01-02",
			Address: &complete.Address{
				Street:     "Rua Teste",
				Number:     "123",
				Complement: "AP 123",
				ZipCode:    "12345987",
				City:       "Rio de Janeiro",
				State:      "RJ",
				Country:    "BRA",
			},
			DeliveryAddress: &complete.DeliveryAddress{
				Street:     "Rua Teste",
				Number:     "123",
				Complement: "AP 123",
				ZipCode:    "12345987",
				City:       "Rio de Janeiro",
				State:      "RJ",
				Country:    "BRA",
			},
		},
		Payment: &complete.Payment{
			Currency:         "BRL",
			Country:          "BRA",
			ServiceTaxAmount: 0,
			Installments:     1,
			Interest:         "ByMerchant",
			Capture:          true,
			Authenticate:     false,
			SoftDescriptor:   "123456789ABCD",
			CreditCard: &complete.CreditCard{
				CardNumber:     "1234123412341231",
				Holder:         "Teste Holder",
				ExpirationDate: "12/2030",
				SecurityCode:   "123",
				SaveCard:       "false",
				Brand:          "Visa",
				CardOnFile: &complete.CardOnFile{
					Usage:  "Used",
					Reason: "Unscheduled",
				},
			},
			IsCryptoCurrencyNegotiation: true,
			Type:                        "CreditCard",
			Amount:                      15700,
			AirlineData: &complete.AirlineData{
				TicketNumber: "AR988983",
			},
		},
	}

	response, err := cielo.Creditcard.Transaction(tc)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
```

### <a name="creditSaleCapture"></a>  Capturando uma venda

```golang

	import (
		"github.com/welittonjr/cielo"
		"github.com/welittonjr/cielo/types/capture"
	)

	captura := capture.Sale{
		PaymentId: "15906ada-6ae1-487b-9101-1f6f5216b698",
		Amount:    2000,
	}

	response, err := cielo.Creditcard.CaptureSaleTransaction(captura)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

```