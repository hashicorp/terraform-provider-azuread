package auth

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/oauth2"
	"strings"
)

type Claims struct {
	Audience          string   `json:"aud"`
	Issuer            string   `json:"iss"`
	IdentityProvider  string   `json:"idp"`
	ObjectId          string   `json:"oid"`
	Roles             []string `json:"roles"`
	Subject           string   `json:"sub"`
	TenantRegionScope string   `json:"tenant_region_scope"`
	TenantId          string   `json:"tid"`
	Version           string   `json:"ver"`

	AppDisplayName string `json:"app_displayname,omitempty"`
	AppId          string `json:"appid,omitempty"`
	IdType         string `json:"idtyp,omitempty"`
}

func ParseClaims(token *oauth2.Token) (claims Claims, err error) {
	if token == nil {
		return
	}
	jwt := strings.Split(token.AccessToken, ".")
	payload, err := base64.StdEncoding.DecodeString(jwt[1])
	if err != nil {
		return
	}
	err = json.Unmarshal(payload, &claims)
	return
}
