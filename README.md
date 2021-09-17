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

### <a name="creditSimpleTransaction"></a>  Criando uma transação


```go
import (
	"github.com/welittonjr/cielo"
	"github.com/welittonjr/cielo/type/transaction"
)

ts := transaction.Simple{
	MerchantOrderID: "2014111703",
	Customer: &transaction.Customer{
		Name: "Comprador crédito simples",
	},
	Payment: &transaction.Payment{
		Type:           "CreditCard",
		Amount:         15700,
		Installments:   1,
		SoftDescriptor: "123456789ABCD",
		CreditCard: &transaction.CreditCard{
			CardNumber:     "1234123412341231",
			Holder:         "Teste Holder",
			ExpirationDate: "12/2030",
			SecurityCode:   "123",
			Brand:          "Visa",
			CardOnFile: &transaction.CardOnFile{
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