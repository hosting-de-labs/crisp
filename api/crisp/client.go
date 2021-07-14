package crisp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/twitchtv/twirp"
)

const (
	// DefaultAuthHeader defaults the default authentication header used to authorize requests
	DefaultAuthHeader string = "Authorization"
)

type Client struct {
	Crisp

	token string
}

func NewCrispClient(baseURL string) *Client {
	cl := &Client{}

	cl.Crisp = NewCrispProtobufClient(
		baseURL,
		&http.Client{},
		twirp.WithClientInterceptors(
			authInterceptor(
				DefaultAuthHeader,
				func() string { return cl.token },
			),
		),
	)
	return cl
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func authInterceptor(header string, tknFnc func() string) twirp.Interceptor {
	return func(next twirp.Method) twirp.Method {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if m, ok := twirp.MethodName(ctx); ok && m != "Register" {
				var err error
				ctx, err = withAuthContext(ctx, header, tknFnc)
				if err != nil {
					return nil, err
				}
			}

			return next(ctx, req)
		}
	}
}

func withAuthContext(ctx context.Context, headerName string, tknFnc func() string) (context.Context, error) {
	header := make(http.Header)
	header.Set(headerName, tknFnc())

	newCtx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		return nil, fmt.Errorf("twirp error setting headers: %s", err)
	}

	return newCtx, nil
}
