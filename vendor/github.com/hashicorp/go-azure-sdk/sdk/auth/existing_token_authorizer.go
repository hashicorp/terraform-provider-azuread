package auth

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type ExistingTokenAuthorizerOptions struct {
	//Api describes the Azure API being used
	Api environments.Api

	//Token is the pre-existing token to use.
	Token string

	//RefreshToken is the pre-existing refresh token to use.
	RefreshToken string
}

func NewExistingTokenAuthorizer(ctx context.Context, options ExistingTokenAuthorizerOptions) (Authorizer, error) {
	resource, err := environments.Resource(options.Api)

	if err != nil {
		return nil, fmt.Errorf("determining resource for api %q: %+v", options.Api.Name(), err)
	}

	conf, err := 

}
