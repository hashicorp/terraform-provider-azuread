package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleIdScopedMemberId{}

// DirectoryRoleIdScopedMemberId is a struct representing the Resource ID for a Directory Role Id Scoped Member
type DirectoryRoleIdScopedMemberId struct {
	DirectoryRoleId        string
	ScopedRoleMembershipId string
}

// NewDirectoryRoleIdScopedMemberID returns a new DirectoryRoleIdScopedMemberId struct
func NewDirectoryRoleIdScopedMemberID(directoryRoleId string, scopedRoleMembershipId string) DirectoryRoleIdScopedMemberId {
	return DirectoryRoleIdScopedMemberId{
		DirectoryRoleId:        directoryRoleId,
		ScopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// ParseDirectoryRoleIdScopedMemberID parses 'input' into a DirectoryRoleIdScopedMemberId
func ParseDirectoryRoleIdScopedMemberID(input string) (*DirectoryRoleIdScopedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRoleIdScopedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRoleIdScopedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryRoleIdScopedMemberIDInsensitively parses 'input' case-insensitively into a DirectoryRoleIdScopedMemberId
// note: this method should only be used for API response data and not user input
func ParseDirectoryRoleIdScopedMemberIDInsensitively(input string) (*DirectoryRoleIdScopedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRoleIdScopedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRoleIdScopedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryRoleIdScopedMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryRoleId, ok = input.Parsed["directoryRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryRoleId", input)
	}

	if id.ScopedRoleMembershipId, ok = input.Parsed["scopedRoleMembershipId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", input)
	}

	return nil
}

// ValidateDirectoryRoleIdScopedMemberID checks that 'input' can be parsed as a Directory Role Id Scoped Member ID
func ValidateDirectoryRoleIdScopedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryRoleIdScopedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Role Id Scoped Member ID
func (id DirectoryRoleIdScopedMemberId) ID() string {
	fmtString := "/directoryRoles/%s/scopedMembers/%s"
	return fmt.Sprintf(fmtString, id.DirectoryRoleId, id.ScopedRoleMembershipId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Role Id Scoped Member ID
func (id DirectoryRoleIdScopedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directoryRoles", "directoryRoles", "directoryRoles"),
		resourceids.UserSpecifiedSegment("directoryRoleId", "directoryRoleId"),
		resourceids.StaticSegment("scopedMembers", "scopedMembers", "scopedMembers"),
		resourceids.UserSpecifiedSegment("scopedRoleMembershipId", "scopedRoleMembershipId"),
	}
}

// String returns a human-readable description of this Directory Role Id Scoped Member ID
func (id DirectoryRoleIdScopedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Directory Role: %q", id.DirectoryRoleId),
		fmt.Sprintf("Scoped Role Membership: %q", id.ScopedRoleMembershipId),
	}
	return fmt.Sprintf("Directory Role Id Scoped Member (%s)", strings.Join(components, "\n"))
}
