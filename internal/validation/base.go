package validation

import (
	"errors"
	"net/http"
)

var methodAllowedError = "Only POST and GET methods allowed"

func Validate(req *http.Request) (bool, error) {
	if req.Method != http.MethodPost && req.Method != http.MethodGet {
		return false, errors.New(methodAllowedError)
	}

	return true, nil
}
