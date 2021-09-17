package cielo

import (
	"github.com/welittonjr/cielo/requests"
	"github.com/welittonjr/cielo/services"
)

type Credentials struct {
	MerchantId  string
	MerchantKey string
	Debug       bool
	Sandbox     bool
	RequestId   string
}

type cielo struct {
	Creditcard services.CreditCard
}

func New(credentials Credentials) cielo {
	hosts := getHostNames(credentials.Debug)
	hostnameTransacao := hosts[0]
	hostnameQuery := hosts[1]

	httpRequest := requests.Request{
		HostnameTransacao: hostnameTransacao,
		HostnameQuery:     hostnameQuery,
		MerchantId:        credentials.MerchantId,
		MerchantKey:       credentials.MerchantKey,
		RequestId:         credentials.RequestId,
		Debug:             credentials.Debug,
	}

	return cielo{
		Creditcard: services.Creditcard(httpRequest),
	}
}

func getHostNames(sandbox bool) []string {
	if sandbox {
		return []string{
			"https://apisandbox.cieloecommerce.cielo.com.br",
			"https://apiquerysandbox.cieloecommerce.cielo.com.br",
		}
	} else {
		return []string{
			"https://api.cieloecommerce.cielo.com.br",
			"https://apiquery.cieloecommerce.cielo.com.br",
		}
	}
}
