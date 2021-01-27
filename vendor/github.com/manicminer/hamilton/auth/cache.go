package auth

import (
	"sync"

	"golang.org/x/oauth2"
)

// cachedAuthorizer caches a token until it expires, then acquires a new token from source
type cachedAuthorizer struct {
	source Authorizer
	mutex  sync.Mutex
	token  *oauth2.Token
}

// Token returns the current token if it's still valid, else will acquire a new token
func (c *cachedAuthorizer) Token() (*oauth2.Token, error) {
	if c.token == nil || !c.token.Valid() {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		token, err := c.source.Token()
		if err != nil {
			return nil, err
		}
		c.token = token
	}
	return c.token, nil
}

// CachedAuthorizer returns an Authorizer that caches an access token for the duration of its validity.
// If the cached token expires, a new one is acquired and cached.
func CachedAuthorizer(src Authorizer) Authorizer {
	return &cachedAuthorizer{
		source: src,
	}
}
