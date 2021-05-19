package msgraph

import (
	"context"
	"fmt"
	"strings"

	"github.com/manicminer/hamilton/msgraph"
)

func GroupCheckNameAvailability(ctx context.Context, client *msgraph.GroupsClient, displayName string, existingID *string) (*string, error) {
	filter := fmt.Sprintf("displayName eq '%s'", displayName)
	result, _, err := client.List(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("unable to list groups: %+v", err)
	}

	for _, r := range *result {
		if r.ID == nil {
			return nil, fmt.Errorf("group returned with nil ID")
		}
		if existingID != nil && *existingID == *r.ID {
			continue
		}
		if r.DisplayName != nil && strings.EqualFold(displayName, *r.DisplayName) {
			return r.ID, nil
		}
	}

	return nil, nil
}
