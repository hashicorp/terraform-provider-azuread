// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import "fmt"

type GroupMemberId struct {
	ObjectSubResourceId
	GroupId  string
	MemberId string
}

func NewGroupMemberID(groupId, memberId string) GroupMemberId {
	return GroupMemberId{
		ObjectSubResourceId: NewObjectSubResourceID(groupId, "member", memberId),
		GroupId:             groupId,
		MemberId:            memberId,
	}
}

func GroupMemberID(idString string) (*GroupMemberId, error) {
	id, err := ObjectSubResourceID(idString, "member")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Member ID: %v", err)
	}

	return &GroupMemberId{
		ObjectSubResourceId: *id,
		GroupId:             id.objectId,
		MemberId:            id.subId,
	}, nil
}
