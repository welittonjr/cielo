package errors

import (
	"fmt"
	"strconv"
)

// CieloError ...
type CieloError []struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

// InternalServerError ...
type InternalServerError struct {
	Message string
	Code    int
}

// ResourceNotFound ...
type ResourceNotFound struct {
	Message string
	Code    int
}

// BadRequestError ...
type BadRequestError struct {
	Message string
	Code    int
}

func handleError(msg string, code int) string {
	str1 := "Message:"
	str2 := "Code:"
	return fmt.Sprintf("%s %s %s %s", str1, msg, str2, strconv.Itoa(code))
}

// Error ...
func (s *InternalServerError) Error() string {
	return handleError(s.Message, s.Code)
}

func (s *ResourceNotFound) Error() string {
	return handleError(s.Message, s.Code)
}

func (s *BadRequestError) Error() string {
	return handleError(s.Message, s.Code)
}
