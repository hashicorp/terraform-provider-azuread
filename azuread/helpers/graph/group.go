package graph

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
)

func GroupAllMembers(client graphrbac.GroupsClient, ctx context.Context, groupId string) ([]string, error) {
	it, err := client.GetGroupMembersComplete(ctx, groupId)

	if err != nil {
		return nil, fmt.Errorf("Error listing existing group members from Azure AD Group with ID %q: %+v", groupId, err)
	}

	existingMembers := make([]string, 0)

	var memberObjectID string
	for it.NotDone() {
		// possible members are users, groups or service principals
		// we try to 'cast' each result as the corresponding type and diff
		// if we found the object we're looking for
		user, _ := it.Value().AsUser()
		if user != nil {
			memberObjectID = *user.ObjectID
		}

		group, _ := it.Value().AsADGroup()
		if group != nil {
			memberObjectID = *group.ObjectID
		}

		servicePrincipal, _ := it.Value().AsServicePrincipal()
		if servicePrincipal != nil {
			memberObjectID = *servicePrincipal.ObjectID
		}

		existingMembers = append(existingMembers, memberObjectID)
		if err := it.NextWithContext(ctx); err != nil {
			return nil, fmt.Errorf("Error during pagination of group members from Azure AD Group with ID %q: %+v", groupId, err)
		}
	}

	log.Printf("[DEBUG] %d members in Azure AD group with ID: %q", len(existingMembers), groupId)

	return existingMembers, nil
}

func GroupAddMember(client graphrbac.GroupsClient, ctx context.Context, groupId string, member string) error {
	memberGraphURL := fmt.Sprintf("https://graph.windows.net/%s/directoryObjects/%s", client.TenantID, member)

	properties := graphrbac.GroupAddMemberParameters{
		URL: &memberGraphURL,
	}

	log.Printf("[DEBUG] Adding member with id %q to Azure AD group with id %q", member, groupId)
	if _, err := client.AddMember(ctx, groupId, properties); err != nil {
		return fmt.Errorf("Error adding group member %q to Azure AD Group with ID %q: %+v", member, groupId, err)
	}

	return nil
}

func GroupAddMembers(client graphrbac.GroupsClient, ctx context.Context, groupId string, members []string) error {
	for _, memberUuid := range members {
		err := GroupAddMember(client, ctx, groupId, memberUuid)

		if err != nil {
			return fmt.Errorf("Error while adding members to Azure AD Group with ID %q: %+v", groupId, err)
		}
	}

	return nil
}
