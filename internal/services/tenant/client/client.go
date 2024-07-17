package tenant

import "github.com/hashicorp/terraform-provider-azuread/internal/common"

type Client struct {
}

func NewClient(o *common.ClientOptions) *Client {
	return &Client{}
}
