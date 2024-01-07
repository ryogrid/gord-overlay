// Code generated by ogen, DO NOT EDIT.

package api_external

import (
	"net/http"

	"github.com/ogen-go/ogen/validate"
)

func decodeExternalServiceDeleteValueResponse(resp *http.Response) (res *ExternalServiceDeleteValueOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ExternalServiceDeleteValueOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeExternalServiceFindHostForKeyResponse(resp *http.Response) (res *ExternalServiceFindHostForKeyOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ExternalServiceFindHostForKeyOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeExternalServiceGetValueResponse(resp *http.Response) (res *ExternalServiceGetValueOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ExternalServiceGetValueOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeExternalServicePutValueResponse(resp *http.Response) (res *ExternalServicePutValueOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &ExternalServicePutValueOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
