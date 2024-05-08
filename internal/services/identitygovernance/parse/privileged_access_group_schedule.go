// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type PrivilegedAccessGroupScheduleId struct {
	GroupId      string
	Relationship string
	ScheduleId   string
}

func NewPrivilegedAccessGroupScheduleID(groupId, relationship, scheduleId string) *PrivilegedAccessGroupScheduleId {
	return &PrivilegedAccessGroupScheduleId{
		GroupId:      groupId,
		Relationship: relationship,
		ScheduleId:   scheduleId,
	}
}

func ParsePrivilegedAccessGroupScheduleID(idString string) (*PrivilegedAccessGroupScheduleId, error) {
	// Parse the Schedule ID into its parts
	parts := strings.Split(idString, "_")

	if len(parts) != 3 {
		return nil, fmt.Errorf("parsing GroupScheduleId: expecting 3 parts, got %d", len(parts))
	}

	if _, err := validation.IsUUID(parts[0], "GroupId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing GroupScheduleId: %+v", err)
	}

	if parts[1] != msgraph.PrivilegedAccessGroupRelationshipOwner &&
		parts[1] != msgraph.PrivilegedAccessGroupRelationshipMember {
		return nil, fmt.Errorf("parsing GroupScheduleId: invalid Relationship")
	}

	if _, err := validation.IsUUID(parts[2], "ScheduleId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing GroupScheduleId: %+v", err)
	}

	id := NewPrivilegedAccessGroupScheduleID(parts[0], parts[1], parts[2])

	return id, nil
}

func ValidatePrivilegedAccessGroupScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	_, err := ParsePrivilegedAccessGroupScheduleID(v)
	if err != nil {
		errors = append(errors, err)
	}

	return
}

func (id *PrivilegedAccessGroupScheduleId) ID() string {
	return strings.Join([]string{id.GroupId, id.Relationship, id.ScheduleId}, "_")
}

func (id *PrivilegedAccessGroupScheduleId) String() string {
	return fmt.Sprintf("Privileged Access Group Assigment Schedule ID: %q", id.ID())
}
