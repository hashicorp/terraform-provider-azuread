// Copyright (c) HashiCorp, Inc.
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

type GroupOwnerId struct {
	ObjectSubResourceId
	GroupId string
	OwnerId string
}

func NewGroupOwnerID(groupId, ownerId string) GroupOwnerId {
	return GroupOwnerId{
		ObjectSubResourceId: NewObjectSubResourceID(groupId, "owner", ownerId),
		GroupId:             groupId,
		OwnerId:             ownerId,
	}
}

func GroupOwnerID(idString string) (*GroupOwnerId, error) {
	id, err := ObjectSubResourceID(idString, "owner")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Owner ID: %v", err)
	}

	return &GroupOwnerId{
		ObjectSubResourceId: *id,
		GroupId:             id.objectId,
		OwnerId:             id.subId,
	}, nil
}
