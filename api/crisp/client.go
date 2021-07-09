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
}

func NewCrispClient(baseURL string, token string) *Client {
	c := NewCrispProtobufClient(
		baseURL,
		&http.Client{},
		twirp.WithClientHooks(hooks(token)),
	)

	return &Client{c}
}

func hooks(token string) *twirp.ClientHooks {
	return &twirp.ClientHooks{
		RequestPrepared: func(ctx context.Context, request *http.Request) (context.Context, error) {
			rqMethod := path.Base(request.RequestURI)
			if rqMethod != "Register" {
				return withAuthContext(ctx, DefaultAuthHeader, token)
			}

			return ctx, nil
		},
		ResponseReceived: nil,
		Error:            nil,
	}
}

func withAuthContext(ctx context.Context, headerName string, token string) (context.Context, error) {
	header := make(http.Header)
	header.Set(headerName, token)

	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)
	if err != nil {
		return nil, fmt.Errorf("twirp error setting headers: %s", err)
	}

	return ctx, nil
}
