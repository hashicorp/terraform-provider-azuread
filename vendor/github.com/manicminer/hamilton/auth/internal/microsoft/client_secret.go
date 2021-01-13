package microsoft

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

type secretSource struct {
	ctx  context.Context
	conf *Config
}

func (a secretSource) Token() (*oauth2.Token, error) {
	hc := oauth2.NewClient(a.ctx, nil)
	v := url.Values{
		"client_id":     {a.conf.ClientID},
		"client_secret": {a.conf.ClientSecret},
		"grant_type":    {"client_credentials"},
	}
	if a.conf.Resource != "" {
		v["resource"] = []string{a.conf.Resource}
	} else {
		v["scope"] = []string{strings.Join(a.conf.Scopes, " ")}
	}
	resp, err := hc.PostForm(a.conf.TokenURL, v)
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}

	return token(resp)
}
