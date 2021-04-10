package msgraph

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// InvitationsClient performs operations on Invitations.
type InvitationsClient struct {
	BaseClient Client
}

// NewInvitationsClient returns a new InvitationsClient.
func NewInvitationsClient(tenantId string) *InvitationsClient {
	return &InvitationsClient{
		BaseClient: NewClient(VersionBeta, tenantId),
	}
}

// Create creates a new Invitation.
func (c *InvitationsClient) Create(ctx context.Context, invitation Invitation) (*Invitation, int, error) {
	var status int
	body, err := json.Marshal(invitation)
	if err != nil {
		return nil, status, err
	}
	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity:      "/invitations",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var newInvitation Invitation
	if err := json.Unmarshal(respBody, &newInvitation); err != nil {
		return nil, status, err
	}
	return &newInvitation, status, nil
}
