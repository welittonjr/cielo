package requests

import (
	"bytes"
	"fmt"
	"net/http"
)

// Request encapsula todas as informações necessárias
// para fazer uma solicitação HTTP para APIs cielo
type Request struct {
	HostnameTransacao string
	HostnameQuery     string
	MerchantId        string
	MerchantKey       string
	RequestId         string
	Debug             bool
}

func (request *Request) addRequestHeadersInternal(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		if key == "Content-Type" {
			continue
		}
		req.Header.Set(key, value)
	}
}

func (request *Request) addRequestHeaders(req *http.Request, headers map[string]string) {
	// Define os padrões primeiro no caso de indisponibilidade
	req.Header.Set("Content-Type", "application/json")

	// Defina os outros meios no cabeçalho
	request.addRequestHeadersInternal(req, headers)
}

// Fazer Request
func (request *Request) doRequest(req *http.Request) {

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	// if response.StatusCode >= constants.HTTP_STATUS_OK {

	// }

	// Debugs
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println()
	fmt.Println("response Status:", response.StatusCode)
	// fmt.Println()
	// fmt.Println("response Headers:", response.Header)
	// fmt.Println()
	// fmt.Println("response JSON Body:", string(body))
}

// Post ...
func (request *Request) Post(url string, payload string, headers map[string]string) {

	jsonStr := []byte(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	// Adicionando recursos extras no cabeçalho
	request.addRequestHeaders(req, headers)

	// Realizar Requisição
	request.doRequest(req)
}
