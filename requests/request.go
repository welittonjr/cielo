package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/welittonjr/cielo/constants"
	"github.com/welittonjr/cielo/errors"
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

func processResponse(response *http.Response) (map[string]interface{}, error) {

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	resp := make(map[string]interface{})

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Fazer Request
func (request *Request) doRequest(req *http.Request) (map[string]interface{}, error) {

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == constants.HTTP_STATUS_OK {
		return processResponse(response)
	}

	// Gera um erro dependendo do tipo de erro na resposta
	var responseError errors.CieloError
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &responseError)
	errorCielo := responseError[0]

	switch response.StatusCode {
	case constants.HTTP_STATUS_INTERVAL_SERVER_ERROR:
		return nil, &errors.InternalServerError{
			Message: errorCielo.Message,
			Code:    errorCielo.Code,
		}

	case constants.HTTP_STATUS_RESOURCE_NOT_FOUND:
		return nil, &errors.ResourceNotFound{
			Message: errorCielo.Message,
			Code:    errorCielo.Code,
		}

	case constants.HTTP_STATUS_BAD_REQUEST:
		return nil, &errors.BadRequestError{
			Message: errorCielo.Message,
			Code:    errorCielo.Code,
		}
	}

	return processResponse(response)
}

// Post ...
func (request *Request) Post(url string, payload string, headers map[string]string) (map[string]interface{}, error) {

	jsonStr := []byte(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	// Adicionando recursos extras no cabeçalho
	request.addRequestHeaders(req, headers)

	// Realizar Requisição
	return request.doRequest(req)
}
