package auth

import (
	"context"
	"go-skeleton-auth/internal/entity/auth"
	"go-skeleton-auth/pkg/errors"
	"go-skeleton-auth/pkg/httpclient"
	"net/http"
)

// Data ...
type Data struct {
	client  *httpclient.Client
	baseURL string
}

// New ...
func New(client *httpclient.Client, baseURL string) Data {
	d := Data{
		client:  client,
		baseURL: baseURL,
	}
	return d
}

// CheckAuth ...
func (d Data) CheckAuth(ctx context.Context, code string) (auth.Auth, error) {
	var auth auth.Auth
	var endpoint = "/checkrights"
	var url = d.baseURL + endpoint
	var body = map[string]interface{}{
		"code": code,
	}

	type ctxKey string
	_token := ctx.Value(ctxKey("_token")).(string)

	headers := make(http.Header)
	headers.Set("Authorization", _token)
	headers.Set("Content-Type", "application/json")

	_, err := d.client.PostJSON(ctx, url, endpoint, headers, body, &auth)
	if err != nil {
		return auth, errors.Wrap(err, "[DATA][CheckAuth]")
	}

	return auth, nil
}
