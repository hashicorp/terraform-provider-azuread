package groups

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
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

func groupGetAdministrativeUnitOfGroup(ctx context.Context, client *msgraph.GroupsClient, auId string, id string) (bool, int, error) {
	resp, status, _, err := client.BaseClient.Get(ctx, msgraph.GetHttpRequestInput{
		ConsistencyFailureFunc: msgraph.RetryOn404ConsistencyFailureFunc,
		OData: odata.Query{
			Select: []string{"id"},
		},
		ValidStatusCodes: []int{http.StatusOK},
		Uri: msgraph.Uri{
			Entity:      fmt.Sprintf("/groups/%s/membersOf/microsoft.graph.administrativeUnit", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return false, status, fmt.Errorf("GroupsClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		Members []struct {
			Id string `json:"id"`
		} `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return false, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	for _, v := range data.Members {
		if v.Id == auId {
			return true, status, nil
		}
	}

	return false, status, nil
}

func groupGetAdditional(ctx context.Context, client *msgraph.GroupsClient, id string) (*msgraph.Group, error) {
	query := odata.Query{Select: []string{"allowExternalSenders", "autoSubscribeNewMembers", "hideFromAddressLists", "hideFromOutlookClients"}}
	groupExtra, _, err := client.Get(ctx, id, query)
	if err != nil {
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
