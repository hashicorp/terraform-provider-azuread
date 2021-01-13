package microsoft

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func token(resp *http.Response) (*oauth2.Token, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}

	if c := resp.StatusCode; c < 200 || c > 299 {
		return nil, &oauth2.RetrieveError{
			Response: resp,
			Body:     body,
		}
	}

	// token response can arrive with numeric values as integers or strings :(
	var tokenRes struct {
		AccessToken string      `json:"access_token"`
		TokenType   string      `json:"token_type"`
		IDToken     string      `json:"id_token"`
		Resource    string      `json:"resource"`
		Scope       string      `json:"scope"`
		ExpiresIn   interface{} `json:"expires_in"` // relative seconds from now
		ExpiresOn   interface{} `json:"expires_on"` // timestamp
	}
	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}

	token := &oauth2.Token{
		AccessToken: tokenRes.AccessToken,
		TokenType:   tokenRes.TokenType,
	}
	var secs time.Duration
	if exp, ok := tokenRes.ExpiresIn.(string); ok && exp != "" {
		if v, err := strconv.Atoi(exp); err == nil {
			secs = time.Duration(v)
		}
	} else if exp, ok := tokenRes.ExpiresIn.(int64); ok {
		secs = time.Duration(exp)
	} else if exp, ok := tokenRes.ExpiresIn.(float64); ok {
		secs = time.Duration(exp)
	}
	if secs > 0 {
		token.Expiry = time.Now().Add(secs * time.Second)
	}

	return token, nil
}
