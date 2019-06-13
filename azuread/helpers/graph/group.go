package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
)

func GroupAllMembers(groupId string, client graphrbac.GroupsClient, ctx context.Context) ([]string, error) {
	it, err := client.GetGroupMembersComplete(ctx, groupId)

	if err != nil {
		return nil, fmt.Errorf("Error listing existing group members from Azure AD Group with ID %q: %+v", groupId, err)
	}

	existingMembers := make([]string, 0)

	for it.NotDone() {
		currUser, _ := it.Value().AsUser()
		existingMembers = append(existingMembers, *currUser.ObjectID)
		if err := it.NextWithContext(ctx); err != nil {
			return nil, fmt.Errorf("Error during pagination of group members from Azure AD Group with ID %q: %+v", groupId, err)
		}
	}

	log.Printf("[DEBUG] %d members in Azure AD group with ID: %q", len(existingMembers), groupId)

	return existingMembers, nil
}

func GroupAddMember(groupId string, member string, client graphrbac.GroupsClient, ctx context.Context) error {
	memberGraphURL := fmt.Sprintf("https://graph.windows.net/%s/directoryObjects/%s", client.TenantID, member)

	properties := graphrbac.GroupAddMemberParameters{
		URL: &memberGraphURL,
	}

	log.Printf("[DEBUG] Adding member with id %q to Azure AD group with id %q", member, groupId)
	if _, err := client.AddMember(ctx, groupId, properties); err != nil {
		return err
	}

	return nil
}
