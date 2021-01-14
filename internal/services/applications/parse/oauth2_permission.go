package parse

import "fmt"

type OAuth2PermissionId struct {
	ObjectId     string
	PermissionId string
}

func NewOAuth2PermissionID(objectId, permissionId string) OAuth2PermissionId {
	return OAuth2PermissionId{
		ObjectId:     objectId,
		PermissionId: permissionId,
	}
}

func (id OAuth2PermissionId) String() string {
	return id.ObjectId + "/scope/" + id.PermissionId
}

func OAuth2PermissionID(idString string) (*OAuth2PermissionId, error) {
	id, err := ObjectSubResourceID(idString, "scope")
	if err != nil {
		return nil, fmt.Errorf("unable to parse OAuth2 Permission ID: %v", err)
	}

	return &OAuth2PermissionId{
		ObjectId:     id.objectId,
		PermissionId: id.subId,
	}, nil
}
