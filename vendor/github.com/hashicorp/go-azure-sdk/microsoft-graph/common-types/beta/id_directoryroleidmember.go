package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRoleIdMemberId{}

// DirectoryRoleIdMemberId is a struct representing the Resource ID for a Directory Role Id Member
type DirectoryRoleIdMemberId struct {
	DirectoryRoleId   string
	DirectoryObjectId string
}

// NewDirectoryRoleIdMemberID returns a new DirectoryRoleIdMemberId struct
func NewDirectoryRoleIdMemberID(directoryRoleId string, directoryObjectId string) DirectoryRoleIdMemberId {
	return DirectoryRoleIdMemberId{
		DirectoryRoleId:   directoryRoleId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseDirectoryRoleIdMemberID parses 'input' into a DirectoryRoleIdMemberId
func ParseDirectoryRoleIdMemberID(input string) (*DirectoryRoleIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRoleIdMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRoleIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryRoleIdMemberIDInsensitively parses 'input' case-insensitively into a DirectoryRoleIdMemberId
// note: this method should only be used for API response data and not user input
func ParseDirectoryRoleIdMemberIDInsensitively(input string) (*DirectoryRoleIdMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRoleIdMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRoleIdMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryRoleIdMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DirectoryRoleId, ok = input.Parsed["directoryRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryRoleId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateDirectoryRoleIdMemberID checks that 'input' can be parsed as a Directory Role Id Member ID
func ValidateDirectoryRoleIdMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryRoleIdMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Role Id Member ID
func (id DirectoryRoleIdMemberId) ID() string {
	fmtString := "/directoryRoles/%s/members/%s"
	return fmt.Sprintf(fmtString, id.DirectoryRoleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Role Id Member ID
func (id DirectoryRoleIdMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directoryRoles", "directoryRoles", "directoryRoles"),
		resourceids.UserSpecifiedSegment("directoryRoleId", "directoryRoleId"),
		resourceids.StaticSegment("members", "members", "members"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Directory Role Id Member ID
func (id DirectoryRoleIdMemberId) String() string {
	components := []string{
		fmt.Sprintf("Directory Role: %q", id.DirectoryRoleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Role Id Member (%s)", strings.Join(components, "\n"))
}
