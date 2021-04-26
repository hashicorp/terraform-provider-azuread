package msgraph

import (
	"context"
	"fmt"

	"github.com/manicminer/hamilton/msgraph"
)

func DirectoryRoleTemplateFindByName(ctx context.Context, client *msgraph.DirectoryRoleTemplatesClient, displayName string) (*msgraph.DirectoryRoleTemplate, error) {
	result, _, err := client.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to list Directory Role Templates: %+v", err)
	}

	if result != nil {
		for _, dirRoleTemplate := range *result {
			if dirRoleTemplate.DisplayName != nil && *dirRoleTemplate.DisplayName == displayName {
				return &dirRoleTemplate, nil
			}
		}
	}

	return nil, nil
}
