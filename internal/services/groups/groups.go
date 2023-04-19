package groups

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/manicminer/hamilton/msgraph"
)

func groupDefaultMailNickname() string {
	charSet := "0123456789abcdef"
	result := make([]byte, 9)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 9; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	resultString := string(result)
	return resultString[:8] + "-" + resultString[8:]
}

func groupFindByName(ctx context.Context, client *msgraph.GroupsClient, displayName string) (*[]msgraph.Group, error) {
	query := odata.Query{
		Filter: fmt.Sprintf("displayName eq '%s'", displayName),
	}
	groups, _, err := client.List(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to list Groups with filter %q: %+v", query.Filter, err)
	}

	result := make([]msgraph.Group, 0)
	if groups != nil {
		for _, group := range *groups {
			if group.DisplayName != nil && *group.DisplayName == displayName {
				result = append(result, group)
			}
		}
	}

	return &result, nil
}

func groupGetAdditional(ctx context.Context, client *msgraph.GroupsClient, id string) (*msgraph.Group, error) {
	query := odata.Query{Select: []string{"allowExternalSenders", "autoSubscribeNewMembers", "hideFromAddressLists", "hideFromOutlookClients"}}
	groupExtra, status, err := client.Get(ctx, id, query)
	if err != nil {
		if status == http.StatusNotFound {
			// API returns 404 when these M365-only fields are requested for a group in a non-M365 tenant, so we
			// don't raise an error in this case and proceed as if they are not set.
			// See https://github.com/microsoftgraph/msgraph-metadata/issues/333
			return nil, nil
		}
		return nil, fmt.Errorf("retrieving additional fields: %+v", err)
	}
	return groupExtra, nil
}

func hasGroupType(groupTypes []msgraph.GroupType, value msgraph.GroupType) bool {
	for _, v := range groupTypes {
		if value == v {
			return true
		}
	}
	return false
}
