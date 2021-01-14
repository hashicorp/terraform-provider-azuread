package parse

import "fmt"

type AppRoleId struct {
	ObjectId string
	RoleId   string
}

func NewAppRoleID(objectId, roleId string) AppRoleId {
	return AppRoleId{
		ObjectId: objectId,
		RoleId:   roleId,
	}
}

func (id AppRoleId) String() string {
	return id.ObjectId + "/role/" + id.RoleId
}

func AppRoleID(idString string) (*AppRoleId, error) {
	id, err := ObjectSubResourceID(idString, "role")
	if err != nil {
		return nil, fmt.Errorf("unable to parse App Role ID: %v", err)
	}

	return &AppRoleId{
		ObjectId: id.objectId,
		RoleId:   id.subId,
	}, nil
}
