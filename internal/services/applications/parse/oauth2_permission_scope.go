package parse

import "fmt"

type OAuth2PermissionScopeId struct {
	ObjectId string
	ScopeId  string
}

func NewOAuth2PermissionScopeID(objectId, permissionId string) OAuth2PermissionScopeId {
	return OAuth2PermissionScopeId{
		ObjectId: objectId,
		ScopeId:  permissionId,
	}
}

func (id OAuth2PermissionScopeId) String() string {
	return id.ObjectId + "/scope/" + id.ScopeId
}

func OAuth2PermissionScopeID(idString string) (*OAuth2PermissionScopeId, error) {
	id, err := ObjectSubResourceID(idString, "scope")
	if err != nil {
		return nil, fmt.Errorf("unable to parse OAuth2 Permission ID: %v", err)
	}

	return &OAuth2PermissionScopeId{
		ObjectId: id.objectId,
		ScopeId:  id.subId,
	}, nil
}
