package custompollers

import (
	"context"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	flexibleFederatedIdentityCredential "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
)

var _ pollers.PollerType = &applicationFlexibleFederatedCredentialCreationPoller{}

type applicationFlexibleFederatedCredentialCreationPoller struct {
	client *flexibleFederatedIdentityCredential.FederatedIdentityCredentialClient
	id     beta.ApplicationIdFederatedIdentityCredentialId
}

func NewApplicationFlexibleFederatedCredentialCreationPoller(client *flexibleFederatedIdentityCredential.FederatedIdentityCredentialClient, id beta.ApplicationIdFederatedIdentityCredentialId) *applicationFlexibleFederatedCredentialCreationPoller {
	return &applicationFlexibleFederatedCredentialCreationPoller{
		client: client,
		id:     id,
	}
}

func (a applicationFlexibleFederatedCredentialCreationPoller) Poll(ctx context.Context) (*pollers.PollResult, error) {
	resp, err := a.client.GetFederatedIdentityCredential(ctx, a.id, flexibleFederatedIdentityCredential.DefaultGetFederatedIdentityCredentialOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return &pollers.PollResult{
				PollInterval: 1 * time.Second,
				Status:       pollers.PollingStatusInProgress,
			}, nil
		}
		return &pollers.PollResult{
			Status: pollers.PollingStatusFailed,
		}, err
	}

	credential := resp.Model
	if credential == nil {
		return &pollers.PollResult{
			PollInterval: 1 * time.Second,
			Status:       pollers.PollingStatusInProgress,
		}, nil
	}

	return &pollers.PollResult{
		PollInterval: 1 * time.Second,
		Status:       pollers.PollingStatusSucceeded,
	}, nil
}
