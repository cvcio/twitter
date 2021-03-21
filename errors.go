package twitter

import (
	"strconv"
	"strings"
)

type APIError struct {
	Code    int
	Message string
}

func NewAPIError(err error) *APIError {
	if err == nil {
		return &APIError{0, ""}
	}
	apiError := &APIError{0, err.Error()}

	if errorParts := strings.Split(err.Error(), " - "); len(errorParts) > 0 {
		code, e := strconv.Atoi(errorParts[len(errorParts)-1])
		if e != nil {
			code = 0
		}

		apiError.Code = code
		apiError.Message = errorParts[0]
	}

	return apiError
}
