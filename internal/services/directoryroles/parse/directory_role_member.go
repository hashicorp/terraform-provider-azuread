// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import "fmt"

type DirectoryRoleMemberId struct {
	ObjectSubResourceId
	DirectoryRoleId string
	MemberId        string
}

func NewDirectoryRoleMemberID(groupId, memberId string) DirectoryRoleMemberId {
	return DirectoryRoleMemberId{
		ObjectSubResourceId: NewObjectSubResourceID(groupId, "member", memberId),
		DirectoryRoleId:     groupId,
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
