package crisp

import (
	"context"
	"fmt"
	"net/http"
	"path"

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
		twirp.WithClientHooks(hooks(func() string { return cl.token })),
	)
	return cl
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func hooks(tknFnc func() string) *twirp.ClientHooks {
	return &twirp.ClientHooks{
		RequestPrepared: func(ctx context.Context, request *http.Request) (context.Context, error) {
			rqMethod := path.Base(request.RequestURI)
			if rqMethod != "Register" {
				return withAuthContext(ctx, DefaultAuthHeader, tknFnc)
			}

			return ctx, nil
		},
		ResponseReceived: nil,
		Error:            nil,
	}
}

func withAuthContext(ctx context.Context, headerName string, tknFnc func() string) (context.Context, error) {
	header := make(http.Header)
	header.Set(headerName, tknFnc())

	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		return nil, fmt.Errorf("twirp error setting headers: %s", err)
	}

	return ctx, nil
}
