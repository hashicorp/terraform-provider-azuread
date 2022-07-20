package helpers

import (
	"github.com/manicminer/hamilton/msgraph"
)

func AppRolesFilterByOrigin(in *[]msgraph.AppRole, filterOrigin string) *[]msgraph.AppRole {
	if in == nil {
		return nil
	}

	var result []msgraph.AppRole
	result = []msgraph.AppRole{}

	for _, appRole := range *in {
		if appRole.Origin != nil && *appRole.Origin == filterOrigin {
			result = append(result, appRole)
		}
	}
	return &result
}
