package parse

import "fmt"

type DirectoryRoleMemberId struct {
	ObjectSubResourceId
	DirectoryRoleId string
	MemberId        string
}

func NewDirectoryRoleMemberID(directoryRoleId, memberId string) DirectoryRoleMemberId {
	return DirectoryRoleMemberId{
		ObjectSubResourceId: NewObjectSubResourceID(directoryRoleId, "member", memberId),
		DirectoryRoleId:     directoryRoleId,
		MemberId:            memberId,
	}
}

func DirectoryRoleMemberID(idString string) (*DirectoryRoleMemberId, error) {
	id, err := ObjectSubResourceID(idString, "member")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Member ID: %v", err)
	}

	return &DirectoryRoleMemberId{
		ObjectSubResourceId: *id,
		DirectoryRoleId:     id.objectId,
		MemberId:            id.subId,
	}, nil
}
