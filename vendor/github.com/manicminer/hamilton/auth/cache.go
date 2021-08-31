package auth

import (
	"sync"

	"golang.org/x/oauth2"
)

// CachedAuthorizer caches a token until it expires, then acquires a new token from Source
type CachedAuthorizer struct {
	// Source contains the underlying Authorizer for obtaining tokens
	Source Authorizer

	mutex sync.RWMutex
	token *oauth2.Token
}

// Token returns the current token if it's still valid, else will acquire a new token
func (c *CachedAuthorizer) Token() (*oauth2.Token, error) {
	c.mutex.RLock()
	valid := c.token != nil && c.token.Valid()
	c.mutex.RUnlock()

	if !valid {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		token, err := c.Source.Token()
		if err != nil {
			return nil, err
		}
		c.token = token
	}

	return c.token, nil
}

// NewCachedAuthorizer returns an Authorizer that caches an access token for the duration of its validity.
// If the cached token expires, a new one is acquired and cached.
func NewCachedAuthorizer(src Authorizer) Authorizer {
	return &CachedAuthorizer{
		Source: src,
	}
}
