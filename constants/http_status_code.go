package constants

import "net/http"

const (
	// HTTP_STATUS_OK [200, 2001] ...
	HTTP_STATUS_OK = http.StatusOK

	// HTTP_STATUS_OK // [400] ...
	HTTP_STATUS_BAD_REQUEST = http.StatusBadRequest

	// HTTP_STATUS_RESOURCE_NOT_FOUND [404] ...
	HTTP_STATUS_RESOURCE_NOT_FOUND = http.StatusNotFound

	// HTTP_STATUS_INTERVAL_SERVER_ERROR [500]...
	HTTP_STATUS_INTERVAL_SERVER_ERROR = http.StatusInternalServerError

	// HTTP_STATUS_REDIRECT [300]...
	HTTP_STATUS_REDIRECT = http.StatusMultipleChoices
)
