package msgraph

import (
	"context"
	"fmt"

	"github.com/manicminer/hamilton/msgraph"
)

func DirectoryRoleFindByName(ctx context.Context, client *msgraph.DirectoryRolesClient, displayName string) ([]*msgraph.DirectoryRole, error) {
	result, _, err := client.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to list Directory Roles: %+v", err)
	}

	var dirRoles []*msgraph.DirectoryRole
	if result != nil {
		for _, dirRole := range *result {
			if dirRole.DisplayName != nil && *dirRole.DisplayName == displayName {
				dirRoles = append(dirRoles, &dirRole)
			}
		}
	}

	return dirRoles, nil
}
