package utils

import "errors"

var (
	ErrorPasswordVerify = errors.New("Password don't match.")
	ErrorTokenExpire    = errors.New("Auth token is expired.")
)

type errorString struct {
	ErrorPasswordVerify string
	ErrorTokenExpire    string
}

func (e *errorString) Error(err error) string {
	switch err {
	case ErrorPasswordVerify:
		return e.ErrorPasswordVerify
	case ErrorTokenExpire:
		return e.ErrorTokenExpire
	}
	return "Unknown error."
}
