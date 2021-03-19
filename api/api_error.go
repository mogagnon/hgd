package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	*http.Response
}

func (response *Response) HandleApiError(message string) error {
	if !isSuccessfulResponse(response.StatusCode) {
		var apiError Error
		_ = json.NewDecoder(response.Body).Decode(&apiError)

		return fmt.Errorf("%s: %d %s", message, response.StatusCode, apiError.Message)
	}

	return nil
}

func isSuccessfulResponse(statusCode int) bool {
	var successCode = []int{200, 201, 202, 204}

	for _, code := range successCode {
		if statusCode == code {
			return true
		}
	}

	return false
}
