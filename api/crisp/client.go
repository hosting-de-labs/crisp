package crisp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/twitchtv/twirp"
)

const (
	DefaultAuthHeader string = "Authorization"
)

type Client struct {
	Crisp
}

func NewCrispClient(baseURL string, token string) *Client {
	c := NewCrispProtobufClient(
		baseURL,
		&http.Client{},
		twirp.WithClientInterceptors(authInterceptor(DefaultAuthHeader, token)),
	)

	return &Client{c}
}

func authInterceptor(headerName string, token string) twirp.Interceptor {
	return func(next twirp.Method) twirp.Method {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			header := make(http.Header)
			header.Set(headerName, token)

			ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
			if err != nil {
				return nil, fmt.Errorf("twirp error setting headers: %s", err)
			}

			return next(ctx, request)
		}
	}
}
