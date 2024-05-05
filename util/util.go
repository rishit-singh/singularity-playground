package util

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type Credentials struct {
	Username string
	Password string
}

type BasicAuthRequest struct {
	Request     *http.Request
	Credentials Credentials
}

func NewBasicAuthRequest(request *http.Request, credentials Credentials) *BasicAuthRequest {
	req := BasicAuthRequest{
		Request:     request,
		Credentials: credentials,
	}

	req.Request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", credentials.Username, credentials.Password)))))

	return &req
}
